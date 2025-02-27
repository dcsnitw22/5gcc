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

// N2InformationTransferResult struct for N2InformationTransferResult
type N2InformationTransferResult struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *N2InformationTransferResult) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(N2InformationTransferResult)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *N2InformationTransferResult) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableN2InformationTransferResult struct {
	value *N2InformationTransferResult
	isSet bool
}

func (v NullableN2InformationTransferResult) Get() *N2InformationTransferResult {
	return v.value
}

func (v *NullableN2InformationTransferResult) Set(val *N2InformationTransferResult) {
	v.value = val
	v.isSet = true
}

func (v NullableN2InformationTransferResult) IsSet() bool {
	return v.isSet
}

func (v *NullableN2InformationTransferResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2InformationTransferResult(val *N2InformationTransferResult) *NullableN2InformationTransferResult {
	return &NullableN2InformationTransferResult{value: val, isSet: true}
}

func (v NullableN2InformationTransferResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2InformationTransferResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


