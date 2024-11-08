/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// EcsServerAddr - Contains the Edge Configuration Server Address Configuration Information as defined in clause 5.2.3.6.1 of 3GPP TS 23.502.
type EcsServerAddr struct {
	EcsFqdnList []string `json:"ecsFqdnList,omitempty"`

	EcsIpAddressList []IpAddr `json:"ecsIpAddressList,omitempty"`

	EcsUriList []string `json:"ecsUriList,omitempty"`

	EcsProviderId string `json:"ecsProviderId,omitempty"`
}

// AssertEcsServerAddrRequired checks if the required fields are not zero-ed
func AssertEcsServerAddrRequired(obj EcsServerAddr) error {
	for _, el := range obj.EcsIpAddressList {
		if err := AssertIpAddrRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertEcsServerAddrConstraints checks if the values respects the defined constraints
func AssertEcsServerAddrConstraints(obj EcsServerAddr) error {
	return nil
}
