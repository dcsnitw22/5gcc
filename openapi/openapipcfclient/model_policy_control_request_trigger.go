/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PolicyControlRequestTrigger Possible values are - PLMN_CH: PLMN Change - RES_MO_RE: A request for resource modification has been received by the SMF. The SMF always reports to the PCF. - AC_TY_CH: Access Type Change - UE_IP_CH: UE IP address change. The SMF always reports to the PCF. - UE_MAC_CH: A new UE MAC address is detected or a used UE MAC address is inactive for a specific period - AN_CH_COR: Access Network Charging Correlation Information - US_RE: The PDU Session or the Monitoring key specific resources consumed by a UE either reached the threshold or needs to be reported for other reasons. - APP_STA: The start of application traffic has been detected. - APP_STO: The stop of application traffic has been detected. - AN_INFO: Access Network Information report - CM_SES_FAIL: Credit management session failure - PS_DA_OFF: The SMF reports when the 3GPP PS Data Off status changes. The SMF always reports to the PCF. - DEF_QOS_CH: Default QoS Change. The SMF always reports to the PCF. - SE_AMBR_CH: Session AMBR Change. The SMF always reports to the PCF. - QOS_NOTIF: The SMF notify the PCF when receiving notification from RAN that QoS targets of the QoS Flow cannot be guranteed or gurateed again. - NO_CREDIT: Out of credit - PRA_CH: Change of UE presence in Presence Reporting Area - SAREA_CH: Location Change with respect to the Serving Area - SCNN_CH: Location Change with respect to the Serving CN node - RE_TIMEOUT: Indicates the SMF generated the request because there has been a PCC revalidation timeout - RES_RELEASE: Indicate that the SMF can inform the PCF of the outcome of the release of resources for those rules that require so. - SUCC_RES_ALLO: Indicates that the requested rule data is the successful resource allocation. - RAT_TY_CH: RAT Type Change. - REF_QOS_IND_CH: Reflective QoS indication Change - NUM_OF_PACKET_FILTER: Indicates that the SMF shall report the number of supported packet filter for signalled QoS rules - UE_STATUS_RESUME: Indicates that the UE’s status is resumed. - UE_TZ_CH: UE Time Zone Change 
type PolicyControlRequestTrigger struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PolicyControlRequestTrigger) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PolicyControlRequestTrigger)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PolicyControlRequestTrigger) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePolicyControlRequestTrigger struct {
	value *PolicyControlRequestTrigger
	isSet bool
}

func (v NullablePolicyControlRequestTrigger) Get() *PolicyControlRequestTrigger {
	return v.value
}

func (v *NullablePolicyControlRequestTrigger) Set(val *PolicyControlRequestTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyControlRequestTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyControlRequestTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyControlRequestTrigger(val *PolicyControlRequestTrigger) *NullablePolicyControlRequestTrigger {
	return &NullablePolicyControlRequestTrigger{value: val, isSet: true}
}

func (v NullablePolicyControlRequestTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyControlRequestTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


