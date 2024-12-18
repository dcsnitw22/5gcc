/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TraceData1 - contains Trace control and configuration parameters.
type TraceData1 struct {

	// Trace Reference (see 3GPP TS 32.422).It shall be encoded as the concatenation of MCC, MNC and Trace ID as follows: 'MCC'<MNC'-'Trace ID'The Trace ID shall be encoded as a 3 octet string in hexadecimal representation. Each character in the Trace ID string shall take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character representing the 4 most significant bits of the Trace ID shall appear first  in the string, and the character representing the 4 least significant bit of the Trace ID shall appear last in the string.
	TraceRef string `json:"traceRef"`

	// TraceDepth TraceDepth `json:"traceDepth"`

	// List of NE Types (see 3GPP TS 32.422).It shall be encoded as an octet string in hexadecimal representation. Each character in the string shall take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character representing the 4 most significant bits shall appear first in the string, and the character representing the 4 least significant bit shall appear last in the string.Octets shall be coded according to 3GPP TS 32.422.
	NeTypeList string `json:"neTypeList"`

	// Triggering events (see 3GPP TS 32.422).It shall be encoded as an octet string in hexadecimal representation. Each character in the string shall take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character representing the 4 most significant bits shall appear first in the string, and the character representing the 4 least significant bit shall appear last in the string. Octets shall be coded according to 3GPP TS 32.422.
	EventList string `json:"eventList"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	CollectionEntityIpv4Addr string `json:"collectionEntityIpv4Addr,omitempty"`

	// CollectionEntityIpv6Addr Ipv6Addr `json:"collectionEntityIpv6Addr,omitempty"`

	// List of Interfaces (see 3GPP TS 32.422).It shall be encoded as an octet string in hexadecimal representation. Each character in the string shall take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character representing the 4 most significant bits shall appear first in the string, and the character representing the  4 least significant bit shall appear last in the string. Octets shall be coded according to 3GPP TS 32.422. If this attribute is not present, all the interfaces applicable to the list of NE types indicated in the neTypeList attribute should be traced.
	InterfaceList string `json:"interfaceList,omitempty"`
}

// AssertTraceData1Required checks if the required fields are not zero-ed
func AssertTraceData1Required(obj TraceData1) error {
	elements := map[string]interface{}{
		"traceRef": obj.TraceRef,
		// "traceDepth": obj.TraceDepth,
		"neTypeList": obj.NeTypeList,
		"eventList":  obj.EventList,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	// if err := AssertTraceDepthRequired(obj.TraceDepth); err != nil {
	// 	return err
	// }
	// if err := AssertIpv6AddrRequired(obj.CollectionEntityIpv6Addr); err != nil {
	// 	return err
	// }
	return nil
}

// AssertTraceData1Constraints checks if the values respects the defined constraints
func AssertTraceData1Constraints(obj TraceData1) error {
	return nil
}
