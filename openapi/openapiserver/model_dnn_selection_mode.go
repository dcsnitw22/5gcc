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

// DnnSelectionMode - Possible values are - VERIFIED - UE_DNN_NOT_VERIFIED - NW_DNN_NOT_VERIFIED
type DnnSelectionMode struct {
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *DnnSelectionMode) UnmarshalJSON(data []byte) error {

	type Alias DnnSelectionMode // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertDnnSelectionModeRequired checks if the required fields are not zero-ed
func AssertDnnSelectionModeRequired(obj DnnSelectionMode) error {
	return nil
}

// AssertDnnSelectionModeConstraints checks if the values respects the defined constraints
func AssertDnnSelectionModeConstraints(obj DnnSelectionMode) error {
	return nil
}