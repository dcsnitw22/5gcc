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
)

type StatusInfo struct {
	ResourceStatus ResourceStatus `json:"resourceStatus"`

	Cause Cause `json:"cause,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *StatusInfo) UnmarshalJSON(data []byte) error {

	type Alias StatusInfo // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertStatusInfoRequired checks if the required fields are not zero-ed
func AssertStatusInfoRequired(obj StatusInfo) error {
	elements := map[string]interface{}{
		"resourceStatus": obj.ResourceStatus,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertResourceStatusRequired(obj.ResourceStatus); err != nil {
		return err
	}
	if err := AssertCauseRequired(obj.Cause); err != nil {
		return err
	}
	return nil
}

// AssertStatusInfoConstraints checks if the values respects the defined constraints
func AssertStatusInfoConstraints(obj StatusInfo) error {
	return nil
}
