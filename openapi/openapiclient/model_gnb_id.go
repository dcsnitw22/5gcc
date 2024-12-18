/*
Nsmf_PDUSession

SMF PDU Session Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
)

// checks if the GNbId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GNbId{}

// GNbId struct for GNbId
type GNbId struct {
	BitLength int32 `json:"bitLength"`
	GNBValue string `json:"gNBValue"`
}

// NewGNbId instantiates a new GNbId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGNbId(bitLength int32, gNBValue string) *GNbId {
	this := GNbId{}
	this.BitLength = bitLength
	this.GNBValue = gNBValue
	return &this
}

// NewGNbIdWithDefaults instantiates a new GNbId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGNbIdWithDefaults() *GNbId {
	this := GNbId{}
	return &this
}

// GetBitLength returns the BitLength field value
func (o *GNbId) GetBitLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BitLength
}

// GetBitLengthOk returns a tuple with the BitLength field value
// and a boolean to check if the value has been set.
func (o *GNbId) GetBitLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BitLength, true
}

// SetBitLength sets field value
func (o *GNbId) SetBitLength(v int32) {
	o.BitLength = v
}

// GetGNBValue returns the GNBValue field value
func (o *GNbId) GetGNBValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GNBValue
}

// GetGNBValueOk returns a tuple with the GNBValue field value
// and a boolean to check if the value has been set.
func (o *GNbId) GetGNBValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GNBValue, true
}

// SetGNBValue sets field value
func (o *GNbId) SetGNBValue(v string) {
	o.GNBValue = v
}

func (o GNbId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GNbId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bitLength"] = o.BitLength
	toSerialize["gNBValue"] = o.GNBValue
	return toSerialize, nil
}

type NullableGNbId struct {
	value *GNbId
	isSet bool
}

func (v NullableGNbId) Get() *GNbId {
	return v.value
}

func (v *NullableGNbId) Set(val *GNbId) {
	v.value = val
	v.isSet = true
}

func (v NullableGNbId) IsSet() bool {
	return v.isSet
}

func (v *NullableGNbId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGNbId(val *GNbId) *NullableGNbId {
	return &NullableGNbId{value: val, isSet: true}
}

func (v NullableGNbId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGNbId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


