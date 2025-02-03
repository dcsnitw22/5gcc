/*
 * Namf_Communication
 *
 * AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_commn_server

import "encoding/json"

type N2RanInformation struct {
	N2InfoContent N2InfoContent `json:"n2InfoContent"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *N2RanInformation) UnmarshalJSON(data []byte) error {

	type Alias N2RanInformation // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertN2RanInformationRequired checks if the required fields are not zero-ed
func AssertN2RanInformationRequired(obj N2RanInformation) error {
	elements := map[string]interface{}{
		"n2InfoContent": obj.N2InfoContent,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertN2InfoContentRequired(obj.N2InfoContent); err != nil {
		return err
	}
	return nil
}

// AssertN2RanInformationConstraints checks if the values respects the defined constraints
func AssertN2RanInformationConstraints(obj N2RanInformation) error {
	return nil
}
