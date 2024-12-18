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

// checks if the AccNetChargingAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccNetChargingAddress{}

// AccNetChargingAddress Describes the network entity within the access network performing charging
type AccNetChargingAddress struct {
	AnChargIpv4Addr *string `json:"anChargIpv4Addr,omitempty"`
	AnChargIpv6Addr *Ipv6Addr `json:"anChargIpv6Addr,omitempty"`
}

// NewAccNetChargingAddress instantiates a new AccNetChargingAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccNetChargingAddress() *AccNetChargingAddress {
	this := AccNetChargingAddress{}
	return &this
}

// NewAccNetChargingAddressWithDefaults instantiates a new AccNetChargingAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccNetChargingAddressWithDefaults() *AccNetChargingAddress {
	this := AccNetChargingAddress{}
	return &this
}

// GetAnChargIpv4Addr returns the AnChargIpv4Addr field value if set, zero value otherwise.
func (o *AccNetChargingAddress) GetAnChargIpv4Addr() string {
	if o == nil || IsNil(o.AnChargIpv4Addr) {
		var ret string
		return ret
	}
	return *o.AnChargIpv4Addr
}

// GetAnChargIpv4AddrOk returns a tuple with the AnChargIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccNetChargingAddress) GetAnChargIpv4AddrOk() (*string, bool) {
	if o == nil || IsNil(o.AnChargIpv4Addr) {
		return nil, false
	}
	return o.AnChargIpv4Addr, true
}

// HasAnChargIpv4Addr returns a boolean if a field has been set.
func (o *AccNetChargingAddress) HasAnChargIpv4Addr() bool {
	if o != nil && !IsNil(o.AnChargIpv4Addr) {
		return true
	}

	return false
}

// SetAnChargIpv4Addr gets a reference to the given string and assigns it to the AnChargIpv4Addr field.
func (o *AccNetChargingAddress) SetAnChargIpv4Addr(v string) {
	o.AnChargIpv4Addr = &v
}

// GetAnChargIpv6Addr returns the AnChargIpv6Addr field value if set, zero value otherwise.
func (o *AccNetChargingAddress) GetAnChargIpv6Addr() Ipv6Addr {
	if o == nil || IsNil(o.AnChargIpv6Addr) {
		var ret Ipv6Addr
		return ret
	}
	return *o.AnChargIpv6Addr
}

// GetAnChargIpv6AddrOk returns a tuple with the AnChargIpv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccNetChargingAddress) GetAnChargIpv6AddrOk() (*Ipv6Addr, bool) {
	if o == nil || IsNil(o.AnChargIpv6Addr) {
		return nil, false
	}
	return o.AnChargIpv6Addr, true
}

// HasAnChargIpv6Addr returns a boolean if a field has been set.
func (o *AccNetChargingAddress) HasAnChargIpv6Addr() bool {
	if o != nil && !IsNil(o.AnChargIpv6Addr) {
		return true
	}

	return false
}

// SetAnChargIpv6Addr gets a reference to the given Ipv6Addr and assigns it to the AnChargIpv6Addr field.
func (o *AccNetChargingAddress) SetAnChargIpv6Addr(v Ipv6Addr) {
	o.AnChargIpv6Addr = &v
}

func (o AccNetChargingAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccNetChargingAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnChargIpv4Addr) {
		toSerialize["anChargIpv4Addr"] = o.AnChargIpv4Addr
	}
	if !IsNil(o.AnChargIpv6Addr) {
		toSerialize["anChargIpv6Addr"] = o.AnChargIpv6Addr
	}
	return toSerialize, nil
}

type NullableAccNetChargingAddress struct {
	value *AccNetChargingAddress
	isSet bool
}

func (v NullableAccNetChargingAddress) Get() *AccNetChargingAddress {
	return v.value
}

func (v *NullableAccNetChargingAddress) Set(val *AccNetChargingAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableAccNetChargingAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAccNetChargingAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccNetChargingAddress(val *AccNetChargingAddress) *NullableAccNetChargingAddress {
	return &NullableAccNetChargingAddress{value: val, isSet: true}
}

func (v NullableAccNetChargingAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccNetChargingAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


