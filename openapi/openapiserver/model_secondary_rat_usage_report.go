/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiserver

import (
	"encoding/json"
)

type SecondaryRatUsageReport struct {
	SecondaryRatType RatType `json:"secondaryRatType"`

	QosFlowsUsageData []QosFlowUsageReport `json:"qosFlowsUsageData"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *SecondaryRatUsageReport) UnmarshalJSON(data []byte) error {

	type Alias SecondaryRatUsageReport // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertSecondaryRatUsageReportRequired checks if the required fields are not zero-ed
func AssertSecondaryRatUsageReportRequired(obj SecondaryRatUsageReport) error {
	elements := map[string]interface{}{
		"secondaryRatType":  obj.SecondaryRatType,
		"qosFlowsUsageData": obj.QosFlowsUsageData,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertRatTypeRequired(obj.SecondaryRatType); err != nil {
		return err
	}
	for _, el := range obj.QosFlowsUsageData {
		if err := AssertQosFlowUsageReportRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertSecondaryRatUsageReportConstraints checks if the values respects the defined constraints
func AssertSecondaryRatUsageReportConstraints(obj SecondaryRatUsageReport) error {
	return nil
}
