/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_pcf_srv

// RequestedRuleDataType - Possible values are - CH_ID: Indicates that the requested rule data is the charging identifier.  - MS_TIME_ZONE: Indicates that the requested access network info type is the UE's timezone. - USER_LOC_INFO: Indicates that the requested access network info type is the UE's location. - RES_RELEASE: Indicates that the requested rule data is the result of the release of resource. - SUCC_RES_ALLO: Indicates that the requested rule data is the successful resource allocation.
type RequestedRuleDataType struct {
}

// AssertRequestedRuleDataTypeRequired checks if the required fields are not zero-ed
func AssertRequestedRuleDataTypeRequired(obj RequestedRuleDataType) error {
	return nil
}

// AssertRequestedRuleDataTypeConstraints checks if the values respects the defined constraints
func AssertRequestedRuleDataTypeConstraints(obj RequestedRuleDataType) error {
	return nil
}
