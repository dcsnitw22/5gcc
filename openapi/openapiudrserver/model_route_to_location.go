/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// RouteToLocation - At least one of the \"routeInfo\" attribute and the \"routeProfId\" attribute shall be included in the \"RouteToLocation\" data type.
type RouteToLocation struct {

	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai string `json:"dnai"`

	RouteInfo *RouteInformation `json:"routeInfo,omitempty"`

	// Identifies the routing profile Id.
	RouteProfId *string `json:"routeProfId,omitempty"`
}

// AssertRouteToLocationRequired checks if the required fields are not zero-ed
func AssertRouteToLocationRequired(obj RouteToLocation) error {
	elements := map[string]interface{}{
		"dnai": obj.Dnai,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if obj.RouteInfo != nil {
		if err := AssertRouteInformationRequired(*obj.RouteInfo); err != nil {
			return err
		}
	}
	return nil
}

// AssertRouteToLocationConstraints checks if the values respects the defined constraints
func AssertRouteToLocationConstraints(obj RouteToLocation) error {
	return nil
}