/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AnGwAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnGwAddress{}

// AnGwAddress describes the address of the access network gateway control node
type AnGwAddress struct {
	AnGwIpv4Addr *string `json:"anGwIpv4Addr,omitempty"`
	AnGwIpv6Addr *Ipv6Addr `json:"anGwIpv6Addr,omitempty"`
}

// NewAnGwAddress instantiates a new AnGwAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnGwAddress() *AnGwAddress {
	this := AnGwAddress{}
	return &this
}

// NewAnGwAddressWithDefaults instantiates a new AnGwAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnGwAddressWithDefaults() *AnGwAddress {
	this := AnGwAddress{}
	return &this
}

// GetAnGwIpv4Addr returns the AnGwIpv4Addr field value if set, zero value otherwise.
func (o *AnGwAddress) GetAnGwIpv4Addr() string {
	if o == nil || IsNil(o.AnGwIpv4Addr) {
		var ret string
		return ret
	}
	return *o.AnGwIpv4Addr
}

// GetAnGwIpv4AddrOk returns a tuple with the AnGwIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnGwAddress) GetAnGwIpv4AddrOk() (*string, bool) {
	if o == nil || IsNil(o.AnGwIpv4Addr) {
		return nil, false
	}
	return o.AnGwIpv4Addr, true
}

// HasAnGwIpv4Addr returns a boolean if a field has been set.
func (o *AnGwAddress) HasAnGwIpv4Addr() bool {
	if o != nil && !IsNil(o.AnGwIpv4Addr) {
		return true
	}

	return false
}

// SetAnGwIpv4Addr gets a reference to the given string and assigns it to the AnGwIpv4Addr field.
func (o *AnGwAddress) SetAnGwIpv4Addr(v string) {
	o.AnGwIpv4Addr = &v
}

// GetAnGwIpv6Addr returns the AnGwIpv6Addr field value if set, zero value otherwise.
func (o *AnGwAddress) GetAnGwIpv6Addr() Ipv6Addr {
	if o == nil || IsNil(o.AnGwIpv6Addr) {
		var ret Ipv6Addr
		return ret
	}
	return *o.AnGwIpv6Addr
}

// GetAnGwIpv6AddrOk returns a tuple with the AnGwIpv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnGwAddress) GetAnGwIpv6AddrOk() (*Ipv6Addr, bool) {
	if o == nil || IsNil(o.AnGwIpv6Addr) {
		return nil, false
	}
	return o.AnGwIpv6Addr, true
}

// HasAnGwIpv6Addr returns a boolean if a field has been set.
func (o *AnGwAddress) HasAnGwIpv6Addr() bool {
	if o != nil && !IsNil(o.AnGwIpv6Addr) {
		return true
	}

	return false
}

// SetAnGwIpv6Addr gets a reference to the given Ipv6Addr and assigns it to the AnGwIpv6Addr field.
func (o *AnGwAddress) SetAnGwIpv6Addr(v Ipv6Addr) {
	o.AnGwIpv6Addr = &v
}

func (o AnGwAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnGwAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnGwIpv4Addr) {
		toSerialize["anGwIpv4Addr"] = o.AnGwIpv4Addr
	}
	if !IsNil(o.AnGwIpv6Addr) {
		toSerialize["anGwIpv6Addr"] = o.AnGwIpv6Addr
	}
	return toSerialize, nil
}

type NullableAnGwAddress struct {
	value *AnGwAddress
	isSet bool
}

func (v NullableAnGwAddress) Get() *AnGwAddress {
	return v.value
}

func (v *NullableAnGwAddress) Set(val *AnGwAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableAnGwAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAnGwAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnGwAddress(val *AnGwAddress) *NullableAnGwAddress {
	return &NullableAnGwAddress{value: val, isSet: true}
}

func (v NullableAnGwAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnGwAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

