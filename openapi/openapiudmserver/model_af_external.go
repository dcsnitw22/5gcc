/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type AfExternal struct {

	AfId string `json:"afId,omitempty"`

	AllowedGeographicArea []GeographicArea `json:"allowedGeographicArea,omitempty"`

	PrivacyCheckRelatedAction PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`

	ValidTimePeriod ValidTimePeriod `json:"validTimePeriod,omitempty"`
}

// AssertAfExternalRequired checks if the required fields are not zero-ed
func AssertAfExternalRequired(obj AfExternal) error {
	for _, el := range obj.AllowedGeographicArea {
		if err := AssertGeographicAreaRequired(el); err != nil {
			return err
		}
	}
	if err := AssertPrivacyCheckRelatedActionRequired(obj.PrivacyCheckRelatedAction); err != nil {
		return err
	}
	if err := AssertValidTimePeriodRequired(obj.ValidTimePeriod); err != nil {
		return err
	}
	return nil
}

// AssertAfExternalConstraints checks if the values respects the defined constraints
func AssertAfExternalConstraints(obj AfExternal) error {
	return nil
}