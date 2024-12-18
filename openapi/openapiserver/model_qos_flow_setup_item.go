/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiserver

import (
	"encoding/json"
	"errors"
)

type QosFlowSetupItem struct {
	Qfi int32 `json:"qfi"`

	QosRules string `json:"qosRules"`

	Ebi int32 `json:"ebi,omitempty"`

	QosFlowDescription string `json:"qosFlowDescription,omitempty"`

	QosFlowProfile QosFlowProfile `json:"qosFlowProfile,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *QosFlowSetupItem) UnmarshalJSON(data []byte) error {

	type Alias QosFlowSetupItem // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertQosFlowSetupItemRequired checks if the required fields are not zero-ed
func AssertQosFlowSetupItemRequired(obj QosFlowSetupItem) error {
	elements := map[string]interface{}{
		"qfi":      obj.Qfi,
		"qosRules": obj.QosRules,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertQosFlowProfileRequired(obj.QosFlowProfile); err != nil {
		return err
	}
	return nil
}

// AssertQosFlowSetupItemConstraints checks if the values respects the defined constraints
func AssertQosFlowSetupItemConstraints(obj QosFlowSetupItem) error {
	if obj.Qfi < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Qfi > 63 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.Ebi < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Ebi > 15 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
