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

// checks if the Tai type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tai{}

// Tai struct for Tai
type Tai struct {
	PlmnId PlmnId `json:"plmnId"`
	Tac string `json:"tac"`
}

// NewTai instantiates a new Tai object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTai(plmnId PlmnId, tac string) *Tai {
	this := Tai{}
	this.PlmnId = plmnId
	this.Tac = tac
	return &this
}

// NewTaiWithDefaults instantiates a new Tai object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaiWithDefaults() *Tai {
	this := Tai{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *Tai) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *Tai) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *Tai) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetTac returns the Tac field value
func (o *Tai) GetTac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tac
}

// GetTacOk returns a tuple with the Tac field value
// and a boolean to check if the value has been set.
func (o *Tai) GetTacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tac, true
}

// SetTac sets field value
func (o *Tai) SetTac(v string) {
	o.Tac = v
}

func (o Tai) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tai) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plmnId"] = o.PlmnId
	toSerialize["tac"] = o.Tac
	return toSerialize, nil
}

type NullableTai struct {
	value *Tai
	isSet bool
}

func (v NullableTai) Get() *Tai {
	return v.value
}

func (v *NullableTai) Set(val *Tai) {
	v.value = val
	v.isSet = true
}

func (v NullableTai) IsSet() bool {
	return v.isSet
}

func (v *NullableTai) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTai(val *Tai) *NullableTai {
	return &NullableTai{value: val, isSet: true}
}

func (v NullableTai) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTai) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


