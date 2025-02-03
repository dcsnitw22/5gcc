/*
 * Namf_Communication
 *
 * AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_commn_server

import (
	"encoding/json"
)

type Guami struct {
	PlmnId PlmnId `json:"plmnId"`

	AmfId string `json:"amfId"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *Guami) UnmarshalJSON(data []byte) error {

	type Alias Guami // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertGuamiRequired checks if the required fields are not zero-ed
func AssertGuamiRequired(obj Guami) error {
	elements := map[string]interface{}{
		"plmnId": obj.PlmnId,
		"amfId":  obj.AmfId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertPlmnIdRequired(obj.PlmnId); err != nil {
		return err
	}
	return nil
}

// AssertGuamiConstraints checks if the values respects the defined constraints
func AssertGuamiConstraints(obj Guami) error {
	return nil
}
