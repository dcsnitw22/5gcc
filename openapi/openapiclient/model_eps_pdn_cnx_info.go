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

// checks if the EpsPdnCnxInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EpsPdnCnxInfo{}

// EpsPdnCnxInfo struct for EpsPdnCnxInfo
type EpsPdnCnxInfo struct {
	PgwS8cFteid string `json:"pgwS8cFteid"`
	PgwNodeName *string `json:"pgwNodeName,omitempty"`
	LinkedBearerId *int32 `json:"linkedBearerId,omitempty"`
}

// NewEpsPdnCnxInfo instantiates a new EpsPdnCnxInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEpsPdnCnxInfo(pgwS8cFteid string) *EpsPdnCnxInfo {
	this := EpsPdnCnxInfo{}
	this.PgwS8cFteid = pgwS8cFteid
	return &this
}

// NewEpsPdnCnxInfoWithDefaults instantiates a new EpsPdnCnxInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEpsPdnCnxInfoWithDefaults() *EpsPdnCnxInfo {
	this := EpsPdnCnxInfo{}
	return &this
}

// GetPgwS8cFteid returns the PgwS8cFteid field value
func (o *EpsPdnCnxInfo) GetPgwS8cFteid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PgwS8cFteid
}

// GetPgwS8cFteidOk returns a tuple with the PgwS8cFteid field value
// and a boolean to check if the value has been set.
func (o *EpsPdnCnxInfo) GetPgwS8cFteidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PgwS8cFteid, true
}

// SetPgwS8cFteid sets field value
func (o *EpsPdnCnxInfo) SetPgwS8cFteid(v string) {
	o.PgwS8cFteid = v
}

// GetPgwNodeName returns the PgwNodeName field value if set, zero value otherwise.
func (o *EpsPdnCnxInfo) GetPgwNodeName() string {
	if o == nil || IsNil(o.PgwNodeName) {
		var ret string
		return ret
	}
	return *o.PgwNodeName
}

// GetPgwNodeNameOk returns a tuple with the PgwNodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpsPdnCnxInfo) GetPgwNodeNameOk() (*string, bool) {
	if o == nil || IsNil(o.PgwNodeName) {
		return nil, false
	}
	return o.PgwNodeName, true
}

// HasPgwNodeName returns a boolean if a field has been set.
func (o *EpsPdnCnxInfo) HasPgwNodeName() bool {
	if o != nil && !IsNil(o.PgwNodeName) {
		return true
	}

	return false
}

// SetPgwNodeName gets a reference to the given string and assigns it to the PgwNodeName field.
func (o *EpsPdnCnxInfo) SetPgwNodeName(v string) {
	o.PgwNodeName = &v
}

// GetLinkedBearerId returns the LinkedBearerId field value if set, zero value otherwise.
func (o *EpsPdnCnxInfo) GetLinkedBearerId() int32 {
	if o == nil || IsNil(o.LinkedBearerId) {
		var ret int32
		return ret
	}
	return *o.LinkedBearerId
}

// GetLinkedBearerIdOk returns a tuple with the LinkedBearerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpsPdnCnxInfo) GetLinkedBearerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LinkedBearerId) {
		return nil, false
	}
	return o.LinkedBearerId, true
}

// HasLinkedBearerId returns a boolean if a field has been set.
func (o *EpsPdnCnxInfo) HasLinkedBearerId() bool {
	if o != nil && !IsNil(o.LinkedBearerId) {
		return true
	}

	return false
}

// SetLinkedBearerId gets a reference to the given int32 and assigns it to the LinkedBearerId field.
func (o *EpsPdnCnxInfo) SetLinkedBearerId(v int32) {
	o.LinkedBearerId = &v
}

func (o EpsPdnCnxInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EpsPdnCnxInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pgwS8cFteid"] = o.PgwS8cFteid
	if !IsNil(o.PgwNodeName) {
		toSerialize["pgwNodeName"] = o.PgwNodeName
	}
	if !IsNil(o.LinkedBearerId) {
		toSerialize["linkedBearerId"] = o.LinkedBearerId
	}
	return toSerialize, nil
}

type NullableEpsPdnCnxInfo struct {
	value *EpsPdnCnxInfo
	isSet bool
}

func (v NullableEpsPdnCnxInfo) Get() *EpsPdnCnxInfo {
	return v.value
}

func (v *NullableEpsPdnCnxInfo) Set(val *EpsPdnCnxInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEpsPdnCnxInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEpsPdnCnxInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpsPdnCnxInfo(val *EpsPdnCnxInfo) *NullableEpsPdnCnxInfo {
	return &NullableEpsPdnCnxInfo{value: val, isSet: true}
}

func (v NullableEpsPdnCnxInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpsPdnCnxInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


