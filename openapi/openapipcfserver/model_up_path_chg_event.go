/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_pcf_srv

type UpPathChgEvent struct {
	NotificationUri string `json:"notificationUri"`

	// It is used to set the value of Notification Correlation ID in the notification sent by the SMF.
	NotifCorreId string `json:"notifCorreId"`

	DnaiChgType DnaiChangeType `json:"dnaiChgType"`
}

// AssertUpPathChgEventRequired checks if the required fields are not zero-ed
func AssertUpPathChgEventRequired(obj UpPathChgEvent) error {
	elements := map[string]interface{}{
		"notificationUri": obj.NotificationUri,
		"notifCorreId":    obj.NotifCorreId,
		"dnaiChgType":     obj.DnaiChgType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertDnaiChangeTypeRequired(obj.DnaiChgType); err != nil {
		return err
	}
	return nil
}

// AssertUpPathChgEventConstraints checks if the values respects the defined constraints
func AssertUpPathChgEventConstraints(obj UpPathChgEvent) error {
	return nil
}