/*
Nsmf_PDUSession

SMF PDU Session Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
	"time"
)

// checks if the QosFlowUsageReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QosFlowUsageReport{}

// QosFlowUsageReport struct for QosFlowUsageReport
type QosFlowUsageReport struct {
	Qfi int32 `json:"qfi"`
	StartTimeStamp time.Time `json:"startTimeStamp"`
	EndTimeStamp time.Time `json:"endTimeStamp"`
	DownlinkVolume int64 `json:"downlinkVolume"`
	UplinkVolume int64 `json:"uplinkVolume"`
}

// NewQosFlowUsageReport instantiates a new QosFlowUsageReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosFlowUsageReport(qfi int32, startTimeStamp time.Time, endTimeStamp time.Time, downlinkVolume int64, uplinkVolume int64) *QosFlowUsageReport {
	this := QosFlowUsageReport{}
	this.Qfi = qfi
	this.StartTimeStamp = startTimeStamp
	this.EndTimeStamp = endTimeStamp
	this.DownlinkVolume = downlinkVolume
	this.UplinkVolume = uplinkVolume
	return &this
}

// NewQosFlowUsageReportWithDefaults instantiates a new QosFlowUsageReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosFlowUsageReportWithDefaults() *QosFlowUsageReport {
	this := QosFlowUsageReport{}
	return &this
}

// GetQfi returns the Qfi field value
func (o *QosFlowUsageReport) GetQfi() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Qfi
}

// GetQfiOk returns a tuple with the Qfi field value
// and a boolean to check if the value has been set.
func (o *QosFlowUsageReport) GetQfiOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Qfi, true
}

// SetQfi sets field value
func (o *QosFlowUsageReport) SetQfi(v int32) {
	o.Qfi = v
}

// GetStartTimeStamp returns the StartTimeStamp field value
func (o *QosFlowUsageReport) GetStartTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTimeStamp
}

// GetStartTimeStampOk returns a tuple with the StartTimeStamp field value
// and a boolean to check if the value has been set.
func (o *QosFlowUsageReport) GetStartTimeStampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTimeStamp, true
}

// SetStartTimeStamp sets field value
func (o *QosFlowUsageReport) SetStartTimeStamp(v time.Time) {
	o.StartTimeStamp = v
}

// GetEndTimeStamp returns the EndTimeStamp field value
func (o *QosFlowUsageReport) GetEndTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTimeStamp
}

// GetEndTimeStampOk returns a tuple with the EndTimeStamp field value
// and a boolean to check if the value has been set.
func (o *QosFlowUsageReport) GetEndTimeStampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTimeStamp, true
}

// SetEndTimeStamp sets field value
func (o *QosFlowUsageReport) SetEndTimeStamp(v time.Time) {
	o.EndTimeStamp = v
}

// GetDownlinkVolume returns the DownlinkVolume field value
func (o *QosFlowUsageReport) GetDownlinkVolume() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DownlinkVolume
}

// GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field value
// and a boolean to check if the value has been set.
func (o *QosFlowUsageReport) GetDownlinkVolumeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownlinkVolume, true
}

// SetDownlinkVolume sets field value
func (o *QosFlowUsageReport) SetDownlinkVolume(v int64) {
	o.DownlinkVolume = v
}

// GetUplinkVolume returns the UplinkVolume field value
func (o *QosFlowUsageReport) GetUplinkVolume() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UplinkVolume
}

// GetUplinkVolumeOk returns a tuple with the UplinkVolume field value
// and a boolean to check if the value has been set.
func (o *QosFlowUsageReport) GetUplinkVolumeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UplinkVolume, true
}

// SetUplinkVolume sets field value
func (o *QosFlowUsageReport) SetUplinkVolume(v int64) {
	o.UplinkVolume = v
}

func (o QosFlowUsageReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QosFlowUsageReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["qfi"] = o.Qfi
	toSerialize["startTimeStamp"] = o.StartTimeStamp
	toSerialize["endTimeStamp"] = o.EndTimeStamp
	toSerialize["downlinkVolume"] = o.DownlinkVolume
	toSerialize["uplinkVolume"] = o.UplinkVolume
	return toSerialize, nil
}

type NullableQosFlowUsageReport struct {
	value *QosFlowUsageReport
	isSet bool
}

func (v NullableQosFlowUsageReport) Get() *QosFlowUsageReport {
	return v.value
}

func (v *NullableQosFlowUsageReport) Set(val *QosFlowUsageReport) {
	v.value = val
	v.isSet = true
}

func (v NullableQosFlowUsageReport) IsSet() bool {
	return v.isSet
}

func (v *NullableQosFlowUsageReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosFlowUsageReport(val *QosFlowUsageReport) *NullableQosFlowUsageReport {
	return &NullableQosFlowUsageReport{value: val, isSet: true}
}

func (v NullableQosFlowUsageReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosFlowUsageReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


