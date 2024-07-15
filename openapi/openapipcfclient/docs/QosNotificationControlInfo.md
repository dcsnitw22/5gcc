# QosNotificationControlInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefPccRuleIds** | **[]string** | An array of PCC rule id references to the PCC rules associated with the QoS notification control info. | 
**NotifType** | [**QosNotifType**](QosNotifType.md) |  | 
**ContVer** | Pointer to **int32** | Represents the content version of some content. | [optional] 

## Methods

### NewQosNotificationControlInfo

`func NewQosNotificationControlInfo(refPccRuleIds []string, notifType QosNotifType, ) *QosNotificationControlInfo`

NewQosNotificationControlInfo instantiates a new QosNotificationControlInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosNotificationControlInfoWithDefaults

`func NewQosNotificationControlInfoWithDefaults() *QosNotificationControlInfo`

NewQosNotificationControlInfoWithDefaults instantiates a new QosNotificationControlInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefPccRuleIds

`func (o *QosNotificationControlInfo) GetRefPccRuleIds() []string`

GetRefPccRuleIds returns the RefPccRuleIds field if non-nil, zero value otherwise.

### GetRefPccRuleIdsOk

`func (o *QosNotificationControlInfo) GetRefPccRuleIdsOk() (*[]string, bool)`

GetRefPccRuleIdsOk returns a tuple with the RefPccRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefPccRuleIds

`func (o *QosNotificationControlInfo) SetRefPccRuleIds(v []string)`

SetRefPccRuleIds sets RefPccRuleIds field to given value.


### GetNotifType

`func (o *QosNotificationControlInfo) GetNotifType() QosNotifType`

GetNotifType returns the NotifType field if non-nil, zero value otherwise.

### GetNotifTypeOk

`func (o *QosNotificationControlInfo) GetNotifTypeOk() (*QosNotifType, bool)`

GetNotifTypeOk returns a tuple with the NotifType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifType

`func (o *QosNotificationControlInfo) SetNotifType(v QosNotifType)`

SetNotifType sets NotifType field to given value.


### GetContVer

`func (o *QosNotificationControlInfo) GetContVer() int32`

GetContVer returns the ContVer field if non-nil, zero value otherwise.

### GetContVerOk

`func (o *QosNotificationControlInfo) GetContVerOk() (*int32, bool)`

GetContVerOk returns a tuple with the ContVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContVer

`func (o *QosNotificationControlInfo) SetContVer(v int32)`

SetContVer sets ContVer field to given value.

### HasContVer

`func (o *QosNotificationControlInfo) HasContVer() bool`

HasContVer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


