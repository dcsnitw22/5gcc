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

type UpSecurity struct {
	UpIntegr UpIntegrity `json:"upIntegr"`

	UpConfid UpConfidentiality `json:"upConfid"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *UpSecurity) UnmarshalJSON(data []byte) error {

	type Alias UpSecurity // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertUpSecurityRequired checks if the required fields are not zero-ed
func AssertUpSecurityRequired(obj UpSecurity) error {
	elements := map[string]interface{}{
		"upIntegr": obj.UpIntegr,
		"upConfid": obj.UpConfid,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertUpIntegrityRequired(obj.UpIntegr); err != nil {
		return err
	}
	if err := AssertUpConfidentialityRequired(obj.UpConfid); err != nil {
		return err
	}
	return nil
}

// AssertUpSecurityConstraints checks if the values respects the defined constraints
func AssertUpSecurityConstraints(obj UpSecurity) error {
	return nil
}