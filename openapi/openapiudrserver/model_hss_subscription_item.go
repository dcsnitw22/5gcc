/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// HssSubscriptionItem - Contains info about a single HSS event subscription
type HssSubscriptionItem struct {

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	HssInstanceId string `json:"hssInstanceId"`

	// String providing an URI formatted according to RFC 3986.
	SubscriptionId string `json:"subscriptionId"`

	ContextInfo ContextInfo `json:"contextInfo,omitempty"`
}

// AssertHssSubscriptionItemRequired checks if the required fields are not zero-ed
func AssertHssSubscriptionItemRequired(obj HssSubscriptionItem) error {
	elements := map[string]interface{}{
		"hssInstanceId":  obj.HssInstanceId,
		"subscriptionId": obj.SubscriptionId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertContextInfoRequired(obj.ContextInfo); err != nil {
		return err
	}
	return nil
}

// AssertHssSubscriptionItemConstraints checks if the values respects the defined constraints
func AssertHssSubscriptionItemConstraints(obj HssSubscriptionItem) error {
	return nil
}