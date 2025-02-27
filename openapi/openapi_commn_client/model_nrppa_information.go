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

// checks if the NrppaInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NrppaInformation{}

// NrppaInformation struct for NrppaInformation
type NrppaInformation struct {
	NfId string `json:"nfId"`
	NrppaPdu N2InfoContent `json:"nrppaPdu"`
	ServiceInstanceId *string `json:"serviceInstanceId,omitempty"`
}

// NewNrppaInformation instantiates a new NrppaInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNrppaInformation(nfId string, nrppaPdu N2InfoContent) *NrppaInformation {
	this := NrppaInformation{}
	this.NfId = nfId
	this.NrppaPdu = nrppaPdu
	return &this
}

// NewNrppaInformationWithDefaults instantiates a new NrppaInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNrppaInformationWithDefaults() *NrppaInformation {
	this := NrppaInformation{}
	return &this
}

// GetNfId returns the NfId field value
func (o *NrppaInformation) GetNfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfId
}

// GetNfIdOk returns a tuple with the NfId field value
// and a boolean to check if the value has been set.
func (o *NrppaInformation) GetNfIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NfId, true
}

// SetNfId sets field value
func (o *NrppaInformation) SetNfId(v string) {
	o.NfId = v
}

// GetNrppaPdu returns the NrppaPdu field value
func (o *NrppaInformation) GetNrppaPdu() N2InfoContent {
	if o == nil {
		var ret N2InfoContent
		return ret
	}

	return o.NrppaPdu
}

// GetNrppaPduOk returns a tuple with the NrppaPdu field value
// and a boolean to check if the value has been set.
func (o *NrppaInformation) GetNrppaPduOk() (*N2InfoContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NrppaPdu, true
}

// SetNrppaPdu sets field value
func (o *NrppaInformation) SetNrppaPdu(v N2InfoContent) {
	o.NrppaPdu = v
}

// GetServiceInstanceId returns the ServiceInstanceId field value if set, zero value otherwise.
func (o *NrppaInformation) GetServiceInstanceId() string {
	if o == nil || IsNil(o.ServiceInstanceId) {
		var ret string
		return ret
	}
	return *o.ServiceInstanceId
}

// GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrppaInformation) GetServiceInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceInstanceId) {
		return nil, false
	}
	return o.ServiceInstanceId, true
}

// HasServiceInstanceId returns a boolean if a field has been set.
func (o *NrppaInformation) HasServiceInstanceId() bool {
	if o != nil && !IsNil(o.ServiceInstanceId) {
		return true
	}

	return false
}

// SetServiceInstanceId gets a reference to the given string and assigns it to the ServiceInstanceId field.
func (o *NrppaInformation) SetServiceInstanceId(v string) {
	o.ServiceInstanceId = &v
}

func (o NrppaInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NrppaInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nfId"] = o.NfId
	toSerialize["nrppaPdu"] = o.NrppaPdu
	if !IsNil(o.ServiceInstanceId) {
		toSerialize["serviceInstanceId"] = o.ServiceInstanceId
	}
	return toSerialize, nil
}

type NullableNrppaInformation struct {
	value *NrppaInformation
	isSet bool
}

func (v NullableNrppaInformation) Get() *NrppaInformation {
	return v.value
}

func (v *NullableNrppaInformation) Set(val *NrppaInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableNrppaInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableNrppaInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrppaInformation(val *NrppaInformation) *NullableNrppaInformation {
	return &NullableNrppaInformation{value: val, isSet: true}
}

func (v NullableNrppaInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrppaInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


