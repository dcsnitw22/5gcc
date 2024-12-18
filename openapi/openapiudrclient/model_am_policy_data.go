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

// checks if the AmPolicyData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmPolicyData{}

// AmPolicyData Contains the AM policy data for a given subscriber.
type AmPolicyData struct {
	// Contains Presence reporting area information. The praId attribute within the PresenceInfo data type is the key of the map. 
	PraInfos *map[string]PresenceInfo `json:"praInfos,omitempty"`
	SubscCats []string `json:"subscCats,omitempty"`
}

// NewAmPolicyData instantiates a new AmPolicyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmPolicyData() *AmPolicyData {
	this := AmPolicyData{}
	return &this
}

// NewAmPolicyDataWithDefaults instantiates a new AmPolicyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmPolicyDataWithDefaults() *AmPolicyData {
	this := AmPolicyData{}
	return &this
}

// GetPraInfos returns the PraInfos field value if set, zero value otherwise.
func (o *AmPolicyData) GetPraInfos() map[string]PresenceInfo {
	if o == nil || IsNil(o.PraInfos) {
		var ret map[string]PresenceInfo
		return ret
	}
	return *o.PraInfos
}

// GetPraInfosOk returns a tuple with the PraInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmPolicyData) GetPraInfosOk() (*map[string]PresenceInfo, bool) {
	if o == nil || IsNil(o.PraInfos) {
		return nil, false
	}
	return o.PraInfos, true
}

// HasPraInfos returns a boolean if a field has been set.
func (o *AmPolicyData) HasPraInfos() bool {
	if o != nil && !IsNil(o.PraInfos) {
		return true
	}

	return false
}

// SetPraInfos gets a reference to the given map[string]PresenceInfo and assigns it to the PraInfos field.
func (o *AmPolicyData) SetPraInfos(v map[string]PresenceInfo) {
	o.PraInfos = &v
}

// GetSubscCats returns the SubscCats field value if set, zero value otherwise.
func (o *AmPolicyData) GetSubscCats() []string {
	if o == nil || IsNil(o.SubscCats) {
		var ret []string
		return ret
	}
	return o.SubscCats
}

// GetSubscCatsOk returns a tuple with the SubscCats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmPolicyData) GetSubscCatsOk() ([]string, bool) {
	if o == nil || IsNil(o.SubscCats) {
		return nil, false
	}
	return o.SubscCats, true
}

// HasSubscCats returns a boolean if a field has been set.
func (o *AmPolicyData) HasSubscCats() bool {
	if o != nil && !IsNil(o.SubscCats) {
		return true
	}

	return false
}

// SetSubscCats gets a reference to the given []string and assigns it to the SubscCats field.
func (o *AmPolicyData) SetSubscCats(v []string) {
	o.SubscCats = v
}

func (o AmPolicyData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmPolicyData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PraInfos) {
		toSerialize["praInfos"] = o.PraInfos
	}
	if !IsNil(o.SubscCats) {
		toSerialize["subscCats"] = o.SubscCats
	}
	return toSerialize, nil
}

type NullableAmPolicyData struct {
	value *AmPolicyData
	isSet bool
}

func (v NullableAmPolicyData) Get() *AmPolicyData {
	return v.value
}

func (v *NullableAmPolicyData) Set(val *AmPolicyData) {
	v.value = val
	v.isSet = true
}

func (v NullableAmPolicyData) IsSet() bool {
	return v.isSet
}

func (v *NullableAmPolicyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmPolicyData(val *AmPolicyData) *NullableAmPolicyData {
	return &NullableAmPolicyData{value: val, isSet: true}
}

func (v NullableAmPolicyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmPolicyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


