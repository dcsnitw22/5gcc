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
	"errors"
)

type N2SmInformation struct {
	PduSessionId int32 `json:"pduSessionId"`

	N2InfoContent N2InfoContent `json:"n2InfoContent,omitempty"`

	SNssai Snssai `json:"sNssai,omitempty"`

	SubjectToHo bool `json:"subjectToHo,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *N2SmInformation) UnmarshalJSON(data []byte) error {

	type Alias N2SmInformation // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertN2SmInformationRequired checks if the required fields are not zero-ed
func AssertN2SmInformationRequired(obj N2SmInformation) error {
	elements := map[string]interface{}{
		"pduSessionId": obj.PduSessionId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertN2InfoContentRequired(obj.N2InfoContent); err != nil {
		return err
	}
	//TODO add later
	// if err := AssertSnssaiRequired(obj.SNssai); err != nil {
	// 	return err
	// }
	return nil
}

// AssertN2SmInformationConstraints checks if the values respects the defined constraints
func AssertN2SmInformationConstraints(obj N2SmInformation) error {
	if obj.PduSessionId < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.PduSessionId > 255 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
