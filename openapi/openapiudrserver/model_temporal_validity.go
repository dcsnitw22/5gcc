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
	"time"
)

// TemporalValidity - Indicates the time interval(s) during which the AF request is to be applied
type TemporalValidity struct {

	// string with format 'date-time' as defined in OpenAPI.
	StartTime time.Time `json:"startTime,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	StopTime time.Time `json:"stopTime,omitempty"`
}

// AssertTemporalValidityRequired checks if the required fields are not zero-ed
func AssertTemporalValidityRequired(obj TemporalValidity) error {
	return nil
}

// AssertTemporalValidityConstraints checks if the values respects the defined constraints
func AssertTemporalValidityConstraints(obj TemporalValidity) error {
	return nil
}