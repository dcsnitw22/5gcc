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

type NgApCause struct {
	Group int32 `json:"group"`

	Value int32 `json:"value"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *NgApCause) UnmarshalJSON(data []byte) error {

	type Alias NgApCause // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertNgApCauseRequired checks if the required fields are not zero-ed
func AssertNgApCauseRequired(obj NgApCause) error {
	elements := map[string]interface{}{
		"group": obj.Group,
		"value": obj.Value,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertNgApCauseConstraints checks if the values respects the defined constraints
func AssertNgApCauseConstraints(obj NgApCause) error {
	if obj.Group < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Value < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
