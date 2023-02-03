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

// checks if the AccountNFTsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountNFTsResponse{}

// AccountNFTsResponse struct for AccountNFTsResponse
type AccountNFTsResponse struct {
	NftIds []string `json:"nftIds"`
}

// NewAccountNFTsResponse instantiates a new AccountNFTsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountNFTsResponse(nftIds []string) *AccountNFTsResponse {
	this := AccountNFTsResponse{}
	this.NftIds = nftIds
	return &this
}

// NewAccountNFTsResponseWithDefaults instantiates a new AccountNFTsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountNFTsResponseWithDefaults() *AccountNFTsResponse {
	this := AccountNFTsResponse{}
	return &this
}

// GetNftIds returns the NftIds field value
func (o *AccountNFTsResponse) GetNftIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NftIds
}

// GetNftIdsOk returns a tuple with the NftIds field value
// and a boolean to check if the value has been set.
func (o *AccountNFTsResponse) GetNftIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NftIds, true
}

// SetNftIds sets field value
func (o *AccountNFTsResponse) SetNftIds(v []string) {
	o.NftIds = v
}

func (o AccountNFTsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountNFTsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nftIds"] = o.NftIds
	return toSerialize, nil
}

type NullableAccountNFTsResponse struct {
	value *AccountNFTsResponse
	isSet bool
}

func (v NullableAccountNFTsResponse) Get() *AccountNFTsResponse {
	return v.value
}

func (v *NullableAccountNFTsResponse) Set(val *AccountNFTsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountNFTsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountNFTsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountNFTsResponse(val *AccountNFTsResponse) *NullableAccountNFTsResponse {
	return &NullableAccountNFTsResponse{value: val, isSet: true}
}

func (v NullableAccountNFTsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountNFTsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


