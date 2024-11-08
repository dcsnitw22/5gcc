/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type Amf3GppAccessRegistration struct {

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	AmfInstanceId string `json:"amfInstanceId"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	PurgeFlag bool `json:"purgeFlag,omitempty"`

	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.
	Pei string `json:"pei,omitempty"`

	ImsVoPs ImsVoPs `json:"imsVoPs,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	DeregCallbackUri string `json:"deregCallbackUri"`

	AmfServiceNameDereg ServiceName `json:"amfServiceNameDereg,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	PcscfRestorationCallbackUri string `json:"pcscfRestorationCallbackUri,omitempty"`

	AmfServiceNamePcscfRest ServiceName `json:"amfServiceNamePcscfRest,omitempty"`

	InitialRegistrationInd bool `json:"initialRegistrationInd,omitempty"`

	EmergencyRegistrationInd bool `json:"emergencyRegistrationInd,omitempty"`

	Guami Guami `json:"guami"`

	BackupAmfInfo []BackupAmfInfo `json:"backupAmfInfo,omitempty"`

	DrFlag bool `json:"drFlag,omitempty"`

	RatType RatType `json:"ratType"`

	UrrpIndicator bool `json:"urrpIndicator,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	AmfEeSubscriptionId string `json:"amfEeSubscriptionId,omitempty"`

	EpsInterworkingInfo EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`

	UeSrvccCapability bool `json:"ueSrvccCapability,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	RegistrationTime time.Time `json:"registrationTime,omitempty"`

	VgmlcAddress VgmlcAddress `json:"vgmlcAddress,omitempty"`

	ContextInfo ContextInfo `json:"contextInfo,omitempty"`

	NoEeSubscriptionInd bool `json:"noEeSubscriptionInd,omitempty"`

	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi string `json:"supi,omitempty"`

	UeReachableInd UeReachableInd `json:"ueReachableInd,omitempty"`

	ReRegistrationRequired bool `json:"reRegistrationRequired,omitempty"`

	AdminDeregSubWithdrawn bool `json:"adminDeregSubWithdrawn,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	DataRestorationCallbackUri string `json:"dataRestorationCallbackUri,omitempty"`

	ResetIds []string `json:"resetIds,omitempty"`

	DisasterRoamingInd bool `json:"disasterRoamingInd,omitempty"`

	UeMINTCapability bool `json:"ueMINTCapability,omitempty"`

	SorSnpnSiSupported bool `json:"sorSnpnSiSupported,omitempty"`

	UdrRestartInd bool `json:"udrRestartInd,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	LastSynchronizationTime time.Time `json:"lastSynchronizationTime,omitempty"`
}

// AssertAmf3GppAccessRegistrationRequired checks if the required fields are not zero-ed
func AssertAmf3GppAccessRegistrationRequired(obj Amf3GppAccessRegistration) error {
	elements := map[string]interface{}{
		"amfInstanceId":    obj.AmfInstanceId,
		"deregCallbackUri": obj.DeregCallbackUri,
		"guami":            obj.Guami,
		"ratType":          obj.RatType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertImsVoPsRequired(obj.ImsVoPs); err != nil {
		return err
	}
	if err := AssertServiceNameRequired(obj.AmfServiceNameDereg); err != nil {
		return err
	}
	if err := AssertServiceNameRequired(obj.AmfServiceNamePcscfRest); err != nil {
		return err
	}
	if err := AssertGuamiRequired(obj.Guami); err != nil {
		return err
	}
	for _, el := range obj.BackupAmfInfo {
		if err := AssertBackupAmfInfoRequired(el); err != nil {
			return err
		}
	}
	if err := AssertRatTypeRequired(obj.RatType); err != nil {
		return err
	}
	if err := AssertEpsInterworkingInfoRequired(obj.EpsInterworkingInfo); err != nil {
		return err
	}
	if err := AssertVgmlcAddressRequired(obj.VgmlcAddress); err != nil {
		return err
	}
	if err := AssertContextInfoRequired(obj.ContextInfo); err != nil {
		return err
	}
	if err := AssertUeReachableIndRequired(obj.UeReachableInd); err != nil {
		return err
	}
	return nil
}

// AssertAmf3GppAccessRegistrationConstraints checks if the values respects the defined constraints
func AssertAmf3GppAccessRegistrationConstraints(obj Amf3GppAccessRegistration) error {
	return nil
}
