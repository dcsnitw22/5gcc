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

// LcsMoServiceClass struct for LcsMoServiceClass
type LcsMoServiceClass struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LcsMoServiceClass) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(LcsMoServiceClass)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LcsMoServiceClass) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLcsMoServiceClass struct {
	value *LcsMoServiceClass
	isSet bool
}

func (v NullableLcsMoServiceClass) Get() *LcsMoServiceClass {
	return v.value
}

func (v *NullableLcsMoServiceClass) Set(val *LcsMoServiceClass) {
	v.value = val
	v.isSet = true
}

func (v NullableLcsMoServiceClass) IsSet() bool {
	return v.isSet
}

func (v *NullableLcsMoServiceClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLcsMoServiceClass(val *LcsMoServiceClass) *NullableLcsMoServiceClass {
	return &NullableLcsMoServiceClass{value: val, isSet: true}
}

func (v NullableLcsMoServiceClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLcsMoServiceClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


