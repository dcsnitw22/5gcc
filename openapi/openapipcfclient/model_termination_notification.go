/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the TerminationNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminationNotification{}

// TerminationNotification struct for TerminationNotification
type TerminationNotification struct {
	ResourceUri string `json:"resourceUri"`
	Cause PolicyAssociationReleaseCause `json:"cause"`
}

type _TerminationNotification TerminationNotification

// NewTerminationNotification instantiates a new TerminationNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminationNotification(resourceUri string, cause PolicyAssociationReleaseCause) *TerminationNotification {
	this := TerminationNotification{}
	this.ResourceUri = resourceUri
	this.Cause = cause
	return &this
}

// NewTerminationNotificationWithDefaults instantiates a new TerminationNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminationNotificationWithDefaults() *TerminationNotification {
	this := TerminationNotification{}
	return &this
}

// GetResourceUri returns the ResourceUri field value
func (o *TerminationNotification) GetResourceUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceUri
}

// GetResourceUriOk returns a tuple with the ResourceUri field value
// and a boolean to check if the value has been set.
func (o *TerminationNotification) GetResourceUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceUri, true
}

// SetResourceUri sets field value
func (o *TerminationNotification) SetResourceUri(v string) {
	o.ResourceUri = v
}

// GetCause returns the Cause field value
func (o *TerminationNotification) GetCause() PolicyAssociationReleaseCause {
	if o == nil {
		var ret PolicyAssociationReleaseCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *TerminationNotification) GetCauseOk() (*PolicyAssociationReleaseCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *TerminationNotification) SetCause(v PolicyAssociationReleaseCause) {
	o.Cause = v
}

func (o TerminationNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminationNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceUri"] = o.ResourceUri
	toSerialize["cause"] = o.Cause
	return toSerialize, nil
}

func (o *TerminationNotification) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceUri",
		"cause",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTerminationNotification := _TerminationNotification{}

	err = json.Unmarshal(bytes, &varTerminationNotification)

	if err != nil {
		return err
	}

	*o = TerminationNotification(varTerminationNotification)

	return err
}

type NullableTerminationNotification struct {
	value *TerminationNotification
	isSet bool
}

func (v NullableTerminationNotification) Get() *TerminationNotification {
	return v.value
}

func (v *NullableTerminationNotification) Set(val *TerminationNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminationNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminationNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminationNotification(val *TerminationNotification) *NullableTerminationNotification {
	return &NullableTerminationNotification{value: val, isSet: true}
}

func (v NullableTerminationNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminationNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

