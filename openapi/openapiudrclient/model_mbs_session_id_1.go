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

// checks if the MbsSessionId1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsSessionId1{}

// MbsSessionId1 MBS Session Identifier
type MbsSessionId1 struct {
	Tmgi *Tmgi1 `json:"tmgi,omitempty"`
	Ssm *Ssm1 `json:"ssm,omitempty"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).  
	Nid *string `json:"nid,omitempty"`
}

// NewMbsSessionId1 instantiates a new MbsSessionId1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsSessionId1() *MbsSessionId1 {
	this := MbsSessionId1{}
	return &this
}

// NewMbsSessionId1WithDefaults instantiates a new MbsSessionId1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsSessionId1WithDefaults() *MbsSessionId1 {
	this := MbsSessionId1{}
	return &this
}

// GetTmgi returns the Tmgi field value if set, zero value otherwise.
func (o *MbsSessionId1) GetTmgi() Tmgi1 {
	if o == nil || IsNil(o.Tmgi) {
		var ret Tmgi1
		return ret
	}
	return *o.Tmgi
}

// GetTmgiOk returns a tuple with the Tmgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionId1) GetTmgiOk() (*Tmgi1, bool) {
	if o == nil || IsNil(o.Tmgi) {
		return nil, false
	}
	return o.Tmgi, true
}

// HasTmgi returns a boolean if a field has been set.
func (o *MbsSessionId1) HasTmgi() bool {
	if o != nil && !IsNil(o.Tmgi) {
		return true
	}

	return false
}

// SetTmgi gets a reference to the given Tmgi1 and assigns it to the Tmgi field.
func (o *MbsSessionId1) SetTmgi(v Tmgi1) {
	o.Tmgi = &v
}

// GetSsm returns the Ssm field value if set, zero value otherwise.
func (o *MbsSessionId1) GetSsm() Ssm1 {
	if o == nil || IsNil(o.Ssm) {
		var ret Ssm1
		return ret
	}
	return *o.Ssm
}

// GetSsmOk returns a tuple with the Ssm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionId1) GetSsmOk() (*Ssm1, bool) {
	if o == nil || IsNil(o.Ssm) {
		return nil, false
	}
	return o.Ssm, true
}

// HasSsm returns a boolean if a field has been set.
func (o *MbsSessionId1) HasSsm() bool {
	if o != nil && !IsNil(o.Ssm) {
		return true
	}

	return false
}

// SetSsm gets a reference to the given Ssm1 and assigns it to the Ssm field.
func (o *MbsSessionId1) SetSsm(v Ssm1) {
	o.Ssm = &v
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *MbsSessionId1) GetNid() string {
	if o == nil || IsNil(o.Nid) {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionId1) GetNidOk() (*string, bool) {
	if o == nil || IsNil(o.Nid) {
		return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *MbsSessionId1) HasNid() bool {
	if o != nil && !IsNil(o.Nid) {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *MbsSessionId1) SetNid(v string) {
	o.Nid = &v
}

func (o MbsSessionId1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsSessionId1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tmgi) {
		toSerialize["tmgi"] = o.Tmgi
	}
	if !IsNil(o.Ssm) {
		toSerialize["ssm"] = o.Ssm
	}
	if !IsNil(o.Nid) {
		toSerialize["nid"] = o.Nid
	}
	return toSerialize, nil
}

type NullableMbsSessionId1 struct {
	value *MbsSessionId1
	isSet bool
}

func (v NullableMbsSessionId1) Get() *MbsSessionId1 {
	return v.value
}

func (v *NullableMbsSessionId1) Set(val *MbsSessionId1) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSessionId1) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSessionId1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSessionId1(val *MbsSessionId1) *NullableMbsSessionId1 {
	return &NullableMbsSessionId1{value: val, isSet: true}
}

func (v NullableMbsSessionId1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSessionId1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


