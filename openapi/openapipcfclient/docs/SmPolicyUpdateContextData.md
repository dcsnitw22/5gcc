# SmPolicyUpdateContextData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepPolicyCtrlReqTriggers** | Pointer to [**[]PolicyControlRequestTrigger**](PolicyControlRequestTrigger.md) | The policy control reqeust trigges which are met. | [optional] 
**AccNetChIds** | Pointer to [**[]AccNetChId**](AccNetChId.md) | Indicates the access network charging identifier for the PCC rule(s) or whole PDU session. | [optional] 
**AccessType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**ServingNetwork** | Pointer to [**NetworkId**](NetworkId.md) |  | [optional] 
**UserLocationInfo** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UeTimeZone** | Pointer to **string** |  | [optional] 
**RelIpv4Address** | Pointer to **string** |  | [optional] 
**Ipv4Address** | Pointer to **string** |  | [optional] 
**IpDomain** | Pointer to **string** | Indicates the IPv4 address domain | [optional] 
**Ipv6AddressPrefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**RelIpv6AddressPrefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**RelUeMac** | Pointer to **string** |  | [optional] 
**UeMac** | Pointer to **string** |  | [optional] 
**SubsSessAmbr** | Pointer to [**Ambr**](Ambr.md) |  | [optional] 
**SubsDefQos** | Pointer to [**SubscribedDefaultQos**](SubscribedDefaultQos.md) |  | [optional] 
**NumOfPackFilter** | Pointer to **int32** | Contains the number of supported packet filter for signalled QoS rules. | [optional] 
**AccuUsageReports** | Pointer to [**[]AccuUsageReport**](AccuUsageReport.md) | Contains the usage report | [optional] 
**Var3gppPsDataOffStatus** | Pointer to **bool** | If it is included and set to true, the 3GPP PS Data Off is activated by the UE. | [optional] 
**AppDetectionInfos** | Pointer to [**[]AppDetectionInfo**](AppDetectionInfo.md) | Report the start/stop of the application traffic and detected SDF descriptions if applicable. | [optional] 
**RuleReports** | Pointer to [**[]RuleReport**](RuleReport.md) | Used to report the PCC rule failure. | [optional] 
**SessRuleReports** | Pointer to [**[]SessionRuleReport**](SessionRuleReport.md) | Used to report the session rule failure. | [optional] 
**QncReports** | Pointer to [**[]QosNotificationControlInfo**](QosNotificationControlInfo.md) | QoS Notification Control information. | [optional] 
**UserLocationInfoTime** | Pointer to **time.Time** |  | [optional] 
**RepPraInfos** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) | Reports the changes of presence reporting area. | [optional] 
**UeInitResReq** | Pointer to [**UeInitiatedResourceRequest**](UeInitiatedResourceRequest.md) |  | [optional] 
**RefQosIndication** | Pointer to **bool** | If it is included and set to true, the reflective QoS is supported by the UE. If it is included and set to false, the reflective QoS is revoked by the UE. | [optional] 
**QosFlowUsage** | Pointer to [**QosFlowUsage**](QosFlowUsage.md) |  | [optional] 
**CreditManageStatus** | Pointer to [**CreditManagementStatus**](CreditManagementStatus.md) |  | [optional] 
**ServNfId** | Pointer to [**ServingNfIdentity**](ServingNfIdentity.md) |  | [optional] 
**TraceReq** | Pointer to [**NullableTraceData**](TraceData.md) |  | [optional] 

## Methods

### NewSmPolicyUpdateContextData

`func NewSmPolicyUpdateContextData() *SmPolicyUpdateContextData`

NewSmPolicyUpdateContextData instantiates a new SmPolicyUpdateContextData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmPolicyUpdateContextDataWithDefaults

`func NewSmPolicyUpdateContextDataWithDefaults() *SmPolicyUpdateContextData`

NewSmPolicyUpdateContextDataWithDefaults instantiates a new SmPolicyUpdateContextData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepPolicyCtrlReqTriggers

`func (o *SmPolicyUpdateContextData) GetRepPolicyCtrlReqTriggers() []PolicyControlRequestTrigger`

GetRepPolicyCtrlReqTriggers returns the RepPolicyCtrlReqTriggers field if non-nil, zero value otherwise.

### GetRepPolicyCtrlReqTriggersOk

`func (o *SmPolicyUpdateContextData) GetRepPolicyCtrlReqTriggersOk() (*[]PolicyControlRequestTrigger, bool)`

GetRepPolicyCtrlReqTriggersOk returns a tuple with the RepPolicyCtrlReqTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPolicyCtrlReqTriggers

`func (o *SmPolicyUpdateContextData) SetRepPolicyCtrlReqTriggers(v []PolicyControlRequestTrigger)`

SetRepPolicyCtrlReqTriggers sets RepPolicyCtrlReqTriggers field to given value.

### HasRepPolicyCtrlReqTriggers

`func (o *SmPolicyUpdateContextData) HasRepPolicyCtrlReqTriggers() bool`

HasRepPolicyCtrlReqTriggers returns a boolean if a field has been set.

### GetAccNetChIds

`func (o *SmPolicyUpdateContextData) GetAccNetChIds() []AccNetChId`

GetAccNetChIds returns the AccNetChIds field if non-nil, zero value otherwise.

### GetAccNetChIdsOk

`func (o *SmPolicyUpdateContextData) GetAccNetChIdsOk() (*[]AccNetChId, bool)`

GetAccNetChIdsOk returns a tuple with the AccNetChIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccNetChIds

`func (o *SmPolicyUpdateContextData) SetAccNetChIds(v []AccNetChId)`

SetAccNetChIds sets AccNetChIds field to given value.

### HasAccNetChIds

`func (o *SmPolicyUpdateContextData) HasAccNetChIds() bool`

HasAccNetChIds returns a boolean if a field has been set.

### GetAccessType

`func (o *SmPolicyUpdateContextData) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *SmPolicyUpdateContextData) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *SmPolicyUpdateContextData) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *SmPolicyUpdateContextData) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetRatType

`func (o *SmPolicyUpdateContextData) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *SmPolicyUpdateContextData) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *SmPolicyUpdateContextData) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *SmPolicyUpdateContextData) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetServingNetwork

`func (o *SmPolicyUpdateContextData) GetServingNetwork() NetworkId`

GetServingNetwork returns the ServingNetwork field if non-nil, zero value otherwise.

### GetServingNetworkOk

`func (o *SmPolicyUpdateContextData) GetServingNetworkOk() (*NetworkId, bool)`

GetServingNetworkOk returns a tuple with the ServingNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetwork

`func (o *SmPolicyUpdateContextData) SetServingNetwork(v NetworkId)`

SetServingNetwork sets ServingNetwork field to given value.

### HasServingNetwork

`func (o *SmPolicyUpdateContextData) HasServingNetwork() bool`

HasServingNetwork returns a boolean if a field has been set.

### GetUserLocationInfo

`func (o *SmPolicyUpdateContextData) GetUserLocationInfo() UserLocation`

GetUserLocationInfo returns the UserLocationInfo field if non-nil, zero value otherwise.

### GetUserLocationInfoOk

`func (o *SmPolicyUpdateContextData) GetUserLocationInfoOk() (*UserLocation, bool)`

GetUserLocationInfoOk returns a tuple with the UserLocationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocationInfo

`func (o *SmPolicyUpdateContextData) SetUserLocationInfo(v UserLocation)`

SetUserLocationInfo sets UserLocationInfo field to given value.

### HasUserLocationInfo

`func (o *SmPolicyUpdateContextData) HasUserLocationInfo() bool`

HasUserLocationInfo returns a boolean if a field has been set.

### GetUeTimeZone

`func (o *SmPolicyUpdateContextData) GetUeTimeZone() string`

GetUeTimeZone returns the UeTimeZone field if non-nil, zero value otherwise.

### GetUeTimeZoneOk

`func (o *SmPolicyUpdateContextData) GetUeTimeZoneOk() (*string, bool)`

GetUeTimeZoneOk returns a tuple with the UeTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTimeZone

`func (o *SmPolicyUpdateContextData) SetUeTimeZone(v string)`

SetUeTimeZone sets UeTimeZone field to given value.

### HasUeTimeZone

`func (o *SmPolicyUpdateContextData) HasUeTimeZone() bool`

HasUeTimeZone returns a boolean if a field has been set.

### GetRelIpv4Address

`func (o *SmPolicyUpdateContextData) GetRelIpv4Address() string`

GetRelIpv4Address returns the RelIpv4Address field if non-nil, zero value otherwise.

### GetRelIpv4AddressOk

`func (o *SmPolicyUpdateContextData) GetRelIpv4AddressOk() (*string, bool)`

GetRelIpv4AddressOk returns a tuple with the RelIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelIpv4Address

`func (o *SmPolicyUpdateContextData) SetRelIpv4Address(v string)`

SetRelIpv4Address sets RelIpv4Address field to given value.

### HasRelIpv4Address

`func (o *SmPolicyUpdateContextData) HasRelIpv4Address() bool`

HasRelIpv4Address returns a boolean if a field has been set.

### GetIpv4Address

`func (o *SmPolicyUpdateContextData) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *SmPolicyUpdateContextData) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *SmPolicyUpdateContextData) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *SmPolicyUpdateContextData) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpDomain

`func (o *SmPolicyUpdateContextData) GetIpDomain() string`

GetIpDomain returns the IpDomain field if non-nil, zero value otherwise.

### GetIpDomainOk

`func (o *SmPolicyUpdateContextData) GetIpDomainOk() (*string, bool)`

GetIpDomainOk returns a tuple with the IpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomain

`func (o *SmPolicyUpdateContextData) SetIpDomain(v string)`

SetIpDomain sets IpDomain field to given value.

### HasIpDomain

`func (o *SmPolicyUpdateContextData) HasIpDomain() bool`

HasIpDomain returns a boolean if a field has been set.

### GetIpv6AddressPrefix

`func (o *SmPolicyUpdateContextData) GetIpv6AddressPrefix() Ipv6Prefix`

GetIpv6AddressPrefix returns the Ipv6AddressPrefix field if non-nil, zero value otherwise.

### GetIpv6AddressPrefixOk

`func (o *SmPolicyUpdateContextData) GetIpv6AddressPrefixOk() (*Ipv6Prefix, bool)`

GetIpv6AddressPrefixOk returns a tuple with the Ipv6AddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6AddressPrefix

`func (o *SmPolicyUpdateContextData) SetIpv6AddressPrefix(v Ipv6Prefix)`

SetIpv6AddressPrefix sets Ipv6AddressPrefix field to given value.

### HasIpv6AddressPrefix

`func (o *SmPolicyUpdateContextData) HasIpv6AddressPrefix() bool`

HasIpv6AddressPrefix returns a boolean if a field has been set.

### GetRelIpv6AddressPrefix

`func (o *SmPolicyUpdateContextData) GetRelIpv6AddressPrefix() Ipv6Prefix`

GetRelIpv6AddressPrefix returns the RelIpv6AddressPrefix field if non-nil, zero value otherwise.

### GetRelIpv6AddressPrefixOk

`func (o *SmPolicyUpdateContextData) GetRelIpv6AddressPrefixOk() (*Ipv6Prefix, bool)`

GetRelIpv6AddressPrefixOk returns a tuple with the RelIpv6AddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelIpv6AddressPrefix

`func (o *SmPolicyUpdateContextData) SetRelIpv6AddressPrefix(v Ipv6Prefix)`

SetRelIpv6AddressPrefix sets RelIpv6AddressPrefix field to given value.

### HasRelIpv6AddressPrefix

`func (o *SmPolicyUpdateContextData) HasRelIpv6AddressPrefix() bool`

HasRelIpv6AddressPrefix returns a boolean if a field has been set.

### GetRelUeMac

`func (o *SmPolicyUpdateContextData) GetRelUeMac() string`

GetRelUeMac returns the RelUeMac field if non-nil, zero value otherwise.

### GetRelUeMacOk

`func (o *SmPolicyUpdateContextData) GetRelUeMacOk() (*string, bool)`

GetRelUeMacOk returns a tuple with the RelUeMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelUeMac

`func (o *SmPolicyUpdateContextData) SetRelUeMac(v string)`

SetRelUeMac sets RelUeMac field to given value.

### HasRelUeMac

`func (o *SmPolicyUpdateContextData) HasRelUeMac() bool`

HasRelUeMac returns a boolean if a field has been set.

### GetUeMac

`func (o *SmPolicyUpdateContextData) GetUeMac() string`

GetUeMac returns the UeMac field if non-nil, zero value otherwise.

### GetUeMacOk

`func (o *SmPolicyUpdateContextData) GetUeMacOk() (*string, bool)`

GetUeMacOk returns a tuple with the UeMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMac

`func (o *SmPolicyUpdateContextData) SetUeMac(v string)`

SetUeMac sets UeMac field to given value.

### HasUeMac

`func (o *SmPolicyUpdateContextData) HasUeMac() bool`

HasUeMac returns a boolean if a field has been set.

### GetSubsSessAmbr

`func (o *SmPolicyUpdateContextData) GetSubsSessAmbr() Ambr`

GetSubsSessAmbr returns the SubsSessAmbr field if non-nil, zero value otherwise.

### GetSubsSessAmbrOk

`func (o *SmPolicyUpdateContextData) GetSubsSessAmbrOk() (*Ambr, bool)`

GetSubsSessAmbrOk returns a tuple with the SubsSessAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsSessAmbr

`func (o *SmPolicyUpdateContextData) SetSubsSessAmbr(v Ambr)`

SetSubsSessAmbr sets SubsSessAmbr field to given value.

### HasSubsSessAmbr

`func (o *SmPolicyUpdateContextData) HasSubsSessAmbr() bool`

HasSubsSessAmbr returns a boolean if a field has been set.

### GetSubsDefQos

`func (o *SmPolicyUpdateContextData) GetSubsDefQos() SubscribedDefaultQos`

GetSubsDefQos returns the SubsDefQos field if non-nil, zero value otherwise.

### GetSubsDefQosOk

`func (o *SmPolicyUpdateContextData) GetSubsDefQosOk() (*SubscribedDefaultQos, bool)`

GetSubsDefQosOk returns a tuple with the SubsDefQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsDefQos

`func (o *SmPolicyUpdateContextData) SetSubsDefQos(v SubscribedDefaultQos)`

SetSubsDefQos sets SubsDefQos field to given value.

### HasSubsDefQos

`func (o *SmPolicyUpdateContextData) HasSubsDefQos() bool`

HasSubsDefQos returns a boolean if a field has been set.

### GetNumOfPackFilter

`func (o *SmPolicyUpdateContextData) GetNumOfPackFilter() int32`

GetNumOfPackFilter returns the NumOfPackFilter field if non-nil, zero value otherwise.

### GetNumOfPackFilterOk

`func (o *SmPolicyUpdateContextData) GetNumOfPackFilterOk() (*int32, bool)`

GetNumOfPackFilterOk returns a tuple with the NumOfPackFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfPackFilter

`func (o *SmPolicyUpdateContextData) SetNumOfPackFilter(v int32)`

SetNumOfPackFilter sets NumOfPackFilter field to given value.

### HasNumOfPackFilter

`func (o *SmPolicyUpdateContextData) HasNumOfPackFilter() bool`

HasNumOfPackFilter returns a boolean if a field has been set.

### GetAccuUsageReports

`func (o *SmPolicyUpdateContextData) GetAccuUsageReports() []AccuUsageReport`

GetAccuUsageReports returns the AccuUsageReports field if non-nil, zero value otherwise.

### GetAccuUsageReportsOk

`func (o *SmPolicyUpdateContextData) GetAccuUsageReportsOk() (*[]AccuUsageReport, bool)`

GetAccuUsageReportsOk returns a tuple with the AccuUsageReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuUsageReports

`func (o *SmPolicyUpdateContextData) SetAccuUsageReports(v []AccuUsageReport)`

SetAccuUsageReports sets AccuUsageReports field to given value.

### HasAccuUsageReports

`func (o *SmPolicyUpdateContextData) HasAccuUsageReports() bool`

HasAccuUsageReports returns a boolean if a field has been set.

### GetVar3gppPsDataOffStatus

`func (o *SmPolicyUpdateContextData) GetVar3gppPsDataOffStatus() bool`

GetVar3gppPsDataOffStatus returns the Var3gppPsDataOffStatus field if non-nil, zero value otherwise.

### GetVar3gppPsDataOffStatusOk

`func (o *SmPolicyUpdateContextData) GetVar3gppPsDataOffStatusOk() (*bool, bool)`

GetVar3gppPsDataOffStatusOk returns a tuple with the Var3gppPsDataOffStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3gppPsDataOffStatus

`func (o *SmPolicyUpdateContextData) SetVar3gppPsDataOffStatus(v bool)`

SetVar3gppPsDataOffStatus sets Var3gppPsDataOffStatus field to given value.

### HasVar3gppPsDataOffStatus

`func (o *SmPolicyUpdateContextData) HasVar3gppPsDataOffStatus() bool`

HasVar3gppPsDataOffStatus returns a boolean if a field has been set.

### GetAppDetectionInfos

`func (o *SmPolicyUpdateContextData) GetAppDetectionInfos() []AppDetectionInfo`

GetAppDetectionInfos returns the AppDetectionInfos field if non-nil, zero value otherwise.

### GetAppDetectionInfosOk

`func (o *SmPolicyUpdateContextData) GetAppDetectionInfosOk() (*[]AppDetectionInfo, bool)`

GetAppDetectionInfosOk returns a tuple with the AppDetectionInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDetectionInfos

`func (o *SmPolicyUpdateContextData) SetAppDetectionInfos(v []AppDetectionInfo)`

SetAppDetectionInfos sets AppDetectionInfos field to given value.

### HasAppDetectionInfos

`func (o *SmPolicyUpdateContextData) HasAppDetectionInfos() bool`

HasAppDetectionInfos returns a boolean if a field has been set.

### GetRuleReports

`func (o *SmPolicyUpdateContextData) GetRuleReports() []RuleReport`

GetRuleReports returns the RuleReports field if non-nil, zero value otherwise.

### GetRuleReportsOk

`func (o *SmPolicyUpdateContextData) GetRuleReportsOk() (*[]RuleReport, bool)`

GetRuleReportsOk returns a tuple with the RuleReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleReports

`func (o *SmPolicyUpdateContextData) SetRuleReports(v []RuleReport)`

SetRuleReports sets RuleReports field to given value.

### HasRuleReports

`func (o *SmPolicyUpdateContextData) HasRuleReports() bool`

HasRuleReports returns a boolean if a field has been set.

### GetSessRuleReports

`func (o *SmPolicyUpdateContextData) GetSessRuleReports() []SessionRuleReport`

GetSessRuleReports returns the SessRuleReports field if non-nil, zero value otherwise.

### GetSessRuleReportsOk

`func (o *SmPolicyUpdateContextData) GetSessRuleReportsOk() (*[]SessionRuleReport, bool)`

GetSessRuleReportsOk returns a tuple with the SessRuleReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessRuleReports

`func (o *SmPolicyUpdateContextData) SetSessRuleReports(v []SessionRuleReport)`

SetSessRuleReports sets SessRuleReports field to given value.

### HasSessRuleReports

`func (o *SmPolicyUpdateContextData) HasSessRuleReports() bool`

HasSessRuleReports returns a boolean if a field has been set.

### GetQncReports

`func (o *SmPolicyUpdateContextData) GetQncReports() []QosNotificationControlInfo`

GetQncReports returns the QncReports field if non-nil, zero value otherwise.

### GetQncReportsOk

`func (o *SmPolicyUpdateContextData) GetQncReportsOk() (*[]QosNotificationControlInfo, bool)`

GetQncReportsOk returns a tuple with the QncReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQncReports

`func (o *SmPolicyUpdateContextData) SetQncReports(v []QosNotificationControlInfo)`

SetQncReports sets QncReports field to given value.

### HasQncReports

`func (o *SmPolicyUpdateContextData) HasQncReports() bool`

HasQncReports returns a boolean if a field has been set.

### GetUserLocationInfoTime

`func (o *SmPolicyUpdateContextData) GetUserLocationInfoTime() time.Time`

GetUserLocationInfoTime returns the UserLocationInfoTime field if non-nil, zero value otherwise.

### GetUserLocationInfoTimeOk

`func (o *SmPolicyUpdateContextData) GetUserLocationInfoTimeOk() (*time.Time, bool)`

GetUserLocationInfoTimeOk returns a tuple with the UserLocationInfoTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocationInfoTime

`func (o *SmPolicyUpdateContextData) SetUserLocationInfoTime(v time.Time)`

SetUserLocationInfoTime sets UserLocationInfoTime field to given value.

### HasUserLocationInfoTime

`func (o *SmPolicyUpdateContextData) HasUserLocationInfoTime() bool`

HasUserLocationInfoTime returns a boolean if a field has been set.

### GetRepPraInfos

`func (o *SmPolicyUpdateContextData) GetRepPraInfos() map[string]PresenceInfo`

GetRepPraInfos returns the RepPraInfos field if non-nil, zero value otherwise.

### GetRepPraInfosOk

`func (o *SmPolicyUpdateContextData) GetRepPraInfosOk() (*map[string]PresenceInfo, bool)`

GetRepPraInfosOk returns a tuple with the RepPraInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPraInfos

`func (o *SmPolicyUpdateContextData) SetRepPraInfos(v map[string]PresenceInfo)`

SetRepPraInfos sets RepPraInfos field to given value.

### HasRepPraInfos

`func (o *SmPolicyUpdateContextData) HasRepPraInfos() bool`

HasRepPraInfos returns a boolean if a field has been set.

### GetUeInitResReq

`func (o *SmPolicyUpdateContextData) GetUeInitResReq() UeInitiatedResourceRequest`

GetUeInitResReq returns the UeInitResReq field if non-nil, zero value otherwise.

### GetUeInitResReqOk

`func (o *SmPolicyUpdateContextData) GetUeInitResReqOk() (*UeInitiatedResourceRequest, bool)`

GetUeInitResReqOk returns a tuple with the UeInitResReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeInitResReq

`func (o *SmPolicyUpdateContextData) SetUeInitResReq(v UeInitiatedResourceRequest)`

SetUeInitResReq sets UeInitResReq field to given value.

### HasUeInitResReq

`func (o *SmPolicyUpdateContextData) HasUeInitResReq() bool`

HasUeInitResReq returns a boolean if a field has been set.

### GetRefQosIndication

`func (o *SmPolicyUpdateContextData) GetRefQosIndication() bool`

GetRefQosIndication returns the RefQosIndication field if non-nil, zero value otherwise.

### GetRefQosIndicationOk

`func (o *SmPolicyUpdateContextData) GetRefQosIndicationOk() (*bool, bool)`

GetRefQosIndicationOk returns a tuple with the RefQosIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefQosIndication

`func (o *SmPolicyUpdateContextData) SetRefQosIndication(v bool)`

SetRefQosIndication sets RefQosIndication field to given value.

### HasRefQosIndication

`func (o *SmPolicyUpdateContextData) HasRefQosIndication() bool`

HasRefQosIndication returns a boolean if a field has been set.

### GetQosFlowUsage

`func (o *SmPolicyUpdateContextData) GetQosFlowUsage() QosFlowUsage`

GetQosFlowUsage returns the QosFlowUsage field if non-nil, zero value otherwise.

### GetQosFlowUsageOk

`func (o *SmPolicyUpdateContextData) GetQosFlowUsageOk() (*QosFlowUsage, bool)`

GetQosFlowUsageOk returns a tuple with the QosFlowUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowUsage

`func (o *SmPolicyUpdateContextData) SetQosFlowUsage(v QosFlowUsage)`

SetQosFlowUsage sets QosFlowUsage field to given value.

### HasQosFlowUsage

`func (o *SmPolicyUpdateContextData) HasQosFlowUsage() bool`

HasQosFlowUsage returns a boolean if a field has been set.

### GetCreditManageStatus

`func (o *SmPolicyUpdateContextData) GetCreditManageStatus() CreditManagementStatus`

GetCreditManageStatus returns the CreditManageStatus field if non-nil, zero value otherwise.

### GetCreditManageStatusOk

`func (o *SmPolicyUpdateContextData) GetCreditManageStatusOk() (*CreditManagementStatus, bool)`

GetCreditManageStatusOk returns a tuple with the CreditManageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditManageStatus

`func (o *SmPolicyUpdateContextData) SetCreditManageStatus(v CreditManagementStatus)`

SetCreditManageStatus sets CreditManageStatus field to given value.

### HasCreditManageStatus

`func (o *SmPolicyUpdateContextData) HasCreditManageStatus() bool`

HasCreditManageStatus returns a boolean if a field has been set.

### GetServNfId

`func (o *SmPolicyUpdateContextData) GetServNfId() ServingNfIdentity`

GetServNfId returns the ServNfId field if non-nil, zero value otherwise.

### GetServNfIdOk

`func (o *SmPolicyUpdateContextData) GetServNfIdOk() (*ServingNfIdentity, bool)`

GetServNfIdOk returns a tuple with the ServNfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServNfId

`func (o *SmPolicyUpdateContextData) SetServNfId(v ServingNfIdentity)`

SetServNfId sets ServNfId field to given value.

### HasServNfId

`func (o *SmPolicyUpdateContextData) HasServNfId() bool`

HasServNfId returns a boolean if a field has been set.

### GetTraceReq

`func (o *SmPolicyUpdateContextData) GetTraceReq() TraceData`

GetTraceReq returns the TraceReq field if non-nil, zero value otherwise.

### GetTraceReqOk

`func (o *SmPolicyUpdateContextData) GetTraceReqOk() (*TraceData, bool)`

GetTraceReqOk returns a tuple with the TraceReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceReq

`func (o *SmPolicyUpdateContextData) SetTraceReq(v TraceData)`

SetTraceReq sets TraceReq field to given value.

### HasTraceReq

`func (o *SmPolicyUpdateContextData) HasTraceReq() bool`

HasTraceReq returns a boolean if a field has been set.

### SetTraceReqNil

`func (o *SmPolicyUpdateContextData) SetTraceReqNil(b bool)`

 SetTraceReqNil sets the value for TraceReq to be an explicit nil

### UnsetTraceReq
`func (o *SmPolicyUpdateContextData) UnsetTraceReq()`

UnsetTraceReq ensures that no value is present for TraceReq, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


