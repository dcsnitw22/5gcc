/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"errors"
)

// UrspRuleRequest - Contains parameters that can be used to guide the URSP.
type UrspRuleRequest struct {
	TrafficDesc *TrafficDescriptorComponents `json:"trafficDesc,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	RelatPrecedence int32 `json:"relatPrecedence,omitempty"`

	// Sets of parameters that may be used to guide the Route Selection Descriptors of the  URSP.
	RouteSelParamSets []RouteSelectionParameterSet `json:"routeSelParamSets,omitempty"`
}

// AssertUrspRuleRequestRequired checks if the required fields are not zero-ed
func AssertUrspRuleRequestRequired(obj UrspRuleRequest) error {
	if obj.TrafficDesc != nil {
		if err := AssertTrafficDescriptorComponentsRequired(*obj.TrafficDesc); err != nil {
			return err
		}
	}
	for _, el := range obj.RouteSelParamSets {
		if err := AssertRouteSelectionParameterSetRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertUrspRuleRequestConstraints checks if the values respects the defined constraints
func AssertUrspRuleRequestConstraints(obj UrspRuleRequest) error {
	if obj.RelatPrecedence < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}