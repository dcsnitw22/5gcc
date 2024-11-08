# ConditionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CondId** | **string** | Uniquely identifies the condition data within a PDU session. | 
**ActivationTime** | Pointer to **NullableTime** |  | [optional] 
**DeactivationTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewConditionData

`func NewConditionData(condId string, ) *ConditionData`

NewConditionData instantiates a new ConditionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionDataWithDefaults

`func NewConditionDataWithDefaults() *ConditionData`

NewConditionDataWithDefaults instantiates a new ConditionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondId

`func (o *ConditionData) GetCondId() string`

GetCondId returns the CondId field if non-nil, zero value otherwise.

### GetCondIdOk

`func (o *ConditionData) GetCondIdOk() (*string, bool)`

GetCondIdOk returns a tuple with the CondId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondId

`func (o *ConditionData) SetCondId(v string)`

SetCondId sets CondId field to given value.


### GetActivationTime

`func (o *ConditionData) GetActivationTime() time.Time`

GetActivationTime returns the ActivationTime field if non-nil, zero value otherwise.

### GetActivationTimeOk

`func (o *ConditionData) GetActivationTimeOk() (*time.Time, bool)`

GetActivationTimeOk returns a tuple with the ActivationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationTime

`func (o *ConditionData) SetActivationTime(v time.Time)`

SetActivationTime sets ActivationTime field to given value.

### HasActivationTime

`func (o *ConditionData) HasActivationTime() bool`

HasActivationTime returns a boolean if a field has been set.

### SetActivationTimeNil

`func (o *ConditionData) SetActivationTimeNil(b bool)`

 SetActivationTimeNil sets the value for ActivationTime to be an explicit nil

### UnsetActivationTime
`func (o *ConditionData) UnsetActivationTime()`

UnsetActivationTime ensures that no value is present for ActivationTime, not even an explicit nil
### GetDeactivationTime

`func (o *ConditionData) GetDeactivationTime() time.Time`

GetDeactivationTime returns the DeactivationTime field if non-nil, zero value otherwise.

### GetDeactivationTimeOk

`func (o *ConditionData) GetDeactivationTimeOk() (*time.Time, bool)`

GetDeactivationTimeOk returns a tuple with the DeactivationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivationTime

`func (o *ConditionData) SetDeactivationTime(v time.Time)`

SetDeactivationTime sets DeactivationTime field to given value.

### HasDeactivationTime

`func (o *ConditionData) HasDeactivationTime() bool`

HasDeactivationTime returns a boolean if a field has been set.

### SetDeactivationTimeNil

`func (o *ConditionData) SetDeactivationTimeNil(b bool)`

 SetDeactivationTimeNil sets the value for DeactivationTime to be an explicit nil

### UnsetDeactivationTime
`func (o *ConditionData) UnsetDeactivationTime()`

UnsetDeactivationTime ensures that no value is present for DeactivationTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


