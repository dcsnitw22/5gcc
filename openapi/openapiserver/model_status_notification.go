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

type StatusNotification struct {
	StatusInfo StatusInfo `json:"statusInfo"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *StatusNotification) UnmarshalJSON(data []byte) error {

	type Alias StatusNotification // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertStatusNotificationRequired checks if the required fields are not zero-ed
func AssertStatusNotificationRequired(obj StatusNotification) error {
	elements := map[string]interface{}{
		"statusInfo": obj.StatusInfo,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertStatusInfoRequired(obj.StatusInfo); err != nil {
		return err
	}
	return nil
}

// AssertStatusNotificationConstraints checks if the values respects the defined constraints
func AssertStatusNotificationConstraints(obj StatusNotification) error {
	return nil
}