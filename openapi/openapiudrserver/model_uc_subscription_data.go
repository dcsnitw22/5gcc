/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UcSubscriptionData - Contains the User Consent Subscription Data.
type UcSubscriptionData struct {

	// A map(list of key-value pairs) where user consent purpose serves as key of user consent
	UserConsentPerPurposeList map[string]UserConsent `json:"userConsentPerPurposeList,omitempty"`
}

// AssertUcSubscriptionDataRequired checks if the required fields are not zero-ed
func AssertUcSubscriptionDataRequired(obj UcSubscriptionData) error {
	return nil
}

// AssertUcSubscriptionDataConstraints checks if the values respects the defined constraints
func AssertUcSubscriptionDataConstraints(obj UcSubscriptionData) error {
	return nil
}
