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

type SmContextUpdatedData struct {
	UpCnxState UpCnxState `json:"upCnxState,omitempty"`

	HoState HoState `json:"hoState,omitempty"`

	ReleaseEbiList []int32 `json:"releaseEbiList,omitempty"`

	AllocatedEbiList []EbiArpMapping `json:"allocatedEbiList,omitempty"`

	ModifiedEbiList []EbiArpMapping `json:"modifiedEbiList,omitempty"`

	N1SmMsg RefToBinaryData `json:"n1SmMsg,omitempty"`

	N2SmInfo RefToBinaryData `json:"n2SmInfo,omitempty"`

	N2SmInfoType N2SmInfoType `json:"n2SmInfoType,omitempty"`

	EpsBearerSetup []string `json:"epsBearerSetup,omitempty"`

	DataForwarding bool `json:"dataForwarding,omitempty"`

	Cause Cause `json:"cause,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *SmContextUpdatedData) UnmarshalJSON(data []byte) error {

	type Alias SmContextUpdatedData // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertSmContextUpdatedDataRequired checks if the required fields are not zero-ed
func AssertSmContextUpdatedDataRequired(obj SmContextUpdatedData) error {
	if err := AssertUpCnxStateRequired(obj.UpCnxState); err != nil {
		return err
	}
	if err := AssertHoStateRequired(obj.HoState); err != nil {
		return err
	}
	for _, el := range obj.AllocatedEbiList {
		if err := AssertEbiArpMappingRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ModifiedEbiList {
		if err := AssertEbiArpMappingRequired(el); err != nil {
			return err
		}
	}
	if err := AssertRefToBinaryDataRequired(obj.N1SmMsg); err != nil {
		return err
	}
	if err := AssertRefToBinaryDataRequired(obj.N2SmInfo); err != nil {
		return err
	}
	if err := AssertN2SmInfoTypeRequired(obj.N2SmInfoType); err != nil {
		return err
	}
	if err := AssertCauseRequired(obj.Cause); err != nil {
		return err
	}
	return nil
}

// AssertSmContextUpdatedDataConstraints checks if the values respects the defined constraints
func AssertSmContextUpdatedDataConstraints(obj SmContextUpdatedData) error {
	return nil
}