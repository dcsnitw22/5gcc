/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiserver

import (
	"encoding/json"
)

// RequestIndication - Possible values are - UE_REQ_PDU_SES_MOD - UE_REQ_PDU_SES_REL - PDU_SES_MOB - NW_REQ_PDU_SES_AUTH - NW_REQ_PDU_SES_MOD - NW_REQ_PDU_SES_REL - EBI_ASSIGNMENT_REQ
type RequestIndication struct {
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *RequestIndication) UnmarshalJSON(data []byte) error {

	type Alias RequestIndication // To avoid infinite recursion
	return json.Unmarshal(data, (*Alias)(m))
}

// AssertRequestIndicationRequired checks if the required fields are not zero-ed
func AssertRequestIndicationRequired(obj RequestIndication) error {
	return nil
}

// AssertRequestIndicationConstraints checks if the values respects the defined constraints
func AssertRequestIndicationConstraints(obj RequestIndication) error {
	return nil
}