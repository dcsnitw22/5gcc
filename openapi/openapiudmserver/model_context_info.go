/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type ContextInfo struct {

	OrigHeaders []string `json:"origHeaders,omitempty"`

	RequestHeaders []string `json:"requestHeaders,omitempty"`
}

// AssertContextInfoRequired checks if the required fields are not zero-ed
func AssertContextInfoRequired(obj ContextInfo) error {
	return nil
}

// AssertContextInfoConstraints checks if the values respects the defined constraints
func AssertContextInfoConstraints(obj ContextInfo) error {
	return nil
}