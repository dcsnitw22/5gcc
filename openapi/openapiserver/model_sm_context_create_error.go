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
	"time"
)

type SmContextCreateError struct {
	Error ProblemDetails `json:"error"`

	N1SmMsg RefToBinaryData `json:"n1SmMsg,omitempty"`

	RecoveryTime time.Time `json:"recoveryTime,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *SmContextCreateError) UnmarshalJSON(data []byte) error {

	type Alias SmContextCreateError // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertSmContextCreateErrorRequired checks if the required fields are not zero-ed
func AssertSmContextCreateErrorRequired(obj SmContextCreateError) error {
	elements := map[string]interface{}{
		"error": obj.Error,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertProblemDetailsRequired(obj.Error); err != nil {
		return err
	}
	if err := AssertRefToBinaryDataRequired(obj.N1SmMsg); err != nil {
		return err
	}
	return nil
}

// AssertSmContextCreateErrorConstraints checks if the values respects the defined constraints
func AssertSmContextCreateErrorConstraints(obj SmContextCreateError) error {
	return nil
}
