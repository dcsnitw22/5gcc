/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// AerialUeSubscriptionInfo1 - Contains the Aerial UE Subscription Information, it at least contains the Aerial UE Indication.
type AerialUeSubscriptionInfo1 struct {
	AerialUeInd AerialUeIndication `json:"aerialUeInd"`

	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Var3gppUavId string `json:"3gppUavId,omitempty"`
}

// AssertAerialUeSubscriptionInfo1Required checks if the required fields are not zero-ed
func AssertAerialUeSubscriptionInfo1Required(obj AerialUeSubscriptionInfo1) error {
	elements := map[string]interface{}{
		"aerialUeInd": obj.AerialUeInd,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertAerialUeIndicationRequired(obj.AerialUeInd); err != nil {
		return err
	}
	return nil
}

// AssertAerialUeSubscriptionInfo1Constraints checks if the values respects the defined constraints
func AssertAerialUeSubscriptionInfo1Constraints(obj AerialUeSubscriptionInfo1) error {
	return nil
}