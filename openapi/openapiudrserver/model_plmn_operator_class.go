/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PlmnOperatorClass struct {
	LcsClientClass LcsClientClass `json:"lcsClientClass"`

	LcsClientIds []string `json:"lcsClientIds"`
}

// AssertPlmnOperatorClassRequired checks if the required fields are not zero-ed
func AssertPlmnOperatorClassRequired(obj PlmnOperatorClass) error {
	elements := map[string]interface{}{
		"lcsClientClass": obj.LcsClientClass,
		"lcsClientIds":   obj.LcsClientIds,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertLcsClientClassRequired(obj.LcsClientClass); err != nil {
		return err
	}
	return nil
}

// AssertPlmnOperatorClassConstraints checks if the values respects the defined constraints
func AssertPlmnOperatorClassConstraints(obj PlmnOperatorClass) error {
	return nil
}
