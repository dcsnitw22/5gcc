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

// EventForMdt The enumeration EventForMdt defines events triggered measurement for logged MDT in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.11-1 
type EventForMdt struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EventForMdt) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(EventForMdt)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EventForMdt) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEventForMdt struct {
	value *EventForMdt
	isSet bool
}

func (v NullableEventForMdt) Get() *EventForMdt {
	return v.value
}

func (v *NullableEventForMdt) Set(val *EventForMdt) {
	v.value = val
	v.isSet = true
}

func (v NullableEventForMdt) IsSet() bool {
	return v.isSet
}

func (v *NullableEventForMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventForMdt(val *EventForMdt) *NullableEventForMdt {
	return &NullableEventForMdt{value: val, isSet: true}
}

func (v NullableEventForMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventForMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


