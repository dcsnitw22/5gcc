/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_pcf_srv

// NetLocAccessSupport - Possible values are - ANR_NOT_SUPPORTED: Indicates that the access network does not support the report of access network information. - TZR_NOT_SUPPORTED: Indicates that the access network does not support the report of UE time zone. - LOC_NOT_SUPPORTED: Indicates that the access network does not support the report of UE Location (or PLMN Id).
type NetLocAccessSupport struct {
}

// AssertNetLocAccessSupportRequired checks if the required fields are not zero-ed
func AssertNetLocAccessSupportRequired(obj NetLocAccessSupport) error {
	return nil
}

// AssertNetLocAccessSupportConstraints checks if the values respects the defined constraints
func AssertNetLocAccessSupportConstraints(obj NetLocAccessSupport) error {
	return nil
}