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

// checks if the Dynamic5Qi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dynamic5Qi{}

// Dynamic5Qi struct for Dynamic5Qi
type Dynamic5Qi struct {
	ResourceType QosResourceType `json:"resourceType"`
	PriorityLevel int32 `json:"priorityLevel"`
	PacketDelayBudget int32 `json:"packetDelayBudget"`
	PacketErrRate string `json:"packetErrRate"`
	AverWindow *int32 `json:"averWindow,omitempty"`
	MaxDataBurstVol *int32 `json:"maxDataBurstVol,omitempty"`
}

// NewDynamic5Qi instantiates a new Dynamic5Qi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamic5Qi(resourceType QosResourceType, priorityLevel int32, packetDelayBudget int32, packetErrRate string) *Dynamic5Qi {
	this := Dynamic5Qi{}
	this.ResourceType = resourceType
	this.PriorityLevel = priorityLevel
	this.PacketDelayBudget = packetDelayBudget
	this.PacketErrRate = packetErrRate
	var averWindow int32 = 2000
	this.AverWindow = &averWindow
	return &this
}

// NewDynamic5QiWithDefaults instantiates a new Dynamic5Qi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamic5QiWithDefaults() *Dynamic5Qi {
	this := Dynamic5Qi{}
	var averWindow int32 = 2000
	this.AverWindow = &averWindow
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *Dynamic5Qi) GetResourceType() QosResourceType {
	if o == nil {
		var ret QosResourceType
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *Dynamic5Qi) GetResourceTypeOk() (*QosResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *Dynamic5Qi) SetResourceType(v QosResourceType) {
	o.ResourceType = v
}

// GetPriorityLevel returns the PriorityLevel field value
func (o *Dynamic5Qi) GetPriorityLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PriorityLevel
}

// GetPriorityLevelOk returns a tuple with the PriorityLevel field value
// and a boolean to check if the value has been set.
func (o *Dynamic5Qi) GetPriorityLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriorityLevel, true
}

// SetPriorityLevel sets field value
func (o *Dynamic5Qi) SetPriorityLevel(v int32) {
	o.PriorityLevel = v
}

// GetPacketDelayBudget returns the PacketDelayBudget field value
func (o *Dynamic5Qi) GetPacketDelayBudget() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PacketDelayBudget
}

// GetPacketDelayBudgetOk returns a tuple with the PacketDelayBudget field value
// and a boolean to check if the value has been set.
func (o *Dynamic5Qi) GetPacketDelayBudgetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PacketDelayBudget, true
}

// SetPacketDelayBudget sets field value
func (o *Dynamic5Qi) SetPacketDelayBudget(v int32) {
	o.PacketDelayBudget = v
}

// GetPacketErrRate returns the PacketErrRate field value
func (o *Dynamic5Qi) GetPacketErrRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PacketErrRate
}

// GetPacketErrRateOk returns a tuple with the PacketErrRate field value
// and a boolean to check if the value has been set.
func (o *Dynamic5Qi) GetPacketErrRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PacketErrRate, true
}

// SetPacketErrRate sets field value
func (o *Dynamic5Qi) SetPacketErrRate(v string) {
	o.PacketErrRate = v
}

// GetAverWindow returns the AverWindow field value if set, zero value otherwise.
func (o *Dynamic5Qi) GetAverWindow() int32 {
	if o == nil || IsNil(o.AverWindow) {
		var ret int32
		return ret
	}
	return *o.AverWindow
}

// GetAverWindowOk returns a tuple with the AverWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dynamic5Qi) GetAverWindowOk() (*int32, bool) {
	if o == nil || IsNil(o.AverWindow) {
		return nil, false
	}
	return o.AverWindow, true
}

// HasAverWindow returns a boolean if a field has been set.
func (o *Dynamic5Qi) HasAverWindow() bool {
	if o != nil && !IsNil(o.AverWindow) {
		return true
	}

	return false
}

// SetAverWindow gets a reference to the given int32 and assigns it to the AverWindow field.
func (o *Dynamic5Qi) SetAverWindow(v int32) {
	o.AverWindow = &v
}

// GetMaxDataBurstVol returns the MaxDataBurstVol field value if set, zero value otherwise.
func (o *Dynamic5Qi) GetMaxDataBurstVol() int32 {
	if o == nil || IsNil(o.MaxDataBurstVol) {
		var ret int32
		return ret
	}
	return *o.MaxDataBurstVol
}

// GetMaxDataBurstVolOk returns a tuple with the MaxDataBurstVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dynamic5Qi) GetMaxDataBurstVolOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxDataBurstVol) {
		return nil, false
	}
	return o.MaxDataBurstVol, true
}

// HasMaxDataBurstVol returns a boolean if a field has been set.
func (o *Dynamic5Qi) HasMaxDataBurstVol() bool {
	if o != nil && !IsNil(o.MaxDataBurstVol) {
		return true
	}

	return false
}

// SetMaxDataBurstVol gets a reference to the given int32 and assigns it to the MaxDataBurstVol field.
func (o *Dynamic5Qi) SetMaxDataBurstVol(v int32) {
	o.MaxDataBurstVol = &v
}

func (o Dynamic5Qi) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dynamic5Qi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceType"] = o.ResourceType
	toSerialize["priorityLevel"] = o.PriorityLevel
	toSerialize["packetDelayBudget"] = o.PacketDelayBudget
	toSerialize["packetErrRate"] = o.PacketErrRate
	if !IsNil(o.AverWindow) {
		toSerialize["averWindow"] = o.AverWindow
	}
	if !IsNil(o.MaxDataBurstVol) {
		toSerialize["maxDataBurstVol"] = o.MaxDataBurstVol
	}
	return toSerialize, nil
}

type NullableDynamic5Qi struct {
	value *Dynamic5Qi
	isSet bool
}

func (v NullableDynamic5Qi) Get() *Dynamic5Qi {
	return v.value
}

func (v *NullableDynamic5Qi) Set(val *Dynamic5Qi) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamic5Qi) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamic5Qi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamic5Qi(val *Dynamic5Qi) *NullableDynamic5Qi {
	return &NullableDynamic5Qi{value: val, isSet: true}
}

func (v NullableDynamic5Qi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamic5Qi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

