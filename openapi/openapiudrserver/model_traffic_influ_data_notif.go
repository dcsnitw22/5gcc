/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TrafficInfluDataNotif - Represents traffic influence data for notification.
type TrafficInfluDataNotif struct {

	// String providing an URI formatted according to RFC 3986.
	ResUri string `json:"resUri"`

	TrafficInfluData TrafficInfluData `json:"trafficInfluData,omitempty"`
}

// AssertTrafficInfluDataNotifRequired checks if the required fields are not zero-ed
func AssertTrafficInfluDataNotifRequired(obj TrafficInfluDataNotif) error {
	elements := map[string]interface{}{
		"resUri": obj.ResUri,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertTrafficInfluDataRequired(obj.TrafficInfluData); err != nil {
		return err
	}
	return nil
}

// AssertTrafficInfluDataNotifConstraints checks if the values respects the defined constraints
func AssertTrafficInfluDataNotifConstraints(obj TrafficInfluDataNotif) error {
	return nil
}
