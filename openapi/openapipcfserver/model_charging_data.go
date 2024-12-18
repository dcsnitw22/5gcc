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
)

type ChargingData struct {

	// Univocally identifies the charging control policy data within a PDU session.
	ChgId string `json:"chgId"`

	MeteringMethod *MeteringMethod `json:"meteringMethod,omitempty"`

	// Indicates the offline charging is applicable to the PCC rule when it is included and set to true.
	Offline bool `json:"offline,omitempty"`

	// Indicates the online charging is applicable to the PCC rule when it is included and set to true.
	Online bool `json:"online,omitempty"`

	// Indicates whether the service data flow is allowed to start while the SMF is waiting for the response to the credit request.
	SdfHandl bool `json:"sdfHandl,omitempty"`

	RatingGroup int32 `json:"ratingGroup,omitempty"`

	ReportingLevel *ReportingLevel `json:"reportingLevel,omitempty"`

	ServiceId int32 `json:"serviceId,omitempty"`

	// Indicates the sponsor identity.
	SponsorId string `json:"sponsorId,omitempty"`

	// Indicates the application service provider identity.
	AppSvcProvId string `json:"appSvcProvId,omitempty"`

	AfChargingIdentifier int32 `json:"afChargingIdentifier,omitempty"`

	AfChargId string `json:"afChargId,omitempty"`
}

// AssertChargingDataRequired checks if the required fields are not zero-ed
func AssertChargingDataRequired(obj ChargingData) error {
	elements := map[string]interface{}{
		"chgId": obj.ChgId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if obj.MeteringMethod != nil {
		if err := AssertMeteringMethodRequired(*obj.MeteringMethod); err != nil {
			return err
		}
	}
	if obj.ReportingLevel != nil {
		if err := AssertReportingLevelRequired(*obj.ReportingLevel); err != nil {
			return err
		}
	}
	return nil
}

// AssertChargingDataConstraints checks if the values respects the defined constraints
func AssertChargingDataConstraints(obj ChargingData) error {
	if obj.RatingGroup < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.ServiceId < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.AfChargingIdentifier < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
