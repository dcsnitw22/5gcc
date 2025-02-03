/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_commn_client

import (
	"encoding/json"
)

// checks if the UEContextRelease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UEContextRelease{}

// UEContextRelease struct for UEContextRelease
type UEContextRelease struct {
	Supi *string `json:"supi,omitempty"`
	UnauthenticatedSupi *bool `json:"unauthenticatedSupi,omitempty"`
	NgapCause NgApCause `json:"ngapCause"`
}

// NewUEContextRelease instantiates a new UEContextRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUEContextRelease(ngapCause NgApCause) *UEContextRelease {
	this := UEContextRelease{}
	var unauthenticatedSupi bool = false
	this.UnauthenticatedSupi = &unauthenticatedSupi
	this.NgapCause = ngapCause
	return &this
}

// NewUEContextReleaseWithDefaults instantiates a new UEContextRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUEContextReleaseWithDefaults() *UEContextRelease {
	this := UEContextRelease{}
	var unauthenticatedSupi bool = false
	this.UnauthenticatedSupi = &unauthenticatedSupi
	return &this
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *UEContextRelease) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UEContextRelease) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *UEContextRelease) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *UEContextRelease) SetSupi(v string) {
	o.Supi = &v
}

// GetUnauthenticatedSupi returns the UnauthenticatedSupi field value if set, zero value otherwise.
func (o *UEContextRelease) GetUnauthenticatedSupi() bool {
	if o == nil || IsNil(o.UnauthenticatedSupi) {
		var ret bool
		return ret
	}
	return *o.UnauthenticatedSupi
}

// GetUnauthenticatedSupiOk returns a tuple with the UnauthenticatedSupi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UEContextRelease) GetUnauthenticatedSupiOk() (*bool, bool) {
	if o == nil || IsNil(o.UnauthenticatedSupi) {
		return nil, false
	}
	return o.UnauthenticatedSupi, true
}

// HasUnauthenticatedSupi returns a boolean if a field has been set.
func (o *UEContextRelease) HasUnauthenticatedSupi() bool {
	if o != nil && !IsNil(o.UnauthenticatedSupi) {
		return true
	}

	return false
}

// SetUnauthenticatedSupi gets a reference to the given bool and assigns it to the UnauthenticatedSupi field.
func (o *UEContextRelease) SetUnauthenticatedSupi(v bool) {
	o.UnauthenticatedSupi = &v
}

// GetNgapCause returns the NgapCause field value
func (o *UEContextRelease) GetNgapCause() NgApCause {
	if o == nil {
		var ret NgApCause
		return ret
	}

	return o.NgapCause
}

// GetNgapCauseOk returns a tuple with the NgapCause field value
// and a boolean to check if the value has been set.
func (o *UEContextRelease) GetNgapCauseOk() (*NgApCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NgapCause, true
}

// SetNgapCause sets field value
func (o *UEContextRelease) SetNgapCause(v NgApCause) {
	o.NgapCause = v
}

func (o UEContextRelease) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UEContextRelease) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !IsNil(o.UnauthenticatedSupi) {
		toSerialize["unauthenticatedSupi"] = o.UnauthenticatedSupi
	}
	toSerialize["ngapCause"] = o.NgapCause
	return toSerialize, nil
}

type NullableUEContextRelease struct {
	value *UEContextRelease
	isSet bool
}

func (v NullableUEContextRelease) Get() *UEContextRelease {
	return v.value
}

func (v *NullableUEContextRelease) Set(val *UEContextRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableUEContextRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableUEContextRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUEContextRelease(val *UEContextRelease) *NullableUEContextRelease {
	return &NullableUEContextRelease{value: val, isSet: true}
}

func (v NullableUEContextRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUEContextRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


