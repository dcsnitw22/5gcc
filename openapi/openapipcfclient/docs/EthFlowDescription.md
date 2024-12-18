# EthFlowDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestMacAddr** | Pointer to **string** |  | [optional] 
**EthType** | **string** |  | 
**FDesc** | Pointer to **string** | Defines a packet filter of an IP flow. | [optional] 
**FDir** | Pointer to [**FlowDirection**](FlowDirection.md) |  | [optional] 
**SourceMacAddr** | Pointer to **string** |  | [optional] 
**VlanTags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEthFlowDescription

`func NewEthFlowDescription(ethType string, ) *EthFlowDescription`

NewEthFlowDescription instantiates a new EthFlowDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthFlowDescriptionWithDefaults

`func NewEthFlowDescriptionWithDefaults() *EthFlowDescription`

NewEthFlowDescriptionWithDefaults instantiates a new EthFlowDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestMacAddr

`func (o *EthFlowDescription) GetDestMacAddr() string`

GetDestMacAddr returns the DestMacAddr field if non-nil, zero value otherwise.

### GetDestMacAddrOk

`func (o *EthFlowDescription) GetDestMacAddrOk() (*string, bool)`

GetDestMacAddrOk returns a tuple with the DestMacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestMacAddr

`func (o *EthFlowDescription) SetDestMacAddr(v string)`

SetDestMacAddr sets DestMacAddr field to given value.

### HasDestMacAddr

`func (o *EthFlowDescription) HasDestMacAddr() bool`

HasDestMacAddr returns a boolean if a field has been set.

### GetEthType

`func (o *EthFlowDescription) GetEthType() string`

GetEthType returns the EthType field if non-nil, zero value otherwise.

### GetEthTypeOk

`func (o *EthFlowDescription) GetEthTypeOk() (*string, bool)`

GetEthTypeOk returns a tuple with the EthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthType

`func (o *EthFlowDescription) SetEthType(v string)`

SetEthType sets EthType field to given value.


### GetFDesc

`func (o *EthFlowDescription) GetFDesc() string`

GetFDesc returns the FDesc field if non-nil, zero value otherwise.

### GetFDescOk

`func (o *EthFlowDescription) GetFDescOk() (*string, bool)`

GetFDescOk returns a tuple with the FDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFDesc

`func (o *EthFlowDescription) SetFDesc(v string)`

SetFDesc sets FDesc field to given value.

### HasFDesc

`func (o *EthFlowDescription) HasFDesc() bool`

HasFDesc returns a boolean if a field has been set.

### GetFDir

`func (o *EthFlowDescription) GetFDir() FlowDirection`

GetFDir returns the FDir field if non-nil, zero value otherwise.

### GetFDirOk

`func (o *EthFlowDescription) GetFDirOk() (*FlowDirection, bool)`

GetFDirOk returns a tuple with the FDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFDir

`func (o *EthFlowDescription) SetFDir(v FlowDirection)`

SetFDir sets FDir field to given value.

### HasFDir

`func (o *EthFlowDescription) HasFDir() bool`

HasFDir returns a boolean if a field has been set.

### GetSourceMacAddr

`func (o *EthFlowDescription) GetSourceMacAddr() string`

GetSourceMacAddr returns the SourceMacAddr field if non-nil, zero value otherwise.

### GetSourceMacAddrOk

`func (o *EthFlowDescription) GetSourceMacAddrOk() (*string, bool)`

GetSourceMacAddrOk returns a tuple with the SourceMacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMacAddr

`func (o *EthFlowDescription) SetSourceMacAddr(v string)`

SetSourceMacAddr sets SourceMacAddr field to given value.

### HasSourceMacAddr

`func (o *EthFlowDescription) HasSourceMacAddr() bool`

HasSourceMacAddr returns a boolean if a field has been set.

### GetVlanTags

`func (o *EthFlowDescription) GetVlanTags() []string`

GetVlanTags returns the VlanTags field if non-nil, zero value otherwise.

### GetVlanTagsOk

`func (o *EthFlowDescription) GetVlanTagsOk() (*[]string, bool)`

GetVlanTagsOk returns a tuple with the VlanTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTags

`func (o *EthFlowDescription) SetVlanTags(v []string)`

SetVlanTags sets VlanTags field to given value.

### HasVlanTags

`func (o *EthFlowDescription) HasVlanTags() bool`

HasVlanTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


