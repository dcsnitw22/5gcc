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

// checks if the MmContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MmContext{}

// MmContext struct for MmContext
type MmContext struct {
	AccessType AccessType `json:"accessType"`
	NasSecurityMode *NasSecurityMode `json:"nasSecurityMode,omitempty"`
	EpsNasSecurityMode *EpsNasSecurityMode `json:"epsNasSecurityMode,omitempty"`
	NasDownlinkCount *int32 `json:"nasDownlinkCount,omitempty"`
	NasUplinkCount *int32 `json:"nasUplinkCount,omitempty"`
	UeSecurityCapability *string `json:"ueSecurityCapability,omitempty"`
	S1UeNetworkCapability *string `json:"s1UeNetworkCapability,omitempty"`
	AllowedNssai []Snssai `json:"allowedNssai,omitempty"`
	NssaiMappingList []NssaiMapping `json:"nssaiMappingList,omitempty"`
	NsInstanceList []string `json:"nsInstanceList,omitempty"`
	ExpectedUEbehavior *ExpectedUeBehavior `json:"expectedUEbehavior,omitempty"`
	N3IwfId interface{} `json:"n3IwfId,omitempty"`
	AnN2ApId *int32 `json:"anN2ApId,omitempty"`
}

// NewMmContext instantiates a new MmContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMmContext(accessType AccessType) *MmContext {
	this := MmContext{}
	this.AccessType = accessType
	return &this
}

// NewMmContextWithDefaults instantiates a new MmContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMmContextWithDefaults() *MmContext {
	this := MmContext{}
	return &this
}

// GetAccessType returns the AccessType field value
func (o *MmContext) GetAccessType() AccessType {
	if o == nil {
		var ret AccessType
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *MmContext) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *MmContext) SetAccessType(v AccessType) {
	o.AccessType = v
}

// GetNasSecurityMode returns the NasSecurityMode field value if set, zero value otherwise.
func (o *MmContext) GetNasSecurityMode() NasSecurityMode {
	if o == nil || IsNil(o.NasSecurityMode) {
		var ret NasSecurityMode
		return ret
	}
	return *o.NasSecurityMode
}

// GetNasSecurityModeOk returns a tuple with the NasSecurityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MmContext) GetNasSecurityModeOk() (*NasSecurityMode, bool) {
	if o == nil || IsNil(o.NasSecurityMode) {
		return nil, false
	}
	return o.NasSecurityMode, true
}

// HasNasSecurityMode returns a boolean if a field has been set.
func (o *MmContext) HasNasSecurityMode() bool {
	if o != nil && !IsNil(o.NasSecurityMode) {
		return true
	}

	return false
}

// SetNasSecurityMode gets a reference to the given NasSecurityMode and assigns it to the NasSecurityMode field.
func (o *MmContext) SetNasSecurityMode(v NasSecurityMode) {
	o.NasSecurityMode = &v
}

// GetEpsNasSecurityMode returns the EpsNasSecurityMode field value if set, zero value otherwise.
func (o *MmContext) GetEpsNasSecurityMode() EpsNasSecurityMode {
	if o == nil || IsNil(o.EpsNasSecurityMode) {
		var ret EpsNasSecurityMode
		return ret
	}
	return *o.EpsNasSecurityMode
}

// GetEpsNasSecurityModeOk returns a tuple with the EpsNasSecurityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MmContext) GetEpsNasSecurityModeOk() (*EpsNasSecurityMode, bool) {
	if o == nil || IsNil(o.EpsNasSecurityMode) {
		return nil, false
	}
	return o.EpsNasSecurityMode, true
}

// HasEpsNasSecurityMode returns a boolean if a field has been set.
func (o *MmContext) HasEpsNasSecurityMode() bool {
	if o != nil && !IsNil(o.EpsNasSecurityMode) {
		return true
	}

	return false
}

// SetEpsNasSecurityMode gets a reference to the given EpsNasSecurityMode and assigns it to the EpsNasSecurityMode field.
func (o *MmContext) SetEpsNasSecurityMode(v EpsNasSecurityMode) {
	o.EpsNasSecurityMode = &v
}

// GetNasDownlinkCount returns the NasDownlinkCount field value if set, zero value otherwise.
func (o *MmContext) GetNasDownlinkCount() int32 {
	if o == nil || IsNil(o.NasDownlinkCount) {
		var ret int32
		return ret
	}
	return *o.NasDownlinkCount
}

// GetNasDownlinkCountOk returns a tuple with the NasDownlinkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MmContext) GetNasDownlinkCountOk() (*int32, bool) {
	if o == nil || IsNil(o.NasDownlinkCount) {
		return nil, false
	}
	return o.NasDownlinkCount, true
}

// HasNasDownlinkCount returns a boolean if a field has been set.
func (o *MmContext) HasNasDownlinkCount() bool {
	if o != nil && !IsNil(o.NasDownlinkCount) {
		return true
	}

	return false
}

// SetNasDownlinkCount gets a reference to the given int32 and assigns it to the NasDownlinkCount field.
func (o *MmContext) SetNasDownlinkCount(v int32) {
	o.NasDownlinkCount = &v
}

// GetNasUplinkCount returns the NasUplinkCount field value if set, zero value otherwise.
func (o *MmContext) GetNasUplinkCount() int32 {
	if o == nil || IsNil(o.NasUplinkCount) {
		var ret int32
		return ret
	}
	return *o.NasUplinkCount
}

// GetNasUplinkCountOk returns a tuple with the NasUplinkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MmContext) GetNasUplinkCountOk() (*int32, bool) {
	if o == nil || IsNil(o.NasUplinkCount) {
		return nil, false
	}
	return o.NasUplinkCount, true
}

// HasNasUplinkCount returns a boolean if a field has been set.
func (o *MmContext) HasNasUplinkCount() bool {
	if o != nil && !IsNil(o.NasUplinkCount) {
		return true
	}

	return false
}

// SetNasUplinkCount gets a reference to the given int32 and assigns it to the NasUplinkCount field.
func (o *MmContext) SetNasUplinkCount(v int32) {
	o.NasUplinkCount = &v
}

// GetUeSecurityCapability returns the UeSecurityCapability field value if set, zero value otherwise.
func (o *MmContext) GetUeSecurityCapability() string {
	if o == nil || IsNil(o.UeSecurityCapability) {
		var ret string
		return ret
	}
	return *o.UeSecurityCapability
}

// GetUeSecurityCapabilityOk returns a tuple with the UeSecurityCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MmContext) GetUeSecurityCapabilityOk() (*string, bool) {
	if o == nil || IsNil(o.UeSecurityCapability) {
		return nil, false
	}
	return o.UeSecurityCapability, true
}

// HasUeSecurityCapability returns a boolean if a field has been set.
func (o *MmContext) HasUeSecurityCapability() bool {
	if o != nil && !IsNil(o.UeSecurityCapability) {
		return true
	}

	return false
}

// SetUeSecurityCapability gets a reference to the given string and assigns it to the UeSecurityCapability field.
func (o *MmContext) SetUeSecurityCapability(v string) {
	o.UeSecurityCapability = &v
}

// GetS1UeNetworkCapability returns the S1UeNetworkCapability field value if set, zero value otherwise.
func (o *MmContext) GetS1UeNetworkCapability() string {
	if o == nil || IsNil(o.S1UeNetworkCapability) {
		var ret string
		return ret
	}
	return *o.S1UeNetworkCapability
}

// GetS1UeNetworkCapabilityOk returns a tuple with the S1UeNetworkCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MmContext) GetS1UeNetworkCapabilityOk() (*string, bool) {
	if o == nil || IsNil(o.S1UeNetworkCapability) {
		return nil, false
	}
	return o.S1UeNetworkCapability, true
}

// HasS1UeNetworkCapability returns a boolean if a field has been set.
func (o *MmContext) HasS1UeNetworkCapability() bool {
	if o != nil && !IsNil(o.S1UeNetworkCapability) {
		return true
	}

	return false
}

// SetS1UeNetworkCapability gets a reference to the given string and assigns it to the S1UeNetworkCapability field.
func (o *MmContext) SetS1UeNetworkCapability(v string) {
	o.S1UeNetworkCapability = &v
}

// GetAllowedNssai returns the AllowedNssai field value if set, zero value otherwise.
func (o *MmContext) GetAllowedNssai() []Snssai {
	if o == nil || IsNil(o.AllowedNssai) {
		var ret []Snssai
		return ret
	}
	return o.AllowedNssai
}

// GetAllowedNssaiOk returns a tuple with the AllowedNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MmContext) GetAllowedNssaiOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.AllowedNssai) {
		return nil, false
	}
	return o.AllowedNssai, true
}

// HasAllowedNssai returns a boolean if a field has been set.
func (o *MmContext) HasAllowedNssai() bool {
	if o != nil && !IsNil(o.AllowedNssai) {
		return true
	}

	return false
}

// SetAllowedNssai gets a reference to the given []Snssai and assigns it to the AllowedNssai field.
func (o *MmContext) SetAllowedNssai(v []Snssai) {
	o.AllowedNssai = v
}

// GetNssaiMappingList returns the NssaiMappingList field value if set, zero value otherwise.
func (o *MmContext) GetNssaiMappingList() []NssaiMapping {
	if o == nil || IsNil(o.NssaiMappingList) {
		var ret []NssaiMapping
		return ret
	}
	return o.NssaiMappingList
}

// GetNssaiMappingListOk returns a tuple with the NssaiMappingList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MmContext) GetNssaiMappingListOk() ([]NssaiMapping, bool) {
	if o == nil || IsNil(o.NssaiMappingList) {
		return nil, false
	}
	return o.NssaiMappingList, true
}

// HasNssaiMappingList returns a boolean if a field has been set.
func (o *MmContext) HasNssaiMappingList() bool {
	if o != nil && !IsNil(o.NssaiMappingList) {
		return true
	}

	return false
}

// SetNssaiMappingList gets a reference to the given []NssaiMapping and assigns it to the NssaiMappingList field.
func (o *MmContext) SetNssaiMappingList(v []NssaiMapping) {
	o.NssaiMappingList = v
}

// GetNsInstanceList returns the NsInstanceList field value if set, zero value otherwise.
func (o *MmContext) GetNsInstanceList() []string {
	if o == nil || IsNil(o.NsInstanceList) {
		var ret []string
		return ret
	}
	return o.NsInstanceList
}

// GetNsInstanceListOk returns a tuple with the NsInstanceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MmContext) GetNsInstanceListOk() ([]string, bool) {
	if o == nil || IsNil(o.NsInstanceList) {
		return nil, false
	}
	return o.NsInstanceList, true
}

// HasNsInstanceList returns a boolean if a field has been set.
func (o *MmContext) HasNsInstanceList() bool {
	if o != nil && !IsNil(o.NsInstanceList) {
		return true
	}

	return false
}

// SetNsInstanceList gets a reference to the given []string and assigns it to the NsInstanceList field.
func (o *MmContext) SetNsInstanceList(v []string) {
	o.NsInstanceList = v
}

// GetExpectedUEbehavior returns the ExpectedUEbehavior field value if set, zero value otherwise.
func (o *MmContext) GetExpectedUEbehavior() ExpectedUeBehavior {
	if o == nil || IsNil(o.ExpectedUEbehavior) {
		var ret ExpectedUeBehavior
		return ret
	}
	return *o.ExpectedUEbehavior
}

// GetExpectedUEbehaviorOk returns a tuple with the ExpectedUEbehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MmContext) GetExpectedUEbehaviorOk() (*ExpectedUeBehavior, bool) {
	if o == nil || IsNil(o.ExpectedUEbehavior) {
		return nil, false
	}
	return o.ExpectedUEbehavior, true
}

// HasExpectedUEbehavior returns a boolean if a field has been set.
func (o *MmContext) HasExpectedUEbehavior() bool {
	if o != nil && !IsNil(o.ExpectedUEbehavior) {
		return true
	}

	return false
}

// SetExpectedUEbehavior gets a reference to the given ExpectedUeBehavior and assigns it to the ExpectedUEbehavior field.
func (o *MmContext) SetExpectedUEbehavior(v ExpectedUeBehavior) {
	o.ExpectedUEbehavior = &v
}

// GetN3IwfId returns the N3IwfId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MmContext) GetN3IwfId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.N3IwfId
}

// GetN3IwfIdOk returns a tuple with the N3IwfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MmContext) GetN3IwfIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.N3IwfId) {
		return nil, false
	}
	return &o.N3IwfId, true
}

// HasN3IwfId returns a boolean if a field has been set.
func (o *MmContext) HasN3IwfId() bool {
	if o != nil && IsNil(o.N3IwfId) {
		return true
	}

	return false
}

// SetN3IwfId gets a reference to the given interface{} and assigns it to the N3IwfId field.
func (o *MmContext) SetN3IwfId(v interface{}) {
	o.N3IwfId = v
}

// GetAnN2ApId returns the AnN2ApId field value if set, zero value otherwise.
func (o *MmContext) GetAnN2ApId() int32 {
	if o == nil || IsNil(o.AnN2ApId) {
		var ret int32
		return ret
	}
	return *o.AnN2ApId
}

// GetAnN2ApIdOk returns a tuple with the AnN2ApId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MmContext) GetAnN2ApIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AnN2ApId) {
		return nil, false
	}
	return o.AnN2ApId, true
}

// HasAnN2ApId returns a boolean if a field has been set.
func (o *MmContext) HasAnN2ApId() bool {
	if o != nil && !IsNil(o.AnN2ApId) {
		return true
	}

	return false
}

// SetAnN2ApId gets a reference to the given int32 and assigns it to the AnN2ApId field.
func (o *MmContext) SetAnN2ApId(v int32) {
	o.AnN2ApId = &v
}

func (o MmContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MmContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accessType"] = o.AccessType
	if !IsNil(o.NasSecurityMode) {
		toSerialize["nasSecurityMode"] = o.NasSecurityMode
	}
	if !IsNil(o.EpsNasSecurityMode) {
		toSerialize["epsNasSecurityMode"] = o.EpsNasSecurityMode
	}
	if !IsNil(o.NasDownlinkCount) {
		toSerialize["nasDownlinkCount"] = o.NasDownlinkCount
	}
	if !IsNil(o.NasUplinkCount) {
		toSerialize["nasUplinkCount"] = o.NasUplinkCount
	}
	if !IsNil(o.UeSecurityCapability) {
		toSerialize["ueSecurityCapability"] = o.UeSecurityCapability
	}
	if !IsNil(o.S1UeNetworkCapability) {
		toSerialize["s1UeNetworkCapability"] = o.S1UeNetworkCapability
	}
	if !IsNil(o.AllowedNssai) {
		toSerialize["allowedNssai"] = o.AllowedNssai
	}
	if !IsNil(o.NssaiMappingList) {
		toSerialize["nssaiMappingList"] = o.NssaiMappingList
	}
	if !IsNil(o.NsInstanceList) {
		toSerialize["nsInstanceList"] = o.NsInstanceList
	}
	if !IsNil(o.ExpectedUEbehavior) {
		toSerialize["expectedUEbehavior"] = o.ExpectedUEbehavior
	}
	if o.N3IwfId != nil {
		toSerialize["n3IwfId"] = o.N3IwfId
	}
	if !IsNil(o.AnN2ApId) {
		toSerialize["anN2ApId"] = o.AnN2ApId
	}
	return toSerialize, nil
}

type NullableMmContext struct {
	value *MmContext
	isSet bool
}

func (v NullableMmContext) Get() *MmContext {
	return v.value
}

func (v *NullableMmContext) Set(val *MmContext) {
	v.value = val
	v.isSet = true
}

func (v NullableMmContext) IsSet() bool {
	return v.isSet
}

func (v *NullableMmContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMmContext(val *MmContext) *NullableMmContext {
	return &NullableMmContext{value: val, isSet: true}
}

func (v NullableMmContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMmContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

