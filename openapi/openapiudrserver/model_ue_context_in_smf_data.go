/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UeContextInSmfData struct {

	// A map (list of key-value pairs where PduSessionId serves as key) of PduSessions
	PduSessions map[string]PduSession `json:"pduSessions,omitempty"`

	PgwInfo []PgwInfo `json:"pgwInfo,omitempty"`

	EmergencyInfo *EmergencyInfo `json:"emergencyInfo,omitempty"`
}

// AssertUeContextInSmfDataRequired checks if the required fields are not zero-ed
func AssertUeContextInSmfDataRequired(obj UeContextInSmfData) error {
	for _, el := range obj.PgwInfo {
		if err := AssertPgwInfoRequired(el); err != nil {
			return err
		}
	}
	if obj.EmergencyInfo != nil {
		if err := AssertEmergencyInfoRequired(*obj.EmergencyInfo); err != nil {
			return err
		}
	}
	return nil
}

// AssertUeContextInSmfDataConstraints checks if the values respects the defined constraints
func AssertUeContextInSmfDataConstraints(obj UeContextInSmfData) error {
	return nil
}
