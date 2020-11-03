package vmcontext

import (
	"fmt"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/sctransaction"
	"github.com/iotaledger/wasp/packages/util"
	"github.com/iotaledger/wasp/packages/vm/vmtypes"
	"github.com/iotaledger/wasp/plugins/publisher"
)

func (vmctx *VMContext) ChainID() coretypes.ChainID {
	return vmctx.chainID
}

func (vmctx *VMContext) ContractIndex() uint16 {
	return vmctx.callStack[len(vmctx.callStack)-1].contractIndex
}

func (vmctx *VMContext) OwnerAddress() *address.Address {
	return &vmctx.ownerAddress
}

func (vmctx *VMContext) Timestamp() int64 {
	return vmctx.timestamp
}

func (vmctx *VMContext) Entropy() hashing.HashValue {
	return vmctx.entropy
}

func (vmctx *VMContext) Log() *logger.Logger {
	return vmctx.log
}

func (vmctx *VMContext) DumpAccount() string {
	return vmctx.txBuilder.Dump()
}

func (vmctx *VMContext) SendRequest(par vmtypes.NewRequestParams) bool {
	if par.IncludeReward > 0 {
		availableIotas := vmctx.txBuilder.GetInputBalance(balance.ColorIOTA)
		if par.IncludeReward+1 > availableIotas {
			return false
		}
		err := vmctx.txBuilder.MoveTokensToAddress((address.Address)(par.TargetContractID.ChainID()), balance.ColorIOTA, par.IncludeReward)
		if err != nil {
			return false
		}
	}
	reqBlock := sctransaction.NewRequestBlock(vmctx.ContractIndex(), par.TargetContractID, par.EntryPoint)
	reqBlock.WithTimelock(par.Timelock)
	reqBlock.SetArgs(par.Params)

	if err := vmctx.txBuilder.AddRequestBlock(reqBlock); err != nil {
		return false
	}
	return true
}

func (vmctx *VMContext) SendRequestToSelf(reqCode coretypes.EntryPointCode, params dict.Dict) bool {
	return vmctx.SendRequest(vmtypes.NewRequestParams{
		TargetContractID: coretypes.NewContractID(vmctx.chainID, vmctx.ContractIndex()),
		EntryPoint:       reqCode,
		Params:           params,
		IncludeReward:    0,
	})
}

func (vmctx *VMContext) SendRequestToSelfWithDelay(entryPoint coretypes.EntryPointCode, args dict.Dict, delaySec uint32) bool {
	timelock := util.NanoSecToUnixSec(vmctx.timestamp) + delaySec

	return vmctx.SendRequest(vmtypes.NewRequestParams{
		TargetContractID: coretypes.NewContractID(vmctx.chainID, vmctx.ContractIndex()),
		EntryPoint:       entryPoint,
		Params:           args,
		Timelock:         timelock,
		IncludeReward:    0,
	})
}

func (vmctx *VMContext) Publish(msg string) {
	publisher.Publish("vmmsg", vmctx.chainID.String(), fmt.Sprintf("%d", vmctx.ContractIndex()), msg)
}

func (vmctx *VMContext) Publishf(format string, args ...interface{}) {
	publisher.Publish("vmmsg", vmctx.chainID.String(), fmt.Sprintf("%d", vmctx.ContractIndex()), fmt.Sprintf(format, args...))
}

func (vmctx *VMContext) Request() *sctransaction.RequestRef {
	return &vmctx.reqRef
}