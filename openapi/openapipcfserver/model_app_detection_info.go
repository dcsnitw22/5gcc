/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_pcf_srv

type AppDetectionInfo struct {

	// A reference to the application detection filter configured at the UPF
	AppId string `json:"appId"`

	// Identifier sent by the SMF in order to allow correlation of application Start and Stop events to the specific service data flow description, if service data flow descriptions are deducible.
	InstanceId string `json:"instanceId,omitempty"`

	// Contains the detected service data flow descriptions if they are deducible.
	SdfDescriptions []FlowInformation `json:"sdfDescriptions,omitempty"`
}

// AssertAppDetectionInfoRequired checks if the required fields are not zero-ed
func AssertAppDetectionInfoRequired(obj AppDetectionInfo) error {
	elements := map[string]interface{}{
		"appId": obj.AppId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.SdfDescriptions {
		if err := AssertFlowInformationRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertAppDetectionInfoConstraints checks if the values respects the defined constraints
func AssertAppDetectionInfoConstraints(obj AppDetectionInfo) error {
	return nil
}
