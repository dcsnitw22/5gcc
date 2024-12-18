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
	"errors"
	"time"
)

type EutraLocation struct {
	Tai Tai `json:"tai"`

	Ecgi Ecgi `json:"ecgi"`

	AgeOfLocationInformation int32 `json:"ageOfLocationInformation,omitempty"`

	UeLocationTimestamp time.Time `json:"ueLocationTimestamp,omitempty"`

	GeographicalInformation string `json:"geographicalInformation,omitempty"`

	GeodeticInformation string `json:"geodeticInformation,omitempty"`

	GlobalNgenbId *GlobalRanNodeId `json:"globalNgenbId,omitempty"`
}

// AssertEutraLocationRequired checks if the required fields are not zero-ed
func AssertEutraLocationRequired(obj EutraLocation) error {
	elements := map[string]interface{}{
		"tai":  obj.Tai,
		"ecgi": obj.Ecgi,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertTaiRequired(obj.Tai); err != nil {
		return err
	}
	if err := AssertEcgiRequired(obj.Ecgi); err != nil {
		return err
	}
	if obj.GlobalNgenbId != nil {
		if err := AssertGlobalRanNodeIdRequired(*obj.GlobalNgenbId); err != nil {
			return err
		}
	}
	return nil
}

// AssertEutraLocationConstraints checks if the values respects the defined constraints
func AssertEutraLocationConstraints(obj EutraLocation) error {
	if obj.AgeOfLocationInformation < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.AgeOfLocationInformation > 32767 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
