// Code generated by schema tool; DO NOT EDIT.

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package coreaccounts

import (
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"
)

type ImmutableFoundryCreateNewParams struct {
	Proxy wasmtypes.Proxy
}

func NewImmutableFoundryCreateNewParams() ImmutableFoundryCreateNewParams {
	return ImmutableFoundryCreateNewParams{Proxy: wasmlib.NewParamsProxy()}
}

// token scheme for the new foundry
func (s ImmutableFoundryCreateNewParams) TokenScheme() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.Proxy.Root(ParamTokenScheme))
}

type MutableFoundryCreateNewParams struct {
	Proxy wasmtypes.Proxy
}

// token scheme for the new foundry
func (s MutableFoundryCreateNewParams) TokenScheme() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.Proxy.Root(ParamTokenScheme))
}

type ImmutableFoundryDestroyParams struct {
	Proxy wasmtypes.Proxy
}

func NewImmutableFoundryDestroyParams() ImmutableFoundryDestroyParams {
	return ImmutableFoundryDestroyParams{Proxy: wasmlib.NewParamsProxy()}
}

// serial number of the foundry
func (s ImmutableFoundryDestroyParams) FoundrySN() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.Proxy.Root(ParamFoundrySN))
}

type MutableFoundryDestroyParams struct {
	Proxy wasmtypes.Proxy
}

// serial number of the foundry
func (s MutableFoundryDestroyParams) FoundrySN() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.Proxy.Root(ParamFoundrySN))
}

type ImmutableFoundryModifySupplyParams struct {
	Proxy wasmtypes.Proxy
}

func NewImmutableFoundryModifySupplyParams() ImmutableFoundryModifySupplyParams {
	return ImmutableFoundryModifySupplyParams{Proxy: wasmlib.NewParamsProxy()}
}

// mint (default) or destroy tokens
func (s ImmutableFoundryModifySupplyParams) DestroyTokens() wasmtypes.ScImmutableBool {
	return wasmtypes.NewScImmutableBool(s.Proxy.Root(ParamDestroyTokens))
}

// serial number of the foundry
func (s ImmutableFoundryModifySupplyParams) FoundrySN() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.Proxy.Root(ParamFoundrySN))
}

// positive nonzero amount to mint or destroy
func (s ImmutableFoundryModifySupplyParams) SupplyDeltaAbs() wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(s.Proxy.Root(ParamSupplyDeltaAbs))
}

type MutableFoundryModifySupplyParams struct {
	Proxy wasmtypes.Proxy
}

// mint (default) or destroy tokens
func (s MutableFoundryModifySupplyParams) DestroyTokens() wasmtypes.ScMutableBool {
	return wasmtypes.NewScMutableBool(s.Proxy.Root(ParamDestroyTokens))
}

// serial number of the foundry
func (s MutableFoundryModifySupplyParams) FoundrySN() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.Proxy.Root(ParamFoundrySN))
}

// positive nonzero amount to mint or destroy
func (s MutableFoundryModifySupplyParams) SupplyDeltaAbs() wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(s.Proxy.Root(ParamSupplyDeltaAbs))
}

type ImmutableTransferAccountToChainParams struct {
	Proxy wasmtypes.Proxy
}

func NewImmutableTransferAccountToChainParams() ImmutableTransferAccountToChainParams {
	return ImmutableTransferAccountToChainParams{Proxy: wasmlib.NewParamsProxy()}
}

// Optional gas amount to reserve in the allowance for the internal
// call to transferAllowanceTo(). Default 100_000 (MinGasFee).
func (s ImmutableTransferAccountToChainParams) GasReserve() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.Proxy.Root(ParamGasReserve))
}

type MutableTransferAccountToChainParams struct {
	Proxy wasmtypes.Proxy
}

// Optional gas amount to reserve in the allowance for the internal
// call to transferAllowanceTo(). Default 100_000 (MinGasFee).
func (s MutableTransferAccountToChainParams) GasReserve() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.Proxy.Root(ParamGasReserve))
}

type ImmutableTransferAllowanceToParams struct {
	Proxy wasmtypes.Proxy
}

func NewImmutableTransferAllowanceToParams() ImmutableTransferAllowanceToParams {
	return ImmutableTransferAllowanceToParams{Proxy: wasmlib.NewParamsProxy()}
}

// The target L2 account
func (s ImmutableTransferAllowanceToParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.Proxy.Root(ParamAgentID))
}

type MutableTransferAllowanceToParams struct {
	Proxy wasmtypes.Proxy
}

// The target L2 account
func (s MutableTransferAllowanceToParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.Proxy.Root(ParamAgentID))
}

type ImmutableAccountFoundriesParams struct {
	Proxy wasmtypes.Proxy
}

func NewImmutableAccountFoundriesParams() ImmutableAccountFoundriesParams {
	return ImmutableAccountFoundriesParams{Proxy: wasmlib.NewParamsProxy()}
}

// account agent ID
func (s ImmutableAccountFoundriesParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.Proxy.Root(ParamAgentID))
}

type MutableAccountFoundriesParams struct {
	Proxy wasmtypes.Proxy
}

// account agent ID
func (s MutableAccountFoundriesParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.Proxy.Root(ParamAgentID))
}

type ImmutableAccountNFTAmountParams struct {
	Proxy wasmtypes.Proxy
}

func NewImmutableAccountNFTAmountParams() ImmutableAccountNFTAmountParams {
	return ImmutableAccountNFTAmountParams{Proxy: wasmlib.NewParamsProxy()}
}

// account agent ID
func (s ImmutableAccountNFTAmountParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.Proxy.Root(ParamAgentID))
}

type MutableAccountNFTAmountParams struct {
	Proxy wasmtypes.Proxy
}

// account agent ID
func (s MutableAccountNFTAmountParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.Proxy.Root(ParamAgentID))
}

type ImmutableAccountNFTAmountInCollectionParams struct {
	Proxy wasmtypes.Proxy
}

func NewImmutableAccountNFTAmountInCollectionParams() ImmutableAccountNFTAmountInCollectionParams {
	return ImmutableAccountNFTAmountInCollectionParams{Proxy: wasmlib.NewParamsProxy()}
}

// account agent ID
func (s ImmutableAccountNFTAmountInCollectionParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.Proxy.Root(ParamAgentID))
}

// NFT ID of collection
func (s ImmutableAccountNFTAmountInCollectionParams) Collection() wasmtypes.ScImmutableNftID {
	return wasmtypes.NewScImmutableNftID(s.Proxy.Root(ParamCollection))
}

type MutableAccountNFTAmountInCollectionParams struct {
	Proxy wasmtypes.Proxy
}

// account agent ID
func (s MutableAccountNFTAmountInCollectionParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.Proxy.Root(ParamAgentID))
}

// NFT ID of collection
func (s MutableAccountNFTAmountInCollectionParams) Collection() wasmtypes.ScMutableNftID {
	return wasmtypes.NewScMutableNftID(s.Proxy.Root(ParamCollection))
}

type ImmutableAccountNFTsParams struct {
	Proxy wasmtypes.Proxy
}

func NewImmutableAccountNFTsParams() ImmutableAccountNFTsParams {
	return ImmutableAccountNFTsParams{Proxy: wasmlib.NewParamsProxy()}
}

// account agent ID
func (s ImmutableAccountNFTsParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.Proxy.Root(ParamAgentID))
}

type MutableAccountNFTsParams struct {
	Proxy wasmtypes.Proxy
}

// account agent ID
func (s MutableAccountNFTsParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.Proxy.Root(ParamAgentID))
}

type ImmutableAccountNFTsInCollectionParams struct {
	Proxy wasmtypes.Proxy
}

func NewImmutableAccountNFTsInCollectionParams() ImmutableAccountNFTsInCollectionParams {
	return ImmutableAccountNFTsInCollectionParams{Proxy: wasmlib.NewParamsProxy()}
}

// account agent ID
func (s ImmutableAccountNFTsInCollectionParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.Proxy.Root(ParamAgentID))
}

// NFT ID of collection
func (s ImmutableAccountNFTsInCollectionParams) Collection() wasmtypes.ScImmutableNftID {
	return wasmtypes.NewScImmutableNftID(s.Proxy.Root(ParamCollection))
}

type MutableAccountNFTsInCollectionParams struct {
	Proxy wasmtypes.Proxy
}

// account agent ID
func (s MutableAccountNFTsInCollectionParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.Proxy.Root(ParamAgentID))
}

// NFT ID of collection
func (s MutableAccountNFTsInCollectionParams) Collection() wasmtypes.ScMutableNftID {
	return wasmtypes.NewScMutableNftID(s.Proxy.Root(ParamCollection))
}

type ImmutableBalanceParams struct {
	Proxy wasmtypes.Proxy
}

func NewImmutableBalanceParams() ImmutableBalanceParams {
	return ImmutableBalanceParams{Proxy: wasmlib.NewParamsProxy()}
}

// account agent ID
func (s ImmutableBalanceParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.Proxy.Root(ParamAgentID))
}

type MutableBalanceParams struct {
	Proxy wasmtypes.Proxy
}

// account agent ID
func (s MutableBalanceParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.Proxy.Root(ParamAgentID))
}

type ImmutableBalanceBaseTokenParams struct {
	Proxy wasmtypes.Proxy
}

func NewImmutableBalanceBaseTokenParams() ImmutableBalanceBaseTokenParams {
	return ImmutableBalanceBaseTokenParams{Proxy: wasmlib.NewParamsProxy()}
}

// account agent ID
func (s ImmutableBalanceBaseTokenParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.Proxy.Root(ParamAgentID))
}

type MutableBalanceBaseTokenParams struct {
	Proxy wasmtypes.Proxy
}

// account agent ID
func (s MutableBalanceBaseTokenParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.Proxy.Root(ParamAgentID))
}

type ImmutableBalanceNativeTokenParams struct {
	Proxy wasmtypes.Proxy
}

func NewImmutableBalanceNativeTokenParams() ImmutableBalanceNativeTokenParams {
	return ImmutableBalanceNativeTokenParams{Proxy: wasmlib.NewParamsProxy()}
}

// account agent ID
func (s ImmutableBalanceNativeTokenParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.Proxy.Root(ParamAgentID))
}

// native token ID
func (s ImmutableBalanceNativeTokenParams) TokenID() wasmtypes.ScImmutableTokenID {
	return wasmtypes.NewScImmutableTokenID(s.Proxy.Root(ParamTokenID))
}

type MutableBalanceNativeTokenParams struct {
	Proxy wasmtypes.Proxy
}

// account agent ID
func (s MutableBalanceNativeTokenParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.Proxy.Root(ParamAgentID))
}

// native token ID
func (s MutableBalanceNativeTokenParams) TokenID() wasmtypes.ScMutableTokenID {
	return wasmtypes.NewScMutableTokenID(s.Proxy.Root(ParamTokenID))
}

type ImmutableFoundryOutputParams struct {
	Proxy wasmtypes.Proxy
}

func NewImmutableFoundryOutputParams() ImmutableFoundryOutputParams {
	return ImmutableFoundryOutputParams{Proxy: wasmlib.NewParamsProxy()}
}

// serial number of the foundry
func (s ImmutableFoundryOutputParams) FoundrySN() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.Proxy.Root(ParamFoundrySN))
}

type MutableFoundryOutputParams struct {
	Proxy wasmtypes.Proxy
}

// serial number of the foundry
func (s MutableFoundryOutputParams) FoundrySN() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.Proxy.Root(ParamFoundrySN))
}

type ImmutableGetAccountNonceParams struct {
	Proxy wasmtypes.Proxy
}

func NewImmutableGetAccountNonceParams() ImmutableGetAccountNonceParams {
	return ImmutableGetAccountNonceParams{Proxy: wasmlib.NewParamsProxy()}
}

// account agent ID
func (s ImmutableGetAccountNonceParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.Proxy.Root(ParamAgentID))
}

type MutableGetAccountNonceParams struct {
	Proxy wasmtypes.Proxy
}

// account agent ID
func (s MutableGetAccountNonceParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.Proxy.Root(ParamAgentID))
}

type ImmutableNftDataParams struct {
	Proxy wasmtypes.Proxy
}

func NewImmutableNftDataParams() ImmutableNftDataParams {
	return ImmutableNftDataParams{Proxy: wasmlib.NewParamsProxy()}
}

// NFT ID
func (s ImmutableNftDataParams) NftID() wasmtypes.ScImmutableNftID {
	return wasmtypes.NewScImmutableNftID(s.Proxy.Root(ParamNftID))
}

type MutableNftDataParams struct {
	Proxy wasmtypes.Proxy
}

// NFT ID
func (s MutableNftDataParams) NftID() wasmtypes.ScMutableNftID {
	return wasmtypes.NewScMutableNftID(s.Proxy.Root(ParamNftID))
}
