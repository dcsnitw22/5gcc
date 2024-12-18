/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// AuthorizationInfo - Represents NIDD authorization information.
type AuthorizationInfo struct {
	Snssai Snssai `json:"snssai"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn"`

	// String uniquely identifying MTC provider information.
	MtcProviderInformation string `json:"mtcProviderInformation"`

	// String providing an URI formatted according to RFC 3986.
	AuthUpdateCallbackUri string `json:"authUpdateCallbackUri"`

	AfId string `json:"afId,omitempty"`

	// Identity of the NEF
	NefId string `json:"nefId,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime time.Time `json:"validityTime,omitempty"`

	ContextInfo ContextInfo `json:"contextInfo,omitempty"`
}

// AssertAuthorizationInfoRequired checks if the required fields are not zero-ed
func AssertAuthorizationInfoRequired(obj AuthorizationInfo) error {
	elements := map[string]interface{}{
		"snssai":                 obj.Snssai,
		"dnn":                    obj.Dnn,
		"mtcProviderInformation": obj.MtcProviderInformation,
		"authUpdateCallbackUri":  obj.AuthUpdateCallbackUri,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertSnssaiRequired(obj.Snssai); err != nil {
		return err
	}
	if err := AssertContextInfoRequired(obj.ContextInfo); err != nil {
		return err
	}
	return nil
}

// AssertAuthorizationInfoConstraints checks if the values respects the defined constraints
func AssertAuthorizationInfoConstraints(obj AuthorizationInfo) error {
	return nil
}
