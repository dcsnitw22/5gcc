/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi


import (
	"time"
)



type PgwInfo struct {

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn"`

	// Fully Qualified Domain Name
	PgwFqdn string `json:"pgwFqdn"`

	PgwIpAddr *IpAddress `json:"pgwIpAddr,omitempty"`

	PlmnId PlmnId `json:"plmnId,omitempty"`

	EpdgInd bool `json:"epdgInd,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	PcfId string `json:"pcfId,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	RegistrationTime time.Time `json:"registrationTime,omitempty"`
}

// AssertPgwInfoRequired checks if the required fields are not zero-ed
func AssertPgwInfoRequired(obj PgwInfo) error {
	elements := map[string]interface{}{
		"dnn": obj.Dnn,
		"pgwFqdn": obj.PgwFqdn,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if obj.PgwIpAddr != nil {
		if err := AssertIpAddressRequired(*obj.PgwIpAddr); err != nil {
			return err
		}
	}
	if err := AssertPlmnIdRequired(obj.PlmnId); err != nil {
		return err
	}
	return nil
}

// AssertPgwInfoConstraints checks if the values respects the defined constraints
func AssertPgwInfoConstraints(obj PgwInfo) error {
	return nil
}