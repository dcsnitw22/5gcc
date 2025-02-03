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
	"errors"
)

type SmContextReleaseData struct {
	Cause Cause `json:"cause,omitempty"`

	NgApCause NgApCause `json:"ngApCause,omitempty"`

	Var5gMmCauseValue int32 `json:"5gMmCauseValue,omitempty"`

	UeLocation UserLocation `json:"ueLocation,omitempty"`

	UeTimeZone string `json:"ueTimeZone,omitempty"`

	AddUeLocation UserLocation `json:"addUeLocation,omitempty"`

	VsmfReleaseOnly bool `json:"vsmfReleaseOnly,omitempty"`

	N2SmInfo RefToBinaryData `json:"n2SmInfo,omitempty"`

	N2SmInfoType N2SmInfoType `json:"n2SmInfoType,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *SmContextReleaseData) UnmarshalJSON(data []byte) error {
	m.VsmfReleaseOnly = false

	type Alias SmContextReleaseData // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertSmContextReleaseDataRequired checks if the required fields are not zero-ed
func AssertSmContextReleaseDataRequired(obj SmContextReleaseData) error {
	if err := AssertCauseRequired(obj.Cause); err != nil {
		return err
	}
	if err := AssertNgApCauseRequired(obj.NgApCause); err != nil {
		return err
	}
	if err := AssertUserLocationRequired(obj.UeLocation); err != nil {
		return err
	}
	if err := AssertUserLocationRequired(obj.AddUeLocation); err != nil {
		return err
	}
	if err := AssertRefToBinaryDataRequired(obj.N2SmInfo); err != nil {
		return err
	}
	if err := AssertN2SmInfoTypeRequired(obj.N2SmInfoType); err != nil {
		return err
	}
	return nil
}

// AssertSmContextReleaseDataConstraints checks if the values respects the defined constraints
func AssertSmContextReleaseDataConstraints(obj SmContextReleaseData) error {
	if obj.Var5gMmCauseValue < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
