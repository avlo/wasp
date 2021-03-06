// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

use fairauction::*;
use schema::*;
use wasmlib::*;

mod fairauction;
mod schema;
mod types;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_FINALIZE_AUCTION, func_finalize_auction_thunk);
    exports.add_func(FUNC_PLACE_BID, func_place_bid_thunk);
    exports.add_func(FUNC_SET_OWNER_MARGIN, func_set_owner_margin_thunk);
    exports.add_func(FUNC_START_AUCTION, func_start_auction_thunk);
    exports.add_view(VIEW_GET_INFO, view_get_info_thunk);
}

pub struct FuncFinalizeAuctionParams {
    pub color: ScImmutableColor, // color identifies the auction
}

fn func_finalize_auction_thunk(ctx: &ScFuncContext) {
    // only SC itself can invoke this function
    ctx.require(ctx.caller() == ctx.contract_id().as_agent_id(), "no permission");

    let p = ctx.params();
    let params = FuncFinalizeAuctionParams {
        color: p.get_color(PARAM_COLOR),
    };
    ctx.require(params.color.exists(), "missing mandatory color");
    func_finalize_auction(ctx, &params);
}

pub struct FuncPlaceBidParams {
    pub color: ScImmutableColor, // color identifies the auction
}

fn func_place_bid_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncPlaceBidParams {
        color: p.get_color(PARAM_COLOR),
    };
    ctx.require(params.color.exists(), "missing mandatory color");
    func_place_bid(ctx, &params);
}

pub struct FuncSetOwnerMarginParams {
    pub owner_margin: ScImmutableInt, // new SC owner margin in promilles
}

fn func_set_owner_margin_thunk(ctx: &ScFuncContext) {
    // only SC creator can set owner margin
    ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

    let p = ctx.params();
    let params = FuncSetOwnerMarginParams {
        owner_margin: p.get_int(PARAM_OWNER_MARGIN),
    };
    ctx.require(params.owner_margin.exists(), "missing mandatory ownerMargin");
    func_set_owner_margin(ctx, &params);
}

//@formatter:off
pub struct FuncStartAuctionParams {
    pub color:       ScImmutableColor,  // color of the tokens being auctioned
    pub description: ScImmutableString, // description of the tokens being auctioned
    pub duration:    ScImmutableInt,    // duration of auction in minutes
    pub minimum_bid: ScImmutableInt,    // minimum required amount for any bid
}
//@formatter:on

fn func_start_auction_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncStartAuctionParams {
        color: p.get_color(PARAM_COLOR),
        description: p.get_string(PARAM_DESCRIPTION),
        duration: p.get_int(PARAM_DURATION),
        minimum_bid: p.get_int(PARAM_MINIMUM_BID),
    };
    ctx.require(params.color.exists(), "missing mandatory color");
    ctx.require(params.minimum_bid.exists(), "missing mandatory minimumBid");
    func_start_auction(ctx, &params);
}

pub struct ViewGetInfoParams {
    pub color: ScImmutableColor, // color identifies the auction
}

fn view_get_info_thunk(ctx: &ScViewContext) {
    let p = ctx.params();
    let params = ViewGetInfoParams {
        color: p.get_color(PARAM_COLOR),
    };
    ctx.require(params.color.exists(), "missing mandatory color");
    view_get_info(ctx, &params);
}
