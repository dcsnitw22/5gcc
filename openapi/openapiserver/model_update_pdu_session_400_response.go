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

type UpdatePduSession400Response struct {
	JsonData HsmfUpdateError `json:"jsonData,omitempty"`

	BinaryDataN1SmInfoToUe *os.File `json:"binaryDataN1SmInfoToUe,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *UpdatePduSession400Response) UnmarshalJSON(data []byte) error {

	type Alias UpdatePduSession400Response // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertUpdatePduSession400ResponseRequired checks if the required fields are not zero-ed
func AssertUpdatePduSession400ResponseRequired(obj UpdatePduSession400Response) error {
	if err := AssertHsmfUpdateErrorRequired(obj.JsonData); err != nil {
		return err
	}
	return nil
}

// AssertUpdatePduSession400ResponseConstraints checks if the values respects the defined constraints
func AssertUpdatePduSession400ResponseConstraints(obj UpdatePduSession400Response) error {
	return nil
}
