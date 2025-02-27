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

// checks if the AssignedEbiData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignedEbiData{}

// AssignedEbiData struct for AssignedEbiData
type AssignedEbiData struct {
	PduSessionId int32 `json:"pduSessionId"`
	AssignedEbiList []EbiArpMapping `json:"assignedEbiList"`
	FailedArpList []Arp `json:"failedArpList,omitempty"`
	ReleasedEbiList []int32 `json:"releasedEbiList,omitempty"`
}

// NewAssignedEbiData instantiates a new AssignedEbiData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignedEbiData(pduSessionId int32, assignedEbiList []EbiArpMapping) *AssignedEbiData {
	this := AssignedEbiData{}
	this.PduSessionId = pduSessionId
	this.AssignedEbiList = assignedEbiList
	return &this
}

// NewAssignedEbiDataWithDefaults instantiates a new AssignedEbiData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignedEbiDataWithDefaults() *AssignedEbiData {
	this := AssignedEbiData{}
	return &this
}

// GetPduSessionId returns the PduSessionId field value
func (o *AssignedEbiData) GetPduSessionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PduSessionId
}

// GetPduSessionIdOk returns a tuple with the PduSessionId field value
// and a boolean to check if the value has been set.
func (o *AssignedEbiData) GetPduSessionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PduSessionId, true
}

// SetPduSessionId sets field value
func (o *AssignedEbiData) SetPduSessionId(v int32) {
	o.PduSessionId = v
}

// GetAssignedEbiList returns the AssignedEbiList field value
func (o *AssignedEbiData) GetAssignedEbiList() []EbiArpMapping {
	if o == nil {
		var ret []EbiArpMapping
		return ret
	}

	return o.AssignedEbiList
}

// GetAssignedEbiListOk returns a tuple with the AssignedEbiList field value
// and a boolean to check if the value has been set.
func (o *AssignedEbiData) GetAssignedEbiListOk() ([]EbiArpMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedEbiList, true
}

// SetAssignedEbiList sets field value
func (o *AssignedEbiData) SetAssignedEbiList(v []EbiArpMapping) {
	o.AssignedEbiList = v
}

// GetFailedArpList returns the FailedArpList field value if set, zero value otherwise.
func (o *AssignedEbiData) GetFailedArpList() []Arp {
	if o == nil || IsNil(o.FailedArpList) {
		var ret []Arp
		return ret
	}
	return o.FailedArpList
}

// GetFailedArpListOk returns a tuple with the FailedArpList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedEbiData) GetFailedArpListOk() ([]Arp, bool) {
	if o == nil || IsNil(o.FailedArpList) {
		return nil, false
	}
	return o.FailedArpList, true
}

// HasFailedArpList returns a boolean if a field has been set.
func (o *AssignedEbiData) HasFailedArpList() bool {
	if o != nil && !IsNil(o.FailedArpList) {
		return true
	}

	return false
}

// SetFailedArpList gets a reference to the given []Arp and assigns it to the FailedArpList field.
func (o *AssignedEbiData) SetFailedArpList(v []Arp) {
	o.FailedArpList = v
}

// GetReleasedEbiList returns the ReleasedEbiList field value if set, zero value otherwise.
func (o *AssignedEbiData) GetReleasedEbiList() []int32 {
	if o == nil || IsNil(o.ReleasedEbiList) {
		var ret []int32
		return ret
	}
	return o.ReleasedEbiList
}

// GetReleasedEbiListOk returns a tuple with the ReleasedEbiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedEbiData) GetReleasedEbiListOk() ([]int32, bool) {
	if o == nil || IsNil(o.ReleasedEbiList) {
		return nil, false
	}
	return o.ReleasedEbiList, true
}

// HasReleasedEbiList returns a boolean if a field has been set.
func (o *AssignedEbiData) HasReleasedEbiList() bool {
	if o != nil && !IsNil(o.ReleasedEbiList) {
		return true
	}

	return false
}

// SetReleasedEbiList gets a reference to the given []int32 and assigns it to the ReleasedEbiList field.
func (o *AssignedEbiData) SetReleasedEbiList(v []int32) {
	o.ReleasedEbiList = v
}

func (o AssignedEbiData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignedEbiData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pduSessionId"] = o.PduSessionId
	toSerialize["assignedEbiList"] = o.AssignedEbiList
	if !IsNil(o.FailedArpList) {
		toSerialize["failedArpList"] = o.FailedArpList
	}
	if !IsNil(o.ReleasedEbiList) {
		toSerialize["releasedEbiList"] = o.ReleasedEbiList
	}
	return toSerialize, nil
}

type NullableAssignedEbiData struct {
	value *AssignedEbiData
	isSet bool
}

func (v NullableAssignedEbiData) Get() *AssignedEbiData {
	return v.value
}

func (v *NullableAssignedEbiData) Set(val *AssignedEbiData) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignedEbiData) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignedEbiData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignedEbiData(val *AssignedEbiData) *NullableAssignedEbiData {
	return &NullableAssignedEbiData{value: val, isSet: true}
}

func (v NullableAssignedEbiData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignedEbiData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


