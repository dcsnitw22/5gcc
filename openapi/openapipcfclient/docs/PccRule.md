# PccRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowInfos** | Pointer to [**[]FlowInformation**](FlowInformation.md) | An array of IP flow packet filter information. | [optional] 
**AppId** | Pointer to **string** | A reference to the application detection filter configured at the UPF. | [optional] 
**ContVer** | Pointer to **int32** | Represents the content version of some content. | [optional] 
**PccRuleId** | **string** | Univocally identifies the PCC rule within a PDU session. | 
**Precedence** | Pointer to **int32** |  | [optional] 
**AfSigProtocol** | Pointer to [**NullableAfSigProtocol**](AfSigProtocol.md) |  | [optional] 
**AppReloc** | Pointer to **bool** | Indication of application relocation possibility. | [optional] 
**RefQosData** | Pointer to **[]string** | A reference to the QosData policy decision type. It is the qosId described in subclause 5.6.2.8. | [optional] 
**RefTcData** | Pointer to **[]string** | A reference to the TrafficControlData policy decision type. It is the tcId described in subclause 5.6.2.10. | [optional] 
**RefChgData** | Pointer to **[]string** | A reference to the ChargingData policy decision type. It is the chgId described in subclause 5.6.2.11. | [optional] 
**RefUmData** | Pointer to **[]string** | A reference to UsageMonitoringData policy decision type. It is the umId described in subclause 5.6.2.12. | [optional] 
**RefCondData** | Pointer to **NullableString** | A reference to the condition data. It is the condId described in subclause 5.6.2.9. | [optional] 

## Methods

### NewPccRule

`func NewPccRule(pccRuleId string, ) *PccRule`

NewPccRule instantiates a new PccRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPccRuleWithDefaults

`func NewPccRuleWithDefaults() *PccRule`

NewPccRuleWithDefaults instantiates a new PccRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowInfos

`func (o *PccRule) GetFlowInfos() []FlowInformation`

GetFlowInfos returns the FlowInfos field if non-nil, zero value otherwise.

### GetFlowInfosOk

`func (o *PccRule) GetFlowInfosOk() (*[]FlowInformation, bool)`

GetFlowInfosOk returns a tuple with the FlowInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfos

`func (o *PccRule) SetFlowInfos(v []FlowInformation)`

SetFlowInfos sets FlowInfos field to given value.

### HasFlowInfos

`func (o *PccRule) HasFlowInfos() bool`

HasFlowInfos returns a boolean if a field has been set.

### GetAppId

`func (o *PccRule) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *PccRule) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *PccRule) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *PccRule) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetContVer

`func (o *PccRule) GetContVer() int32`

GetContVer returns the ContVer field if non-nil, zero value otherwise.

### GetContVerOk

`func (o *PccRule) GetContVerOk() (*int32, bool)`

GetContVerOk returns a tuple with the ContVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContVer

`func (o *PccRule) SetContVer(v int32)`

SetContVer sets ContVer field to given value.

### HasContVer

`func (o *PccRule) HasContVer() bool`

HasContVer returns a boolean if a field has been set.

### GetPccRuleId

`func (o *PccRule) GetPccRuleId() string`

GetPccRuleId returns the PccRuleId field if non-nil, zero value otherwise.

### GetPccRuleIdOk

`func (o *PccRule) GetPccRuleIdOk() (*string, bool)`

GetPccRuleIdOk returns a tuple with the PccRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPccRuleId

`func (o *PccRule) SetPccRuleId(v string)`

SetPccRuleId sets PccRuleId field to given value.


### GetPrecedence

`func (o *PccRule) GetPrecedence() int32`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *PccRule) GetPrecedenceOk() (*int32, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *PccRule) SetPrecedence(v int32)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *PccRule) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetAfSigProtocol

`func (o *PccRule) GetAfSigProtocol() AfSigProtocol`

GetAfSigProtocol returns the AfSigProtocol field if non-nil, zero value otherwise.

### GetAfSigProtocolOk

`func (o *PccRule) GetAfSigProtocolOk() (*AfSigProtocol, bool)`

GetAfSigProtocolOk returns a tuple with the AfSigProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfSigProtocol

`func (o *PccRule) SetAfSigProtocol(v AfSigProtocol)`

SetAfSigProtocol sets AfSigProtocol field to given value.

### HasAfSigProtocol

`func (o *PccRule) HasAfSigProtocol() bool`

HasAfSigProtocol returns a boolean if a field has been set.

### SetAfSigProtocolNil

`func (o *PccRule) SetAfSigProtocolNil(b bool)`

 SetAfSigProtocolNil sets the value for AfSigProtocol to be an explicit nil

### UnsetAfSigProtocol
`func (o *PccRule) UnsetAfSigProtocol()`

UnsetAfSigProtocol ensures that no value is present for AfSigProtocol, not even an explicit nil
### GetAppReloc

`func (o *PccRule) GetAppReloc() bool`

GetAppReloc returns the AppReloc field if non-nil, zero value otherwise.

### GetAppRelocOk

`func (o *PccRule) GetAppRelocOk() (*bool, bool)`

GetAppRelocOk returns a tuple with the AppReloc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppReloc

`func (o *PccRule) SetAppReloc(v bool)`

SetAppReloc sets AppReloc field to given value.

### HasAppReloc

`func (o *PccRule) HasAppReloc() bool`

HasAppReloc returns a boolean if a field has been set.

### GetRefQosData

`func (o *PccRule) GetRefQosData() []string`

GetRefQosData returns the RefQosData field if non-nil, zero value otherwise.

### GetRefQosDataOk

`func (o *PccRule) GetRefQosDataOk() (*[]string, bool)`

GetRefQosDataOk returns a tuple with the RefQosData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefQosData

`func (o *PccRule) SetRefQosData(v []string)`

SetRefQosData sets RefQosData field to given value.

### HasRefQosData

`func (o *PccRule) HasRefQosData() bool`

HasRefQosData returns a boolean if a field has been set.

### GetRefTcData

`func (o *PccRule) GetRefTcData() []string`

GetRefTcData returns the RefTcData field if non-nil, zero value otherwise.

### GetRefTcDataOk

`func (o *PccRule) GetRefTcDataOk() (*[]string, bool)`

GetRefTcDataOk returns a tuple with the RefTcData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefTcData

`func (o *PccRule) SetRefTcData(v []string)`

SetRefTcData sets RefTcData field to given value.

### HasRefTcData

`func (o *PccRule) HasRefTcData() bool`

HasRefTcData returns a boolean if a field has been set.

### GetRefChgData

`func (o *PccRule) GetRefChgData() []string`

GetRefChgData returns the RefChgData field if non-nil, zero value otherwise.

### GetRefChgDataOk

`func (o *PccRule) GetRefChgDataOk() (*[]string, bool)`

GetRefChgDataOk returns a tuple with the RefChgData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefChgData

`func (o *PccRule) SetRefChgData(v []string)`

SetRefChgData sets RefChgData field to given value.

### HasRefChgData

`func (o *PccRule) HasRefChgData() bool`

HasRefChgData returns a boolean if a field has been set.

### SetRefChgDataNil

`func (o *PccRule) SetRefChgDataNil(b bool)`

 SetRefChgDataNil sets the value for RefChgData to be an explicit nil

### UnsetRefChgData
`func (o *PccRule) UnsetRefChgData()`

UnsetRefChgData ensures that no value is present for RefChgData, not even an explicit nil
### GetRefUmData

`func (o *PccRule) GetRefUmData() []string`

GetRefUmData returns the RefUmData field if non-nil, zero value otherwise.

### GetRefUmDataOk

`func (o *PccRule) GetRefUmDataOk() (*[]string, bool)`

GetRefUmDataOk returns a tuple with the RefUmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUmData

`func (o *PccRule) SetRefUmData(v []string)`

SetRefUmData sets RefUmData field to given value.

### HasRefUmData

`func (o *PccRule) HasRefUmData() bool`

HasRefUmData returns a boolean if a field has been set.

### SetRefUmDataNil

`func (o *PccRule) SetRefUmDataNil(b bool)`

 SetRefUmDataNil sets the value for RefUmData to be an explicit nil

### UnsetRefUmData
`func (o *PccRule) UnsetRefUmData()`

UnsetRefUmData ensures that no value is present for RefUmData, not even an explicit nil
### GetRefCondData

`func (o *PccRule) GetRefCondData() string`

GetRefCondData returns the RefCondData field if non-nil, zero value otherwise.

### GetRefCondDataOk

`func (o *PccRule) GetRefCondDataOk() (*string, bool)`

GetRefCondDataOk returns a tuple with the RefCondData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCondData

`func (o *PccRule) SetRefCondData(v string)`

SetRefCondData sets RefCondData field to given value.

### HasRefCondData

`func (o *PccRule) HasRefCondData() bool`

HasRefCondData returns a boolean if a field has been set.

### SetRefCondDataNil

`func (o *PccRule) SetRefCondDataNil(b bool)`

 SetRefCondDataNil sets the value for RefCondData to be an explicit nil

### UnsetRefCondData
`func (o *PccRule) UnsetRefCondData()`

UnsetRefCondData ensures that no value is present for RefCondData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


