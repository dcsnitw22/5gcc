/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type SscModes struct {
	DefaultSscMode SscMode `json:"defaultSscMode"`

	AllowedSscModes []SscMode `json:"allowedSscModes,omitempty"`
}

// AssertSscModesRequired checks if the required fields are not zero-ed
func AssertSscModesRequired(obj SscModes) error {
	elements := map[string]interface{}{
		"defaultSscMode": obj.DefaultSscMode,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertSscModeRequired(obj.DefaultSscMode); err != nil {
		return err
	}
	for _, el := range obj.AllowedSscModes {
		if err := AssertSscModeRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertSscModesConstraints checks if the values respects the defined constraints
func AssertSscModesConstraints(obj SscModes) error {
	return nil
}
