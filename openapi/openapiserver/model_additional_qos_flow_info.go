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

type AdditionalQosFlowInfo struct {
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *AdditionalQosFlowInfo) UnmarshalJSON(data []byte) error {

	type Alias AdditionalQosFlowInfo // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertAdditionalQosFlowInfoRequired checks if the required fields are not zero-ed
func AssertAdditionalQosFlowInfoRequired(obj AdditionalQosFlowInfo) error {
	return nil
}

// AssertAdditionalQosFlowInfoConstraints checks if the values respects the defined constraints
func AssertAdditionalQosFlowInfoConstraints(obj AdditionalQosFlowInfo) error {
	return nil
}