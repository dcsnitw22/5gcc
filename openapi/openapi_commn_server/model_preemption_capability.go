/*
 * Namf_Communication
 *
 * AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_commn_server

import "encoding/json"

type PreemptionCapability struct {
	String *string
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *PreemptionCapability) UnmarshalJSON(data []byte) error {

	type Alias PreemptionCapability // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertPreemptionCapabilityRequired checks if the required fields are not zero-ed
func AssertPreemptionCapabilityRequired(obj PreemptionCapability) error {
	return nil
}

// AssertPreemptionCapabilityConstraints checks if the values respects the defined constraints
func AssertPreemptionCapabilityConstraints(obj PreemptionCapability) error {
	return nil
}
