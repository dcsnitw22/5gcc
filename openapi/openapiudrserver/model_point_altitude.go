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

// PointAltitude - Ellipsoid point with altitude.
type PointAltitude struct {
	GadShape

	Point GeographicalCoordinates `json:"point"`

	// Indicates value of altitude.
	Altitude float64 `json:"altitude"`
}

// AssertPointAltitudeRequired checks if the required fields are not zero-ed
func AssertPointAltitudeRequired(obj PointAltitude) error {
	elements := map[string]interface{}{
		"point":    obj.Point,
		"altitude": obj.Altitude,
		"shape":    obj.Shape,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertGadShapeRequired(obj.GadShape); err != nil {
		return err
	}

	if err := AssertGeographicalCoordinatesRequired(obj.Point); err != nil {
		return err
	}
	return nil
}

// AssertPointAltitudeConstraints checks if the values respects the defined constraints
func AssertPointAltitudeConstraints(obj PointAltitude) error {
	if obj.Altitude < -32767 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Altitude > 32767 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
