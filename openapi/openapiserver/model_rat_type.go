/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiserver

// import (
// 	"encoding/json"
// )

type RatType struct {
}

//Commented because it's causing errors in update sm context request
// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
//func (m *RatType) UnmarshalJSON(data []byte) error {

//	type Alias RatType // To avoid infinite recursion
//	return json.Unmarshal(data, (*Alias)(m))
//}

// AssertRatTypeRequired checks if the required fields are not zero-ed
func AssertRatTypeRequired(obj RatType) error {
	return nil
}

// AssertRatTypeConstraints checks if the values respects the defined constraints
func AssertRatTypeConstraints(obj RatType) error {
	return nil
}