/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// MessageWaitingData - Message Waiting Data list.
type MessageWaitingData struct {
	MwdList []SmscData `json:"mwdList,omitempty"`
}

// AssertMessageWaitingDataRequired checks if the required fields are not zero-ed
func AssertMessageWaitingDataRequired(obj MessageWaitingData) error {
	for _, el := range obj.MwdList {
		if err := AssertSmscDataRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertMessageWaitingDataConstraints checks if the values respects the defined constraints
func AssertMessageWaitingDataConstraints(obj MessageWaitingData) error {
	return nil
}
