/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_pcf_srv

import (
	"time"
)

type SmPolicyDeleteData struct {
	UserLocationInfo UserLocation `json:"userLocationInfo,omitempty"`

	UeTimeZone string `json:"ueTimeZone,omitempty"`

	ServingNetwork NetworkId `json:"servingNetwork,omitempty"`

	UserLocationInfoTime time.Time `json:"userLocationInfoTime,omitempty"`

	// Contains the RAN and/or NAS release cause.
	RanNasRelCauses []RanNasRelCause `json:"ranNasRelCauses,omitempty"`

	// Contains the usage report
	AccuUsageReports []AccuUsageReport `json:"accuUsageReports,omitempty"`
}

// AssertSmPolicyDeleteDataRequired checks if the required fields are not zero-ed
func AssertSmPolicyDeleteDataRequired(obj SmPolicyDeleteData) error {
	if err := AssertUserLocationRequired(obj.UserLocationInfo); err != nil {
		return err
	}
	if err := AssertNetworkIdRequired(obj.ServingNetwork); err != nil {
		return err
	}
	for _, el := range obj.RanNasRelCauses {
		if err := AssertRanNasRelCauseRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.AccuUsageReports {
		if err := AssertAccuUsageReportRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertSmPolicyDeleteDataConstraints checks if the values respects the defined constraints
func AssertSmPolicyDeleteDataConstraints(obj SmPolicyDeleteData) error {
	return nil
}