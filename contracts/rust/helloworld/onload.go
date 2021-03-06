// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package helloworld

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncHelloWorld, funcHelloWorldThunk)
	exports.AddView(ViewGetHelloWorld, viewGetHelloWorldThunk)
}

type FuncHelloWorldParams struct {
}

func funcHelloWorldThunk(ctx *wasmlib.ScFuncContext) {
	params := &FuncHelloWorldParams{
	}
	funcHelloWorld(ctx, params)
}

type ViewGetHelloWorldParams struct {
}

func viewGetHelloWorldThunk(ctx *wasmlib.ScViewContext) {
	params := &ViewGetHelloWorldParams{
	}
	viewGetHelloWorld(ctx, params)
}
