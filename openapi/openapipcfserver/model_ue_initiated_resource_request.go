/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_pcf_srv

type UeInitiatedResourceRequest struct {
	PccRuleId string `json:"pccRuleId,omitempty"`

	RuleOp RuleOperation `json:"ruleOp"`

	Precedence int32 `json:"precedence,omitempty"`

	PackFiltInfo []PacketFilterInfo `json:"packFiltInfo"`

	ReqQos RequestedQos `json:"reqQos,omitempty"`
}

// AssertUeInitiatedResourceRequestRequired checks if the required fields are not zero-ed
func AssertUeInitiatedResourceRequestRequired(obj UeInitiatedResourceRequest) error {
	elements := map[string]interface{}{
		"ruleOp":       obj.RuleOp,
		"packFiltInfo": obj.PackFiltInfo,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertRuleOperationRequired(obj.RuleOp); err != nil {
		return err
	}
	for _, el := range obj.PackFiltInfo {
		if err := AssertPacketFilterInfoRequired(el); err != nil {
			return err
		}
	}
	if err := AssertRequestedQosRequired(obj.ReqQos); err != nil {
		return err
	}
	return nil
}

// AssertUeInitiatedResourceRequestConstraints checks if the values respects the defined constraints
func AssertUeInitiatedResourceRequestConstraints(obj UeInitiatedResourceRequest) error {
	return nil
}