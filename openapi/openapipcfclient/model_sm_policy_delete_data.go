/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the SmPolicyDeleteData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmPolicyDeleteData{}

// SmPolicyDeleteData struct for SmPolicyDeleteData
type SmPolicyDeleteData struct {
	UserLocationInfo *UserLocation `json:"userLocationInfo,omitempty"`
	UeTimeZone *string `json:"ueTimeZone,omitempty"`
	ServingNetwork *NetworkId `json:"servingNetwork,omitempty"`
	UserLocationInfoTime *time.Time `json:"userLocationInfoTime,omitempty"`
	// Contains the RAN and/or NAS release cause.
	RanNasRelCauses []RanNasRelCause `json:"ranNasRelCauses,omitempty"`
	// Contains the usage report
	AccuUsageReports []AccuUsageReport `json:"accuUsageReports,omitempty"`
}

// NewSmPolicyDeleteData instantiates a new SmPolicyDeleteData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmPolicyDeleteData() *SmPolicyDeleteData {
	this := SmPolicyDeleteData{}
	return &this
}

// NewSmPolicyDeleteDataWithDefaults instantiates a new SmPolicyDeleteData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmPolicyDeleteDataWithDefaults() *SmPolicyDeleteData {
	this := SmPolicyDeleteData{}
	return &this
}

// GetUserLocationInfo returns the UserLocationInfo field value if set, zero value otherwise.
func (o *SmPolicyDeleteData) GetUserLocationInfo() UserLocation {
	if o == nil || IsNil(o.UserLocationInfo) {
		var ret UserLocation
		return ret
	}
	return *o.UserLocationInfo
}

// GetUserLocationInfoOk returns a tuple with the UserLocationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyDeleteData) GetUserLocationInfoOk() (*UserLocation, bool) {
	if o == nil || IsNil(o.UserLocationInfo) {
		return nil, false
	}
	return o.UserLocationInfo, true
}

// HasUserLocationInfo returns a boolean if a field has been set.
func (o *SmPolicyDeleteData) HasUserLocationInfo() bool {
	if o != nil && !IsNil(o.UserLocationInfo) {
		return true
	}

	return false
}

// SetUserLocationInfo gets a reference to the given UserLocation and assigns it to the UserLocationInfo field.
func (o *SmPolicyDeleteData) SetUserLocationInfo(v UserLocation) {
	o.UserLocationInfo = &v
}

// GetUeTimeZone returns the UeTimeZone field value if set, zero value otherwise.
func (o *SmPolicyDeleteData) GetUeTimeZone() string {
	if o == nil || IsNil(o.UeTimeZone) {
		var ret string
		return ret
	}
	return *o.UeTimeZone
}

// GetUeTimeZoneOk returns a tuple with the UeTimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyDeleteData) GetUeTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.UeTimeZone) {
		return nil, false
	}
	return o.UeTimeZone, true
}

// HasUeTimeZone returns a boolean if a field has been set.
func (o *SmPolicyDeleteData) HasUeTimeZone() bool {
	if o != nil && !IsNil(o.UeTimeZone) {
		return true
	}

	return false
}

// SetUeTimeZone gets a reference to the given string and assigns it to the UeTimeZone field.
func (o *SmPolicyDeleteData) SetUeTimeZone(v string) {
	o.UeTimeZone = &v
}

// GetServingNetwork returns the ServingNetwork field value if set, zero value otherwise.
func (o *SmPolicyDeleteData) GetServingNetwork() NetworkId {
	if o == nil || IsNil(o.ServingNetwork) {
		var ret NetworkId
		return ret
	}
	return *o.ServingNetwork
}

// GetServingNetworkOk returns a tuple with the ServingNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyDeleteData) GetServingNetworkOk() (*NetworkId, bool) {
	if o == nil || IsNil(o.ServingNetwork) {
		return nil, false
	}
	return o.ServingNetwork, true
}

// HasServingNetwork returns a boolean if a field has been set.
func (o *SmPolicyDeleteData) HasServingNetwork() bool {
	if o != nil && !IsNil(o.ServingNetwork) {
		return true
	}

	return false
}

// SetServingNetwork gets a reference to the given NetworkId and assigns it to the ServingNetwork field.
func (o *SmPolicyDeleteData) SetServingNetwork(v NetworkId) {
	o.ServingNetwork = &v
}

// GetUserLocationInfoTime returns the UserLocationInfoTime field value if set, zero value otherwise.
func (o *SmPolicyDeleteData) GetUserLocationInfoTime() time.Time {
	if o == nil || IsNil(o.UserLocationInfoTime) {
		var ret time.Time
		return ret
	}
	return *o.UserLocationInfoTime
}

// GetUserLocationInfoTimeOk returns a tuple with the UserLocationInfoTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyDeleteData) GetUserLocationInfoTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UserLocationInfoTime) {
		return nil, false
	}
	return o.UserLocationInfoTime, true
}

// HasUserLocationInfoTime returns a boolean if a field has been set.
func (o *SmPolicyDeleteData) HasUserLocationInfoTime() bool {
	if o != nil && !IsNil(o.UserLocationInfoTime) {
		return true
	}

	return false
}

// SetUserLocationInfoTime gets a reference to the given time.Time and assigns it to the UserLocationInfoTime field.
func (o *SmPolicyDeleteData) SetUserLocationInfoTime(v time.Time) {
	o.UserLocationInfoTime = &v
}

// GetRanNasRelCauses returns the RanNasRelCauses field value if set, zero value otherwise.
func (o *SmPolicyDeleteData) GetRanNasRelCauses() []RanNasRelCause {
	if o == nil || IsNil(o.RanNasRelCauses) {
		var ret []RanNasRelCause
		return ret
	}
	return o.RanNasRelCauses
}

// GetRanNasRelCausesOk returns a tuple with the RanNasRelCauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyDeleteData) GetRanNasRelCausesOk() ([]RanNasRelCause, bool) {
	if o == nil || IsNil(o.RanNasRelCauses) {
		return nil, false
	}
	return o.RanNasRelCauses, true
}

// HasRanNasRelCauses returns a boolean if a field has been set.
func (o *SmPolicyDeleteData) HasRanNasRelCauses() bool {
	if o != nil && !IsNil(o.RanNasRelCauses) {
		return true
	}

	return false
}

// SetRanNasRelCauses gets a reference to the given []RanNasRelCause and assigns it to the RanNasRelCauses field.
func (o *SmPolicyDeleteData) SetRanNasRelCauses(v []RanNasRelCause) {
	o.RanNasRelCauses = v
}

// GetAccuUsageReports returns the AccuUsageReports field value if set, zero value otherwise.
func (o *SmPolicyDeleteData) GetAccuUsageReports() []AccuUsageReport {
	if o == nil || IsNil(o.AccuUsageReports) {
		var ret []AccuUsageReport
		return ret
	}
	return o.AccuUsageReports
}

// GetAccuUsageReportsOk returns a tuple with the AccuUsageReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyDeleteData) GetAccuUsageReportsOk() ([]AccuUsageReport, bool) {
	if o == nil || IsNil(o.AccuUsageReports) {
		return nil, false
	}
	return o.AccuUsageReports, true
}

// HasAccuUsageReports returns a boolean if a field has been set.
func (o *SmPolicyDeleteData) HasAccuUsageReports() bool {
	if o != nil && !IsNil(o.AccuUsageReports) {
		return true
	}

	return false
}

// SetAccuUsageReports gets a reference to the given []AccuUsageReport and assigns it to the AccuUsageReports field.
func (o *SmPolicyDeleteData) SetAccuUsageReports(v []AccuUsageReport) {
	o.AccuUsageReports = v
}

func (o SmPolicyDeleteData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmPolicyDeleteData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserLocationInfo) {
		toSerialize["userLocationInfo"] = o.UserLocationInfo
	}
	if !IsNil(o.UeTimeZone) {
		toSerialize["ueTimeZone"] = o.UeTimeZone
	}
	if !IsNil(o.ServingNetwork) {
		toSerialize["servingNetwork"] = o.ServingNetwork
	}
	if !IsNil(o.UserLocationInfoTime) {
		toSerialize["userLocationInfoTime"] = o.UserLocationInfoTime
	}
	if !IsNil(o.RanNasRelCauses) {
		toSerialize["ranNasRelCauses"] = o.RanNasRelCauses
	}
	if !IsNil(o.AccuUsageReports) {
		toSerialize["accuUsageReports"] = o.AccuUsageReports
	}
	return toSerialize, nil
}

type NullableSmPolicyDeleteData struct {
	value *SmPolicyDeleteData
	isSet bool
}

func (v NullableSmPolicyDeleteData) Get() *SmPolicyDeleteData {
	return v.value
}

func (v *NullableSmPolicyDeleteData) Set(val *SmPolicyDeleteData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmPolicyDeleteData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmPolicyDeleteData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmPolicyDeleteData(val *SmPolicyDeleteData) *NullableSmPolicyDeleteData {
	return &NullableSmPolicyDeleteData{value: val, isSet: true}
}

func (v NullableSmPolicyDeleteData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmPolicyDeleteData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

