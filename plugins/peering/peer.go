package peering

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/iotaledger/goshimmer/dapps/waspconn/packages/chopper"
	"github.com/iotaledger/goshimmer/packages/tangle"
	"github.com/iotaledger/hive.go/backoff"
	"go.uber.org/atomic"
)

// represents point-to-point TCP connection between two qnodes and another
// it is used as transport for message exchange
// Another end is always using the same connection
// the Peer takes care about exchanging heartbeat messages.
// It keeps last several received heartbeats as "lad" data to be able to calculate how synced/unsynced
// clocks of peer are.
type Peer struct {
	*sync.RWMutex
	isDismissed atomic.Bool       // to be GC-ed
	peerconn    *peeredConnection // nil means not connected
	handshakeOk bool
	// network locations as taken from the SC data
	remoteNetid string

	startOnce *sync.Once
	numUsers  int
}

// retry net.Dial once, on fail after 0.5s
var dialRetryPolicy = backoff.ConstantBackOff(backoffDelay).With(backoff.MaxRetries(dialRetries))

func isInbound(remoteLocation string) bool {
	if remoteLocation == MyNetworkId() {
		panic("remoteNetid == myLocation")
	}
	return remoteLocation < MyNetworkId()
}

func (peer *Peer) isInbound() bool {
	return isInbound(peer.remoteNetid)
}

func peeringId(remoteLocation string) string {
	if isInbound(remoteLocation) {
		return remoteLocation + "<" + MyNetworkId()
	} else {
		return MyNetworkId() + "<" + remoteLocation
	}
}

func (peer *Peer) PeeringId() string {
	return peeringId(peer.remoteNetid)
}

// return true if is alive and average latencyRingBuf in nanosec
func (peer *Peer) IsAlive() bool {
	peer.RLock()
	defer peer.RUnlock()
	return peer.peerconn != nil && peer.handshakeOk
}

func (peer *Peer) NumUsers() int {
	peer.RLock()
	defer peer.RUnlock()
	return peer.numUsers
}

func (peer *Peer) connStatus() (bool, bool) {
	peer.RLock()
	defer peer.RUnlock()
	if peer.isDismissed.Load() {
		return false, false
	}
	return peer.peerconn != nil, peer.handshakeOk
}

func (peer *Peer) closeConn() {
	peer.Lock()
	defer peer.Unlock()

	if peer.isDismissed.Load() {
		return
	}
	if peer.peerconn != nil {
		_ = peer.peerconn.Close()
	}
}

// dials outbound address and established connection
func (peer *Peer) runOutbound() {
	if peer.isDismissed.Load() {
		return
	}
	if peer.isInbound() {
		return
	}
	if peer.peerconn != nil {
		panic("peer.peerconn != nil")
	}
	log.Debugf("runOutbound %s", peer.remoteNetid)

	// always try to reconnect
	defer func() {
		go func() {
			time.Sleep(restartAfter)
			peer.Lock()
			if !peer.isDismissed.Load() {
				peer.startOnce = &sync.Once{}
				log.Debugf("will run again: %s", peer.PeeringId())
			}
			peer.Unlock()
		}()
	}()

	var conn net.Conn

	if err := backoff.Retry(dialRetryPolicy, func() error {
		var err error
		conn, err = net.DialTimeout("tcp", peer.remoteNetid, dialTimeout)
		if err != nil {
			return fmt.Errorf("dial %s failed: %w", peer.remoteNetid, err)
		}
		return nil
	}); err != nil {
		log.Warn(err)
		return
	}
	peer.peerconn = newPeeredConnection(conn, peer)
	if err := peer.sendHandshake(); err != nil {
		log.Errorf("error during sendHandshake: %v", err)
		return
	}
	log.Debugf("starting reading outbound %s", peer.remoteNetid)
	err := peer.peerconn.Read()
	log.Debugw("stopped reading outbound. Closing", "remote", peer.remoteNetid, "err", err)
	peer.closeConn()
}

// sends handshake message. It contains myLocation
func (peer *Peer) sendHandshake() error {
	data := encodeMessage(&PeerMessage{
		MsgType: MsgTypeHandshake,
		MsgData: []byte(peer.PeeringId()),
	}, time.Now().UnixNano())
	_, err := peer.peerconn.Write(data)
	log.Debugf("sendHandshake '%s' --> '%s', id = %s", MyNetworkId(), peer.remoteNetid, peer.PeeringId())
	return err
}

func (peer *Peer) SendMsg(msg *PeerMessage) error {
	if msg.MsgType < FirstCommitteeMsgCode {
		return errors.New("reserved message code")
	}
	data := encodeMessage(msg, time.Now().UnixNano())

	choppedData, chopped := chopper.ChopData(data, tangle.MaxMessageSize-chunkMessageOverhead)

	peer.RLock()
	defer peer.RUnlock()

	if !chopped {
		return peer.sendData(data)
	}
	return peer.sendChunks(choppedData)
}

func (peer *Peer) sendChunks(chopped [][]byte) error {
	ts := time.Now().UnixNano()
	for _, piece := range chopped {
		d := encodeMessage(&PeerMessage{
			MsgType: MsgTypeMsgChunk,
			MsgData: piece,
		}, ts)
		if err := peer.sendData(d); err != nil {
			return err
		}
	}
	return nil
}

// SendMsgToPeers sends same msg to all peers in the slice which are not nil
// with the same timestamp
// return number of successfully sent messages and timestamp
func SendMsgToPeers(msg *PeerMessage, ts int64, peers ...*Peer) uint16 {
	if msg.MsgType < FirstCommitteeMsgCode {
		return 0
	}
	// timestamped here, once
	data := encodeMessage(msg, ts)
	choppedData, chopped := chopper.ChopData(data, tangle.MaxMessageSize-chunkMessageOverhead)

	numSent := uint16(0)
	for _, peer := range peers {
		if peer == nil {
			continue
		}
		peer.RLock()
		if !chopped {
			if err := peer.sendData(data); err == nil {
				numSent++
			}
		} else {
			if err := peer.sendChunks(choppedData); err == nil {
				numSent++
			}
		}
		peer.RUnlock()
	}
	return numSent
}

func (peer *Peer) sendData(data []byte) error {
	if peer.peerconn == nil {
		return fmt.Errorf("no connection with %s", peer.remoteNetid)
	}
	num, err := peer.peerconn.Write(data)
	if num != len(data) {
		return fmt.Errorf("not all bytes were written. err = %v", err)
	}
	return nil
}
