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

// checks if the N3gaLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &N3gaLocation{}

// N3gaLocation struct for N3gaLocation
type N3gaLocation struct {
	N3gppTai *Tai `json:"n3gppTai,omitempty"`
	N3IwfId *string `json:"n3IwfId,omitempty"`
	UeIpv4Addr *string `json:"ueIpv4Addr,omitempty"`
	UeIpv6Addr *Ipv6Addr `json:"ueIpv6Addr,omitempty"`
	PortNumber *int32 `json:"portNumber,omitempty"`
}

// NewN3gaLocation instantiates a new N3gaLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN3gaLocation() *N3gaLocation {
	this := N3gaLocation{}
	return &this
}

// NewN3gaLocationWithDefaults instantiates a new N3gaLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN3gaLocationWithDefaults() *N3gaLocation {
	this := N3gaLocation{}
	return &this
}

// GetN3gppTai returns the N3gppTai field value if set, zero value otherwise.
func (o *N3gaLocation) GetN3gppTai() Tai {
	if o == nil || IsNil(o.N3gppTai) {
		var ret Tai
		return ret
	}
	return *o.N3gppTai
}

// GetN3gppTaiOk returns a tuple with the N3gppTai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gaLocation) GetN3gppTaiOk() (*Tai, bool) {
	if o == nil || IsNil(o.N3gppTai) {
		return nil, false
	}
	return o.N3gppTai, true
}

// HasN3gppTai returns a boolean if a field has been set.
func (o *N3gaLocation) HasN3gppTai() bool {
	if o != nil && !IsNil(o.N3gppTai) {
		return true
	}

	return false
}

// SetN3gppTai gets a reference to the given Tai and assigns it to the N3gppTai field.
func (o *N3gaLocation) SetN3gppTai(v Tai) {
	o.N3gppTai = &v
}

// GetN3IwfId returns the N3IwfId field value if set, zero value otherwise.
func (o *N3gaLocation) GetN3IwfId() string {
	if o == nil || IsNil(o.N3IwfId) {
		var ret string
		return ret
	}
	return *o.N3IwfId
}

// GetN3IwfIdOk returns a tuple with the N3IwfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gaLocation) GetN3IwfIdOk() (*string, bool) {
	if o == nil || IsNil(o.N3IwfId) {
		return nil, false
	}
	return o.N3IwfId, true
}

// HasN3IwfId returns a boolean if a field has been set.
func (o *N3gaLocation) HasN3IwfId() bool {
	if o != nil && !IsNil(o.N3IwfId) {
		return true
	}

	return false
}

// SetN3IwfId gets a reference to the given string and assigns it to the N3IwfId field.
func (o *N3gaLocation) SetN3IwfId(v string) {
	o.N3IwfId = &v
}

// GetUeIpv4Addr returns the UeIpv4Addr field value if set, zero value otherwise.
func (o *N3gaLocation) GetUeIpv4Addr() string {
	if o == nil || IsNil(o.UeIpv4Addr) {
		var ret string
		return ret
	}
	return *o.UeIpv4Addr
}

// GetUeIpv4AddrOk returns a tuple with the UeIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gaLocation) GetUeIpv4AddrOk() (*string, bool) {
	if o == nil || IsNil(o.UeIpv4Addr) {
		return nil, false
	}
	return o.UeIpv4Addr, true
}

// HasUeIpv4Addr returns a boolean if a field has been set.
func (o *N3gaLocation) HasUeIpv4Addr() bool {
	if o != nil && !IsNil(o.UeIpv4Addr) {
		return true
	}

	return false
}

// SetUeIpv4Addr gets a reference to the given string and assigns it to the UeIpv4Addr field.
func (o *N3gaLocation) SetUeIpv4Addr(v string) {
	o.UeIpv4Addr = &v
}

// GetUeIpv6Addr returns the UeIpv6Addr field value if set, zero value otherwise.
func (o *N3gaLocation) GetUeIpv6Addr() Ipv6Addr {
	if o == nil || IsNil(o.UeIpv6Addr) {
		var ret Ipv6Addr
		return ret
	}
	return *o.UeIpv6Addr
}

// GetUeIpv6AddrOk returns a tuple with the UeIpv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gaLocation) GetUeIpv6AddrOk() (*Ipv6Addr, bool) {
	if o == nil || IsNil(o.UeIpv6Addr) {
		return nil, false
	}
	return o.UeIpv6Addr, true
}

// HasUeIpv6Addr returns a boolean if a field has been set.
func (o *N3gaLocation) HasUeIpv6Addr() bool {
	if o != nil && !IsNil(o.UeIpv6Addr) {
		return true
	}

	return false
}

// SetUeIpv6Addr gets a reference to the given Ipv6Addr and assigns it to the UeIpv6Addr field.
func (o *N3gaLocation) SetUeIpv6Addr(v Ipv6Addr) {
	o.UeIpv6Addr = &v
}

// GetPortNumber returns the PortNumber field value if set, zero value otherwise.
func (o *N3gaLocation) GetPortNumber() int32 {
	if o == nil || IsNil(o.PortNumber) {
		var ret int32
		return ret
	}
	return *o.PortNumber
}

// GetPortNumberOk returns a tuple with the PortNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gaLocation) GetPortNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PortNumber) {
		return nil, false
	}
	return o.PortNumber, true
}

// HasPortNumber returns a boolean if a field has been set.
func (o *N3gaLocation) HasPortNumber() bool {
	if o != nil && !IsNil(o.PortNumber) {
		return true
	}

	return false
}

// SetPortNumber gets a reference to the given int32 and assigns it to the PortNumber field.
func (o *N3gaLocation) SetPortNumber(v int32) {
	o.PortNumber = &v
}

func (o N3gaLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o N3gaLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.N3gppTai) {
		toSerialize["n3gppTai"] = o.N3gppTai
	}
	if !IsNil(o.N3IwfId) {
		toSerialize["n3IwfId"] = o.N3IwfId
	}
	if !IsNil(o.UeIpv4Addr) {
		toSerialize["ueIpv4Addr"] = o.UeIpv4Addr
	}
	if !IsNil(o.UeIpv6Addr) {
		toSerialize["ueIpv6Addr"] = o.UeIpv6Addr
	}
	if !IsNil(o.PortNumber) {
		toSerialize["portNumber"] = o.PortNumber
	}
	return toSerialize, nil
}

type NullableN3gaLocation struct {
	value *N3gaLocation
	isSet bool
}

func (v NullableN3gaLocation) Get() *N3gaLocation {
	return v.value
}

func (v *NullableN3gaLocation) Set(val *N3gaLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableN3gaLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableN3gaLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN3gaLocation(val *N3gaLocation) *NullableN3gaLocation {
	return &NullableN3gaLocation{value: val, isSet: true}
}

func (v NullableN3gaLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN3gaLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


