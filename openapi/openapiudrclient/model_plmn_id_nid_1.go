/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_udr_cli

import (
	"encoding/json"
	"fmt"
)

// checks if the PlmnIdNid1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlmnIdNid1{}

// PlmnIdNid1 Contains the serving core network operator PLMN ID and, for an SNPN, the NID that together with the PLMN ID identifies the SNPN. 
type PlmnIdNid1 struct {
	// Mobile Country Code part of the PLMN, comprising 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.  
	Mcc string `json:"mcc"`
	// Mobile Network Code part of the PLMN, comprising 2 or 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.
	Mnc string `json:"mnc"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).  
	Nid *string `json:"nid,omitempty"`
}

type _PlmnIdNid1 PlmnIdNid1

// NewPlmnIdNid1 instantiates a new PlmnIdNid1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnIdNid1(mcc string, mnc string) *PlmnIdNid1 {
	this := PlmnIdNid1{}
	this.Mcc = mcc
	this.Mnc = mnc
	return &this
}

// NewPlmnIdNid1WithDefaults instantiates a new PlmnIdNid1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnIdNid1WithDefaults() *PlmnIdNid1 {
	this := PlmnIdNid1{}
	return &this
}

// GetMcc returns the Mcc field value
func (o *PlmnIdNid1) GetMcc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value
// and a boolean to check if the value has been set.
func (o *PlmnIdNid1) GetMccOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mcc, true
}

// SetMcc sets field value
func (o *PlmnIdNid1) SetMcc(v string) {
	o.Mcc = v
}

// GetMnc returns the Mnc field value
func (o *PlmnIdNid1) GetMnc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value
// and a boolean to check if the value has been set.
func (o *PlmnIdNid1) GetMncOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mnc, true
}

// SetMnc sets field value
func (o *PlmnIdNid1) SetMnc(v string) {
	o.Mnc = v
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *PlmnIdNid1) GetNid() string {
	if o == nil || IsNil(o.Nid) {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnIdNid1) GetNidOk() (*string, bool) {
	if o == nil || IsNil(o.Nid) {
		return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *PlmnIdNid1) HasNid() bool {
	if o != nil && !IsNil(o.Nid) {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *PlmnIdNid1) SetNid(v string) {
	o.Nid = &v
}

func (o PlmnIdNid1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlmnIdNid1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mcc"] = o.Mcc
	toSerialize["mnc"] = o.Mnc
	if !IsNil(o.Nid) {
		toSerialize["nid"] = o.Nid
	}
	return toSerialize, nil
}

func (o *PlmnIdNid1) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mcc",
		"mnc",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPlmnIdNid1 := _PlmnIdNid1{}

	err = json.Unmarshal(bytes, &varPlmnIdNid1)

	if err != nil {
		return err
	}

	*o = PlmnIdNid1(varPlmnIdNid1)

	return err
}

type NullablePlmnIdNid1 struct {
	value *PlmnIdNid1
	isSet bool
}

func (v NullablePlmnIdNid1) Get() *PlmnIdNid1 {
	return v.value
}

func (v *NullablePlmnIdNid1) Set(val *PlmnIdNid1) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnIdNid1) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnIdNid1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnIdNid1(val *PlmnIdNid1) *NullablePlmnIdNid1 {
	return &NullablePlmnIdNid1{value: val, isSet: true}
}

func (v NullablePlmnIdNid1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnIdNid1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


