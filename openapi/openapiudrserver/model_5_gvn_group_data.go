/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Model5GvnGroupData struct {

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn"`

	SNssai Snssai `json:"sNssai"`

	PduSessionTypes []PduSessionType `json:"pduSessionTypes,omitempty"`

	AppDescriptors []AppDescriptor `json:"appDescriptors,omitempty"`

	SecondaryAuth bool `json:"secondaryAuth,omitempty"`

	DnAaaIpAddressAllocation bool `json:"dnAaaIpAddressAllocation,omitempty"`

	DnAaaAddress *IpAddress1 `json:"dnAaaAddress,omitempty"`

	AdditionalDnAaaAddresses []IpAddress1 `json:"additionalDnAaaAddresses,omitempty"`

	// Fully Qualified Domain Name
	DnAaaFqdn string `json:"dnAaaFqdn,omitempty"`
}

// AssertModel5GvnGroupDataRequired checks if the required fields are not zero-ed
func AssertModel5GvnGroupDataRequired(obj Model5GvnGroupData) error {
	elements := map[string]interface{}{
		"dnn":    obj.Dnn,
		"sNssai": obj.SNssai,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertSnssaiRequired(obj.SNssai); err != nil {
		return err
	}
	for _, el := range obj.PduSessionTypes {
		if err := AssertPduSessionTypeRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.AppDescriptors {
		if err := AssertAppDescriptorRequired(el); err != nil {
			return err
		}
	}
	if obj.DnAaaAddress != nil {
		if err := AssertIpAddress1Required(*obj.DnAaaAddress); err != nil {
			return err
		}
	}
	for _, el := range obj.AdditionalDnAaaAddresses {
		if err := AssertIpAddress1Required(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertModel5GvnGroupDataConstraints checks if the values respects the defined constraints
func AssertModel5GvnGroupDataConstraints(obj Model5GvnGroupData) error {
	return nil
}