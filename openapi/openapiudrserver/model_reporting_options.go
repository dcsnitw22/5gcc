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
	"time"
)

type ReportingOptions struct {
	ReportMode EventReportMode `json:"reportMode,omitempty"`

	MaxNumOfReports int32 `json:"maxNumOfReports,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	Expiry time.Time `json:"expiry,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	SamplingRatio int32 `json:"samplingRatio,omitempty"`

	// indicating a time in seconds.
	GuardTime int32 `json:"guardTime,omitempty"`

	// indicating a time in seconds.
	ReportPeriod int32 `json:"reportPeriod,omitempty"`

	NotifFlag NotificationFlag `json:"notifFlag,omitempty"`
}

// AssertReportingOptionsRequired checks if the required fields are not zero-ed
func AssertReportingOptionsRequired(obj ReportingOptions) error {
	if err := AssertEventReportModeRequired(obj.ReportMode); err != nil {
		return err
	}
	if err := AssertNotificationFlagRequired(obj.NotifFlag); err != nil {
		return err
	}
	return nil
}

// AssertReportingOptionsConstraints checks if the values respects the defined constraints
func AssertReportingOptionsConstraints(obj ReportingOptions) error {
	if obj.SamplingRatio < 1 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.SamplingRatio > 100 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}