# SessionRuleReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleIds** | **[]string** | Contains the identifier of the affected session rule(s). | 
**RuleStatus** | [**RuleStatus**](RuleStatus.md) |  | 
**SessRuleFailureCode** | Pointer to [**SessionRuleFailureCode**](SessionRuleFailureCode.md) |  | [optional] 

## Methods

### NewSessionRuleReport

`func NewSessionRuleReport(ruleIds []string, ruleStatus RuleStatus, ) *SessionRuleReport`

NewSessionRuleReport instantiates a new SessionRuleReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionRuleReportWithDefaults

`func NewSessionRuleReportWithDefaults() *SessionRuleReport`

NewSessionRuleReportWithDefaults instantiates a new SessionRuleReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleIds

`func (o *SessionRuleReport) GetRuleIds() []string`

GetRuleIds returns the RuleIds field if non-nil, zero value otherwise.

### GetRuleIdsOk

`func (o *SessionRuleReport) GetRuleIdsOk() (*[]string, bool)`

GetRuleIdsOk returns a tuple with the RuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIds

`func (o *SessionRuleReport) SetRuleIds(v []string)`

SetRuleIds sets RuleIds field to given value.


### GetRuleStatus

`func (o *SessionRuleReport) GetRuleStatus() RuleStatus`

GetRuleStatus returns the RuleStatus field if non-nil, zero value otherwise.

### GetRuleStatusOk

`func (o *SessionRuleReport) GetRuleStatusOk() (*RuleStatus, bool)`

GetRuleStatusOk returns a tuple with the RuleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleStatus

`func (o *SessionRuleReport) SetRuleStatus(v RuleStatus)`

SetRuleStatus sets RuleStatus field to given value.


### GetSessRuleFailureCode

`func (o *SessionRuleReport) GetSessRuleFailureCode() SessionRuleFailureCode`

GetSessRuleFailureCode returns the SessRuleFailureCode field if non-nil, zero value otherwise.

### GetSessRuleFailureCodeOk

`func (o *SessionRuleReport) GetSessRuleFailureCodeOk() (*SessionRuleFailureCode, bool)`

GetSessRuleFailureCodeOk returns a tuple with the SessRuleFailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessRuleFailureCode

`func (o *SessionRuleReport) SetSessRuleFailureCode(v SessionRuleFailureCode)`

SetSessRuleFailureCode sets SessRuleFailureCode field to given value.

### HasSessRuleFailureCode

`func (o *SessionRuleReport) HasSessRuleFailureCode() bool`

HasSessRuleFailureCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


