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

// checks if the IpAddr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpAddr{}

// IpAddr Contains an IP adresse.
type IpAddr struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Ipv4Addr *string `json:"ipv4Addr,omitempty"`
	Ipv6Addr *Ipv6Addr `json:"ipv6Addr,omitempty"`
	Ipv6Prefix *Ipv6Prefix `json:"ipv6Prefix,omitempty"`
}

// NewIpAddr instantiates a new IpAddr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpAddr() *IpAddr {
	this := IpAddr{}
	return &this
}

// NewIpAddrWithDefaults instantiates a new IpAddr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpAddrWithDefaults() *IpAddr {
	this := IpAddr{}
	return &this
}

// GetIpv4Addr returns the Ipv4Addr field value if set, zero value otherwise.
func (o *IpAddr) GetIpv4Addr() string {
	if o == nil || IsNil(o.Ipv4Addr) {
		var ret string
		return ret
	}
	return *o.Ipv4Addr
}

// GetIpv4AddrOk returns a tuple with the Ipv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpAddr) GetIpv4AddrOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv4Addr) {
		return nil, false
	}
	return o.Ipv4Addr, true
}

// HasIpv4Addr returns a boolean if a field has been set.
func (o *IpAddr) HasIpv4Addr() bool {
	if o != nil && !IsNil(o.Ipv4Addr) {
		return true
	}

	return false
}

// SetIpv4Addr gets a reference to the given string and assigns it to the Ipv4Addr field.
func (o *IpAddr) SetIpv4Addr(v string) {
	o.Ipv4Addr = &v
}

// GetIpv6Addr returns the Ipv6Addr field value if set, zero value otherwise.
func (o *IpAddr) GetIpv6Addr() Ipv6Addr {
	if o == nil || IsNil(o.Ipv6Addr) {
		var ret Ipv6Addr
		return ret
	}
	return *o.Ipv6Addr
}

// GetIpv6AddrOk returns a tuple with the Ipv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpAddr) GetIpv6AddrOk() (*Ipv6Addr, bool) {
	if o == nil || IsNil(o.Ipv6Addr) {
		return nil, false
	}
	return o.Ipv6Addr, true
}

// HasIpv6Addr returns a boolean if a field has been set.
func (o *IpAddr) HasIpv6Addr() bool {
	if o != nil && !IsNil(o.Ipv6Addr) {
		return true
	}

	return false
}

// SetIpv6Addr gets a reference to the given Ipv6Addr and assigns it to the Ipv6Addr field.
func (o *IpAddr) SetIpv6Addr(v Ipv6Addr) {
	o.Ipv6Addr = &v
}

// GetIpv6Prefix returns the Ipv6Prefix field value if set, zero value otherwise.
func (o *IpAddr) GetIpv6Prefix() Ipv6Prefix {
	if o == nil || IsNil(o.Ipv6Prefix) {
		var ret Ipv6Prefix
		return ret
	}
	return *o.Ipv6Prefix
}

// GetIpv6PrefixOk returns a tuple with the Ipv6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpAddr) GetIpv6PrefixOk() (*Ipv6Prefix, bool) {
	if o == nil || IsNil(o.Ipv6Prefix) {
		return nil, false
	}
	return o.Ipv6Prefix, true
}

// HasIpv6Prefix returns a boolean if a field has been set.
func (o *IpAddr) HasIpv6Prefix() bool {
	if o != nil && !IsNil(o.Ipv6Prefix) {
		return true
	}

	return false
}

// SetIpv6Prefix gets a reference to the given Ipv6Prefix and assigns it to the Ipv6Prefix field.
func (o *IpAddr) SetIpv6Prefix(v Ipv6Prefix) {
	o.Ipv6Prefix = &v
}

func (o IpAddr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpAddr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ipv4Addr) {
		toSerialize["ipv4Addr"] = o.Ipv4Addr
	}
	if !IsNil(o.Ipv6Addr) {
		toSerialize["ipv6Addr"] = o.Ipv6Addr
	}
	if !IsNil(o.Ipv6Prefix) {
		toSerialize["ipv6Prefix"] = o.Ipv6Prefix
	}
	return toSerialize, nil
}

type NullableIpAddr struct {
	value *IpAddr
	isSet bool
}

func (v NullableIpAddr) Get() *IpAddr {
	return v.value
}

func (v *NullableIpAddr) Set(val *IpAddr) {
	v.value = val
	v.isSet = true
}

func (v NullableIpAddr) IsSet() bool {
	return v.isSet
}

func (v *NullableIpAddr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpAddr(val *IpAddr) *NullableIpAddr {
	return &NullableIpAddr{value: val, isSet: true}
}

func (v NullableIpAddr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpAddr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

