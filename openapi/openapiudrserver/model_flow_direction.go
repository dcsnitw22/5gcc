/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// FlowDirection - Possible values are - DOWNLINK: The corresponding filter applies for traffic to the UE. - UPLINK: The corresponding filter applies for traffic from the UE. - BIDIRECTIONAL: The corresponding filter applies for traffic both to and from the UE. - UNSPECIFIED: The corresponding filter applies for traffic to the UE (downlink), but has no specific direction declared. The service data flow detection shall apply the filter for uplink traffic as if the filter was bidirectional. The PCF shall not use the value UNSPECIFIED in filters created by the network in NW-initiated procedures. The PCF shall only include the value UNSPECIFIED in filters in UE-initiated procedures if the same value is received from the SMF.
type FlowDirection struct {
}

// AssertFlowDirectionRequired checks if the required fields are not zero-ed
func AssertFlowDirectionRequired(obj FlowDirection) error {
	return nil
}

// AssertFlowDirectionConstraints checks if the values respects the defined constraints
func AssertFlowDirectionConstraints(obj FlowDirection) error {
	return nil
}