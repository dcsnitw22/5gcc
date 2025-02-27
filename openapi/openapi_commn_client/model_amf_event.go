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

// checks if the AmfEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmfEvent{}

// AmfEvent struct for AmfEvent
type AmfEvent struct {
	Type AmfEventType `json:"type"`
	ImmediateFlag *bool `json:"immediateFlag,omitempty"`
	AreaList []AmfEventArea `json:"areaList,omitempty"`
	LocationFilterList []LocationFilter `json:"locationFilterList,omitempty"`
	RefId *int32 `json:"refId,omitempty"`
	ReachabilityFilter *ReachabilityFilter `json:"reachabilityFilter,omitempty"`
}

// NewAmfEvent instantiates a new AmfEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfEvent(type_ AmfEventType) *AmfEvent {
	this := AmfEvent{}
	this.Type = type_
	return &this
}

// NewAmfEventWithDefaults instantiates a new AmfEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfEventWithDefaults() *AmfEvent {
	this := AmfEvent{}
	return &this
}

// GetType returns the Type field value
func (o *AmfEvent) GetType() AmfEventType {
	if o == nil {
		var ret AmfEventType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetTypeOk() (*AmfEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AmfEvent) SetType(v AmfEventType) {
	o.Type = v
}

// GetImmediateFlag returns the ImmediateFlag field value if set, zero value otherwise.
func (o *AmfEvent) GetImmediateFlag() bool {
	if o == nil || IsNil(o.ImmediateFlag) {
		var ret bool
		return ret
	}
	return *o.ImmediateFlag
}

// GetImmediateFlagOk returns a tuple with the ImmediateFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetImmediateFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.ImmediateFlag) {
		return nil, false
	}
	return o.ImmediateFlag, true
}

// HasImmediateFlag returns a boolean if a field has been set.
func (o *AmfEvent) HasImmediateFlag() bool {
	if o != nil && !IsNil(o.ImmediateFlag) {
		return true
	}

	return false
}

// SetImmediateFlag gets a reference to the given bool and assigns it to the ImmediateFlag field.
func (o *AmfEvent) SetImmediateFlag(v bool) {
	o.ImmediateFlag = &v
}

// GetAreaList returns the AreaList field value if set, zero value otherwise.
func (o *AmfEvent) GetAreaList() []AmfEventArea {
	if o == nil || IsNil(o.AreaList) {
		var ret []AmfEventArea
		return ret
	}
	return o.AreaList
}

// GetAreaListOk returns a tuple with the AreaList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetAreaListOk() ([]AmfEventArea, bool) {
	if o == nil || IsNil(o.AreaList) {
		return nil, false
	}
	return o.AreaList, true
}

// HasAreaList returns a boolean if a field has been set.
func (o *AmfEvent) HasAreaList() bool {
	if o != nil && !IsNil(o.AreaList) {
		return true
	}

	return false
}

// SetAreaList gets a reference to the given []AmfEventArea and assigns it to the AreaList field.
func (o *AmfEvent) SetAreaList(v []AmfEventArea) {
	o.AreaList = v
}

// GetLocationFilterList returns the LocationFilterList field value if set, zero value otherwise.
func (o *AmfEvent) GetLocationFilterList() []LocationFilter {
	if o == nil || IsNil(o.LocationFilterList) {
		var ret []LocationFilter
		return ret
	}
	return o.LocationFilterList
}

// GetLocationFilterListOk returns a tuple with the LocationFilterList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetLocationFilterListOk() ([]LocationFilter, bool) {
	if o == nil || IsNil(o.LocationFilterList) {
		return nil, false
	}
	return o.LocationFilterList, true
}

// HasLocationFilterList returns a boolean if a field has been set.
func (o *AmfEvent) HasLocationFilterList() bool {
	if o != nil && !IsNil(o.LocationFilterList) {
		return true
	}

	return false
}

// SetLocationFilterList gets a reference to the given []LocationFilter and assigns it to the LocationFilterList field.
func (o *AmfEvent) SetLocationFilterList(v []LocationFilter) {
	o.LocationFilterList = v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *AmfEvent) GetRefId() int32 {
	if o == nil || IsNil(o.RefId) {
		var ret int32
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetRefIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RefId) {
		return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *AmfEvent) HasRefId() bool {
	if o != nil && !IsNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given int32 and assigns it to the RefId field.
func (o *AmfEvent) SetRefId(v int32) {
	o.RefId = &v
}

// GetReachabilityFilter returns the ReachabilityFilter field value if set, zero value otherwise.
func (o *AmfEvent) GetReachabilityFilter() ReachabilityFilter {
	if o == nil || IsNil(o.ReachabilityFilter) {
		var ret ReachabilityFilter
		return ret
	}
	return *o.ReachabilityFilter
}

// GetReachabilityFilterOk returns a tuple with the ReachabilityFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetReachabilityFilterOk() (*ReachabilityFilter, bool) {
	if o == nil || IsNil(o.ReachabilityFilter) {
		return nil, false
	}
	return o.ReachabilityFilter, true
}

// HasReachabilityFilter returns a boolean if a field has been set.
func (o *AmfEvent) HasReachabilityFilter() bool {
	if o != nil && !IsNil(o.ReachabilityFilter) {
		return true
	}

	return false
}

// SetReachabilityFilter gets a reference to the given ReachabilityFilter and assigns it to the ReachabilityFilter field.
func (o *AmfEvent) SetReachabilityFilter(v ReachabilityFilter) {
	o.ReachabilityFilter = &v
}

func (o AmfEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmfEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.ImmediateFlag) {
		toSerialize["immediateFlag"] = o.ImmediateFlag
	}
	if !IsNil(o.AreaList) {
		toSerialize["areaList"] = o.AreaList
	}
	if !IsNil(o.LocationFilterList) {
		toSerialize["locationFilterList"] = o.LocationFilterList
	}
	if !IsNil(o.RefId) {
		toSerialize["refId"] = o.RefId
	}
	if !IsNil(o.ReachabilityFilter) {
		toSerialize["reachabilityFilter"] = o.ReachabilityFilter
	}
	return toSerialize, nil
}

type NullableAmfEvent struct {
	value *AmfEvent
	isSet bool
}

func (v NullableAmfEvent) Get() *AmfEvent {
	return v.value
}

func (v *NullableAmfEvent) Set(val *AmfEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEvent(val *AmfEvent) *NullableAmfEvent {
	return &NullableAmfEvent{value: val, isSet: true}
}

func (v NullableAmfEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


