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

// checks if the WirelineServiceAreaRestriction1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WirelineServiceAreaRestriction1{}

// WirelineServiceAreaRestriction1 The \"restrictionType\" attribute and the \"areas\" attribute shall be either both present or absent.  The empty array of areas is used when service is allowed/restricted nowhere. 
type WirelineServiceAreaRestriction1 struct {
	RestrictionType *RestrictionType `json:"restrictionType,omitempty"`
	Areas []WirelineArea1 `json:"areas,omitempty"`
}

// NewWirelineServiceAreaRestriction1 instantiates a new WirelineServiceAreaRestriction1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWirelineServiceAreaRestriction1() *WirelineServiceAreaRestriction1 {
	this := WirelineServiceAreaRestriction1{}
	return &this
}

// NewWirelineServiceAreaRestriction1WithDefaults instantiates a new WirelineServiceAreaRestriction1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWirelineServiceAreaRestriction1WithDefaults() *WirelineServiceAreaRestriction1 {
	this := WirelineServiceAreaRestriction1{}
	return &this
}

// GetRestrictionType returns the RestrictionType field value if set, zero value otherwise.
func (o *WirelineServiceAreaRestriction1) GetRestrictionType() RestrictionType {
	if o == nil || IsNil(o.RestrictionType) {
		var ret RestrictionType
		return ret
	}
	return *o.RestrictionType
}

// GetRestrictionTypeOk returns a tuple with the RestrictionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelineServiceAreaRestriction1) GetRestrictionTypeOk() (*RestrictionType, bool) {
	if o == nil || IsNil(o.RestrictionType) {
		return nil, false
	}
	return o.RestrictionType, true
}

// HasRestrictionType returns a boolean if a field has been set.
func (o *WirelineServiceAreaRestriction1) HasRestrictionType() bool {
	if o != nil && !IsNil(o.RestrictionType) {
		return true
	}

	return false
}

// SetRestrictionType gets a reference to the given RestrictionType and assigns it to the RestrictionType field.
func (o *WirelineServiceAreaRestriction1) SetRestrictionType(v RestrictionType) {
	o.RestrictionType = &v
}

// GetAreas returns the Areas field value if set, zero value otherwise.
func (o *WirelineServiceAreaRestriction1) GetAreas() []WirelineArea1 {
	if o == nil || IsNil(o.Areas) {
		var ret []WirelineArea1
		return ret
	}
	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelineServiceAreaRestriction1) GetAreasOk() ([]WirelineArea1, bool) {
	if o == nil || IsNil(o.Areas) {
		return nil, false
	}
	return o.Areas, true
}

// HasAreas returns a boolean if a field has been set.
func (o *WirelineServiceAreaRestriction1) HasAreas() bool {
	if o != nil && !IsNil(o.Areas) {
		return true
	}

	return false
}

// SetAreas gets a reference to the given []WirelineArea1 and assigns it to the Areas field.
func (o *WirelineServiceAreaRestriction1) SetAreas(v []WirelineArea1) {
	o.Areas = v
}

func (o WirelineServiceAreaRestriction1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WirelineServiceAreaRestriction1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RestrictionType) {
		toSerialize["restrictionType"] = o.RestrictionType
	}
	if !IsNil(o.Areas) {
		toSerialize["areas"] = o.Areas
	}
	return toSerialize, nil
}

type NullableWirelineServiceAreaRestriction1 struct {
	value *WirelineServiceAreaRestriction1
	isSet bool
}

func (v NullableWirelineServiceAreaRestriction1) Get() *WirelineServiceAreaRestriction1 {
	return v.value
}

func (v *NullableWirelineServiceAreaRestriction1) Set(val *WirelineServiceAreaRestriction1) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelineServiceAreaRestriction1) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelineServiceAreaRestriction1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelineServiceAreaRestriction1(val *WirelineServiceAreaRestriction1) *NullableWirelineServiceAreaRestriction1 {
	return &NullableWirelineServiceAreaRestriction1{value: val, isSet: true}
}

func (v NullableWirelineServiceAreaRestriction1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelineServiceAreaRestriction1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

