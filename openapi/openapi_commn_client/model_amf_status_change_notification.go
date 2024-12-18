/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_commn_client

import (
	"encoding/json"
)

// checks if the AmfStatusChangeNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmfStatusChangeNotification{}

// AmfStatusChangeNotification struct for AmfStatusChangeNotification
type AmfStatusChangeNotification struct {
	AmfStatusInfoList []AmfStatusInfo `json:"amfStatusInfoList"`
}

// NewAmfStatusChangeNotification instantiates a new AmfStatusChangeNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfStatusChangeNotification(amfStatusInfoList []AmfStatusInfo) *AmfStatusChangeNotification {
	this := AmfStatusChangeNotification{}
	this.AmfStatusInfoList = amfStatusInfoList
	return &this
}

// NewAmfStatusChangeNotificationWithDefaults instantiates a new AmfStatusChangeNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfStatusChangeNotificationWithDefaults() *AmfStatusChangeNotification {
	this := AmfStatusChangeNotification{}
	return &this
}

// GetAmfStatusInfoList returns the AmfStatusInfoList field value
func (o *AmfStatusChangeNotification) GetAmfStatusInfoList() []AmfStatusInfo {
	if o == nil {
		var ret []AmfStatusInfo
		return ret
	}

	return o.AmfStatusInfoList
}

// GetAmfStatusInfoListOk returns a tuple with the AmfStatusInfoList field value
// and a boolean to check if the value has been set.
func (o *AmfStatusChangeNotification) GetAmfStatusInfoListOk() ([]AmfStatusInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmfStatusInfoList, true
}

// SetAmfStatusInfoList sets field value
func (o *AmfStatusChangeNotification) SetAmfStatusInfoList(v []AmfStatusInfo) {
	o.AmfStatusInfoList = v
}

func (o AmfStatusChangeNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmfStatusChangeNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amfStatusInfoList"] = o.AmfStatusInfoList
	return toSerialize, nil
}

type NullableAmfStatusChangeNotification struct {
	value *AmfStatusChangeNotification
	isSet bool
}

func (v NullableAmfStatusChangeNotification) Get() *AmfStatusChangeNotification {
	return v.value
}

func (v *NullableAmfStatusChangeNotification) Set(val *AmfStatusChangeNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfStatusChangeNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfStatusChangeNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfStatusChangeNotification(val *AmfStatusChangeNotification) *NullableAmfStatusChangeNotification {
	return &NullableAmfStatusChangeNotification{value: val, isSet: true}
}

func (v NullableAmfStatusChangeNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfStatusChangeNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


