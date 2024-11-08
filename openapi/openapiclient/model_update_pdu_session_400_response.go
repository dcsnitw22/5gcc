/*
Nsmf_PDUSession

SMF PDU Session Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
	"os"
)

// checks if the UpdatePduSession400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePduSession400Response{}

// UpdatePduSession400Response struct for UpdatePduSession400Response
type UpdatePduSession400Response struct {
	JsonData *HsmfUpdateError `json:"jsonData,omitempty"`
	BinaryDataN1SmInfoToUe **os.File `json:"binaryDataN1SmInfoToUe,omitempty"`
}

// NewUpdatePduSession400Response instantiates a new UpdatePduSession400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePduSession400Response() *UpdatePduSession400Response {
	this := UpdatePduSession400Response{}
	return &this
}

// NewUpdatePduSession400ResponseWithDefaults instantiates a new UpdatePduSession400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePduSession400ResponseWithDefaults() *UpdatePduSession400Response {
	this := UpdatePduSession400Response{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *UpdatePduSession400Response) GetJsonData() HsmfUpdateError {
	if o == nil || IsNil(o.JsonData) {
		var ret HsmfUpdateError
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePduSession400Response) GetJsonDataOk() (*HsmfUpdateError, bool) {
	if o == nil || IsNil(o.JsonData) {
		return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *UpdatePduSession400Response) HasJsonData() bool {
	if o != nil && !IsNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given HsmfUpdateError and assigns it to the JsonData field.
func (o *UpdatePduSession400Response) SetJsonData(v HsmfUpdateError) {
	o.JsonData = &v
}

// GetBinaryDataN1SmInfoToUe returns the BinaryDataN1SmInfoToUe field value if set, zero value otherwise.
func (o *UpdatePduSession400Response) GetBinaryDataN1SmInfoToUe() *os.File {
	if o == nil || IsNil(o.BinaryDataN1SmInfoToUe) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN1SmInfoToUe
}

// GetBinaryDataN1SmInfoToUeOk returns a tuple with the BinaryDataN1SmInfoToUe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePduSession400Response) GetBinaryDataN1SmInfoToUeOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN1SmInfoToUe) {
		return nil, false
	}
	return o.BinaryDataN1SmInfoToUe, true
}

// HasBinaryDataN1SmInfoToUe returns a boolean if a field has been set.
func (o *UpdatePduSession400Response) HasBinaryDataN1SmInfoToUe() bool {
	if o != nil && !IsNil(o.BinaryDataN1SmInfoToUe) {
		return true
	}

	return false
}

// SetBinaryDataN1SmInfoToUe gets a reference to the given *os.File and assigns it to the BinaryDataN1SmInfoToUe field.
func (o *UpdatePduSession400Response) SetBinaryDataN1SmInfoToUe(v *os.File) {
	o.BinaryDataN1SmInfoToUe = &v
}

func (o UpdatePduSession400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePduSession400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !IsNil(o.BinaryDataN1SmInfoToUe) {
		toSerialize["binaryDataN1SmInfoToUe"] = o.BinaryDataN1SmInfoToUe
	}
	return toSerialize, nil
}

type NullableUpdatePduSession400Response struct {
	value *UpdatePduSession400Response
	isSet bool
}

func (v NullableUpdatePduSession400Response) Get() *UpdatePduSession400Response {
	return v.value
}

func (v *NullableUpdatePduSession400Response) Set(val *UpdatePduSession400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePduSession400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePduSession400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePduSession400Response(val *UpdatePduSession400Response) *NullableUpdatePduSession400Response {
	return &NullableUpdatePduSession400Response{value: val, isSet: true}
}

func (v NullableUpdatePduSession400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePduSession400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


