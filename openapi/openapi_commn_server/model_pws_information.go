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

type PwsInformation struct {
	MessageIdentifier int32 `json:"messageIdentifier"`

	SerialNumber int32 `json:"serialNumber"`

	PwsContainer N2InfoContent `json:"pwsContainer"`

	SendRanResponse bool `json:"sendRanResponse,omitempty"`

	OmcId string `json:"omcId,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *PwsInformation) UnmarshalJSON(data []byte) error {
	m.SendRanResponse = false

	type Alias PwsInformation // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertPwsInformationRequired checks if the required fields are not zero-ed
func AssertPwsInformationRequired(obj PwsInformation) error {
	elements := map[string]interface{}{
		"messageIdentifier": obj.MessageIdentifier,
		"serialNumber":      obj.SerialNumber,
		"pwsContainer":      obj.PwsContainer,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertN2InfoContentRequired(obj.PwsContainer); err != nil {
		return err
	}
	return nil
}

// AssertPwsInformationConstraints checks if the values respects the defined constraints
func AssertPwsInformationConstraints(obj PwsInformation) error {
	if obj.MessageIdentifier < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.MessageIdentifier > 65535 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.SerialNumber < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.SerialNumber > 65535 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}