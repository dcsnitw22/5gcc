/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type EpsInterworkingInfo struct {

	// A map (list of key-value pairs where Dnn serves as key) of EpsIwkPgws
	EpsIwkPgws map[string]EpsIwkPgw `json:"epsIwkPgws,omitempty"`
}

// AssertEpsInterworkingInfoRequired checks if the required fields are not zero-ed
func AssertEpsInterworkingInfoRequired(obj EpsInterworkingInfo) error {
	return nil
}

// AssertEpsInterworkingInfoConstraints checks if the values respects the defined constraints
func AssertEpsInterworkingInfoConstraints(obj EpsInterworkingInfo) error {
	return nil
}