/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




// AccessTokenErr - Error returned in the access token response message
type AccessTokenErr struct {

	Error string `json:"error"`

	ErrorDescription string `json:"error_description,omitempty"`

	ErrorUri string `json:"error_uri,omitempty"`
}

// AssertAccessTokenErrRequired checks if the required fields are not zero-ed
func AssertAccessTokenErrRequired(obj AccessTokenErr) error {
	elements := map[string]interface{}{
		"error": obj.Error,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertAccessTokenErrConstraints checks if the values respects the defined constraints
func AssertAccessTokenErrConstraints(obj AccessTokenErr) error {
	return nil
}