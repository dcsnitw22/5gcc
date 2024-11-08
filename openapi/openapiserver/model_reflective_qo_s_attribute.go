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

type ReflectiveQoSAttribute struct {
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *ReflectiveQoSAttribute) UnmarshalJSON(data []byte) error {

	type Alias ReflectiveQoSAttribute // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertReflectiveQoSAttributeRequired checks if the required fields are not zero-ed
func AssertReflectiveQoSAttributeRequired(obj ReflectiveQoSAttribute) error {
	return nil
}

// AssertReflectiveQoSAttributeConstraints checks if the values respects the defined constraints
func AssertReflectiveQoSAttributeConstraints(obj ReflectiveQoSAttribute) error {
	return nil
}
