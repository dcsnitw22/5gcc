/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Polygon - Polygon.
type Polygon struct {
	GadShape

	// List of points.
	PointList []GeographicalCoordinates `json:"pointList"`
}

// AssertPolygonRequired checks if the required fields are not zero-ed
func AssertPolygonRequired(obj Polygon) error {
	elements := map[string]interface{}{
		"pointList": obj.PointList,
		"shape":     obj.Shape,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertGadShapeRequired(obj.GadShape); err != nil {
		return err
	}

	for _, el := range obj.PointList {
		if err := AssertGeographicalCoordinatesRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertPolygonConstraints checks if the values respects the defined constraints
func AssertPolygonConstraints(obj Polygon) error {
	return nil
}
