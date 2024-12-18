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

type SorInfo1 struct {
	SteeringContainer SteeringContainer `json:"steeringContainer,omitempty"`

	// Contains indication whether the acknowledgement from UE is needed.
	AckInd bool `json:"ackInd"`

	// MAC value for protecting SOR procedure (SoR-MAC-IAUSF and SoR-XMAC-IUE).
	SorMacIausf string `json:"sorMacIausf,omitempty"`

	// CounterSoR.
	Countersor string `json:"countersor,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime time.Time `json:"provisioningTime"`

	// string with format 'bytes' as defined in OpenAPI
	SorTransparentContainer string `json:"sorTransparentContainer,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	SorCmci string `json:"sorCmci,omitempty"`

	StoreSorCmciInMe bool `json:"storeSorCmciInMe,omitempty"`

	UsimSupportOfSorCmci bool `json:"usimSupportOfSorCmci,omitempty"`
}

// AssertSorInfo1Required checks if the required fields are not zero-ed
func AssertSorInfo1Required(obj SorInfo1) error {
	elements := map[string]interface{}{
		"ackInd":           obj.AckInd,
		"provisioningTime": obj.ProvisioningTime,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertSteeringContainerRequired(obj.SteeringContainer); err != nil {
		return err
	}
	return nil
}

// AssertSorInfo1Constraints checks if the values respects the defined constraints
func AssertSorInfo1Constraints(obj SorInfo1) error {
	return nil
}
