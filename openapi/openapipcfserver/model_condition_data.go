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

type ConditionData struct {

	// Uniquely identifies the condition data within a PDU session.
	CondId string `json:"condId"`

	ActivationTime *time.Time `json:"activationTime,omitempty"`

	DeactivationTime *time.Time `json:"deactivationTime,omitempty"`
}

// AssertConditionDataRequired checks if the required fields are not zero-ed
func AssertConditionDataRequired(obj ConditionData) error {
	elements := map[string]interface{}{
		"condId": obj.CondId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertConditionDataConstraints checks if the values respects the defined constraints
func AssertConditionDataConstraints(obj ConditionData) error {
	return nil
}
