/*
Wasp API

REST API for the Wasp node

API version: 0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the EstimateGasRequestOffledger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EstimateGasRequestOffledger{}

// EstimateGasRequestOffledger struct for EstimateGasRequestOffledger
type EstimateGasRequestOffledger struct {
	// Offledger Request (Hex)
	RequestBytes string `json:"requestBytes"`
}

// NewEstimateGasRequestOffledger instantiates a new EstimateGasRequestOffledger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimateGasRequestOffledger(requestBytes string) *EstimateGasRequestOffledger {
	this := EstimateGasRequestOffledger{}
	this.RequestBytes = requestBytes
	return &this
}

// NewEstimateGasRequestOffledgerWithDefaults instantiates a new EstimateGasRequestOffledger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimateGasRequestOffledgerWithDefaults() *EstimateGasRequestOffledger {
	this := EstimateGasRequestOffledger{}
	return &this
}

// GetRequestBytes returns the RequestBytes field value
func (o *EstimateGasRequestOffledger) GetRequestBytes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestBytes
}

// GetRequestBytesOk returns a tuple with the RequestBytes field value
// and a boolean to check if the value has been set.
func (o *EstimateGasRequestOffledger) GetRequestBytesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestBytes, true
}

// SetRequestBytes sets field value
func (o *EstimateGasRequestOffledger) SetRequestBytes(v string) {
	o.RequestBytes = v
}

func (o EstimateGasRequestOffledger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EstimateGasRequestOffledger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requestBytes"] = o.RequestBytes
	return toSerialize, nil
}

type NullableEstimateGasRequestOffledger struct {
	value *EstimateGasRequestOffledger
	isSet bool
}

func (v NullableEstimateGasRequestOffledger) Get() *EstimateGasRequestOffledger {
	return v.value
}

func (v *NullableEstimateGasRequestOffledger) Set(val *EstimateGasRequestOffledger) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimateGasRequestOffledger) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimateGasRequestOffledger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimateGasRequestOffledger(val *EstimateGasRequestOffledger) *NullableEstimateGasRequestOffledger {
	return &NullableEstimateGasRequestOffledger{value: val, isSet: true}
}

func (v NullableEstimateGasRequestOffledger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimateGasRequestOffledger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


