/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SponsorConnectivityData - Contains the sponsored data connectivity related information for a sponsor identifier.
type SponsorConnectivityData struct {
	AspIds []string `json:"aspIds"`
}

// AssertSponsorConnectivityDataRequired checks if the required fields are not zero-ed
func AssertSponsorConnectivityDataRequired(obj SponsorConnectivityData) error {
	elements := map[string]interface{}{
		"aspIds": obj.AspIds,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertSponsorConnectivityDataConstraints checks if the values respects the defined constraints
func AssertSponsorConnectivityDataConstraints(obj SponsorConnectivityData) error {
	return nil
}