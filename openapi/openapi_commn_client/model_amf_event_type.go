/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_commn_client

import (
	"encoding/json"
	"fmt"
)

// AmfEventType struct for AmfEventType
type AmfEventType struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AmfEventType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AmfEventType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AmfEventType) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAmfEventType struct {
	value *AmfEventType
	isSet bool
}

func (v NullableAmfEventType) Get() *AmfEventType {
	return v.value
}

func (v *NullableAmfEventType) Set(val *AmfEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEventType(val *AmfEventType) *NullableAmfEventType {
	return &NullableAmfEventType{value: val, isSet: true}
}

func (v NullableAmfEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


