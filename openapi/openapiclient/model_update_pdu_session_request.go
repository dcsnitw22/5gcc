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

// checks if the UpdatePduSessionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePduSessionRequest{}

// UpdatePduSessionRequest struct for UpdatePduSessionRequest
type UpdatePduSessionRequest struct {
	JsonData *HsmfUpdateData `json:"jsonData,omitempty"`
	BinaryDataN1SmInfoFromUe **os.File `json:"binaryDataN1SmInfoFromUe,omitempty"`
	BinaryDataUnknownN1SmInfo **os.File `json:"binaryDataUnknownN1SmInfo,omitempty"`
}

// NewUpdatePduSessionRequest instantiates a new UpdatePduSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePduSessionRequest() *UpdatePduSessionRequest {
	this := UpdatePduSessionRequest{}
	return &this
}

// NewUpdatePduSessionRequestWithDefaults instantiates a new UpdatePduSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePduSessionRequestWithDefaults() *UpdatePduSessionRequest {
	this := UpdatePduSessionRequest{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *UpdatePduSessionRequest) GetJsonData() HsmfUpdateData {
	if o == nil || IsNil(o.JsonData) {
		var ret HsmfUpdateData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePduSessionRequest) GetJsonDataOk() (*HsmfUpdateData, bool) {
	if o == nil || IsNil(o.JsonData) {
		return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *UpdatePduSessionRequest) HasJsonData() bool {
	if o != nil && !IsNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given HsmfUpdateData and assigns it to the JsonData field.
func (o *UpdatePduSessionRequest) SetJsonData(v HsmfUpdateData) {
	o.JsonData = &v
}

// GetBinaryDataN1SmInfoFromUe returns the BinaryDataN1SmInfoFromUe field value if set, zero value otherwise.
func (o *UpdatePduSessionRequest) GetBinaryDataN1SmInfoFromUe() *os.File {
	if o == nil || IsNil(o.BinaryDataN1SmInfoFromUe) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN1SmInfoFromUe
}

// GetBinaryDataN1SmInfoFromUeOk returns a tuple with the BinaryDataN1SmInfoFromUe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePduSessionRequest) GetBinaryDataN1SmInfoFromUeOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN1SmInfoFromUe) {
		return nil, false
	}
	return o.BinaryDataN1SmInfoFromUe, true
}

// HasBinaryDataN1SmInfoFromUe returns a boolean if a field has been set.
func (o *UpdatePduSessionRequest) HasBinaryDataN1SmInfoFromUe() bool {
	if o != nil && !IsNil(o.BinaryDataN1SmInfoFromUe) {
		return true
	}

	return false
}

// SetBinaryDataN1SmInfoFromUe gets a reference to the given *os.File and assigns it to the BinaryDataN1SmInfoFromUe field.
func (o *UpdatePduSessionRequest) SetBinaryDataN1SmInfoFromUe(v *os.File) {
	o.BinaryDataN1SmInfoFromUe = &v
}

// GetBinaryDataUnknownN1SmInfo returns the BinaryDataUnknownN1SmInfo field value if set, zero value otherwise.
func (o *UpdatePduSessionRequest) GetBinaryDataUnknownN1SmInfo() *os.File {
	if o == nil || IsNil(o.BinaryDataUnknownN1SmInfo) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataUnknownN1SmInfo
}

// GetBinaryDataUnknownN1SmInfoOk returns a tuple with the BinaryDataUnknownN1SmInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePduSessionRequest) GetBinaryDataUnknownN1SmInfoOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataUnknownN1SmInfo) {
		return nil, false
	}
	return o.BinaryDataUnknownN1SmInfo, true
}

// HasBinaryDataUnknownN1SmInfo returns a boolean if a field has been set.
func (o *UpdatePduSessionRequest) HasBinaryDataUnknownN1SmInfo() bool {
	if o != nil && !IsNil(o.BinaryDataUnknownN1SmInfo) {
		return true
	}

	return false
}

// SetBinaryDataUnknownN1SmInfo gets a reference to the given *os.File and assigns it to the BinaryDataUnknownN1SmInfo field.
func (o *UpdatePduSessionRequest) SetBinaryDataUnknownN1SmInfo(v *os.File) {
	o.BinaryDataUnknownN1SmInfo = &v
}

func (o UpdatePduSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePduSessionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !IsNil(o.BinaryDataN1SmInfoFromUe) {
		toSerialize["binaryDataN1SmInfoFromUe"] = o.BinaryDataN1SmInfoFromUe
	}
	if !IsNil(o.BinaryDataUnknownN1SmInfo) {
		toSerialize["binaryDataUnknownN1SmInfo"] = o.BinaryDataUnknownN1SmInfo
	}
	return toSerialize, nil
}

type NullableUpdatePduSessionRequest struct {
	value *UpdatePduSessionRequest
	isSet bool
}

func (v NullableUpdatePduSessionRequest) Get() *UpdatePduSessionRequest {
	return v.value
}

func (v *NullableUpdatePduSessionRequest) Set(val *UpdatePduSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePduSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePduSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePduSessionRequest(val *UpdatePduSessionRequest) *NullableUpdatePduSessionRequest {
	return &NullableUpdatePduSessionRequest{value: val, isSet: true}
}

func (v NullableUpdatePduSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePduSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


