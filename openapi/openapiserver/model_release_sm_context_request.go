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

type ReleaseSmContextRequest struct {
	JsonData SmContextReleaseData `json:"jsonData,omitempty"`

	BinaryDataN2SmInformation *os.File `json:"binaryDataN2SmInformation,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *ReleaseSmContextRequest) UnmarshalJSON(data []byte) error {

	type Alias ReleaseSmContextRequest // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertReleaseSmContextRequestRequired checks if the required fields are not zero-ed
func AssertReleaseSmContextRequestRequired(obj ReleaseSmContextRequest) error {
	if err := AssertSmContextReleaseDataRequired(obj.JsonData); err != nil {
		return err
	}
	return nil
}

// AssertReleaseSmContextRequestConstraints checks if the values respects the defined constraints
func AssertReleaseSmContextRequestConstraints(obj ReleaseSmContextRequest) error {
	return nil
}
