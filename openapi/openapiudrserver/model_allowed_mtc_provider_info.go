/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type AllowedMtcProviderInfo struct {

	// String uniquely identifying MTC provider information.
	MtcProviderInformation string `json:"mtcProviderInformation,omitempty"`

	AfId string `json:"afId,omitempty"`
}

// AssertAllowedMtcProviderInfoRequired checks if the required fields are not zero-ed
func AssertAllowedMtcProviderInfoRequired(obj AllowedMtcProviderInfo) error {
	return nil
}

// AssertAllowedMtcProviderInfoConstraints checks if the values respects the defined constraints
func AssertAllowedMtcProviderInfoConstraints(obj AllowedMtcProviderInfo) error {
	return nil
}
