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

// checks if the NrLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NrLocation{}

// NrLocation struct for NrLocation
type NrLocation struct {
	Tai Tai `json:"tai"`
	Ncgi Ncgi `json:"ncgi"`
	AgeOfLocationInformation *int32 `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp *time.Time `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation *string `json:"geographicalInformation,omitempty"`
	GeodeticInformation *string `json:"geodeticInformation,omitempty"`
	GlobalGnbId interface{} `json:"globalGnbId,omitempty"`
}

// NewNrLocation instantiates a new NrLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNrLocation(tai Tai, ncgi Ncgi) *NrLocation {
	this := NrLocation{}
	this.Tai = tai
	this.Ncgi = ncgi
	return &this
}

// NewNrLocationWithDefaults instantiates a new NrLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNrLocationWithDefaults() *NrLocation {
	this := NrLocation{}
	return &this
}

// GetTai returns the Tai field value
func (o *NrLocation) GetTai() Tai {
	if o == nil {
		var ret Tai
		return ret
	}

	return o.Tai
}

// GetTaiOk returns a tuple with the Tai field value
// and a boolean to check if the value has been set.
func (o *NrLocation) GetTaiOk() (*Tai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tai, true
}

// SetTai sets field value
func (o *NrLocation) SetTai(v Tai) {
	o.Tai = v
}

// GetNcgi returns the Ncgi field value
func (o *NrLocation) GetNcgi() Ncgi {
	if o == nil {
		var ret Ncgi
		return ret
	}

	return o.Ncgi
}

// GetNcgiOk returns a tuple with the Ncgi field value
// and a boolean to check if the value has been set.
func (o *NrLocation) GetNcgiOk() (*Ncgi, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ncgi, true
}

// SetNcgi sets field value
func (o *NrLocation) SetNcgi(v Ncgi) {
	o.Ncgi = v
}

// GetAgeOfLocationInformation returns the AgeOfLocationInformation field value if set, zero value otherwise.
func (o *NrLocation) GetAgeOfLocationInformation() int32 {
	if o == nil || IsNil(o.AgeOfLocationInformation) {
		var ret int32
		return ret
	}
	return *o.AgeOfLocationInformation
}

// GetAgeOfLocationInformationOk returns a tuple with the AgeOfLocationInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrLocation) GetAgeOfLocationInformationOk() (*int32, bool) {
	if o == nil || IsNil(o.AgeOfLocationInformation) {
		return nil, false
	}
	return o.AgeOfLocationInformation, true
}

// HasAgeOfLocationInformation returns a boolean if a field has been set.
func (o *NrLocation) HasAgeOfLocationInformation() bool {
	if o != nil && !IsNil(o.AgeOfLocationInformation) {
		return true
	}

	return false
}

// SetAgeOfLocationInformation gets a reference to the given int32 and assigns it to the AgeOfLocationInformation field.
func (o *NrLocation) SetAgeOfLocationInformation(v int32) {
	o.AgeOfLocationInformation = &v
}

// GetUeLocationTimestamp returns the UeLocationTimestamp field value if set, zero value otherwise.
func (o *NrLocation) GetUeLocationTimestamp() time.Time {
	if o == nil || IsNil(o.UeLocationTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.UeLocationTimestamp
}

// GetUeLocationTimestampOk returns a tuple with the UeLocationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrLocation) GetUeLocationTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UeLocationTimestamp) {
		return nil, false
	}
	return o.UeLocationTimestamp, true
}

// HasUeLocationTimestamp returns a boolean if a field has been set.
func (o *NrLocation) HasUeLocationTimestamp() bool {
	if o != nil && !IsNil(o.UeLocationTimestamp) {
		return true
	}

	return false
}

// SetUeLocationTimestamp gets a reference to the given time.Time and assigns it to the UeLocationTimestamp field.
func (o *NrLocation) SetUeLocationTimestamp(v time.Time) {
	o.UeLocationTimestamp = &v
}

// GetGeographicalInformation returns the GeographicalInformation field value if set, zero value otherwise.
func (o *NrLocation) GetGeographicalInformation() string {
	if o == nil || IsNil(o.GeographicalInformation) {
		var ret string
		return ret
	}
	return *o.GeographicalInformation
}

// GetGeographicalInformationOk returns a tuple with the GeographicalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrLocation) GetGeographicalInformationOk() (*string, bool) {
	if o == nil || IsNil(o.GeographicalInformation) {
		return nil, false
	}
	return o.GeographicalInformation, true
}

// HasGeographicalInformation returns a boolean if a field has been set.
func (o *NrLocation) HasGeographicalInformation() bool {
	if o != nil && !IsNil(o.GeographicalInformation) {
		return true
	}

	return false
}

// SetGeographicalInformation gets a reference to the given string and assigns it to the GeographicalInformation field.
func (o *NrLocation) SetGeographicalInformation(v string) {
	o.GeographicalInformation = &v
}

// GetGeodeticInformation returns the GeodeticInformation field value if set, zero value otherwise.
func (o *NrLocation) GetGeodeticInformation() string {
	if o == nil || IsNil(o.GeodeticInformation) {
		var ret string
		return ret
	}
	return *o.GeodeticInformation
}

// GetGeodeticInformationOk returns a tuple with the GeodeticInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrLocation) GetGeodeticInformationOk() (*string, bool) {
	if o == nil || IsNil(o.GeodeticInformation) {
		return nil, false
	}
	return o.GeodeticInformation, true
}

// HasGeodeticInformation returns a boolean if a field has been set.
func (o *NrLocation) HasGeodeticInformation() bool {
	if o != nil && !IsNil(o.GeodeticInformation) {
		return true
	}

	return false
}

// SetGeodeticInformation gets a reference to the given string and assigns it to the GeodeticInformation field.
func (o *NrLocation) SetGeodeticInformation(v string) {
	o.GeodeticInformation = &v
}

// GetGlobalGnbId returns the GlobalGnbId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NrLocation) GetGlobalGnbId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.GlobalGnbId
}

// GetGlobalGnbIdOk returns a tuple with the GlobalGnbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NrLocation) GetGlobalGnbIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.GlobalGnbId) {
		return nil, false
	}
	return &o.GlobalGnbId, true
}

// HasGlobalGnbId returns a boolean if a field has been set.
func (o *NrLocation) HasGlobalGnbId() bool {
	if o != nil && IsNil(o.GlobalGnbId) {
		return true
	}

	return false
}

// SetGlobalGnbId gets a reference to the given interface{} and assigns it to the GlobalGnbId field.
func (o *NrLocation) SetGlobalGnbId(v interface{}) {
	o.GlobalGnbId = v
}

func (o NrLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NrLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tai"] = o.Tai
	toSerialize["ncgi"] = o.Ncgi
	if !IsNil(o.AgeOfLocationInformation) {
		toSerialize["ageOfLocationInformation"] = o.AgeOfLocationInformation
	}
	if !IsNil(o.UeLocationTimestamp) {
		toSerialize["ueLocationTimestamp"] = o.UeLocationTimestamp
	}
	if !IsNil(o.GeographicalInformation) {
		toSerialize["geographicalInformation"] = o.GeographicalInformation
	}
	if !IsNil(o.GeodeticInformation) {
		toSerialize["geodeticInformation"] = o.GeodeticInformation
	}
	if o.GlobalGnbId != nil {
		toSerialize["globalGnbId"] = o.GlobalGnbId
	}
	return toSerialize, nil
}

type NullableNrLocation struct {
	value *NrLocation
	isSet bool
}

func (v NullableNrLocation) Get() *NrLocation {
	return v.value
}

func (v *NullableNrLocation) Set(val *NrLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableNrLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableNrLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrLocation(val *NrLocation) *NullableNrLocation {
	return &NullableNrLocation{value: val, isSet: true}
}

func (v NullableNrLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


