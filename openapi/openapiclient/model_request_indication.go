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

// RequestIndication Possible values are - UE_REQ_PDU_SES_MOD - UE_REQ_PDU_SES_REL - PDU_SES_MOB - NW_REQ_PDU_SES_AUTH - NW_REQ_PDU_SES_MOD - NW_REQ_PDU_SES_REL - EBI_ASSIGNMENT_REQ 
type RequestIndication struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RequestIndication) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(RequestIndication)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RequestIndication) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRequestIndication struct {
	value *RequestIndication
	isSet bool
}

func (v NullableRequestIndication) Get() *RequestIndication {
	return v.value
}

func (v *NullableRequestIndication) Set(val *RequestIndication) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestIndication) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestIndication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestIndication(val *RequestIndication) *NullableRequestIndication {
	return &NullableRequestIndication{value: val, isSet: true}
}

func (v NullableRequestIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestIndication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


