/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_pcf_srv

type QosNotificationControlInfo struct {

	// An array of PCC rule id references to the PCC rules associated with the QoS notification control info.
	RefPccRuleIds []string `json:"refPccRuleIds"`

	NotifType QosNotifType `json:"notifType"`

	// Represents the content version of some content.
	ContVer int32 `json:"contVer,omitempty"`
}

// AssertQosNotificationControlInfoRequired checks if the required fields are not zero-ed
func AssertQosNotificationControlInfoRequired(obj QosNotificationControlInfo) error {
	elements := map[string]interface{}{
		"refPccRuleIds": obj.RefPccRuleIds,
		"notifType":     obj.NotifType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertQosNotifTypeRequired(obj.NotifType); err != nil {
		return err
	}
	return nil
}

// AssertQosNotificationControlInfoConstraints checks if the values respects the defined constraints
func AssertQosNotificationControlInfoConstraints(obj QosNotificationControlInfo) error {
	return nil
}