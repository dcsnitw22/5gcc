/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_pcf_srv

// AnGwAddress - describes the address of the access network gateway control node
type AnGwAddress struct {
	AnGwIpv4Addr string `json:"anGwIpv4Addr,omitempty"`

	AnGwIpv6Addr Ipv6Addr `json:"anGwIpv6Addr,omitempty"`
}

// AssertAnGwAddressRequired checks if the required fields are not zero-ed
func AssertAnGwAddressRequired(obj AnGwAddress) error {
	if err := AssertIpv6AddrRequired(obj.AnGwIpv6Addr); err != nil {
		return err
	}
	return nil
}

// AssertAnGwAddressConstraints checks if the values respects the defined constraints
func AssertAnGwAddressConstraints(obj AnGwAddress) error {
	return nil
}