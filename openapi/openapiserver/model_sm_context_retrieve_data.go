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

type SmContextRetrieveData struct {
	TargetMmeCap MmeCapabilities `json:"targetMmeCap,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *SmContextRetrieveData) UnmarshalJSON(data []byte) error {

	type Alias SmContextRetrieveData // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertSmContextRetrieveDataRequired checks if the required fields are not zero-ed
func AssertSmContextRetrieveDataRequired(obj SmContextRetrieveData) error {
	if err := AssertMmeCapabilitiesRequired(obj.TargetMmeCap); err != nil {
		return err
	}
	return nil
}

// AssertSmContextRetrieveDataConstraints checks if the values respects the defined constraints
func AssertSmContextRetrieveDataConstraints(obj SmContextRetrieveData) error {
	return nil
}