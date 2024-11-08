/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type EdrxParameters1 struct {
	RatType RatType `json:"ratType"`

	EdrxValue string `json:"edrxValue"`
}

// AssertEdrxParameters1Required checks if the required fields are not zero-ed
func AssertEdrxParameters1Required(obj EdrxParameters1) error {
	elements := map[string]interface{}{
		"ratType":   obj.RatType,
		"edrxValue": obj.EdrxValue,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertRatTypeRequired(obj.RatType); err != nil {
		return err
	}
	return nil
}

// AssertEdrxParameters1Constraints checks if the values respects the defined constraints
func AssertEdrxParameters1Constraints(obj EdrxParameters1) error {
	return nil
}
