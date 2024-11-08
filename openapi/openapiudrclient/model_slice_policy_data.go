/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_udr_cli

import (
	"encoding/json"
)

// checks if the SlicePolicyData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlicePolicyData{}

// SlicePolicyData Contains the network slice specific policy control information.
type SlicePolicyData struct {
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MbrUl *string `json:"mbrUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MbrDl *string `json:"mbrDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	RemainMbrUl *string `json:"remainMbrUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	RemainMbrDl *string `json:"remainMbrDl,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
}

// NewSlicePolicyData instantiates a new SlicePolicyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlicePolicyData() *SlicePolicyData {
	this := SlicePolicyData{}
	return &this
}

// NewSlicePolicyDataWithDefaults instantiates a new SlicePolicyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlicePolicyDataWithDefaults() *SlicePolicyData {
	this := SlicePolicyData{}
	return &this
}

// GetMbrUl returns the MbrUl field value if set, zero value otherwise.
func (o *SlicePolicyData) GetMbrUl() string {
	if o == nil || IsNil(o.MbrUl) {
		var ret string
		return ret
	}
	return *o.MbrUl
}

// GetMbrUlOk returns a tuple with the MbrUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlicePolicyData) GetMbrUlOk() (*string, bool) {
	if o == nil || IsNil(o.MbrUl) {
		return nil, false
	}
	return o.MbrUl, true
}

// HasMbrUl returns a boolean if a field has been set.
func (o *SlicePolicyData) HasMbrUl() bool {
	if o != nil && !IsNil(o.MbrUl) {
		return true
	}

	return false
}

// SetMbrUl gets a reference to the given string and assigns it to the MbrUl field.
func (o *SlicePolicyData) SetMbrUl(v string) {
	o.MbrUl = &v
}

// GetMbrDl returns the MbrDl field value if set, zero value otherwise.
func (o *SlicePolicyData) GetMbrDl() string {
	if o == nil || IsNil(o.MbrDl) {
		var ret string
		return ret
	}
	return *o.MbrDl
}

// GetMbrDlOk returns a tuple with the MbrDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlicePolicyData) GetMbrDlOk() (*string, bool) {
	if o == nil || IsNil(o.MbrDl) {
		return nil, false
	}
	return o.MbrDl, true
}

// HasMbrDl returns a boolean if a field has been set.
func (o *SlicePolicyData) HasMbrDl() bool {
	if o != nil && !IsNil(o.MbrDl) {
		return true
	}

	return false
}

// SetMbrDl gets a reference to the given string and assigns it to the MbrDl field.
func (o *SlicePolicyData) SetMbrDl(v string) {
	o.MbrDl = &v
}

// GetRemainMbrUl returns the RemainMbrUl field value if set, zero value otherwise.
func (o *SlicePolicyData) GetRemainMbrUl() string {
	if o == nil || IsNil(o.RemainMbrUl) {
		var ret string
		return ret
	}
	return *o.RemainMbrUl
}

// GetRemainMbrUlOk returns a tuple with the RemainMbrUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlicePolicyData) GetRemainMbrUlOk() (*string, bool) {
	if o == nil || IsNil(o.RemainMbrUl) {
		return nil, false
	}
	return o.RemainMbrUl, true
}

// HasRemainMbrUl returns a boolean if a field has been set.
func (o *SlicePolicyData) HasRemainMbrUl() bool {
	if o != nil && !IsNil(o.RemainMbrUl) {
		return true
	}

	return false
}

// SetRemainMbrUl gets a reference to the given string and assigns it to the RemainMbrUl field.
func (o *SlicePolicyData) SetRemainMbrUl(v string) {
	o.RemainMbrUl = &v
}

// GetRemainMbrDl returns the RemainMbrDl field value if set, zero value otherwise.
func (o *SlicePolicyData) GetRemainMbrDl() string {
	if o == nil || IsNil(o.RemainMbrDl) {
		var ret string
		return ret
	}
	return *o.RemainMbrDl
}

// GetRemainMbrDlOk returns a tuple with the RemainMbrDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlicePolicyData) GetRemainMbrDlOk() (*string, bool) {
	if o == nil || IsNil(o.RemainMbrDl) {
		return nil, false
	}
	return o.RemainMbrDl, true
}

// HasRemainMbrDl returns a boolean if a field has been set.
func (o *SlicePolicyData) HasRemainMbrDl() bool {
	if o != nil && !IsNil(o.RemainMbrDl) {
		return true
	}

	return false
}

// SetRemainMbrDl gets a reference to the given string and assigns it to the RemainMbrDl field.
func (o *SlicePolicyData) SetRemainMbrDl(v string) {
	o.RemainMbrDl = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *SlicePolicyData) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlicePolicyData) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *SlicePolicyData) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *SlicePolicyData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

// GetResetIds returns the ResetIds field value if set, zero value otherwise.
func (o *SlicePolicyData) GetResetIds() []string {
	if o == nil || IsNil(o.ResetIds) {
		var ret []string
		return ret
	}
	return o.ResetIds
}

// GetResetIdsOk returns a tuple with the ResetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlicePolicyData) GetResetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResetIds) {
		return nil, false
	}
	return o.ResetIds, true
}

// HasResetIds returns a boolean if a field has been set.
func (o *SlicePolicyData) HasResetIds() bool {
	if o != nil && !IsNil(o.ResetIds) {
		return true
	}

	return false
}

// SetResetIds gets a reference to the given []string and assigns it to the ResetIds field.
func (o *SlicePolicyData) SetResetIds(v []string) {
	o.ResetIds = v
}

func (o SlicePolicyData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlicePolicyData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MbrUl) {
		toSerialize["mbrUl"] = o.MbrUl
	}
	if !IsNil(o.MbrDl) {
		toSerialize["mbrDl"] = o.MbrDl
	}
	if !IsNil(o.RemainMbrUl) {
		toSerialize["remainMbrUl"] = o.RemainMbrUl
	}
	if !IsNil(o.RemainMbrDl) {
		toSerialize["remainMbrDl"] = o.RemainMbrDl
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	if !IsNil(o.ResetIds) {
		toSerialize["resetIds"] = o.ResetIds
	}
	return toSerialize, nil
}

type NullableSlicePolicyData struct {
	value *SlicePolicyData
	isSet bool
}

func (v NullableSlicePolicyData) Get() *SlicePolicyData {
	return v.value
}

func (v *NullableSlicePolicyData) Set(val *SlicePolicyData) {
	v.value = val
	v.isSet = true
}

func (v NullableSlicePolicyData) IsSet() bool {
	return v.isSet
}

func (v *NullableSlicePolicyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlicePolicyData(val *SlicePolicyData) *NullableSlicePolicyData {
	return &NullableSlicePolicyData{value: val, isSet: true}
}

func (v NullableSlicePolicyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlicePolicyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


