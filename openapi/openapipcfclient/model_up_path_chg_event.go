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

// checks if the UpPathChgEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpPathChgEvent{}

// UpPathChgEvent struct for UpPathChgEvent
type UpPathChgEvent struct {
	NotificationUri string `json:"notificationUri"`
	// It is used to set the value of Notification Correlation ID in the notification sent by the SMF.
	NotifCorreId string `json:"notifCorreId"`
	DnaiChgType DnaiChangeType `json:"dnaiChgType"`
}

type _UpPathChgEvent UpPathChgEvent

// NewUpPathChgEvent instantiates a new UpPathChgEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpPathChgEvent(notificationUri string, notifCorreId string, dnaiChgType DnaiChangeType) *UpPathChgEvent {
	this := UpPathChgEvent{}
	this.NotificationUri = notificationUri
	this.NotifCorreId = notifCorreId
	this.DnaiChgType = dnaiChgType
	return &this
}

// NewUpPathChgEventWithDefaults instantiates a new UpPathChgEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpPathChgEventWithDefaults() *UpPathChgEvent {
	this := UpPathChgEvent{}
	return &this
}

// GetNotificationUri returns the NotificationUri field value
func (o *UpPathChgEvent) GetNotificationUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationUri
}

// GetNotificationUriOk returns a tuple with the NotificationUri field value
// and a boolean to check if the value has been set.
func (o *UpPathChgEvent) GetNotificationUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationUri, true
}

// SetNotificationUri sets field value
func (o *UpPathChgEvent) SetNotificationUri(v string) {
	o.NotificationUri = v
}

// GetNotifCorreId returns the NotifCorreId field value
func (o *UpPathChgEvent) GetNotifCorreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifCorreId
}

// GetNotifCorreIdOk returns a tuple with the NotifCorreId field value
// and a boolean to check if the value has been set.
func (o *UpPathChgEvent) GetNotifCorreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifCorreId, true
}

// SetNotifCorreId sets field value
func (o *UpPathChgEvent) SetNotifCorreId(v string) {
	o.NotifCorreId = v
}

// GetDnaiChgType returns the DnaiChgType field value
func (o *UpPathChgEvent) GetDnaiChgType() DnaiChangeType {
	if o == nil {
		var ret DnaiChangeType
		return ret
	}

	return o.DnaiChgType
}

// GetDnaiChgTypeOk returns a tuple with the DnaiChgType field value
// and a boolean to check if the value has been set.
func (o *UpPathChgEvent) GetDnaiChgTypeOk() (*DnaiChangeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DnaiChgType, true
}

// SetDnaiChgType sets field value
func (o *UpPathChgEvent) SetDnaiChgType(v DnaiChangeType) {
	o.DnaiChgType = v
}

func (o UpPathChgEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpPathChgEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notificationUri"] = o.NotificationUri
	toSerialize["notifCorreId"] = o.NotifCorreId
	toSerialize["dnaiChgType"] = o.DnaiChgType
	return toSerialize, nil
}

func (o *UpPathChgEvent) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"notificationUri",
		"notifCorreId",
		"dnaiChgType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpPathChgEvent := _UpPathChgEvent{}

	err = json.Unmarshal(bytes, &varUpPathChgEvent)

	if err != nil {
		return err
	}

	*o = UpPathChgEvent(varUpPathChgEvent)

	return err
}

type NullableUpPathChgEvent struct {
	value *UpPathChgEvent
	isSet bool
}

func (v NullableUpPathChgEvent) Get() *UpPathChgEvent {
	return v.value
}

func (v *NullableUpPathChgEvent) Set(val *UpPathChgEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpPathChgEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpPathChgEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpPathChgEvent(val *UpPathChgEvent) *NullableUpPathChgEvent {
	return &NullableUpPathChgEvent{value: val, isSet: true}
}

func (v NullableUpPathChgEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpPathChgEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

