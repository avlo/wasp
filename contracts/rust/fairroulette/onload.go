// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package fairroulette

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncLockBets, funcLockBetsThunk)
	exports.AddFunc(FuncPayWinners, funcPayWinnersThunk)
	exports.AddFunc(FuncPlaceBet, funcPlaceBetThunk)
	exports.AddFunc(FuncPlayPeriod, funcPlayPeriodThunk)
}

type FuncLockBetsParams struct {
}

func funcLockBetsThunk(ctx *wasmlib.ScFuncContext) {
	// only SC itself can invoke this function
	ctx.Require(ctx.From(ctx.ContractId().AsAgentId()), "no permission")

	params := &FuncLockBetsParams{
	}
	funcLockBets(ctx, params)
}

type FuncPayWinnersParams struct {
}

func funcPayWinnersThunk(ctx *wasmlib.ScFuncContext) {
	// only SC itself can invoke this function
	ctx.Require(ctx.From(ctx.ContractId().AsAgentId()), "no permission")

	params := &FuncPayWinnersParams{
	}
	funcPayWinners(ctx, params)
}

type FuncPlaceBetParams struct {
	Number wasmlib.ScImmutableInt // the number a better bets on
}

func funcPlaceBetThunk(ctx *wasmlib.ScFuncContext) {
	p := ctx.Params()
	params := &FuncPlaceBetParams{
		Number: p.GetInt(ParamNumber),
	}
	ctx.Require(params.Number.Exists(), "missing mandatory number")
	funcPlaceBet(ctx, params)
}

type FuncPlayPeriodParams struct {
	PlayPeriod wasmlib.ScImmutableInt // number of minutes in one playing round
}

func funcPlayPeriodThunk(ctx *wasmlib.ScFuncContext) {
	// only SC creator can update the play period
	ctx.Require(ctx.From(ctx.ContractCreator()), "no permission")

	p := ctx.Params()
	params := &FuncPlayPeriodParams{
		PlayPeriod: p.GetInt(ParamPlayPeriod),
	}
	ctx.Require(params.PlayPeriod.Exists(), "missing mandatory playPeriod")
	funcPlayPeriod(ctx, params)
}
