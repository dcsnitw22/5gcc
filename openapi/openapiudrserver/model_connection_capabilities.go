/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ConnectionCapabilities - Possible values are   - IMS: Indicates the connection capability to support IMS service.   - MMS: Indicates the connection capability to support MMS service.   - SUPL: Indicates the connection capability to support SUPL service.   - INTERNET: Indicates the connection capability to support Internet service.
type ConnectionCapabilities struct {
}

// AssertConnectionCapabilitiesRequired checks if the required fields are not zero-ed
func AssertConnectionCapabilitiesRequired(obj ConnectionCapabilities) error {
	return nil
}

// AssertConnectionCapabilitiesConstraints checks if the values respects the defined constraints
func AssertConnectionCapabilitiesConstraints(obj ConnectionCapabilities) error {
	return nil
}