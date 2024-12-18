/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// LteV2xAuth - Contains LTE V2X services authorized information.
type LteV2xAuth struct {
	VehicleUeAuth UeAuth `json:"vehicleUeAuth,omitempty"`

	PedestrianUeAuth UeAuth `json:"pedestrianUeAuth,omitempty"`
}

// AssertLteV2xAuthRequired checks if the required fields are not zero-ed
func AssertLteV2xAuthRequired(obj LteV2xAuth) error {
	if err := AssertUeAuthRequired(obj.VehicleUeAuth); err != nil {
		return err
	}
	if err := AssertUeAuthRequired(obj.PedestrianUeAuth); err != nil {
		return err
	}
	return nil
}

// AssertLteV2xAuthConstraints checks if the values respects the defined constraints
func AssertLteV2xAuthConstraints(obj LteV2xAuth) error {
	return nil
}
