/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_udr_cli

import (
	"encoding/json"
	"fmt"
)

// UsageMonLevel Represents the usage monitoring level.
type UsageMonLevel struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UsageMonLevel) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(UsageMonLevel)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UsageMonLevel) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUsageMonLevel struct {
	value *UsageMonLevel
	isSet bool
}

func (v NullableUsageMonLevel) Get() *UsageMonLevel {
	return v.value
}

func (v *NullableUsageMonLevel) Set(val *UsageMonLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageMonLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageMonLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageMonLevel(val *UsageMonLevel) *NullableUsageMonLevel {
	return &NullableUsageMonLevel{value: val, isSet: true}
}

func (v NullableUsageMonLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageMonLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


