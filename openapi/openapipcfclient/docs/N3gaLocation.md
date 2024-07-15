# N3gaLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N3gppTai** | Pointer to [**Tai**](Tai.md) |  | [optional] 
**N3IwfId** | Pointer to **string** |  | [optional] 
**UeIpv4Addr** | Pointer to **string** |  | [optional] 
**UeIpv6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**PortNumber** | Pointer to **int32** |  | [optional] 

## Methods

### NewN3gaLocation

`func NewN3gaLocation() *N3gaLocation`

NewN3gaLocation instantiates a new N3gaLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN3gaLocationWithDefaults

`func NewN3gaLocationWithDefaults() *N3gaLocation`

NewN3gaLocationWithDefaults instantiates a new N3gaLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN3gppTai

`func (o *N3gaLocation) GetN3gppTai() Tai`

GetN3gppTai returns the N3gppTai field if non-nil, zero value otherwise.

### GetN3gppTaiOk

`func (o *N3gaLocation) GetN3gppTaiOk() (*Tai, bool)`

GetN3gppTaiOk returns a tuple with the N3gppTai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN3gppTai

`func (o *N3gaLocation) SetN3gppTai(v Tai)`

SetN3gppTai sets N3gppTai field to given value.

### HasN3gppTai

`func (o *N3gaLocation) HasN3gppTai() bool`

HasN3gppTai returns a boolean if a field has been set.

### GetN3IwfId

`func (o *N3gaLocation) GetN3IwfId() string`

GetN3IwfId returns the N3IwfId field if non-nil, zero value otherwise.

### GetN3IwfIdOk

`func (o *N3gaLocation) GetN3IwfIdOk() (*string, bool)`

GetN3IwfIdOk returns a tuple with the N3IwfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN3IwfId

`func (o *N3gaLocation) SetN3IwfId(v string)`

SetN3IwfId sets N3IwfId field to given value.

### HasN3IwfId

`func (o *N3gaLocation) HasN3IwfId() bool`

HasN3IwfId returns a boolean if a field has been set.

### GetUeIpv4Addr

`func (o *N3gaLocation) GetUeIpv4Addr() string`

GetUeIpv4Addr returns the UeIpv4Addr field if non-nil, zero value otherwise.

### GetUeIpv4AddrOk

`func (o *N3gaLocation) GetUeIpv4AddrOk() (*string, bool)`

GetUeIpv4AddrOk returns a tuple with the UeIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv4Addr

`func (o *N3gaLocation) SetUeIpv4Addr(v string)`

SetUeIpv4Addr sets UeIpv4Addr field to given value.

### HasUeIpv4Addr

`func (o *N3gaLocation) HasUeIpv4Addr() bool`

HasUeIpv4Addr returns a boolean if a field has been set.

### GetUeIpv6Addr

`func (o *N3gaLocation) GetUeIpv6Addr() Ipv6Addr`

GetUeIpv6Addr returns the UeIpv6Addr field if non-nil, zero value otherwise.

### GetUeIpv6AddrOk

`func (o *N3gaLocation) GetUeIpv6AddrOk() (*Ipv6Addr, bool)`

GetUeIpv6AddrOk returns a tuple with the UeIpv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6Addr

`func (o *N3gaLocation) SetUeIpv6Addr(v Ipv6Addr)`

SetUeIpv6Addr sets UeIpv6Addr field to given value.

### HasUeIpv6Addr

`func (o *N3gaLocation) HasUeIpv6Addr() bool`

HasUeIpv6Addr returns a boolean if a field has been set.

### GetPortNumber

`func (o *N3gaLocation) GetPortNumber() int32`

GetPortNumber returns the PortNumber field if non-nil, zero value otherwise.

### GetPortNumberOk

`func (o *N3gaLocation) GetPortNumberOk() (*int32, bool)`

GetPortNumberOk returns a tuple with the PortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNumber

`func (o *N3gaLocation) SetPortNumber(v int32)`

SetPortNumber sets PortNumber field to given value.

### HasPortNumber

`func (o *N3gaLocation) HasPortNumber() bool`

HasPortNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


