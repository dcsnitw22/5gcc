/*
Nsmf_PDUSession

SMF PDU Session Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
	"fmt"
)

// EpsInterworkingIndication Possible values are - NONE - WITH_N26 - WITHOUT_N26 
type EpsInterworkingIndication struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EpsInterworkingIndication) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(EpsInterworkingIndication)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EpsInterworkingIndication) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEpsInterworkingIndication struct {
	value *EpsInterworkingIndication
	isSet bool
}

func (v NullableEpsInterworkingIndication) Get() *EpsInterworkingIndication {
	return v.value
}

func (v *NullableEpsInterworkingIndication) Set(val *EpsInterworkingIndication) {
	v.value = val
	v.isSet = true
}

func (v NullableEpsInterworkingIndication) IsSet() bool {
	return v.isSet
}

func (v *NullableEpsInterworkingIndication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpsInterworkingIndication(val *EpsInterworkingIndication) *NullableEpsInterworkingIndication {
	return &NullableEpsInterworkingIndication{value: val, isSet: true}
}

func (v NullableEpsInterworkingIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpsInterworkingIndication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

