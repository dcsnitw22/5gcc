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

type ExpectedUeBehaviour struct {
	AfInstanceId string `json:"afInstanceId"`

	ReferenceId int32 `json:"referenceId"`

	StationaryIndication *StationaryIndication `json:"stationaryIndication,omitempty"`

	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	CommunicationDurationTime *int32 `json:"communicationDurationTime,omitempty"`

	ScheduledCommunicationType *ScheduledCommunicationType `json:"scheduledCommunicationType,omitempty"`

	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	PeriodicTime *int32 `json:"periodicTime,omitempty"`

	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`

	// Identifies the UE's expected geographical movement. The attribute is only applicable in 5G.
	ExpectedUmts *[]LocationArea `json:"expectedUmts,omitempty"`

	TrafficProfile *TrafficProfile `json:"trafficProfile,omitempty"`

	BatteryIndication *BatteryIndication `json:"batteryIndication,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime time.Time `json:"validityTime,omitempty"`

	// String uniquely identifying MTC provider information.
	MtcProviderInformation string `json:"mtcProviderInformation,omitempty"`
}

// AssertExpectedUeBehaviourRequired checks if the required fields are not zero-ed
func AssertExpectedUeBehaviourRequired(obj ExpectedUeBehaviour) error {
	elements := map[string]interface{}{
		"afInstanceId": obj.AfInstanceId,
		"referenceId":  obj.ReferenceId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if obj.StationaryIndication != nil {
		if err := AssertStationaryIndicationRequired(*obj.StationaryIndication); err != nil {
			return err
		}
	}
	if obj.ScheduledCommunicationType != nil {
		if err := AssertScheduledCommunicationTypeRequired(*obj.ScheduledCommunicationType); err != nil {
			return err
		}
	}
	if obj.ScheduledCommunicationTime != nil {
		if err := AssertScheduledCommunicationTimeRequired(*obj.ScheduledCommunicationTime); err != nil {
			return err
		}
	}
	if obj.ExpectedUmts != nil {
		for _, el := range *obj.ExpectedUmts {
			if err := AssertLocationAreaRequired(el); err != nil {
				return err
			}
		}
	}
	if obj.TrafficProfile != nil {
		if err := AssertTrafficProfileRequired(*obj.TrafficProfile); err != nil {
			return err
		}
	}
	if obj.BatteryIndication != nil {
		if err := AssertBatteryIndicationRequired(*obj.BatteryIndication); err != nil {
			return err
		}
	}
	return nil
}

// AssertExpectedUeBehaviourConstraints checks if the values respects the defined constraints
func AssertExpectedUeBehaviourConstraints(obj ExpectedUeBehaviour) error {
	return nil
}
