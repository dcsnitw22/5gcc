/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_udr_cli

import (
	"encoding/json"
)

// checks if the EcsAddrConfigInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EcsAddrConfigInfo{}

// EcsAddrConfigInfo struct for EcsAddrConfigInfo
type EcsAddrConfigInfo struct {
	EcsServerAddr *EcsServerAddr `json:"ecsServerAddr,omitempty"`
	SpatialValidityCond *SpatialValidityCond `json:"spatialValidityCond,omitempty"`
}

// NewEcsAddrConfigInfo instantiates a new EcsAddrConfigInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEcsAddrConfigInfo() *EcsAddrConfigInfo {
	this := EcsAddrConfigInfo{}
	return &this
}

// NewEcsAddrConfigInfoWithDefaults instantiates a new EcsAddrConfigInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEcsAddrConfigInfoWithDefaults() *EcsAddrConfigInfo {
	this := EcsAddrConfigInfo{}
	return &this
}

// GetEcsServerAddr returns the EcsServerAddr field value if set, zero value otherwise.
func (o *EcsAddrConfigInfo) GetEcsServerAddr() EcsServerAddr {
	if o == nil || IsNil(o.EcsServerAddr) {
		var ret EcsServerAddr
		return ret
	}
	return *o.EcsServerAddr
}

// GetEcsServerAddrOk returns a tuple with the EcsServerAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcsAddrConfigInfo) GetEcsServerAddrOk() (*EcsServerAddr, bool) {
	if o == nil || IsNil(o.EcsServerAddr) {
		return nil, false
	}
	return o.EcsServerAddr, true
}

// HasEcsServerAddr returns a boolean if a field has been set.
func (o *EcsAddrConfigInfo) HasEcsServerAddr() bool {
	if o != nil && !IsNil(o.EcsServerAddr) {
		return true
	}

	return false
}

// SetEcsServerAddr gets a reference to the given EcsServerAddr and assigns it to the EcsServerAddr field.
func (o *EcsAddrConfigInfo) SetEcsServerAddr(v EcsServerAddr) {
	o.EcsServerAddr = &v
}

// GetSpatialValidityCond returns the SpatialValidityCond field value if set, zero value otherwise.
func (o *EcsAddrConfigInfo) GetSpatialValidityCond() SpatialValidityCond {
	if o == nil || IsNil(o.SpatialValidityCond) {
		var ret SpatialValidityCond
		return ret
	}
	return *o.SpatialValidityCond
}

// GetSpatialValidityCondOk returns a tuple with the SpatialValidityCond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcsAddrConfigInfo) GetSpatialValidityCondOk() (*SpatialValidityCond, bool) {
	if o == nil || IsNil(o.SpatialValidityCond) {
		return nil, false
	}
	return o.SpatialValidityCond, true
}

// HasSpatialValidityCond returns a boolean if a field has been set.
func (o *EcsAddrConfigInfo) HasSpatialValidityCond() bool {
	if o != nil && !IsNil(o.SpatialValidityCond) {
		return true
	}

	return false
}

// SetSpatialValidityCond gets a reference to the given SpatialValidityCond and assigns it to the SpatialValidityCond field.
func (o *EcsAddrConfigInfo) SetSpatialValidityCond(v SpatialValidityCond) {
	o.SpatialValidityCond = &v
}

func (o EcsAddrConfigInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EcsAddrConfigInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EcsServerAddr) {
		toSerialize["ecsServerAddr"] = o.EcsServerAddr
	}
	if !IsNil(o.SpatialValidityCond) {
		toSerialize["spatialValidityCond"] = o.SpatialValidityCond
	}
	return toSerialize, nil
}

type NullableEcsAddrConfigInfo struct {
	value *EcsAddrConfigInfo
	isSet bool
}

func (v NullableEcsAddrConfigInfo) Get() *EcsAddrConfigInfo {
	return v.value
}

func (v *NullableEcsAddrConfigInfo) Set(val *EcsAddrConfigInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEcsAddrConfigInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEcsAddrConfigInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEcsAddrConfigInfo(val *EcsAddrConfigInfo) *NullableEcsAddrConfigInfo {
	return &NullableEcsAddrConfigInfo{value: val, isSet: true}
}

func (v NullableEcsAddrConfigInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEcsAddrConfigInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


