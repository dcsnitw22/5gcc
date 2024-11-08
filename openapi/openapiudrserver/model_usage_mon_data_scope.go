/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UsageMonDataScope - Contains a SNSSAI and DNN combinations to which the UsageMonData instance belongs to.
type UsageMonDataScope struct {
	Snssai Snssai `json:"snssai"`

	Dnn []string `json:"dnn,omitempty"`
}

// AssertUsageMonDataScopeRequired checks if the required fields are not zero-ed
func AssertUsageMonDataScopeRequired(obj UsageMonDataScope) error {
	elements := map[string]interface{}{
		"snssai": obj.Snssai,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertSnssaiRequired(obj.Snssai); err != nil {
		return err
	}
	return nil
}

// AssertUsageMonDataScopeConstraints checks if the values respects the defined constraints
func AssertUsageMonDataScopeConstraints(obj UsageMonDataScope) error {
	return nil
}
