/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ProblemDetails1 - Provides additional information in an error response.
type ProblemDetails1 struct {

	// String providing an URI formatted according to RFC 3986.
	Type string `json:"type,omitempty"`

	Title string `json:"title,omitempty"`

	Status int32 `json:"status,omitempty"`

	// A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	Instance string `json:"instance,omitempty"`

	// A machine-readable application error cause specific to this occurrence of the problem.  This IE should be present and provide application-related error information, if available.
	Cause string `json:"cause,omitempty"`

	InvalidParams []InvalidParam `json:"invalidParams,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	AccessTokenError AccessTokenErr `json:"accessTokenError,omitempty"`

	AccessTokenRequest AccessTokenReq `json:"accessTokenRequest,omitempty"`

	// Fully Qualified Domain Name
	NrfId string `json:"nrfId,omitempty"`
}

// AssertProblemDetails1Required checks if the required fields are not zero-ed
func AssertProblemDetails1Required(obj ProblemDetails1) error {
	for _, el := range obj.InvalidParams {
		if err := AssertInvalidParamRequired(el); err != nil {
			return err
		}
	}
	if err := AssertAccessTokenErrRequired(obj.AccessTokenError); err != nil {
		return err
	}
	if err := AssertAccessTokenReqRequired(obj.AccessTokenRequest); err != nil {
		return err
	}
	return nil
}

// AssertProblemDetails1Constraints checks if the values respects the defined constraints
func AssertProblemDetails1Constraints(obj ProblemDetails1) error {
	return nil
}