/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type AdditionalSnssaiData struct {
	RequiredAuthnAuthz bool `json:"requiredAuthnAuthz,omitempty"`

	SubscribedUeSliceMbr *SliceMbr `json:"subscribedUeSliceMbr,omitempty"`

	SubscribedNsSrgList []string `json:"subscribedNsSrgList,omitempty"`
}

// AssertAdditionalSnssaiDataRequired checks if the required fields are not zero-ed
func AssertAdditionalSnssaiDataRequired(obj AdditionalSnssaiData) error {
	if obj.SubscribedUeSliceMbr != nil {
		if err := AssertSliceMbrRequired(*obj.SubscribedUeSliceMbr); err != nil {
			return err
		}
	}
	return nil
}

// AssertAdditionalSnssaiDataConstraints checks if the values respects the defined constraints
func AssertAdditionalSnssaiDataConstraints(obj AdditionalSnssaiData) error {
	return nil
}
