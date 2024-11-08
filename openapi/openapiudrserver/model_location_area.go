/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type LocationArea struct {

	// Identifies a list of geographic area of the user where the UE is located.
	GeographicAreas []GeographicArea `json:"geographicAreas,omitempty"`

	// Identifies a list of civic addresses of the user where the UE is located.
	CivicAddresses []CivicAddress `json:"civicAddresses,omitempty"`

	NwAreaInfo NetworkAreaInfo `json:"nwAreaInfo,omitempty"`

	UmtTime UmtTime `json:"umtTime,omitempty"`
}

// AssertLocationAreaRequired checks if the required fields are not zero-ed
func AssertLocationAreaRequired(obj LocationArea) error {
	for _, el := range obj.GeographicAreas {
		if err := AssertGeographicAreaRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.CivicAddresses {
		if err := AssertCivicAddressRequired(el); err != nil {
			return err
		}
	}
	if err := AssertNetworkAreaInfoRequired(obj.NwAreaInfo); err != nil {
		return err
	}
	if err := AssertUmtTimeRequired(obj.UmtTime); err != nil {
		return err
	}
	return nil
}

// AssertLocationAreaConstraints checks if the values respects the defined constraints
func AssertLocationAreaConstraints(obj LocationArea) error {
	return nil
}
