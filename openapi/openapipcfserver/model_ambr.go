/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_pcf_srv

type Ambr struct {
	Uplink string `json:"uplink"`

	Downlink string `json:"downlink"`
}

// AssertAmbrRequired checks if the required fields are not zero-ed
func AssertAmbrRequired(obj Ambr) error {
	elements := map[string]interface{}{
		"uplink":   obj.Uplink,
		"downlink": obj.Downlink,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertAmbrConstraints checks if the values respects the defined constraints
func AssertAmbrConstraints(obj Ambr) error {
	return nil
}