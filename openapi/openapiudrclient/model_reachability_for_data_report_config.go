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

// ReachabilityForDataReportConfig struct for ReachabilityForDataReportConfig
type ReachabilityForDataReportConfig struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReachabilityForDataReportConfig) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReachabilityForDataReportConfig)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReachabilityForDataReportConfig) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReachabilityForDataReportConfig struct {
	value *ReachabilityForDataReportConfig
	isSet bool
}

func (v NullableReachabilityForDataReportConfig) Get() *ReachabilityForDataReportConfig {
	return v.value
}

func (v *NullableReachabilityForDataReportConfig) Set(val *ReachabilityForDataReportConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableReachabilityForDataReportConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableReachabilityForDataReportConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReachabilityForDataReportConfig(val *ReachabilityForDataReportConfig) *NullableReachabilityForDataReportConfig {
	return &NullableReachabilityForDataReportConfig{value: val, isSet: true}
}

func (v NullableReachabilityForDataReportConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReachabilityForDataReportConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


