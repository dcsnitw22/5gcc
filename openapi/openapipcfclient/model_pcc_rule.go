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

// checks if the PccRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PccRule{}

// PccRule struct for PccRule
type PccRule struct {
	// An array of IP flow packet filter information.
	FlowInfos []FlowInformation `json:"flowInfos,omitempty"`
	// A reference to the application detection filter configured at the UPF.
	AppId *string `json:"appId,omitempty"`
	// Represents the content version of some content.
	ContVer *int32 `json:"contVer,omitempty"`
	// Univocally identifies the PCC rule within a PDU session.
	PccRuleId string `json:"pccRuleId"`
	Precedence *int32 `json:"precedence,omitempty"`
	AfSigProtocol NullableAfSigProtocol `json:"afSigProtocol,omitempty"`
	// Indication of application relocation possibility.
	AppReloc *bool `json:"appReloc,omitempty"`
	// A reference to the QosData policy decision type. It is the qosId described in subclause 5.6.2.8.
	RefQosData []string `json:"refQosData,omitempty"`
	// A reference to the TrafficControlData policy decision type. It is the tcId described in subclause 5.6.2.10.
	RefTcData []string `json:"refTcData,omitempty"`
	// A reference to the ChargingData policy decision type. It is the chgId described in subclause 5.6.2.11.
	RefChgData []string `json:"refChgData,omitempty"`
	// A reference to UsageMonitoringData policy decision type. It is the umId described in subclause 5.6.2.12.
	RefUmData []string `json:"refUmData,omitempty"`
	// A reference to the condition data. It is the condId described in subclause 5.6.2.9.
	RefCondData NullableString `json:"refCondData,omitempty"`
}

type _PccRule PccRule

// NewPccRule instantiates a new PccRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPccRule(pccRuleId string) *PccRule {
	this := PccRule{}
	this.PccRuleId = pccRuleId
	return &this
}

// NewPccRuleWithDefaults instantiates a new PccRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPccRuleWithDefaults() *PccRule {
	this := PccRule{}
	return &this
}

// GetFlowInfos returns the FlowInfos field value if set, zero value otherwise.
func (o *PccRule) GetFlowInfos() []FlowInformation {
	if o == nil || IsNil(o.FlowInfos) {
		var ret []FlowInformation
		return ret
	}
	return o.FlowInfos
}

// GetFlowInfosOk returns a tuple with the FlowInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetFlowInfosOk() ([]FlowInformation, bool) {
	if o == nil || IsNil(o.FlowInfos) {
		return nil, false
	}
	return o.FlowInfos, true
}

// HasFlowInfos returns a boolean if a field has been set.
func (o *PccRule) HasFlowInfos() bool {
	if o != nil && !IsNil(o.FlowInfos) {
		return true
	}

	return false
}

// SetFlowInfos gets a reference to the given []FlowInformation and assigns it to the FlowInfos field.
func (o *PccRule) SetFlowInfos(v []FlowInformation) {
	o.FlowInfos = v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *PccRule) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *PccRule) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *PccRule) SetAppId(v string) {
	o.AppId = &v
}

// GetContVer returns the ContVer field value if set, zero value otherwise.
func (o *PccRule) GetContVer() int32 {
	if o == nil || IsNil(o.ContVer) {
		var ret int32
		return ret
	}
	return *o.ContVer
}

// GetContVerOk returns a tuple with the ContVer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetContVerOk() (*int32, bool) {
	if o == nil || IsNil(o.ContVer) {
		return nil, false
	}
	return o.ContVer, true
}

// HasContVer returns a boolean if a field has been set.
func (o *PccRule) HasContVer() bool {
	if o != nil && !IsNil(o.ContVer) {
		return true
	}

	return false
}

// SetContVer gets a reference to the given int32 and assigns it to the ContVer field.
func (o *PccRule) SetContVer(v int32) {
	o.ContVer = &v
}

// GetPccRuleId returns the PccRuleId field value
func (o *PccRule) GetPccRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PccRuleId
}

// GetPccRuleIdOk returns a tuple with the PccRuleId field value
// and a boolean to check if the value has been set.
func (o *PccRule) GetPccRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PccRuleId, true
}

// SetPccRuleId sets field value
func (o *PccRule) SetPccRuleId(v string) {
	o.PccRuleId = v
}

// GetPrecedence returns the Precedence field value if set, zero value otherwise.
func (o *PccRule) GetPrecedence() int32 {
	if o == nil || IsNil(o.Precedence) {
		var ret int32
		return ret
	}
	return *o.Precedence
}

// GetPrecedenceOk returns a tuple with the Precedence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetPrecedenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Precedence) {
		return nil, false
	}
	return o.Precedence, true
}

// HasPrecedence returns a boolean if a field has been set.
func (o *PccRule) HasPrecedence() bool {
	if o != nil && !IsNil(o.Precedence) {
		return true
	}

	return false
}

// SetPrecedence gets a reference to the given int32 and assigns it to the Precedence field.
func (o *PccRule) SetPrecedence(v int32) {
	o.Precedence = &v
}

// GetAfSigProtocol returns the AfSigProtocol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PccRule) GetAfSigProtocol() AfSigProtocol {
	if o == nil || IsNil(o.AfSigProtocol.Get()) {
		var ret AfSigProtocol
		return ret
	}
	return *o.AfSigProtocol.Get()
}

// GetAfSigProtocolOk returns a tuple with the AfSigProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PccRule) GetAfSigProtocolOk() (*AfSigProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return o.AfSigProtocol.Get(), o.AfSigProtocol.IsSet()
}

// HasAfSigProtocol returns a boolean if a field has been set.
func (o *PccRule) HasAfSigProtocol() bool {
	if o != nil && o.AfSigProtocol.IsSet() {
		return true
	}

	return false
}

// SetAfSigProtocol gets a reference to the given NullableAfSigProtocol and assigns it to the AfSigProtocol field.
func (o *PccRule) SetAfSigProtocol(v AfSigProtocol) {
	o.AfSigProtocol.Set(&v)
}
// SetAfSigProtocolNil sets the value for AfSigProtocol to be an explicit nil
func (o *PccRule) SetAfSigProtocolNil() {
	o.AfSigProtocol.Set(nil)
}

// UnsetAfSigProtocol ensures that no value is present for AfSigProtocol, not even an explicit nil
func (o *PccRule) UnsetAfSigProtocol() {
	o.AfSigProtocol.Unset()
}

// GetAppReloc returns the AppReloc field value if set, zero value otherwise.
func (o *PccRule) GetAppReloc() bool {
	if o == nil || IsNil(o.AppReloc) {
		var ret bool
		return ret
	}
	return *o.AppReloc
}

// GetAppRelocOk returns a tuple with the AppReloc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetAppRelocOk() (*bool, bool) {
	if o == nil || IsNil(o.AppReloc) {
		return nil, false
	}
	return o.AppReloc, true
}

// HasAppReloc returns a boolean if a field has been set.
func (o *PccRule) HasAppReloc() bool {
	if o != nil && !IsNil(o.AppReloc) {
		return true
	}

	return false
}

// SetAppReloc gets a reference to the given bool and assigns it to the AppReloc field.
func (o *PccRule) SetAppReloc(v bool) {
	o.AppReloc = &v
}

// GetRefQosData returns the RefQosData field value if set, zero value otherwise.
func (o *PccRule) GetRefQosData() []string {
	if o == nil || IsNil(o.RefQosData) {
		var ret []string
		return ret
	}
	return o.RefQosData
}

// GetRefQosDataOk returns a tuple with the RefQosData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetRefQosDataOk() ([]string, bool) {
	if o == nil || IsNil(o.RefQosData) {
		return nil, false
	}
	return o.RefQosData, true
}

// HasRefQosData returns a boolean if a field has been set.
func (o *PccRule) HasRefQosData() bool {
	if o != nil && !IsNil(o.RefQosData) {
		return true
	}

	return false
}

// SetRefQosData gets a reference to the given []string and assigns it to the RefQosData field.
func (o *PccRule) SetRefQosData(v []string) {
	o.RefQosData = v
}

// GetRefTcData returns the RefTcData field value if set, zero value otherwise.
func (o *PccRule) GetRefTcData() []string {
	if o == nil || IsNil(o.RefTcData) {
		var ret []string
		return ret
	}
	return o.RefTcData
}

// GetRefTcDataOk returns a tuple with the RefTcData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetRefTcDataOk() ([]string, bool) {
	if o == nil || IsNil(o.RefTcData) {
		return nil, false
	}
	return o.RefTcData, true
}

// HasRefTcData returns a boolean if a field has been set.
func (o *PccRule) HasRefTcData() bool {
	if o != nil && !IsNil(o.RefTcData) {
		return true
	}

	return false
}

// SetRefTcData gets a reference to the given []string and assigns it to the RefTcData field.
func (o *PccRule) SetRefTcData(v []string) {
	o.RefTcData = v
}

// GetRefChgData returns the RefChgData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PccRule) GetRefChgData() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RefChgData
}

// GetRefChgDataOk returns a tuple with the RefChgData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PccRule) GetRefChgDataOk() ([]string, bool) {
	if o == nil || IsNil(o.RefChgData) {
		return nil, false
	}
	return o.RefChgData, true
}

// HasRefChgData returns a boolean if a field has been set.
func (o *PccRule) HasRefChgData() bool {
	if o != nil && IsNil(o.RefChgData) {
		return true
	}

	return false
}

// SetRefChgData gets a reference to the given []string and assigns it to the RefChgData field.
func (o *PccRule) SetRefChgData(v []string) {
	o.RefChgData = v
}

// GetRefUmData returns the RefUmData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PccRule) GetRefUmData() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RefUmData
}

// GetRefUmDataOk returns a tuple with the RefUmData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PccRule) GetRefUmDataOk() ([]string, bool) {
	if o == nil || IsNil(o.RefUmData) {
		return nil, false
	}
	return o.RefUmData, true
}

// HasRefUmData returns a boolean if a field has been set.
func (o *PccRule) HasRefUmData() bool {
	if o != nil && IsNil(o.RefUmData) {
		return true
	}

	return false
}

// SetRefUmData gets a reference to the given []string and assigns it to the RefUmData field.
func (o *PccRule) SetRefUmData(v []string) {
	o.RefUmData = v
}

// GetRefCondData returns the RefCondData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PccRule) GetRefCondData() string {
	if o == nil || IsNil(o.RefCondData.Get()) {
		var ret string
		return ret
	}
	return *o.RefCondData.Get()
}

// GetRefCondDataOk returns a tuple with the RefCondData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PccRule) GetRefCondDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefCondData.Get(), o.RefCondData.IsSet()
}

// HasRefCondData returns a boolean if a field has been set.
func (o *PccRule) HasRefCondData() bool {
	if o != nil && o.RefCondData.IsSet() {
		return true
	}

	return false
}

// SetRefCondData gets a reference to the given NullableString and assigns it to the RefCondData field.
func (o *PccRule) SetRefCondData(v string) {
	o.RefCondData.Set(&v)
}
// SetRefCondDataNil sets the value for RefCondData to be an explicit nil
func (o *PccRule) SetRefCondDataNil() {
	o.RefCondData.Set(nil)
}

// UnsetRefCondData ensures that no value is present for RefCondData, not even an explicit nil
func (o *PccRule) UnsetRefCondData() {
	o.RefCondData.Unset()
}

func (o PccRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PccRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FlowInfos) {
		toSerialize["flowInfos"] = o.FlowInfos
	}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.ContVer) {
		toSerialize["contVer"] = o.ContVer
	}
	toSerialize["pccRuleId"] = o.PccRuleId
	if !IsNil(o.Precedence) {
		toSerialize["precedence"] = o.Precedence
	}
	if o.AfSigProtocol.IsSet() {
		toSerialize["afSigProtocol"] = o.AfSigProtocol.Get()
	}
	if !IsNil(o.AppReloc) {
		toSerialize["appReloc"] = o.AppReloc
	}
	if !IsNil(o.RefQosData) {
		toSerialize["refQosData"] = o.RefQosData
	}
	if !IsNil(o.RefTcData) {
		toSerialize["refTcData"] = o.RefTcData
	}
	if o.RefChgData != nil {
		toSerialize["refChgData"] = o.RefChgData
	}
	if o.RefUmData != nil {
		toSerialize["refUmData"] = o.RefUmData
	}
	if o.RefCondData.IsSet() {
		toSerialize["refCondData"] = o.RefCondData.Get()
	}
	return toSerialize, nil
}

func (o *PccRule) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pccRuleId",
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

	varPccRule := _PccRule{}

	err = json.Unmarshal(bytes, &varPccRule)

	if err != nil {
		return err
	}

	*o = PccRule(varPccRule)

	return err
}

type NullablePccRule struct {
	value *PccRule
	isSet bool
}

func (v NullablePccRule) Get() *PccRule {
	return v.value
}

func (v *NullablePccRule) Set(val *PccRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePccRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePccRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePccRule(val *PccRule) *NullablePccRule {
	return &NullablePccRule{value: val, isSet: true}
}

func (v NullablePccRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePccRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

