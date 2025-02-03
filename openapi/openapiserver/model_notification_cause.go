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

// NotificationCause - Possible values are - QOS_FULFILLED - QOS_NOT_FULFILLED - UP_SEC_FULFILLED - UP_SEC_NOT_FULFILLED
type NotificationCause struct {
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *NotificationCause) UnmarshalJSON(data []byte) error {

	type Alias NotificationCause // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertNotificationCauseRequired checks if the required fields are not zero-ed
func AssertNotificationCauseRequired(obj NotificationCause) error {
	return nil
}

// AssertNotificationCauseConstraints checks if the values respects the defined constraints
func AssertNotificationCauseConstraints(obj NotificationCause) error {
	return nil
}
