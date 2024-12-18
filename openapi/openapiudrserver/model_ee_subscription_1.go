/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type EeSubscription1 struct {

	// String providing an URI formatted according to RFC 3986.
	CallbackReference string `json:"callbackReference"`

	// A map (list of key-value pairs where ReferenceId serves as key) of MonitoringConfigurations
	MonitoringConfigurations map[string]MonitoringConfiguration1 `json:"monitoringConfigurations"`

	ReportingOptions ReportingOptions1 `json:"reportingOptions,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	SubscriptionId string `json:"subscriptionId,omitempty"`

	ContextInfo ContextInfo `json:"contextInfo,omitempty"`

	EpcAppliedInd bool `json:"epcAppliedInd,omitempty"`

	// Fully Qualified Domain Name
	ScefDiamHost string `json:"scefDiamHost,omitempty"`

	// Fully Qualified Domain Name
	ScefDiamRealm string `json:"scefDiamRealm,omitempty"`

	NotifyCorrelationId string `json:"notifyCorrelationId,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	SecondCallbackRef string `json:"secondCallbackRef,omitempty"`

	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi string `json:"gpsi,omitempty"`

	ExcludeGpsiList []string `json:"excludeGpsiList,omitempty"`

	IncludeGpsiList []string `json:"includeGpsiList,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	DataRestorationCallbackUri string `json:"dataRestorationCallbackUri,omitempty"`

	UdrRestartInd bool `json:"udrRestartInd,omitempty"`
}

// AssertEeSubscription1Required checks if the required fields are not zero-ed
func AssertEeSubscription1Required(obj EeSubscription1) error {
	elements := map[string]interface{}{
		"callbackReference":        obj.CallbackReference,
		"monitoringConfigurations": obj.MonitoringConfigurations,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertReportingOptions1Required(obj.ReportingOptions); err != nil {
		return err
	}
	if err := AssertContextInfoRequired(obj.ContextInfo); err != nil {
		return err
	}
	return nil
}

// AssertEeSubscription1Constraints checks if the values respects the defined constraints
func AssertEeSubscription1Constraints(obj EeSubscription1) error {
	return nil
}
