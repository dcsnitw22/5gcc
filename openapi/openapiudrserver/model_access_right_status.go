/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// AccessRightStatus - Possible values are: - FULLY_ALLOWED: The User is fully allowed to access to the channel. - PREVIEW_ALLOWED: The User is preview allowed to access to the channel. - NO_ALLOWED: The User is not allowed to access to the channel.
type AccessRightStatus struct {
}

// AssertAccessRightStatusRequired checks if the required fields are not zero-ed
func AssertAccessRightStatusRequired(obj AccessRightStatus) error {
	return nil
}

// AssertAccessRightStatusConstraints checks if the values respects the defined constraints
func AssertAccessRightStatusConstraints(obj AccessRightStatus) error {
	return nil
}
