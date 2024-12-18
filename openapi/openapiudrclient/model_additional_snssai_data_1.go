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

// checks if the AdditionalSnssaiData1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalSnssaiData1{}

// AdditionalSnssaiData1 struct for AdditionalSnssaiData1
type AdditionalSnssaiData1 struct {
	RequiredAuthnAuthz *bool `json:"requiredAuthnAuthz,omitempty"`
	SubscribedUeSliceMbr NullableSliceMbr `json:"subscribedUeSliceMbr,omitempty"`
	SubscribedNsSrgList []string `json:"subscribedNsSrgList,omitempty"`
}

// NewAdditionalSnssaiData1 instantiates a new AdditionalSnssaiData1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalSnssaiData1() *AdditionalSnssaiData1 {
	this := AdditionalSnssaiData1{}
	return &this
}

// NewAdditionalSnssaiData1WithDefaults instantiates a new AdditionalSnssaiData1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalSnssaiData1WithDefaults() *AdditionalSnssaiData1 {
	this := AdditionalSnssaiData1{}
	return &this
}

// GetRequiredAuthnAuthz returns the RequiredAuthnAuthz field value if set, zero value otherwise.
func (o *AdditionalSnssaiData1) GetRequiredAuthnAuthz() bool {
	if o == nil || IsNil(o.RequiredAuthnAuthz) {
		var ret bool
		return ret
	}
	return *o.RequiredAuthnAuthz
}

// GetRequiredAuthnAuthzOk returns a tuple with the RequiredAuthnAuthz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalSnssaiData1) GetRequiredAuthnAuthzOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiredAuthnAuthz) {
		return nil, false
	}
	return o.RequiredAuthnAuthz, true
}

// HasRequiredAuthnAuthz returns a boolean if a field has been set.
func (o *AdditionalSnssaiData1) HasRequiredAuthnAuthz() bool {
	if o != nil && !IsNil(o.RequiredAuthnAuthz) {
		return true
	}

	return false
}

// SetRequiredAuthnAuthz gets a reference to the given bool and assigns it to the RequiredAuthnAuthz field.
func (o *AdditionalSnssaiData1) SetRequiredAuthnAuthz(v bool) {
	o.RequiredAuthnAuthz = &v
}

// GetSubscribedUeSliceMbr returns the SubscribedUeSliceMbr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdditionalSnssaiData1) GetSubscribedUeSliceMbr() SliceMbr {
	if o == nil || IsNil(o.SubscribedUeSliceMbr.Get()) {
		var ret SliceMbr
		return ret
	}
	return *o.SubscribedUeSliceMbr.Get()
}

// GetSubscribedUeSliceMbrOk returns a tuple with the SubscribedUeSliceMbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdditionalSnssaiData1) GetSubscribedUeSliceMbrOk() (*SliceMbr, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscribedUeSliceMbr.Get(), o.SubscribedUeSliceMbr.IsSet()
}

// HasSubscribedUeSliceMbr returns a boolean if a field has been set.
func (o *AdditionalSnssaiData1) HasSubscribedUeSliceMbr() bool {
	if o != nil && o.SubscribedUeSliceMbr.IsSet() {
		return true
	}

	return false
}

// SetSubscribedUeSliceMbr gets a reference to the given NullableSliceMbr and assigns it to the SubscribedUeSliceMbr field.
func (o *AdditionalSnssaiData1) SetSubscribedUeSliceMbr(v SliceMbr) {
	o.SubscribedUeSliceMbr.Set(&v)
}
// SetSubscribedUeSliceMbrNil sets the value for SubscribedUeSliceMbr to be an explicit nil
func (o *AdditionalSnssaiData1) SetSubscribedUeSliceMbrNil() {
	o.SubscribedUeSliceMbr.Set(nil)
}

// UnsetSubscribedUeSliceMbr ensures that no value is present for SubscribedUeSliceMbr, not even an explicit nil
func (o *AdditionalSnssaiData1) UnsetSubscribedUeSliceMbr() {
	o.SubscribedUeSliceMbr.Unset()
}

// GetSubscribedNsSrgList returns the SubscribedNsSrgList field value if set, zero value otherwise.
func (o *AdditionalSnssaiData1) GetSubscribedNsSrgList() []string {
	if o == nil || IsNil(o.SubscribedNsSrgList) {
		var ret []string
		return ret
	}
	return o.SubscribedNsSrgList
}

// GetSubscribedNsSrgListOk returns a tuple with the SubscribedNsSrgList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalSnssaiData1) GetSubscribedNsSrgListOk() ([]string, bool) {
	if o == nil || IsNil(o.SubscribedNsSrgList) {
		return nil, false
	}
	return o.SubscribedNsSrgList, true
}

// HasSubscribedNsSrgList returns a boolean if a field has been set.
func (o *AdditionalSnssaiData1) HasSubscribedNsSrgList() bool {
	if o != nil && !IsNil(o.SubscribedNsSrgList) {
		return true
	}

	return false
}

// SetSubscribedNsSrgList gets a reference to the given []string and assigns it to the SubscribedNsSrgList field.
func (o *AdditionalSnssaiData1) SetSubscribedNsSrgList(v []string) {
	o.SubscribedNsSrgList = v
}

func (o AdditionalSnssaiData1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalSnssaiData1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequiredAuthnAuthz) {
		toSerialize["requiredAuthnAuthz"] = o.RequiredAuthnAuthz
	}
	if o.SubscribedUeSliceMbr.IsSet() {
		toSerialize["subscribedUeSliceMbr"] = o.SubscribedUeSliceMbr.Get()
	}
	if !IsNil(o.SubscribedNsSrgList) {
		toSerialize["subscribedNsSrgList"] = o.SubscribedNsSrgList
	}
	return toSerialize, nil
}

type NullableAdditionalSnssaiData1 struct {
	value *AdditionalSnssaiData1
	isSet bool
}

func (v NullableAdditionalSnssaiData1) Get() *AdditionalSnssaiData1 {
	return v.value
}

func (v *NullableAdditionalSnssaiData1) Set(val *AdditionalSnssaiData1) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalSnssaiData1) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalSnssaiData1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalSnssaiData1(val *AdditionalSnssaiData1) *NullableAdditionalSnssaiData1 {
	return &NullableAdditionalSnssaiData1{value: val, isSet: true}
}

func (v NullableAdditionalSnssaiData1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalSnssaiData1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


