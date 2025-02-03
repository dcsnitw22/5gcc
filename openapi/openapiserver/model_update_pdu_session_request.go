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
	"os"
)

type UpdatePduSessionRequest struct {
	JsonData HsmfUpdateData `json:"jsonData,omitempty"`

	BinaryDataN1SmInfoFromUe *os.File `json:"binaryDataN1SmInfoFromUe,omitempty"`

	BinaryDataUnknownN1SmInfo *os.File `json:"binaryDataUnknownN1SmInfo,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *UpdatePduSessionRequest) UnmarshalJSON(data []byte) error {

	type Alias UpdatePduSessionRequest // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertUpdatePduSessionRequestRequired checks if the required fields are not zero-ed
func AssertUpdatePduSessionRequestRequired(obj UpdatePduSessionRequest) error {
	if err := AssertHsmfUpdateDataRequired(obj.JsonData); err != nil {
		return err
	}
	return nil
}

// AssertUpdatePduSessionRequestConstraints checks if the values respects the defined constraints
func AssertUpdatePduSessionRequestConstraints(obj UpdatePduSessionRequest) error {
	return nil
}
