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

// N2InfoNotifyReason struct for N2InfoNotifyReason
type N2InfoNotifyReason struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *N2InfoNotifyReason) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(N2InfoNotifyReason)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *N2InfoNotifyReason) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableN2InfoNotifyReason struct {
	value *N2InfoNotifyReason
	isSet bool
}

func (v NullableN2InfoNotifyReason) Get() *N2InfoNotifyReason {
	return v.value
}

func (v *NullableN2InfoNotifyReason) Set(val *N2InfoNotifyReason) {
	v.value = val
	v.isSet = true
}

func (v NullableN2InfoNotifyReason) IsSet() bool {
	return v.isSet
}

func (v *NullableN2InfoNotifyReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2InfoNotifyReason(val *N2InfoNotifyReason) *NullableN2InfoNotifyReason {
	return &NullableN2InfoNotifyReason{value: val, isSet: true}
}

func (v NullableN2InfoNotifyReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2InfoNotifyReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


