/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// NotificationItem - Identifies a data change notification when the change occurs in a fragment (subset of resource data) of a given resource.
type NotificationItem struct {

	// String providing an URI formatted according to RFC 3986.
	ResourceId string `json:"resourceId"`

	NotifItems []UpdatedItem `json:"notifItems"`
}

// AssertNotificationItemRequired checks if the required fields are not zero-ed
func AssertNotificationItemRequired(obj NotificationItem) error {
	elements := map[string]interface{}{
		"resourceId": obj.ResourceId,
		"notifItems": obj.NotifItems,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.NotifItems {
		if err := AssertUpdatedItemRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertNotificationItemConstraints checks if the values respects the defined constraints
func AssertNotificationItemConstraints(obj NotificationItem) error {
	return nil
}