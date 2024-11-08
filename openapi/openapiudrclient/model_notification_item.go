/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_udr_cli

import (
	"encoding/json"
	"fmt"
)

// checks if the NotificationItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationItem{}

// NotificationItem Identifies a data change notification when the change occurs in a fragment (subset of resource data) of a given resource. 
type NotificationItem struct {
	// String providing an URI formatted according to RFC 3986.
	ResourceId string `json:"resourceId"`
	NotifItems []UpdatedItem `json:"notifItems"`
}

type _NotificationItem NotificationItem

// NewNotificationItem instantiates a new NotificationItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationItem(resourceId string, notifItems []UpdatedItem) *NotificationItem {
	this := NotificationItem{}
	this.ResourceId = resourceId
	this.NotifItems = notifItems
	return &this
}

// NewNotificationItemWithDefaults instantiates a new NotificationItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationItemWithDefaults() *NotificationItem {
	this := NotificationItem{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *NotificationItem) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *NotificationItem) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *NotificationItem) SetResourceId(v string) {
	o.ResourceId = v
}

// GetNotifItems returns the NotifItems field value
func (o *NotificationItem) GetNotifItems() []UpdatedItem {
	if o == nil {
		var ret []UpdatedItem
		return ret
	}

	return o.NotifItems
}

// GetNotifItemsOk returns a tuple with the NotifItems field value
// and a boolean to check if the value has been set.
func (o *NotificationItem) GetNotifItemsOk() ([]UpdatedItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotifItems, true
}

// SetNotifItems sets field value
func (o *NotificationItem) SetNotifItems(v []UpdatedItem) {
	o.NotifItems = v
}

func (o NotificationItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["notifItems"] = o.NotifItems
	return toSerialize, nil
}

func (o *NotificationItem) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceId",
		"notifItems",
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

	varNotificationItem := _NotificationItem{}

	err = json.Unmarshal(bytes, &varNotificationItem)

	if err != nil {
		return err
	}

	*o = NotificationItem(varNotificationItem)

	return err
}

type NullableNotificationItem struct {
	value *NotificationItem
	isSet bool
}

func (v NullableNotificationItem) Get() *NotificationItem {
	return v.value
}

func (v *NullableNotificationItem) Set(val *NotificationItem) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationItem) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationItem(val *NotificationItem) *NullableNotificationItem {
	return &NullableNotificationItem{value: val, isSet: true}
}

func (v NullableNotificationItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


