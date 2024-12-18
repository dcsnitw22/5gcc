/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ApplicationDataChangeNotif - Contains changed application data for which notification was requested.
type ApplicationDataChangeNotif struct {
	IptvConfigData *IptvConfigData `json:"iptvConfigData,omitempty"`

	PfdData PfdChangeNotification `json:"pfdData,omitempty"`

	BdtPolicyData BdtPolicyData `json:"bdtPolicyData,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	ResUri string `json:"resUri"`

	SerParamData ServiceParameterData `json:"serParamData,omitempty"`

	AmInfluData AmInfluData `json:"amInfluData,omitempty"`
}

// AssertApplicationDataChangeNotifRequired checks if the required fields are not zero-ed
func AssertApplicationDataChangeNotifRequired(obj ApplicationDataChangeNotif) error {
	elements := map[string]interface{}{
		"resUri": obj.ResUri,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if obj.IptvConfigData != nil {
		if err := AssertIptvConfigDataRequired(*obj.IptvConfigData); err != nil {
			return err
		}
	}
	if err := AssertPfdChangeNotificationRequired(obj.PfdData); err != nil {
		return err
	}
	if err := AssertBdtPolicyDataRequired(obj.BdtPolicyData); err != nil {
		return err
	}
	if err := AssertServiceParameterDataRequired(obj.SerParamData); err != nil {
		return err
	}
	if err := AssertAmInfluDataRequired(obj.AmInfluData); err != nil {
		return err
	}
	return nil
}

// AssertApplicationDataChangeNotifConstraints checks if the values respects the defined constraints
func AssertApplicationDataChangeNotifConstraints(obj ApplicationDataChangeNotif) error {
	return nil
}
