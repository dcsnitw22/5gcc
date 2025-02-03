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

type RoamingChargingProfile struct {
	Triggers []Trigger `json:"triggers,omitempty"`

	PartialRecordMethod PartialRecordMethod `json:"partialRecordMethod,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *RoamingChargingProfile) UnmarshalJSON(data []byte) error {

	type Alias RoamingChargingProfile // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertRoamingChargingProfileRequired checks if the required fields are not zero-ed
func AssertRoamingChargingProfileRequired(obj RoamingChargingProfile) error {
	for _, el := range obj.Triggers {
		if err := AssertTriggerRequired(el); err != nil {
			return err
		}
	}
	if err := AssertPartialRecordMethodRequired(obj.PartialRecordMethod); err != nil {
		return err
	}
	return nil
}

// AssertRoamingChargingProfileConstraints checks if the values respects the defined constraints
func AssertRoamingChargingProfileConstraints(obj RoamingChargingProfile) error {
	return nil
}
