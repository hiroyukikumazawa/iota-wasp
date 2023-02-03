/*
Wasp API

REST API for the Wasp node

API version: 0.4.0-alpha.8-16-g83edf92b9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the L1Params type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &L1Params{}

// L1Params struct for L1Params
type L1Params struct {
	BaseToken BaseToken `json:"baseToken"`
	// The max payload size
	MaxPayloadSize int32 `json:"maxPayloadSize"`
	Protocol ProtocolParameters `json:"protocol"`
}

// NewL1Params instantiates a new L1Params object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewL1Params(baseToken BaseToken, maxPayloadSize int32, protocol ProtocolParameters) *L1Params {
	this := L1Params{}
	this.BaseToken = baseToken
	this.MaxPayloadSize = maxPayloadSize
	this.Protocol = protocol
	return &this
}

// NewL1ParamsWithDefaults instantiates a new L1Params object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewL1ParamsWithDefaults() *L1Params {
	this := L1Params{}
	return &this
}

// GetBaseToken returns the BaseToken field value
func (o *L1Params) GetBaseToken() BaseToken {
	if o == nil {
		var ret BaseToken
		return ret
	}

	return o.BaseToken
}

// GetBaseTokenOk returns a tuple with the BaseToken field value
// and a boolean to check if the value has been set.
func (o *L1Params) GetBaseTokenOk() (*BaseToken, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseToken, true
}

// SetBaseToken sets field value
func (o *L1Params) SetBaseToken(v BaseToken) {
	o.BaseToken = v
}

// GetMaxPayloadSize returns the MaxPayloadSize field value
func (o *L1Params) GetMaxPayloadSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxPayloadSize
}

// GetMaxPayloadSizeOk returns a tuple with the MaxPayloadSize field value
// and a boolean to check if the value has been set.
func (o *L1Params) GetMaxPayloadSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxPayloadSize, true
}

// SetMaxPayloadSize sets field value
func (o *L1Params) SetMaxPayloadSize(v int32) {
	o.MaxPayloadSize = v
}

// GetProtocol returns the Protocol field value
func (o *L1Params) GetProtocol() ProtocolParameters {
	if o == nil {
		var ret ProtocolParameters
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *L1Params) GetProtocolOk() (*ProtocolParameters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *L1Params) SetProtocol(v ProtocolParameters) {
	o.Protocol = v
}

func (o L1Params) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o L1Params) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["baseToken"] = o.BaseToken
	toSerialize["maxPayloadSize"] = o.MaxPayloadSize
	toSerialize["protocol"] = o.Protocol
	return toSerialize, nil
}

type NullableL1Params struct {
	value *L1Params
	isSet bool
}

func (v NullableL1Params) Get() *L1Params {
	return v.value
}

func (v *NullableL1Params) Set(val *L1Params) {
	v.value = val
	v.isSet = true
}

func (v NullableL1Params) IsSet() bool {
	return v.isSet
}

func (v *NullableL1Params) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableL1Params(val *L1Params) *NullableL1Params {
	return &NullableL1Params{value: val, isSet: true}
}

func (v NullableL1Params) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableL1Params) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


