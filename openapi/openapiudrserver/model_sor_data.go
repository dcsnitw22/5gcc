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

// SorData - Used to store the status of the latest SOR data update.
type SorData struct {

	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime time.Time `json:"provisioningTime"`

	UeUpdateStatus UeUpdateStatus `json:"ueUpdateStatus"`

	// MAC value for protecting SOR procedure (SoR-MAC-IAUSF and SoR-XMAC-IUE).
	SorXmacIue string `json:"sorXmacIue,omitempty"`

	// MAC value for protecting SOR procedure (SoR-MAC-IAUSF and SoR-XMAC-IUE).
	SorMacIue string `json:"sorMacIue,omitempty"`

	MeSupportOfSorCmci bool `json:"meSupportOfSorCmci,omitempty"`
}

// AssertSorDataRequired checks if the required fields are not zero-ed
func AssertSorDataRequired(obj SorData) error {
	elements := map[string]interface{}{
		"provisioningTime": obj.ProvisioningTime,
		"ueUpdateStatus":   obj.UeUpdateStatus,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertSorDataConstraints checks if the values respects the defined constraints
func AssertSorDataConstraints(obj SorData) error {
	return nil
}