/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// DnaiChangeType - Possible values are: - EARLY: Early notification of UP path reconfiguration. - EARLY_LATE: Early and late notification of UP path reconfiguration. This value shall   only be present in the subscription to the DNAI change event. - LATE: Late notification of UP path reconfiguration.
type DnaiChangeType struct {
}

// AssertDnaiChangeTypeRequired checks if the required fields are not zero-ed
func AssertDnaiChangeTypeRequired(obj DnaiChangeType) error {
	return nil
}

// AssertDnaiChangeTypeConstraints checks if the values respects the defined constraints
func AssertDnaiChangeTypeConstraints(obj DnaiChangeType) error {
	return nil
}
