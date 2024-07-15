# NetworkId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mnc** | Pointer to **string** |  | [optional] 
**Mcc** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkId

`func NewNetworkId() *NetworkId`

NewNetworkId instantiates a new NetworkId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkIdWithDefaults

`func NewNetworkIdWithDefaults() *NetworkId`

NewNetworkIdWithDefaults instantiates a new NetworkId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMnc

`func (o *NetworkId) GetMnc() string`

GetMnc returns the Mnc field if non-nil, zero value otherwise.

### GetMncOk

`func (o *NetworkId) GetMncOk() (*string, bool)`

GetMncOk returns a tuple with the Mnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnc

`func (o *NetworkId) SetMnc(v string)`

SetMnc sets Mnc field to given value.

### HasMnc

`func (o *NetworkId) HasMnc() bool`

HasMnc returns a boolean if a field has been set.

### GetMcc

`func (o *NetworkId) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *NetworkId) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *NetworkId) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *NetworkId) HasMcc() bool`

HasMcc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


