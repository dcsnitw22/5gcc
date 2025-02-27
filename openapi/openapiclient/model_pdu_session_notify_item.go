/*
Nsmf_PDUSession

SMF PDU Session Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
)

// checks if the PduSessionNotifyItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PduSessionNotifyItem{}

// PduSessionNotifyItem struct for PduSessionNotifyItem
type PduSessionNotifyItem struct {
	NotificationCause NotificationCause `json:"notificationCause"`
}

// NewPduSessionNotifyItem instantiates a new PduSessionNotifyItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPduSessionNotifyItem(notificationCause NotificationCause) *PduSessionNotifyItem {
	this := PduSessionNotifyItem{}
	this.NotificationCause = notificationCause
	return &this
}

// NewPduSessionNotifyItemWithDefaults instantiates a new PduSessionNotifyItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPduSessionNotifyItemWithDefaults() *PduSessionNotifyItem {
	this := PduSessionNotifyItem{}
	return &this
}

// GetNotificationCause returns the NotificationCause field value
func (o *PduSessionNotifyItem) GetNotificationCause() NotificationCause {
	if o == nil {
		var ret NotificationCause
		return ret
	}

	return o.NotificationCause
}

// GetNotificationCauseOk returns a tuple with the NotificationCause field value
// and a boolean to check if the value has been set.
func (o *PduSessionNotifyItem) GetNotificationCauseOk() (*NotificationCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationCause, true
}

// SetNotificationCause sets field value
func (o *PduSessionNotifyItem) SetNotificationCause(v NotificationCause) {
	o.NotificationCause = v
}

func (o PduSessionNotifyItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PduSessionNotifyItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notificationCause"] = o.NotificationCause
	return toSerialize, nil
}

type NullablePduSessionNotifyItem struct {
	value *PduSessionNotifyItem
	isSet bool
}

func (v NullablePduSessionNotifyItem) Get() *PduSessionNotifyItem {
	return v.value
}

func (v *NullablePduSessionNotifyItem) Set(val *PduSessionNotifyItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionNotifyItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionNotifyItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionNotifyItem(val *PduSessionNotifyItem) *NullablePduSessionNotifyItem {
	return &NullablePduSessionNotifyItem{value: val, isSet: true}
}

func (v NullablePduSessionNotifyItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionNotifyItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


