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

// checks if the SmContextRetrievedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmContextRetrievedData{}

// SmContextRetrievedData struct for SmContextRetrievedData
type SmContextRetrievedData struct {
	UeEpsPdnConnection string `json:"ueEpsPdnConnection"`
}

// NewSmContextRetrievedData instantiates a new SmContextRetrievedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmContextRetrievedData(ueEpsPdnConnection string) *SmContextRetrievedData {
	this := SmContextRetrievedData{}
	this.UeEpsPdnConnection = ueEpsPdnConnection
	return &this
}

// NewSmContextRetrievedDataWithDefaults instantiates a new SmContextRetrievedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmContextRetrievedDataWithDefaults() *SmContextRetrievedData {
	this := SmContextRetrievedData{}
	return &this
}

// GetUeEpsPdnConnection returns the UeEpsPdnConnection field value
func (o *SmContextRetrievedData) GetUeEpsPdnConnection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UeEpsPdnConnection
}

// GetUeEpsPdnConnectionOk returns a tuple with the UeEpsPdnConnection field value
// and a boolean to check if the value has been set.
func (o *SmContextRetrievedData) GetUeEpsPdnConnectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UeEpsPdnConnection, true
}

// SetUeEpsPdnConnection sets field value
func (o *SmContextRetrievedData) SetUeEpsPdnConnection(v string) {
	o.UeEpsPdnConnection = v
}

func (o SmContextRetrievedData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmContextRetrievedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ueEpsPdnConnection"] = o.UeEpsPdnConnection
	return toSerialize, nil
}

type NullableSmContextRetrievedData struct {
	value *SmContextRetrievedData
	isSet bool
}

func (v NullableSmContextRetrievedData) Get() *SmContextRetrievedData {
	return v.value
}

func (v *NullableSmContextRetrievedData) Set(val *SmContextRetrievedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmContextRetrievedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmContextRetrievedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmContextRetrievedData(val *SmContextRetrievedData) *NullableSmContextRetrievedData {
	return &NullableSmContextRetrievedData{value: val, isSet: true}
}

func (v NullableSmContextRetrievedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmContextRetrievedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


