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

type MmeCapabilities struct {
	NonIpSupported bool `json:"nonIpSupported,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *MmeCapabilities) UnmarshalJSON(data []byte) error {
	m.NonIpSupported = false

	type Alias MmeCapabilities // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertMmeCapabilitiesRequired checks if the required fields are not zero-ed
func AssertMmeCapabilitiesRequired(obj MmeCapabilities) error {
	return nil
}

// AssertMmeCapabilitiesConstraints checks if the values respects the defined constraints
func AssertMmeCapabilitiesConstraints(obj MmeCapabilities) error {
	return nil
}
