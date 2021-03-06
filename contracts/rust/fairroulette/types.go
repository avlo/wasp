// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package fairroulette

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

type Bet struct {
	Amount int64
	Better *wasmlib.ScAgentId
	Number int64
}

func NewBetFromBytes(bytes []byte) *Bet {
	decode := wasmlib.NewBytesDecoder(bytes)
	data := &Bet{}
	data.Amount = decode.Int()
	data.Better = decode.AgentId()
	data.Number = decode.Int()
	return data
}

func (o *Bet) Bytes() []byte {
	return wasmlib.NewBytesEncoder().
		Int(o.Amount).
		AgentId(o.Better).
		Int(o.Number).
		Data()
}
