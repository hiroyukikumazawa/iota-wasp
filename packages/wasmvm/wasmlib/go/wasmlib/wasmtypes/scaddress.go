// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmtypes

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

const (
	ScAddressAlias   byte = 8
	ScAddressEd25519 byte = 0
	ScAddressNFT     byte = 16
	ScAddressEth     byte = 32

	ScLengthAlias   = 33
	ScLengthEd25519 = 33
	ScLengthNFT     = 33
	ScLengthEth     = 20

	ScAddressLength = ScLengthEd25519
)

type ScAddress struct {
	id [ScAddressLength]byte
}

func (o ScAddress) AsAgentID() ScAgentID {
	return NewScAgentIDFromAddress(o)
}

func (o ScAddress) Bytes() []byte {
	return AddressToBytes(o)
}

func (o ScAddress) String() string {
	return AddressToString(o)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// TODO address type-dependent encoding/decoding?
func AddressDecode(dec *WasmDecoder) ScAddress {
	addr := ScAddress{}
	copy(addr.id[:], dec.FixedBytes(ScAddressLength))
	return addr
}

func AddressEncode(enc *WasmEncoder, value ScAddress) {
	enc.FixedBytes(value.id[:], ScAddressLength)
}

func AddressFromBytes(buf []byte) ScAddress {
	addr := ScAddress{}
	if len(buf) == 0 {
		return addr
	}

	// special case, ETH address has no type byte but different length
	if len(buf) == ScLengthEth {
		addr.id[0] = ScAddressEth
		copy(addr.id[1:], buf)
		return addr
	}

	switch buf[0] {
	case ScAddressAlias:
		if len(buf) != ScLengthAlias {
			panic("invalid Address length: Alias")
		}
	case ScAddressEd25519:
		if len(buf) != ScLengthEd25519 {
			panic("invalid Address length: Ed25519")
		}
	case ScAddressNFT:
		if len(buf) != ScLengthNFT {
			panic("invalid Address length: NFT")
		}
	default:
		panic("invalid Address type")
	}
	copy(addr.id[:], buf)
	return addr
}

func AddressToBytes(value ScAddress) []byte {
	switch value.id[0] {
	case ScAddressAlias:
		return value.id[:ScLengthAlias]
	case ScAddressEd25519:
		return value.id[:ScLengthEd25519]
	case ScAddressNFT:
		return value.id[:ScLengthNFT]
	case ScAddressEth:
		return value.id[1 : ScLengthEth+1]
	default:
		panic("unexpected Address type")
	}
}

func AddressFromString(value string) ScAddress {
	if value[:2] == "0x" {
		return AddressFromBytes(HexDecode(value))
	}
	return Bech32Decode(value)
}

func AddressToString(value ScAddress) string {
	if value.id[0] == ScAddressEth {
		return HexEncode(AddressToBytes(value))
	}
	return Bech32Encode(value)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableAddress struct {
	proxy Proxy
}

func NewScImmutableAddress(proxy Proxy) ScImmutableAddress {
	return ScImmutableAddress{proxy: proxy}
}

func (o ScImmutableAddress) Exists() bool {
	return o.proxy.Exists()
}

func (o ScImmutableAddress) String() string {
	return AddressToString(o.Value())
}

func (o ScImmutableAddress) Value() ScAddress {
	return AddressFromBytes(o.proxy.Get())
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableAddress struct {
	ScImmutableAddress
}

func NewScMutableAddress(proxy Proxy) ScMutableAddress {
	return ScMutableAddress{ScImmutableAddress{proxy: proxy}}
}

func (o ScMutableAddress) Delete() {
	o.proxy.Delete()
}

func (o ScMutableAddress) SetValue(value ScAddress) {
	o.proxy.Set(AddressToBytes(value))
}
