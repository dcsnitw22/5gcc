/*
Npcf_SMPolicyControl API

Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the PartialSuccessReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartialSuccessReport{}

// PartialSuccessReport struct for PartialSuccessReport
type PartialSuccessReport struct {
	FailureCause FailureCause `json:"failureCause"`
	// Information about the PCC rules provisioned by the PCF not successfully installed/activated.
	RuleReports []RuleReport `json:"ruleReports,omitempty"`
	// Information about the session rules provisioned by the PCF not successfully installed.
	SessRuleReports []SessionRuleReport `json:"sessRuleReports,omitempty"`
	UeCampingRep *UeCampingRep `json:"ueCampingRep,omitempty"`
}

type _PartialSuccessReport PartialSuccessReport

// NewPartialSuccessReport instantiates a new PartialSuccessReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialSuccessReport(failureCause FailureCause) *PartialSuccessReport {
	this := PartialSuccessReport{}
	this.FailureCause = failureCause
	return &this
}

// NewPartialSuccessReportWithDefaults instantiates a new PartialSuccessReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialSuccessReportWithDefaults() *PartialSuccessReport {
	this := PartialSuccessReport{}
	return &this
}

// GetFailureCause returns the FailureCause field value
func (o *PartialSuccessReport) GetFailureCause() FailureCause {
	if o == nil {
		var ret FailureCause
		return ret
	}

	return o.FailureCause
}

// GetFailureCauseOk returns a tuple with the FailureCause field value
// and a boolean to check if the value has been set.
func (o *PartialSuccessReport) GetFailureCauseOk() (*FailureCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailureCause, true
}

// SetFailureCause sets field value
func (o *PartialSuccessReport) SetFailureCause(v FailureCause) {
	o.FailureCause = v
}

// GetRuleReports returns the RuleReports field value if set, zero value otherwise.
func (o *PartialSuccessReport) GetRuleReports() []RuleReport {
	if o == nil || IsNil(o.RuleReports) {
		var ret []RuleReport
		return ret
	}
	return o.RuleReports
}

// GetRuleReportsOk returns a tuple with the RuleReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialSuccessReport) GetRuleReportsOk() ([]RuleReport, bool) {
	if o == nil || IsNil(o.RuleReports) {
		return nil, false
	}
	return o.RuleReports, true
}

// HasRuleReports returns a boolean if a field has been set.
func (o *PartialSuccessReport) HasRuleReports() bool {
	if o != nil && !IsNil(o.RuleReports) {
		return true
	}

	return false
}

// SetRuleReports gets a reference to the given []RuleReport and assigns it to the RuleReports field.
func (o *PartialSuccessReport) SetRuleReports(v []RuleReport) {
	o.RuleReports = v
}

// GetSessRuleReports returns the SessRuleReports field value if set, zero value otherwise.
func (o *PartialSuccessReport) GetSessRuleReports() []SessionRuleReport {
	if o == nil || IsNil(o.SessRuleReports) {
		var ret []SessionRuleReport
		return ret
	}
	return o.SessRuleReports
}

// GetSessRuleReportsOk returns a tuple with the SessRuleReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialSuccessReport) GetSessRuleReportsOk() ([]SessionRuleReport, bool) {
	if o == nil || IsNil(o.SessRuleReports) {
		return nil, false
	}
	return o.SessRuleReports, true
}

// HasSessRuleReports returns a boolean if a field has been set.
func (o *PartialSuccessReport) HasSessRuleReports() bool {
	if o != nil && !IsNil(o.SessRuleReports) {
		return true
	}

	return false
}

// SetSessRuleReports gets a reference to the given []SessionRuleReport and assigns it to the SessRuleReports field.
func (o *PartialSuccessReport) SetSessRuleReports(v []SessionRuleReport) {
	o.SessRuleReports = v
}

// GetUeCampingRep returns the UeCampingRep field value if set, zero value otherwise.
func (o *PartialSuccessReport) GetUeCampingRep() UeCampingRep {
	if o == nil || IsNil(o.UeCampingRep) {
		var ret UeCampingRep
		return ret
	}
	return *o.UeCampingRep
}

// GetUeCampingRepOk returns a tuple with the UeCampingRep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialSuccessReport) GetUeCampingRepOk() (*UeCampingRep, bool) {
	if o == nil || IsNil(o.UeCampingRep) {
		return nil, false
	}
	return o.UeCampingRep, true
}

// HasUeCampingRep returns a boolean if a field has been set.
func (o *PartialSuccessReport) HasUeCampingRep() bool {
	if o != nil && !IsNil(o.UeCampingRep) {
		return true
	}

	return false
}

// SetUeCampingRep gets a reference to the given UeCampingRep and assigns it to the UeCampingRep field.
func (o *PartialSuccessReport) SetUeCampingRep(v UeCampingRep) {
	o.UeCampingRep = &v
}

func (o PartialSuccessReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartialSuccessReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["failureCause"] = o.FailureCause
	if !IsNil(o.RuleReports) {
		toSerialize["ruleReports"] = o.RuleReports
	}
	if !IsNil(o.SessRuleReports) {
		toSerialize["sessRuleReports"] = o.SessRuleReports
	}
	if !IsNil(o.UeCampingRep) {
		toSerialize["ueCampingRep"] = o.UeCampingRep
	}
	return toSerialize, nil
}

func (o *PartialSuccessReport) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"failureCause",
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

	varPartialSuccessReport := _PartialSuccessReport{}

	err = json.Unmarshal(bytes, &varPartialSuccessReport)

	if err != nil {
		return err
	}

	*o = PartialSuccessReport(varPartialSuccessReport)

	return err
}

type NullablePartialSuccessReport struct {
	value *PartialSuccessReport
	isSet bool
}

func (v NullablePartialSuccessReport) Get() *PartialSuccessReport {
	return v.value
}

func (v *NullablePartialSuccessReport) Set(val *PartialSuccessReport) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialSuccessReport) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialSuccessReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialSuccessReport(val *PartialSuccessReport) *NullablePartialSuccessReport {
	return &NullablePartialSuccessReport{value: val, isSet: true}
}

func (v NullablePartialSuccessReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialSuccessReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

