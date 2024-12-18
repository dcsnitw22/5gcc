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

// QosFlowUsage Possible values are - GENERAL: Indicate no specific QoS flow usage information is available.  - IMS_SIG: Indicate that the QoS flow is used for IMS signalling only. 
type QosFlowUsage struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *QosFlowUsage) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(QosFlowUsage)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *QosFlowUsage) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableQosFlowUsage struct {
	value *QosFlowUsage
	isSet bool
}

func (v NullableQosFlowUsage) Get() *QosFlowUsage {
	return v.value
}

func (v *NullableQosFlowUsage) Set(val *QosFlowUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableQosFlowUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableQosFlowUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosFlowUsage(val *QosFlowUsage) *NullableQosFlowUsage {
	return &NullableQosFlowUsage{value: val, isSet: true}
}

func (v NullableQosFlowUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosFlowUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


