/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// AmPolicyData - Contains the AM policy data for a given subscriber.
type AmPolicyData struct {

	// Contains Presence reporting area information. The praId attribute within the PresenceInfo data type is the key of the map.
	PraInfos map[string]PresenceInfo `json:"praInfos,omitempty"`

	SubscCats []string `json:"subscCats,omitempty"`
}

// AssertAmPolicyDataRequired checks if the required fields are not zero-ed
func AssertAmPolicyDataRequired(obj AmPolicyData) error {
	return nil
}

// AssertAmPolicyDataConstraints checks if the values respects the defined constraints
func AssertAmPolicyDataConstraints(obj AmPolicyData) error {
	return nil
}