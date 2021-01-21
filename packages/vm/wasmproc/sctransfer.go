// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmproc

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes/cbalances"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
)

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScTransfers struct {
	ScSandboxObject
}

func NewScTransfers(vm *wasmProcessor) *ScTransfers {
	a := &ScTransfers{}
	a.vm = vm
	return a
}

func (a *ScTransfers) GetObjectId(keyId int32, typeId int32) int32 {
	return GetArrayObjectId(a, keyId, typeId, func() WaspObject {
		return NewScTransferInfo(a.vm)
	})
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScTransferInfo struct {
	ScSandboxObject
	address address.Address
}

func NewScTransferInfo(vm *wasmProcessor) *ScTransferInfo {
	o := &ScTransferInfo{}
	o.vm = vm
	return o
}

func (o *ScTransferInfo) Invoke(balances int32) {
	transfers := make(map[balance.Color]int64)
	balancesDict := o.host.FindObject(balances).(*ScDict).kvStore.(dict.Dict)
	balancesDict.MustIterate("", func(key kv.Key, value []byte) bool {
		color, _, err := codec.DecodeColor([]byte(key))
		if err != nil {
			o.Panic(err.Error())
		}
		amount, _, err := codec.DecodeInt64(value)
		if err != nil {
			o.Panic(err.Error())
		}
		o.Trace("TRANSFER #%d c'%s' a'%s'", value, color.String(), o.address.String())
		transfers[color] = amount
		return true
	})
	if !o.vm.ctx.TransferToAddress(o.address, cbalances.NewFromMap(transfers)) {
		o.Panic("failed to transfer to %s", o.address.String())
	}
}

func (o *ScTransferInfo) SetBytes(keyId int32, value []byte) {
	var err error
	switch keyId {
	case wasmhost.KeyAddress:
		o.address, _, err = address.FromBytes(value)
		if err != nil {
			o.Panic("SetBytes: invalid address: " + err.Error())
		}
	default:
		o.invalidKey(keyId)
	}
}

func (o *ScTransferInfo) SetInt(keyId int32, value int64) {
	switch keyId {
	case wasmhost.KeyBalances:
		o.Invoke(int32(value))
	default:
		o.invalidKey(keyId)
	}
}
