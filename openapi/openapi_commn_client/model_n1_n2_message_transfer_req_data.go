/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_commn_client

import (
	"encoding/json"
	"k8s.io/klog"
)

// checks if the N1N2MessageTransferReqData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &N1N2MessageTransferReqData{}

// N1N2MessageTransferReqData struct for N1N2MessageTransferReqData
type N1N2MessageTransferReqData struct {
	N1MessageContainer *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	N2InfoContainer *N2InfoContainer `json:"n2InfoContainer,omitempty"`
	SkipInd *bool `json:"skipInd,omitempty"`
	LastMsgIndication *bool `json:"lastMsgIndication,omitempty"`
	PduSessionId *int32 `json:"pduSessionId,omitempty"`
	LcsCorrelationId *string `json:"lcsCorrelationId,omitempty"`
	Ppi *int32 `json:"ppi,omitempty"`
	Arp *Arp `json:"arp,omitempty"`
	Var5qi *int32 `json:"5qi,omitempty"`
	N1n2FailureTxfNotifURI *string `json:"n1n2FailureTxfNotifURI,omitempty"`
	SmfReallocationInd *bool `json:"smfReallocationInd,omitempty"`
	AreaOfValidity *AreaOfValidity `json:"areaOfValidity,omitempty"`
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	OldGuami *Guami `json:"oldGuami,omitempty"`
}

// NewN1N2MessageTransferReqData instantiates a new N1N2MessageTransferReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN1N2MessageTransferReqData() *N1N2MessageTransferReqData {
	this := N1N2MessageTransferReqData{}
	var skipInd bool = false
	this.SkipInd = &skipInd
	var smfReallocationInd bool = false
	this.SmfReallocationInd = &smfReallocationInd
	return &this
}

// NewN1N2MessageTransferReqDataWithDefaults instantiates a new N1N2MessageTransferReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN1N2MessageTransferReqDataWithDefaults() *N1N2MessageTransferReqData {
	this := N1N2MessageTransferReqData{}
	var skipInd bool = false
	this.SkipInd = &skipInd
	var smfReallocationInd bool = false
	this.SmfReallocationInd = &smfReallocationInd
	return &this
}

// GetN1MessageContainer returns the N1MessageContainer field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetN1MessageContainer() N1MessageContainer {
	if o == nil || IsNil(o.N1MessageContainer) {
		var ret N1MessageContainer
		return ret
	}
	klog.Info(*o.N1MessageContainer)
	return *o.N1MessageContainer
}

// GetN1MessageContainerOk returns a tuple with the N1MessageContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetN1MessageContainerOk() (*N1MessageContainer, bool) {
	if o == nil || IsNil(o.N1MessageContainer) {
		return nil, false
	}
	klog.Infof("N1MessageContainerOK %+v",*o.N1MessageContainer)
	return o.N1MessageContainer, true
}

// HasN1MessageContainer returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasN1MessageContainer() bool {
	if o != nil && !IsNil(o.N1MessageContainer) {
		return true
	}

	return false
}

// SetN1MessageContainer gets a reference to the given N1MessageContainer and assigns it to the N1MessageContainer field.
func (o *N1N2MessageTransferReqData) SetN1MessageContainer(v N1MessageContainer) {
	o.N1MessageContainer = &v
}

// GetN2InfoContainer returns the N2InfoContainer field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetN2InfoContainer() N2InfoContainer {
	if o == nil || IsNil(o.N2InfoContainer) {
		var ret N2InfoContainer
		return ret
	}
	return *o.N2InfoContainer
}

// GetN2InfoContainerOk returns a tuple with the N2InfoContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetN2InfoContainerOk() (*N2InfoContainer, bool) {
	if o == nil || IsNil(o.N2InfoContainer) {
		return nil, false
	}
	return o.N2InfoContainer, true
}

// HasN2InfoContainer returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasN2InfoContainer() bool {
	if o != nil && !IsNil(o.N2InfoContainer) {
		return true
	}

	return false
}

// SetN2InfoContainer gets a reference to the given N2InfoContainer and assigns it to the N2InfoContainer field.
func (o *N1N2MessageTransferReqData) SetN2InfoContainer(v N2InfoContainer) {
	o.N2InfoContainer = &v
}

// GetSkipInd returns the SkipInd field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetSkipInd() bool {
	if o == nil || IsNil(o.SkipInd) {
		var ret bool
		return ret
	}
	return *o.SkipInd
}

// GetSkipIndOk returns a tuple with the SkipInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetSkipIndOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipInd) {
		return nil, false
	}
	return o.SkipInd, true
}

// HasSkipInd returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasSkipInd() bool {
	if o != nil && !IsNil(o.SkipInd) {
		return true
	}

	return false
}

// SetSkipInd gets a reference to the given bool and assigns it to the SkipInd field.
func (o *N1N2MessageTransferReqData) SetSkipInd(v bool) {
	o.SkipInd = &v
}

// GetLastMsgIndication returns the LastMsgIndication field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetLastMsgIndication() bool {
	if o == nil || IsNil(o.LastMsgIndication) {
		var ret bool
		return ret
	}
	return *o.LastMsgIndication
}

// GetLastMsgIndicationOk returns a tuple with the LastMsgIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetLastMsgIndicationOk() (*bool, bool) {
	if o == nil || IsNil(o.LastMsgIndication) {
		return nil, false
	}
	return o.LastMsgIndication, true
}

// HasLastMsgIndication returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasLastMsgIndication() bool {
	if o != nil && !IsNil(o.LastMsgIndication) {
		return true
	}

	return false
}

// SetLastMsgIndication gets a reference to the given bool and assigns it to the LastMsgIndication field.
func (o *N1N2MessageTransferReqData) SetLastMsgIndication(v bool) {
	o.LastMsgIndication = &v
}

// GetPduSessionId returns the PduSessionId field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetPduSessionId() int32 {
	if o == nil || IsNil(o.PduSessionId) {
		var ret int32
		return ret
	}
	return *o.PduSessionId
}

// GetPduSessionIdOk returns a tuple with the PduSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetPduSessionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PduSessionId) {
		return nil, false
	}
	return o.PduSessionId, true
}

// HasPduSessionId returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasPduSessionId() bool {
	if o != nil && !IsNil(o.PduSessionId) {
		return true
	}

	return false
}

// SetPduSessionId gets a reference to the given int32 and assigns it to the PduSessionId field.
func (o *N1N2MessageTransferReqData) SetPduSessionId(v int32) {
	o.PduSessionId = &v
}

// GetLcsCorrelationId returns the LcsCorrelationId field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetLcsCorrelationId() string {
	if o == nil || IsNil(o.LcsCorrelationId) {
		var ret string
		return ret
	}
	return *o.LcsCorrelationId
}

// GetLcsCorrelationIdOk returns a tuple with the LcsCorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetLcsCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.LcsCorrelationId) {
		return nil, false
	}
	return o.LcsCorrelationId, true
}

// HasLcsCorrelationId returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasLcsCorrelationId() bool {
	if o != nil && !IsNil(o.LcsCorrelationId) {
		return true
	}

	return false
}

// SetLcsCorrelationId gets a reference to the given string and assigns it to the LcsCorrelationId field.
func (o *N1N2MessageTransferReqData) SetLcsCorrelationId(v string) {
	o.LcsCorrelationId = &v
}

// GetPpi returns the Ppi field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetPpi() int32 {
	if o == nil || IsNil(o.Ppi) {
		var ret int32
		return ret
	}
	return *o.Ppi
}

// GetPpiOk returns a tuple with the Ppi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetPpiOk() (*int32, bool) {
	if o == nil || IsNil(o.Ppi) {
		return nil, false
	}
	return o.Ppi, true
}

// HasPpi returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasPpi() bool {
	if o != nil && !IsNil(o.Ppi) {
		return true
	}

	return false
}

// SetPpi gets a reference to the given int32 and assigns it to the Ppi field.
func (o *N1N2MessageTransferReqData) SetPpi(v int32) {
	o.Ppi = &v
}

// GetArp returns the Arp field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetArp() Arp {
	if o == nil || IsNil(o.Arp) {
		var ret Arp
		return ret
	}
	return *o.Arp
}

// GetArpOk returns a tuple with the Arp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetArpOk() (*Arp, bool) {
	if o == nil || IsNil(o.Arp) {
		return nil, false
	}
	return o.Arp, true
}

// HasArp returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasArp() bool {
	if o != nil && !IsNil(o.Arp) {
		return true
	}

	return false
}

// SetArp gets a reference to the given Arp and assigns it to the Arp field.
func (o *N1N2MessageTransferReqData) SetArp(v Arp) {
	o.Arp = &v
}

// GetVar5qi returns the Var5qi field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetVar5qi() int32 {
	if o == nil || IsNil(o.Var5qi) {
		var ret int32
		return ret
	}
	return *o.Var5qi
}

// GetVar5qiOk returns a tuple with the Var5qi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetVar5qiOk() (*int32, bool) {
	if o == nil || IsNil(o.Var5qi) {
		return nil, false
	}
	return o.Var5qi, true
}

// HasVar5qi returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasVar5qi() bool {
	if o != nil && !IsNil(o.Var5qi) {
		return true
	}

	return false
}

// SetVar5qi gets a reference to the given int32 and assigns it to the Var5qi field.
func (o *N1N2MessageTransferReqData) SetVar5qi(v int32) {
	o.Var5qi = &v
}

// GetN1n2FailureTxfNotifURI returns the N1n2FailureTxfNotifURI field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetN1n2FailureTxfNotifURI() string {
	if o == nil || IsNil(o.N1n2FailureTxfNotifURI) {
		var ret string
		return ret
	}
	return *o.N1n2FailureTxfNotifURI
}

// GetN1n2FailureTxfNotifURIOk returns a tuple with the N1n2FailureTxfNotifURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetN1n2FailureTxfNotifURIOk() (*string, bool) {
	if o == nil || IsNil(o.N1n2FailureTxfNotifURI) {
		return nil, false
	}
	return o.N1n2FailureTxfNotifURI, true
}

// HasN1n2FailureTxfNotifURI returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasN1n2FailureTxfNotifURI() bool {
	if o != nil && !IsNil(o.N1n2FailureTxfNotifURI) {
		return true
	}

	return false
}

// SetN1n2FailureTxfNotifURI gets a reference to the given string and assigns it to the N1n2FailureTxfNotifURI field.
func (o *N1N2MessageTransferReqData) SetN1n2FailureTxfNotifURI(v string) {
	o.N1n2FailureTxfNotifURI = &v
}

// GetSmfReallocationInd returns the SmfReallocationInd field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetSmfReallocationInd() bool {
	if o == nil || IsNil(o.SmfReallocationInd) {
		var ret bool
		return ret
	}
	return *o.SmfReallocationInd
}

// GetSmfReallocationIndOk returns a tuple with the SmfReallocationInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetSmfReallocationIndOk() (*bool, bool) {
	if o == nil || IsNil(o.SmfReallocationInd) {
		return nil, false
	}
	return o.SmfReallocationInd, true
}

// HasSmfReallocationInd returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasSmfReallocationInd() bool {
	if o != nil && !IsNil(o.SmfReallocationInd) {
		return true
	}

	return false
}

// SetSmfReallocationInd gets a reference to the given bool and assigns it to the SmfReallocationInd field.
func (o *N1N2MessageTransferReqData) SetSmfReallocationInd(v bool) {
	o.SmfReallocationInd = &v
}

// GetAreaOfValidity returns the AreaOfValidity field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetAreaOfValidity() AreaOfValidity {
	if o == nil || IsNil(o.AreaOfValidity) {
		var ret AreaOfValidity
		return ret
	}
	return *o.AreaOfValidity
}

// GetAreaOfValidityOk returns a tuple with the AreaOfValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetAreaOfValidityOk() (*AreaOfValidity, bool) {
	if o == nil || IsNil(o.AreaOfValidity) {
		return nil, false
	}
	return o.AreaOfValidity, true
}

// HasAreaOfValidity returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasAreaOfValidity() bool {
	if o != nil && !IsNil(o.AreaOfValidity) {
		return true
	}

	return false
}

// SetAreaOfValidity gets a reference to the given AreaOfValidity and assigns it to the AreaOfValidity field.
func (o *N1N2MessageTransferReqData) SetAreaOfValidity(v AreaOfValidity) {
	o.AreaOfValidity = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *N1N2MessageTransferReqData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetOldGuami returns the OldGuami field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetOldGuami() Guami {
	if o == nil || IsNil(o.OldGuami) {
		var ret Guami
		return ret
	}
	return *o.OldGuami
}

// GetOldGuamiOk returns a tuple with the OldGuami field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetOldGuamiOk() (*Guami, bool) {
	if o == nil || IsNil(o.OldGuami) {
		return nil, false
	}
	return o.OldGuami, true
}

// HasOldGuami returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasOldGuami() bool {
	if o != nil && !IsNil(o.OldGuami) {
		return true
	}

	return false
}

// SetOldGuami gets a reference to the given Guami and assigns it to the OldGuami field.
func (o *N1N2MessageTransferReqData) SetOldGuami(v Guami) {
	o.OldGuami = &v
}

func (o N1N2MessageTransferReqData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o N1N2MessageTransferReqData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.N1MessageContainer) {
		toSerialize["n1MessageContainer"] = o.N1MessageContainer
	}
	if !IsNil(o.N2InfoContainer) {
		toSerialize["n2InfoContainer"] = o.N2InfoContainer
	}
	if !IsNil(o.SkipInd) {
		toSerialize["skipInd"] = o.SkipInd
	}
	if !IsNil(o.LastMsgIndication) {
		toSerialize["lastMsgIndication"] = o.LastMsgIndication
	}
	if !IsNil(o.PduSessionId) {
		toSerialize["pduSessionId"] = o.PduSessionId
	}
	if !IsNil(o.LcsCorrelationId) {
		toSerialize["lcsCorrelationId"] = o.LcsCorrelationId
	}
	if !IsNil(o.Ppi) {
		toSerialize["ppi"] = o.Ppi
	}
	if !IsNil(o.Arp) {
		toSerialize["arp"] = o.Arp
	}
	if !IsNil(o.Var5qi) {
		toSerialize["5qi"] = o.Var5qi
	}
	if !IsNil(o.N1n2FailureTxfNotifURI) {
		toSerialize["n1n2FailureTxfNotifURI"] = o.N1n2FailureTxfNotifURI
	}
	if !IsNil(o.SmfReallocationInd) {
		toSerialize["smfReallocationInd"] = o.SmfReallocationInd
	}
	if !IsNil(o.AreaOfValidity) {
		toSerialize["areaOfValidity"] = o.AreaOfValidity
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.OldGuami) {
		toSerialize["oldGuami"] = o.OldGuami
	}
	return toSerialize, nil
}

type NullableN1N2MessageTransferReqData struct {
	value *N1N2MessageTransferReqData
	isSet bool
}

func (v NullableN1N2MessageTransferReqData) Get() *N1N2MessageTransferReqData {
	return v.value
}

func (v *NullableN1N2MessageTransferReqData) Set(val *N1N2MessageTransferReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableN1N2MessageTransferReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableN1N2MessageTransferReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN1N2MessageTransferReqData(val *N1N2MessageTransferReqData) *NullableN1N2MessageTransferReqData {
	return &NullableN1N2MessageTransferReqData{value: val, isSet: true}
}

func (v NullableN1N2MessageTransferReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN1N2MessageTransferReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	klog.Info(src,*v.value.N1MessageContainer)
	return json.Unmarshal(src, &v.value)
}


