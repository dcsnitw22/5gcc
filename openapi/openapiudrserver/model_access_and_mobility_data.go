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

// AccessAndMobilityData - Represents Access and Mobility data for a UE.
type AccessAndMobilityData struct {
	Location UserLocation `json:"location,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	LocationTs time.Time `json:"locationTs,omitempty"`

	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.
	TimeZone string `json:"timeZone,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeZoneTs time.Time `json:"timeZoneTs,omitempty"`

	AccessType AccessType `json:"accessType,omitempty"`

	RegStates []RmInfo `json:"regStates,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	RegStatesTs time.Time `json:"regStatesTs,omitempty"`

	ConnStates []CmInfo `json:"connStates,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ConnStatesTs time.Time `json:"connStatesTs,omitempty"`

	ReachabilityStatus UeReachability `json:"reachabilityStatus,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ReachabilityStatusTs time.Time `json:"reachabilityStatusTs,omitempty"`

	SmsOverNasStatus SmsSupport `json:"smsOverNasStatus,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	SmsOverNasStatusTs time.Time `json:"smsOverNasStatusTs,omitempty"`

	// True  The serving PLMN of the UE is different from the HPLMN of the UE; False The serving PLMN of the UE is the HPLMN of the UE.
	RoamingStatus bool `json:"roamingStatus,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	RoamingStatusTs time.Time `json:"roamingStatusTs,omitempty"`

	CurrentPlmn PlmnId1 `json:"currentPlmn,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	CurrentPlmnTs time.Time `json:"currentPlmnTs,omitempty"`

	RatType []RatType `json:"ratType,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	RatTypesTs time.Time `json:"ratTypesTs,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat string `json:"suppFeat,omitempty"`

	ResetIds []string `json:"resetIds,omitempty"`
}

// AssertAccessAndMobilityDataRequired checks if the required fields are not zero-ed
func AssertAccessAndMobilityDataRequired(obj AccessAndMobilityData) error {
	if err := AssertUserLocationRequired(obj.Location); err != nil {
		return err
	}
	for _, el := range obj.RegStates {
		if err := AssertRmInfoRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ConnStates {
		if err := AssertCmInfoRequired(el); err != nil {
			return err
		}
	}
	if err := AssertUeReachabilityRequired(obj.ReachabilityStatus); err != nil {
		return err
	}
	if err := AssertSmsSupportRequired(obj.SmsOverNasStatus); err != nil {
		return err
	}
	if err := AssertPlmnId1Required(obj.CurrentPlmn); err != nil {
		return err
	}
	for _, el := range obj.RatType {
		if err := AssertRatTypeRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertAccessAndMobilityDataConstraints checks if the values respects the defined constraints
func AssertAccessAndMobilityDataConstraints(obj AccessAndMobilityData) error {
	return nil
}
