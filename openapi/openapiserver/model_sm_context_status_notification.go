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

type SmContextStatusNotification struct {
	StatusInfo StatusInfo `json:"statusInfo"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *SmContextStatusNotification) UnmarshalJSON(data []byte) error {

	type Alias SmContextStatusNotification // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertSmContextStatusNotificationRequired checks if the required fields are not zero-ed
func AssertSmContextStatusNotificationRequired(obj SmContextStatusNotification) error {
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

// AssertSmContextStatusNotificationConstraints checks if the values respects the defined constraints
func AssertSmContextStatusNotificationConstraints(obj SmContextStatusNotification) error {
	return nil
}
