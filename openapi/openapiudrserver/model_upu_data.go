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

// UpuData - Used to store the status of the latest UPU data update.
type UpuData struct {

	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime time.Time `json:"provisioningTime"`

	UeUpdateStatus UeUpdateStatus `json:"ueUpdateStatus"`

	// MAC value for protecting UPU procedure (UPU-MAC-IAUSF and UPU-MAC-IUE).
	UpuXmacIue string `json:"upuXmacIue,omitempty"`

	// MAC value for protecting UPU procedure (UPU-MAC-IAUSF and UPU-MAC-IUE).
	UpuMacIue string `json:"upuMacIue,omitempty"`
}

// AssertUpuDataRequired checks if the required fields are not zero-ed
func AssertUpuDataRequired(obj UpuData) error {
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

// AssertUpuDataConstraints checks if the values respects the defined constraints
func AssertUpuDataConstraints(obj UpuData) error {
	return nil
}