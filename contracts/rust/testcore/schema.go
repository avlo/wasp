// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package testcore

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

const ScName = "testcore"
const ScDescription = "Core test for ISCP wasmlib Rust/Wasm library"
const ScHname = wasmlib.ScHname(0x370d33ad)

const ParamAddress = wasmlib.Key("address")
const ParamAgentId = wasmlib.Key("agentID")
const ParamCaller = wasmlib.Key("caller")
const ParamChainId = wasmlib.Key("chainid")
const ParamChainOwnerId = wasmlib.Key("chainOwnerID")
const ParamContractCreator = wasmlib.Key("contractCreator")
const ParamContractId = wasmlib.Key("contractID")
const ParamCounter = wasmlib.Key("counter")
const ParamHash = wasmlib.Key("Hash")
const ParamHname = wasmlib.Key("Hname")
const ParamHnameContract = wasmlib.Key("hnameContract")
const ParamHnameEP = wasmlib.Key("hnameEP")
const ParamHnameZero = wasmlib.Key("Hname-0")
const ParamInt64 = wasmlib.Key("int64")
const ParamInt64Zero = wasmlib.Key("int64-0")
const ParamIntValue = wasmlib.Key("intParamValue")
const ParamName = wasmlib.Key("intParamName")
const ParamString = wasmlib.Key("string")
const ParamStringZero = wasmlib.Key("string-0")

const VarCounter = wasmlib.Key("counter")
const VarHnameEP = wasmlib.Key("hnameEP")

const FuncCallOnChain = "callOnChain"
const FuncCheckContextFromFullEP = "checkContextFromFullEP"
const FuncDoNothing = "doNothing"
const FuncInit = "init"
const FuncPassTypesFull = "passTypesFull"
const FuncRunRecursion = "runRecursion"
const FuncSendToAddress = "sendToAddress"
const FuncSetInt = "setInt"
const FuncTestCallPanicFullEP = "testCallPanicFullEP"
const FuncTestCallPanicViewEPFromFull = "testCallPanicViewEPFromFull"
const FuncTestChainOwnerIDFull = "testChainOwnerIDFull"
const FuncTestContractIDFull = "testContractIDFull"
const FuncTestEventLogDeploy = "testEventLogDeploy"
const FuncTestEventLogEventData = "testEventLogEventData"
const FuncTestEventLogGenericData = "testEventLogGenericData"
const FuncTestPanicFullEP = "testPanicFullEP"
const FuncWithdrawToChain = "withdrawToChain"
const ViewCheckContextFromViewEP = "checkContextFromViewEP"
const ViewFibonacci = "fibonacci"
const ViewGetCounter = "getCounter"
const ViewGetInt = "getInt"
const ViewJustView = "justView"
const ViewPassTypesView = "passTypesView"
const ViewTestCallPanicViewEPFromView = "testCallPanicViewEPFromView"
const ViewTestChainOwnerIDView = "testChainOwnerIDView"
const ViewTestContractIDView = "testContractIDView"
const ViewTestPanicViewEP = "testPanicViewEP"
const ViewTestSandboxCall = "testSandboxCall"

const HFuncCallOnChain = wasmlib.ScHname(0x95a3d123)
const HFuncCheckContextFromFullEP = wasmlib.ScHname(0xa56c24ba)
const HFuncDoNothing = wasmlib.ScHname(0xdda4a6de)
const HFuncInit = wasmlib.ScHname(0x1f44d644)
const HFuncPassTypesFull = wasmlib.ScHname(0x733ea0ea)
const HFuncRunRecursion = wasmlib.ScHname(0x833425fd)
const HFuncSendToAddress = wasmlib.ScHname(0x63ce4634)
const HFuncSetInt = wasmlib.ScHname(0x62056f74)
const HFuncTestCallPanicFullEP = wasmlib.ScHname(0x4c878834)
const HFuncTestCallPanicViewEPFromFull = wasmlib.ScHname(0xfd7e8c1d)
const HFuncTestChainOwnerIDFull = wasmlib.ScHname(0x2aff1167)
const HFuncTestContractIDFull = wasmlib.ScHname(0x95934282)
const HFuncTestEventLogDeploy = wasmlib.ScHname(0x96ff760a)
const HFuncTestEventLogEventData = wasmlib.ScHname(0x0efcf939)
const HFuncTestEventLogGenericData = wasmlib.ScHname(0x6a16629d)
const HFuncTestPanicFullEP = wasmlib.ScHname(0x24fdef07)
const HFuncWithdrawToChain = wasmlib.ScHname(0x437bc026)
const HViewCheckContextFromViewEP = wasmlib.ScHname(0x88ff0167)
const HViewFibonacci = wasmlib.ScHname(0x7940873c)
const HViewGetCounter = wasmlib.ScHname(0xb423e607)
const HViewGetInt = wasmlib.ScHname(0x1887e5ef)
const HViewJustView = wasmlib.ScHname(0x33b8972e)
const HViewPassTypesView = wasmlib.ScHname(0x1a5b87ea)
const HViewTestCallPanicViewEPFromView = wasmlib.ScHname(0x91b10c99)
const HViewTestChainOwnerIDView = wasmlib.ScHname(0x26586c33)
const HViewTestContractIDView = wasmlib.ScHname(0x28a02913)
const HViewTestPanicViewEP = wasmlib.ScHname(0x22bc4d72)
const HViewTestSandboxCall = wasmlib.ScHname(0x42d72b63)
