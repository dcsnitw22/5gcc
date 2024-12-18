/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PolicyDataForIndividualUe - Contains policy data for a given subscriber.
type PolicyDataForIndividualUe struct {
	UePolicyDataSet UePolicySet `json:"uePolicyDataSet,omitempty"`

	SmPolicyDataSet SmPolicyData `json:"smPolicyDataSet,omitempty"`

	AmPolicyDataSet AmPolicyData `json:"amPolicyDataSet,omitempty"`

	// Contains UM policies. The value of the limit identifier is used as the key of the map.
	UmData map[string]UsageMonData `json:"umData,omitempty"`

	// Contains Operator Specific Data resource data. The key of the map is operator specific data element name and the value is the operator specific data of the UE.
	OperatorSpecificDataSet map[string]OperatorSpecificDataContainer `json:"operatorSpecificDataSet,omitempty"`
}

// AssertPolicyDataForIndividualUeRequired checks if the required fields are not zero-ed
func AssertPolicyDataForIndividualUeRequired(obj PolicyDataForIndividualUe) error {
	if err := AssertUePolicySetRequired(obj.UePolicyDataSet); err != nil {
		return err
	}
	if err := AssertSmPolicyDataRequired(obj.SmPolicyDataSet); err != nil {
		return err
	}
	if err := AssertAmPolicyDataRequired(obj.AmPolicyDataSet); err != nil {
		return err
	}
	return nil
}

// AssertPolicyDataForIndividualUeConstraints checks if the values respects the defined constraints
func AssertPolicyDataForIndividualUeConstraints(obj PolicyDataForIndividualUe) error {
	return nil
}
