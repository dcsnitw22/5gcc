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

// checks if the UserLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserLocation{}

// UserLocation struct for UserLocation
type UserLocation struct {
	EutraLocation *EutraLocation `json:"eutraLocation,omitempty"`
	NrLocation *NrLocation `json:"nrLocation,omitempty"`
	N3gaLocation *N3gaLocation `json:"n3gaLocation,omitempty"`
}

// NewUserLocation instantiates a new UserLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLocation() *UserLocation {
	this := UserLocation{}
	return &this
}

// NewUserLocationWithDefaults instantiates a new UserLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLocationWithDefaults() *UserLocation {
	this := UserLocation{}
	return &this
}

// GetEutraLocation returns the EutraLocation field value if set, zero value otherwise.
func (o *UserLocation) GetEutraLocation() EutraLocation {
	if o == nil || IsNil(o.EutraLocation) {
		var ret EutraLocation
		return ret
	}
	return *o.EutraLocation
}

// GetEutraLocationOk returns a tuple with the EutraLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLocation) GetEutraLocationOk() (*EutraLocation, bool) {
	if o == nil || IsNil(o.EutraLocation) {
		return nil, false
	}
	return o.EutraLocation, true
}

// HasEutraLocation returns a boolean if a field has been set.
func (o *UserLocation) HasEutraLocation() bool {
	if o != nil && !IsNil(o.EutraLocation) {
		return true
	}

	return false
}

// SetEutraLocation gets a reference to the given EutraLocation and assigns it to the EutraLocation field.
func (o *UserLocation) SetEutraLocation(v EutraLocation) {
	o.EutraLocation = &v
}

// GetNrLocation returns the NrLocation field value if set, zero value otherwise.
func (o *UserLocation) GetNrLocation() NrLocation {
	if o == nil || IsNil(o.NrLocation) {
		var ret NrLocation
		return ret
	}
	return *o.NrLocation
}

// GetNrLocationOk returns a tuple with the NrLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLocation) GetNrLocationOk() (*NrLocation, bool) {
	if o == nil || IsNil(o.NrLocation) {
		return nil, false
	}
	return o.NrLocation, true
}

// HasNrLocation returns a boolean if a field has been set.
func (o *UserLocation) HasNrLocation() bool {
	if o != nil && !IsNil(o.NrLocation) {
		return true
	}

	return false
}

// SetNrLocation gets a reference to the given NrLocation and assigns it to the NrLocation field.
func (o *UserLocation) SetNrLocation(v NrLocation) {
	o.NrLocation = &v
}

// GetN3gaLocation returns the N3gaLocation field value if set, zero value otherwise.
func (o *UserLocation) GetN3gaLocation() N3gaLocation {
	if o == nil || IsNil(o.N3gaLocation) {
		var ret N3gaLocation
		return ret
	}
	return *o.N3gaLocation
}

// GetN3gaLocationOk returns a tuple with the N3gaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLocation) GetN3gaLocationOk() (*N3gaLocation, bool) {
	if o == nil || IsNil(o.N3gaLocation) {
		return nil, false
	}
	return o.N3gaLocation, true
}

// HasN3gaLocation returns a boolean if a field has been set.
func (o *UserLocation) HasN3gaLocation() bool {
	if o != nil && !IsNil(o.N3gaLocation) {
		return true
	}

	return false
}

// SetN3gaLocation gets a reference to the given N3gaLocation and assigns it to the N3gaLocation field.
func (o *UserLocation) SetN3gaLocation(v N3gaLocation) {
	o.N3gaLocation = &v
}

func (o UserLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EutraLocation) {
		toSerialize["eutraLocation"] = o.EutraLocation
	}
	if !IsNil(o.NrLocation) {
		toSerialize["nrLocation"] = o.NrLocation
	}
	if !IsNil(o.N3gaLocation) {
		toSerialize["n3gaLocation"] = o.N3gaLocation
	}
	return toSerialize, nil
}

type NullableUserLocation struct {
	value *UserLocation
	isSet bool
}

func (v NullableUserLocation) Get() *UserLocation {
	return v.value
}

func (v *NullableUserLocation) Set(val *UserLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLocation(val *UserLocation) *NullableUserLocation {
	return &NullableUserLocation{value: val, isSet: true}
}

func (v NullableUserLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


