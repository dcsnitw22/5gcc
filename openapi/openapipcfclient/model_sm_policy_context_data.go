/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the SmPolicyContextData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmPolicyContextData{}

// SmPolicyContextData struct for SmPolicyContextData
type SmPolicyContextData struct {
	AccNetChId *AccNetChId `json:"accNetChId,omitempty"`
	ChargEntityAddr NullableAccNetChargingAddress `json:"chargEntityAddr,omitempty"`
	Gpsi *string `json:"gpsi,omitempty"`
	Supi string `json:"supi"`
	// When this attribute is included and set to true, it indicates that the supi attribute contains an invalid value.This attribute shall be present if the SUPI is not available in the SMF or the SUPI is unauthenticated. When present it shall be set to true for an invalid SUPI and false (default) for a valid SUPI.
	InvalidSupi *bool `json:"invalidSupi,omitempty"`
	InterGrpIds []string `json:"interGrpIds,omitempty"`
	PduSessionId int32 `json:"pduSessionId"`
	PduSessionType PduSessionType `json:"pduSessionType"`
	Chargingcharacteristics *string `json:"chargingcharacteristics,omitempty"`
	Dnn string `json:"dnn"`
	NotificationUri string `json:"notificationUri"`
	AccessType *AccessType `json:"accessType,omitempty"`
	RatType *RatType `json:"ratType,omitempty"`
	ServingNetwork *NetworkId `json:"servingNetwork,omitempty"`
	UserLocationInfo *UserLocation `json:"userLocationInfo,omitempty"`
	UeTimeZone *string `json:"ueTimeZone,omitempty"`
	Pei *string `json:"pei,omitempty"`
	Ipv4Address *string `json:"ipv4Address,omitempty"`
	Ipv6AddressPrefix *Ipv6Prefix `json:"ipv6AddressPrefix,omitempty"`
	// Indicates the IPv4 address domain
	IpDomain *string `json:"ipDomain,omitempty"`
	SubsSessAmbr *Ambr `json:"subsSessAmbr,omitempty"`
	SubsDefQos *SubscribedDefaultQos `json:"subsDefQos,omitempty"`
	// Contains the number of supported packet filter for signalled QoS rules.
	NumOfPackFilter *int32 `json:"numOfPackFilter,omitempty"`
	// If it is included and set to true, the online charging is applied to the PDU session.
	Online *bool `json:"online,omitempty"`
	// If it is included and set to true, the offline charging is applied to the PDU session.
	Offline *bool `json:"offline,omitempty"`
	// If it is included and set to true, the 3GPP PS Data Off is activated by the UE.
	Var3gppPsDataOffStatus *bool `json:"3gppPsDataOffStatus,omitempty"`
	// If it is included and set to true, the reflective QoS is supported by the UE.
	RefQosIndication *bool `json:"refQosIndication,omitempty"`
	TraceReq NullableTraceData `json:"traceReq,omitempty"`
	SliceInfo Snssai `json:"sliceInfo"`
	// QosFlowUsage *QosFlowUsage `json:"qosFlowUsage,omitempty"`
	ServNfId *ServingNfIdentity `json:"servNfId,omitempty"`
	SuppFeat *string `json:"suppFeat,omitempty"`
	SmfId *string `json:"smfId,omitempty"`
	RecoveryTime *time.Time `json:"recoveryTime,omitempty"`
}

type _SmPolicyContextData SmPolicyContextData

// NewSmPolicyContextData instantiates a new SmPolicyContextData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmPolicyContextData(supi string, pduSessionId int32, pduSessionType PduSessionType, dnn string, notificationUri string, sliceInfo Snssai) *SmPolicyContextData {
	this := SmPolicyContextData{}
	this.Supi = supi
	this.PduSessionId = pduSessionId
	this.PduSessionType = pduSessionType
	this.Dnn = dnn
	this.NotificationUri = notificationUri
	this.SliceInfo = sliceInfo
	return &this
}

// NewSmPolicyContextDataWithDefaults instantiates a new SmPolicyContextData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmPolicyContextDataWithDefaults() *SmPolicyContextData {
	this := SmPolicyContextData{}
	return &this
}

// GetAccNetChId returns the AccNetChId field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetAccNetChId() AccNetChId {
	if o == nil || IsNil(o.AccNetChId) {
		var ret AccNetChId
		return ret
	}
	return *o.AccNetChId
}

// GetAccNetChIdOk returns a tuple with the AccNetChId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetAccNetChIdOk() (*AccNetChId, bool) {
	if o == nil || IsNil(o.AccNetChId) {
		return nil, false
	}
	return o.AccNetChId, true
}

// HasAccNetChId returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasAccNetChId() bool {
	if o != nil && !IsNil(o.AccNetChId) {
		return true
	}

	return false
}

// SetAccNetChId gets a reference to the given AccNetChId and assigns it to the AccNetChId field.
func (o *SmPolicyContextData) SetAccNetChId(v AccNetChId) {
	o.AccNetChId = &v
}

// GetChargEntityAddr returns the ChargEntityAddr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmPolicyContextData) GetChargEntityAddr() AccNetChargingAddress {
	if o == nil || IsNil(o.ChargEntityAddr.Get()) {
		var ret AccNetChargingAddress
		return ret
	}
	return *o.ChargEntityAddr.Get()
}

// GetChargEntityAddrOk returns a tuple with the ChargEntityAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmPolicyContextData) GetChargEntityAddrOk() (*AccNetChargingAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChargEntityAddr.Get(), o.ChargEntityAddr.IsSet()
}

// HasChargEntityAddr returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasChargEntityAddr() bool {
	if o != nil && o.ChargEntityAddr.IsSet() {
		return true
	}

	return false
}

// SetChargEntityAddr gets a reference to the given NullableAccNetChargingAddress and assigns it to the ChargEntityAddr field.
func (o *SmPolicyContextData) SetChargEntityAddr(v AccNetChargingAddress) {
	o.ChargEntityAddr.Set(&v)
}
// SetChargEntityAddrNil sets the value for ChargEntityAddr to be an explicit nil
func (o *SmPolicyContextData) SetChargEntityAddrNil() {
	o.ChargEntityAddr.Set(nil)
}

// UnsetChargEntityAddr ensures that no value is present for ChargEntityAddr, not even an explicit nil
func (o *SmPolicyContextData) UnsetChargEntityAddr() {
	o.ChargEntityAddr.Unset()
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *SmPolicyContextData) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetSupi returns the Supi field value
func (o *SmPolicyContextData) GetSupi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supi
}

// GetSupiOk returns a tuple with the Supi field value
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetSupiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supi, true
}

// SetSupi sets field value
func (o *SmPolicyContextData) SetSupi(v string) {
	o.Supi = v
}

// GetInvalidSupi returns the InvalidSupi field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetInvalidSupi() bool {
	if o == nil || IsNil(o.InvalidSupi) {
		var ret bool
		return ret
	}
	return *o.InvalidSupi
}

// GetInvalidSupiOk returns a tuple with the InvalidSupi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetInvalidSupiOk() (*bool, bool) {
	if o == nil || IsNil(o.InvalidSupi) {
		return nil, false
	}
	return o.InvalidSupi, true
}

// HasInvalidSupi returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasInvalidSupi() bool {
	if o != nil && !IsNil(o.InvalidSupi) {
		return true
	}

	return false
}

// SetInvalidSupi gets a reference to the given bool and assigns it to the InvalidSupi field.
func (o *SmPolicyContextData) SetInvalidSupi(v bool) {
	o.InvalidSupi = &v
}

// GetInterGrpIds returns the InterGrpIds field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetInterGrpIds() []string {
	if o == nil || IsNil(o.InterGrpIds) {
		var ret []string
		return ret
	}
	return o.InterGrpIds
}

// GetInterGrpIdsOk returns a tuple with the InterGrpIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetInterGrpIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.InterGrpIds) {
		return nil, false
	}
	return o.InterGrpIds, true
}

// HasInterGrpIds returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasInterGrpIds() bool {
	if o != nil && !IsNil(o.InterGrpIds) {
		return true
	}

	return false
}

// SetInterGrpIds gets a reference to the given []string and assigns it to the InterGrpIds field.
func (o *SmPolicyContextData) SetInterGrpIds(v []string) {
	o.InterGrpIds = v
}

// GetPduSessionId returns the PduSessionId field value
func (o *SmPolicyContextData) GetPduSessionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PduSessionId
}

// GetPduSessionIdOk returns a tuple with the PduSessionId field value
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetPduSessionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PduSessionId, true
}

// SetPduSessionId sets field value
func (o *SmPolicyContextData) SetPduSessionId(v int32) {
	o.PduSessionId = v
}

// GetPduSessionType returns the PduSessionType field value
func (o *SmPolicyContextData) GetPduSessionType() PduSessionType {
	if o == nil {
		var ret PduSessionType
		return ret
	}

	return o.PduSessionType
}

// GetPduSessionTypeOk returns a tuple with the PduSessionType field value
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetPduSessionTypeOk() (*PduSessionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PduSessionType, true
}

// SetPduSessionType sets field value
func (o *SmPolicyContextData) SetPduSessionType(v PduSessionType) {
	o.PduSessionType = v
}

// GetChargingcharacteristics returns the Chargingcharacteristics field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetChargingcharacteristics() string {
	if o == nil || IsNil(o.Chargingcharacteristics) {
		var ret string
		return ret
	}
	return *o.Chargingcharacteristics
}

// GetChargingcharacteristicsOk returns a tuple with the Chargingcharacteristics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetChargingcharacteristicsOk() (*string, bool) {
	if o == nil || IsNil(o.Chargingcharacteristics) {
		return nil, false
	}
	return o.Chargingcharacteristics, true
}

// HasChargingcharacteristics returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasChargingcharacteristics() bool {
	if o != nil && !IsNil(o.Chargingcharacteristics) {
		return true
	}

	return false
}

// SetChargingcharacteristics gets a reference to the given string and assigns it to the Chargingcharacteristics field.
func (o *SmPolicyContextData) SetChargingcharacteristics(v string) {
	o.Chargingcharacteristics = &v
}

// GetDnn returns the Dnn field value
func (o *SmPolicyContextData) GetDnn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetDnnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *SmPolicyContextData) SetDnn(v string) {
	o.Dnn = v
}

// GetNotificationUri returns the NotificationUri field value
func (o *SmPolicyContextData) GetNotificationUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationUri
}

// GetNotificationUriOk returns a tuple with the NotificationUri field value
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetNotificationUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationUri, true
}

// SetNotificationUri sets field value
func (o *SmPolicyContextData) SetNotificationUri(v string) {
	o.NotificationUri = v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetAccessType() AccessType {
	if o == nil || IsNil(o.AccessType) {
		var ret AccessType
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given AccessType and assigns it to the AccessType field.
func (o *SmPolicyContextData) SetAccessType(v AccessType) {
	o.AccessType = &v
}

// GetRatType returns the RatType field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetRatType() RatType {
	if o == nil || IsNil(o.RatType) {
		var ret RatType
		return ret
	}
	return *o.RatType
}

// GetRatTypeOk returns a tuple with the RatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetRatTypeOk() (*RatType, bool) {
	if o == nil || IsNil(o.RatType) {
		return nil, false
	}
	return o.RatType, true
}

// HasRatType returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasRatType() bool {
	if o != nil && !IsNil(o.RatType) {
		return true
	}

	return false
}

// SetRatType gets a reference to the given RatType and assigns it to the RatType field.
func (o *SmPolicyContextData) SetRatType(v RatType) {
	o.RatType = &v
}

// GetServingNetwork returns the ServingNetwork field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetServingNetwork() NetworkId {
	if o == nil || IsNil(o.ServingNetwork) {
		var ret NetworkId
		return ret
	}
	return *o.ServingNetwork
}

// GetServingNetworkOk returns a tuple with the ServingNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetServingNetworkOk() (*NetworkId, bool) {
	if o == nil || IsNil(o.ServingNetwork) {
		return nil, false
	}
	return o.ServingNetwork, true
}

// HasServingNetwork returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasServingNetwork() bool {
	if o != nil && !IsNil(o.ServingNetwork) {
		return true
	}

	return false
}

// SetServingNetwork gets a reference to the given NetworkId and assigns it to the ServingNetwork field.
func (o *SmPolicyContextData) SetServingNetwork(v NetworkId) {
	o.ServingNetwork = &v
}

// GetUserLocationInfo returns the UserLocationInfo field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetUserLocationInfo() UserLocation {
	if o == nil || IsNil(o.UserLocationInfo) {
		var ret UserLocation
		return ret
	}
	return *o.UserLocationInfo
}

// GetUserLocationInfoOk returns a tuple with the UserLocationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetUserLocationInfoOk() (*UserLocation, bool) {
	if o == nil || IsNil(o.UserLocationInfo) {
		return nil, false
	}
	return o.UserLocationInfo, true
}

// HasUserLocationInfo returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasUserLocationInfo() bool {
	if o != nil && !IsNil(o.UserLocationInfo) {
		return true
	}

	return false
}

// SetUserLocationInfo gets a reference to the given UserLocation and assigns it to the UserLocationInfo field.
func (o *SmPolicyContextData) SetUserLocationInfo(v UserLocation) {
	o.UserLocationInfo = &v
}

// GetUeTimeZone returns the UeTimeZone field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetUeTimeZone() string {
	if o == nil || IsNil(o.UeTimeZone) {
		var ret string
		return ret
	}
	return *o.UeTimeZone
}

// GetUeTimeZoneOk returns a tuple with the UeTimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetUeTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.UeTimeZone) {
		return nil, false
	}
	return o.UeTimeZone, true
}

// HasUeTimeZone returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasUeTimeZone() bool {
	if o != nil && !IsNil(o.UeTimeZone) {
		return true
	}

	return false
}

// SetUeTimeZone gets a reference to the given string and assigns it to the UeTimeZone field.
func (o *SmPolicyContextData) SetUeTimeZone(v string) {
	o.UeTimeZone = &v
}

// GetPei returns the Pei field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetPei() string {
	if o == nil || IsNil(o.Pei) {
		var ret string
		return ret
	}
	return *o.Pei
}

// GetPeiOk returns a tuple with the Pei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetPeiOk() (*string, bool) {
	if o == nil || IsNil(o.Pei) {
		return nil, false
	}
	return o.Pei, true
}

// HasPei returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasPei() bool {
	if o != nil && !IsNil(o.Pei) {
		return true
	}

	return false
}

// SetPei gets a reference to the given string and assigns it to the Pei field.
func (o *SmPolicyContextData) SetPei(v string) {
	o.Pei = &v
}

// GetIpv4Address returns the Ipv4Address field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetIpv4Address() string {
	if o == nil || IsNil(o.Ipv4Address) {
		var ret string
		return ret
	}
	return *o.Ipv4Address
}

// GetIpv4AddressOk returns a tuple with the Ipv4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetIpv4AddressOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv4Address) {
		return nil, false
	}
	return o.Ipv4Address, true
}

// HasIpv4Address returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasIpv4Address() bool {
	if o != nil && !IsNil(o.Ipv4Address) {
		return true
	}

	return false
}

// SetIpv4Address gets a reference to the given string and assigns it to the Ipv4Address field.
func (o *SmPolicyContextData) SetIpv4Address(v string) {
	o.Ipv4Address = &v
}

// GetIpv6AddressPrefix returns the Ipv6AddressPrefix field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetIpv6AddressPrefix() Ipv6Prefix {
	if o == nil || IsNil(o.Ipv6AddressPrefix) {
		var ret Ipv6Prefix
		return ret
	}
	return *o.Ipv6AddressPrefix
}

// GetIpv6AddressPrefixOk returns a tuple with the Ipv6AddressPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetIpv6AddressPrefixOk() (*Ipv6Prefix, bool) {
	if o == nil || IsNil(o.Ipv6AddressPrefix) {
		return nil, false
	}
	return o.Ipv6AddressPrefix, true
}

// HasIpv6AddressPrefix returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasIpv6AddressPrefix() bool {
	if o != nil && !IsNil(o.Ipv6AddressPrefix) {
		return true
	}

	return false
}

// SetIpv6AddressPrefix gets a reference to the given Ipv6Prefix and assigns it to the Ipv6AddressPrefix field.
func (o *SmPolicyContextData) SetIpv6AddressPrefix(v Ipv6Prefix) {
	o.Ipv6AddressPrefix = &v
}

// GetIpDomain returns the IpDomain field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetIpDomain() string {
	if o == nil || IsNil(o.IpDomain) {
		var ret string
		return ret
	}
	return *o.IpDomain
}

// GetIpDomainOk returns a tuple with the IpDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetIpDomainOk() (*string, bool) {
	if o == nil || IsNil(o.IpDomain) {
		return nil, false
	}
	return o.IpDomain, true
}

// HasIpDomain returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasIpDomain() bool {
	if o != nil && !IsNil(o.IpDomain) {
		return true
	}

	return false
}

// SetIpDomain gets a reference to the given string and assigns it to the IpDomain field.
func (o *SmPolicyContextData) SetIpDomain(v string) {
	o.IpDomain = &v
}

// GetSubsSessAmbr returns the SubsSessAmbr field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetSubsSessAmbr() Ambr {
	if o == nil || IsNil(o.SubsSessAmbr) {
		var ret Ambr
		return ret
	}
	return *o.SubsSessAmbr
}

// GetSubsSessAmbrOk returns a tuple with the SubsSessAmbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetSubsSessAmbrOk() (*Ambr, bool) {
	if o == nil || IsNil(o.SubsSessAmbr) {
		return nil, false
	}
	return o.SubsSessAmbr, true
}

// HasSubsSessAmbr returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasSubsSessAmbr() bool {
	if o != nil && !IsNil(o.SubsSessAmbr) {
		return true
	}

	return false
}

// SetSubsSessAmbr gets a reference to the given Ambr and assigns it to the SubsSessAmbr field.
func (o *SmPolicyContextData) SetSubsSessAmbr(v Ambr) {
	o.SubsSessAmbr = &v
}

// GetSubsDefQos returns the SubsDefQos field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetSubsDefQos() SubscribedDefaultQos {
	if o == nil || IsNil(o.SubsDefQos) {
		var ret SubscribedDefaultQos
		return ret
	}
	return *o.SubsDefQos
}

// GetSubsDefQosOk returns a tuple with the SubsDefQos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetSubsDefQosOk() (*SubscribedDefaultQos, bool) {
	if o == nil || IsNil(o.SubsDefQos) {
		return nil, false
	}
	return o.SubsDefQos, true
}

// HasSubsDefQos returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasSubsDefQos() bool {
	if o != nil && !IsNil(o.SubsDefQos) {
		return true
	}

	return false
}

// SetSubsDefQos gets a reference to the given SubscribedDefaultQos and assigns it to the SubsDefQos field.
func (o *SmPolicyContextData) SetSubsDefQos(v SubscribedDefaultQos) {
	o.SubsDefQos = &v
}

// GetNumOfPackFilter returns the NumOfPackFilter field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetNumOfPackFilter() int32 {
	if o == nil || IsNil(o.NumOfPackFilter) {
		var ret int32
		return ret
	}
	return *o.NumOfPackFilter
}

// GetNumOfPackFilterOk returns a tuple with the NumOfPackFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetNumOfPackFilterOk() (*int32, bool) {
	if o == nil || IsNil(o.NumOfPackFilter) {
		return nil, false
	}
	return o.NumOfPackFilter, true
}

// HasNumOfPackFilter returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasNumOfPackFilter() bool {
	if o != nil && !IsNil(o.NumOfPackFilter) {
		return true
	}

	return false
}

// SetNumOfPackFilter gets a reference to the given int32 and assigns it to the NumOfPackFilter field.
func (o *SmPolicyContextData) SetNumOfPackFilter(v int32) {
	o.NumOfPackFilter = &v
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetOnline() bool {
	if o == nil || IsNil(o.Online) {
		var ret bool
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetOnlineOk() (*bool, bool) {
	if o == nil || IsNil(o.Online) {
		return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasOnline() bool {
	if o != nil && !IsNil(o.Online) {
		return true
	}

	return false
}

// SetOnline gets a reference to the given bool and assigns it to the Online field.
func (o *SmPolicyContextData) SetOnline(v bool) {
	o.Online = &v
}

// GetOffline returns the Offline field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetOffline() bool {
	if o == nil || IsNil(o.Offline) {
		var ret bool
		return ret
	}
	return *o.Offline
}

// GetOfflineOk returns a tuple with the Offline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetOfflineOk() (*bool, bool) {
	if o == nil || IsNil(o.Offline) {
		return nil, false
	}
	return o.Offline, true
}

// HasOffline returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasOffline() bool {
	if o != nil && !IsNil(o.Offline) {
		return true
	}

	return false
}

// SetOffline gets a reference to the given bool and assigns it to the Offline field.
func (o *SmPolicyContextData) SetOffline(v bool) {
	o.Offline = &v
}

// GetVar3gppPsDataOffStatus returns the Var3gppPsDataOffStatus field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetVar3gppPsDataOffStatus() bool {
	if o == nil || IsNil(o.Var3gppPsDataOffStatus) {
		var ret bool
		return ret
	}
	return *o.Var3gppPsDataOffStatus
}

// GetVar3gppPsDataOffStatusOk returns a tuple with the Var3gppPsDataOffStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetVar3gppPsDataOffStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Var3gppPsDataOffStatus) {
		return nil, false
	}
	return o.Var3gppPsDataOffStatus, true
}

// HasVar3gppPsDataOffStatus returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasVar3gppPsDataOffStatus() bool {
	if o != nil && !IsNil(o.Var3gppPsDataOffStatus) {
		return true
	}

	return false
}

// SetVar3gppPsDataOffStatus gets a reference to the given bool and assigns it to the Var3gppPsDataOffStatus field.
func (o *SmPolicyContextData) SetVar3gppPsDataOffStatus(v bool) {
	o.Var3gppPsDataOffStatus = &v
}

// GetRefQosIndication returns the RefQosIndication field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetRefQosIndication() bool {
	if o == nil || IsNil(o.RefQosIndication) {
		var ret bool
		return ret
	}
	return *o.RefQosIndication
}

// GetRefQosIndicationOk returns a tuple with the RefQosIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetRefQosIndicationOk() (*bool, bool) {
	if o == nil || IsNil(o.RefQosIndication) {
		return nil, false
	}
	return o.RefQosIndication, true
}

// HasRefQosIndication returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasRefQosIndication() bool {
	if o != nil && !IsNil(o.RefQosIndication) {
		return true
	}

	return false
}

// SetRefQosIndication gets a reference to the given bool and assigns it to the RefQosIndication field.
func (o *SmPolicyContextData) SetRefQosIndication(v bool) {
	o.RefQosIndication = &v
}

// GetTraceReq returns the TraceReq field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmPolicyContextData) GetTraceReq() TraceData {
	if o == nil || IsNil(o.TraceReq.Get()) {
		var ret TraceData
		return ret
	}
	return *o.TraceReq.Get()
}

// GetTraceReqOk returns a tuple with the TraceReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmPolicyContextData) GetTraceReqOk() (*TraceData, bool) {
	if o == nil {
		return nil, false
	}
	return o.TraceReq.Get(), o.TraceReq.IsSet()
}

// HasTraceReq returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasTraceReq() bool {
	if o != nil && o.TraceReq.IsSet() {
		return true
	}

	return false
}

// SetTraceReq gets a reference to the given NullableTraceData and assigns it to the TraceReq field.
func (o *SmPolicyContextData) SetTraceReq(v TraceData) {
	o.TraceReq.Set(&v)
}
// SetTraceReqNil sets the value for TraceReq to be an explicit nil
func (o *SmPolicyContextData) SetTraceReqNil() {
	o.TraceReq.Set(nil)
}

// UnsetTraceReq ensures that no value is present for TraceReq, not even an explicit nil
func (o *SmPolicyContextData) UnsetTraceReq() {
	o.TraceReq.Unset()
}

// GetSliceInfo returns the SliceInfo field value
func (o *SmPolicyContextData) GetSliceInfo() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.SliceInfo
}

// GetSliceInfoOk returns a tuple with the SliceInfo field value
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetSliceInfoOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SliceInfo, true
}

// SetSliceInfo sets field value
func (o *SmPolicyContextData) SetSliceInfo(v Snssai) {
	o.SliceInfo = v
}

// GetQosFlowUsage returns the QosFlowUsage field value if set, zero value otherwise.
// func (o *SmPolicyContextData) GetQosFlowUsage() QosFlowUsage {
// 	if o == nil || IsNil(o.QosFlowUsage) {
// 		var ret QosFlowUsage
// 		return ret
// 	}
// 	return *o.QosFlowUsage
// }

// GetQosFlowUsageOk returns a tuple with the QosFlowUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// func (o *SmPolicyContextData) GetQosFlowUsageOk() (*QosFlowUsage, bool) {
// 	if o == nil || IsNil(o.QosFlowUsage) {
// 		return nil, false
// 	}
// 	return o.QosFlowUsage, true
// }

// HasQosFlowUsage returns a boolean if a field has been set.
// func (o *SmPolicyContextData) HasQosFlowUsage() bool {
// 	if o != nil && !IsNil(o.QosFlowUsage) {
// 		return true
// 	}

// 	return false
// }

// // SetQosFlowUsage gets a reference to the given QosFlowUsage and assigns it to the QosFlowUsage field.
// func (o *SmPolicyContextData) SetQosFlowUsage(v QosFlowUsage) {
// 	o.QosFlowUsage = &v
// }

// GetServNfId returns the ServNfId field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetServNfId() ServingNfIdentity {
	if o == nil || IsNil(o.ServNfId) {
		var ret ServingNfIdentity
		return ret
	}
	return *o.ServNfId
}

// GetServNfIdOk returns a tuple with the ServNfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetServNfIdOk() (*ServingNfIdentity, bool) {
	if o == nil || IsNil(o.ServNfId) {
		return nil, false
	}
	return o.ServNfId, true
}

// HasServNfId returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasServNfId() bool {
	if o != nil && !IsNil(o.ServNfId) {
		return true
	}

	return false
}

// SetServNfId gets a reference to the given ServingNfIdentity and assigns it to the ServNfId field.
func (o *SmPolicyContextData) SetServNfId(v ServingNfIdentity) {
	o.ServNfId = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *SmPolicyContextData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

// GetSmfId returns the SmfId field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetSmfId() string {
	if o == nil || IsNil(o.SmfId) {
		var ret string
		return ret
	}
	return *o.SmfId
}

// GetSmfIdOk returns a tuple with the SmfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetSmfIdOk() (*string, bool) {
	if o == nil || IsNil(o.SmfId) {
		return nil, false
	}
	return o.SmfId, true
}

// HasSmfId returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasSmfId() bool {
	if o != nil && !IsNil(o.SmfId) {
		return true
	}

	return false
}

// SetSmfId gets a reference to the given string and assigns it to the SmfId field.
func (o *SmPolicyContextData) SetSmfId(v string) {
	o.SmfId = &v
}

// GetRecoveryTime returns the RecoveryTime field value if set, zero value otherwise.
func (o *SmPolicyContextData) GetRecoveryTime() time.Time {
	if o == nil || IsNil(o.RecoveryTime) {
		var ret time.Time
		return ret
	}
	return *o.RecoveryTime
}

// GetRecoveryTimeOk returns a tuple with the RecoveryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyContextData) GetRecoveryTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RecoveryTime) {
		return nil, false
	}
	return o.RecoveryTime, true
}

// HasRecoveryTime returns a boolean if a field has been set.
func (o *SmPolicyContextData) HasRecoveryTime() bool {
	if o != nil && !IsNil(o.RecoveryTime) {
		return true
	}

	return false
}

// SetRecoveryTime gets a reference to the given time.Time and assigns it to the RecoveryTime field.
func (o *SmPolicyContextData) SetRecoveryTime(v time.Time) {
	o.RecoveryTime = &v
}

func (o SmPolicyContextData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmPolicyContextData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccNetChId) {
		toSerialize["accNetChId"] = o.AccNetChId
	}
	if o.ChargEntityAddr.IsSet() {
		toSerialize["chargEntityAddr"] = o.ChargEntityAddr.Get()
	}
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	toSerialize["supi"] = o.Supi
	if !IsNil(o.InvalidSupi) {
		toSerialize["invalidSupi"] = o.InvalidSupi
	}
	if !IsNil(o.InterGrpIds) {
		toSerialize["interGrpIds"] = o.InterGrpIds
	}
	toSerialize["pduSessionId"] = o.PduSessionId
	toSerialize["pduSessionType"] = o.PduSessionType
	if !IsNil(o.Chargingcharacteristics) {
		toSerialize["chargingcharacteristics"] = o.Chargingcharacteristics
	}
	toSerialize["dnn"] = o.Dnn
	toSerialize["notificationUri"] = o.NotificationUri
	if !IsNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !IsNil(o.RatType) {
		toSerialize["ratType"] = o.RatType
	}
	if !IsNil(o.ServingNetwork) {
		toSerialize["servingNetwork"] = o.ServingNetwork
	}
	if !IsNil(o.UserLocationInfo) {
		toSerialize["userLocationInfo"] = o.UserLocationInfo
	}
	if !IsNil(o.UeTimeZone) {
		toSerialize["ueTimeZone"] = o.UeTimeZone
	}
	if !IsNil(o.Pei) {
		toSerialize["pei"] = o.Pei
	}
	if !IsNil(o.Ipv4Address) {
		toSerialize["ipv4Address"] = o.Ipv4Address
	}
	if !IsNil(o.Ipv6AddressPrefix) {
		toSerialize["ipv6AddressPrefix"] = o.Ipv6AddressPrefix
	}
	if !IsNil(o.IpDomain) {
		toSerialize["ipDomain"] = o.IpDomain
	}
	if !IsNil(o.SubsSessAmbr) {
		toSerialize["subsSessAmbr"] = o.SubsSessAmbr
	}
	if !IsNil(o.SubsDefQos) {
		toSerialize["subsDefQos"] = o.SubsDefQos
	}
	if !IsNil(o.NumOfPackFilter) {
		toSerialize["numOfPackFilter"] = o.NumOfPackFilter
	}
	if !IsNil(o.Online) {
		toSerialize["online"] = o.Online
	}
	if !IsNil(o.Offline) {
		toSerialize["offline"] = o.Offline
	}
	if !IsNil(o.Var3gppPsDataOffStatus) {
		toSerialize["3gppPsDataOffStatus"] = o.Var3gppPsDataOffStatus
	}
	if !IsNil(o.RefQosIndication) {
		toSerialize["refQosIndication"] = o.RefQosIndication
	}
	if o.TraceReq.IsSet() {
		toSerialize["traceReq"] = o.TraceReq.Get()
	}
	toSerialize["sliceInfo"] = o.SliceInfo
	// if !IsNil(o.QosFlowUsage) {
	// 	toSerialize["qosFlowUsage"] = o.QosFlowUsage
	// }
	if !IsNil(o.ServNfId) {
		toSerialize["servNfId"] = o.ServNfId
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	if !IsNil(o.SmfId) {
		toSerialize["smfId"] = o.SmfId
	}
	if !IsNil(o.RecoveryTime) {
		toSerialize["recoveryTime"] = o.RecoveryTime
	}
	return toSerialize, nil
}

func (o *SmPolicyContextData) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"supi",
		"pduSessionId",
		"pduSessionType",
		"dnn",
		"notificationUri",
		"sliceInfo",
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

	varSmPolicyContextData := _SmPolicyContextData{}

	err = json.Unmarshal(bytes, &varSmPolicyContextData)

	if err != nil {
		return err
	}

	*o = SmPolicyContextData(varSmPolicyContextData)

	return err
}

type NullableSmPolicyContextData struct {
	value *SmPolicyContextData
	isSet bool
}

func (v NullableSmPolicyContextData) Get() *SmPolicyContextData {
	return v.value
}

func (v *NullableSmPolicyContextData) Set(val *SmPolicyContextData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmPolicyContextData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmPolicyContextData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmPolicyContextData(val *SmPolicyContextData) *NullableSmPolicyContextData {
	return &NullableSmPolicyContextData{value: val, isSet: true}
}

func (v NullableSmPolicyContextData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmPolicyContextData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


