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

type Snssai struct {
	Sst int32 `json:"sst"`

	Sd string `json:"sd,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *Snssai) UnmarshalJSON(data []byte) error {

	type Alias Snssai // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertSnssaiRequired checks if the required fields are not zero-ed
func AssertSnssaiRequired(obj Snssai) error {
	elements := map[string]interface{}{
		"sst": obj.Sst,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertSnssaiConstraints checks if the values respects the defined constraints
func AssertSnssaiConstraints(obj Snssai) error {
	if obj.Sst < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Sst > 255 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
