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
	"context"
	"errors"
	"net/http"
)

// SMSFNon3GPPRegistrationDocumentAPIService is a service that implements the logic for the SMSFNon3GPPRegistrationDocumentAPIServicer
// This service should implement the business logic for every endpoint for the SMSFNon3GPPRegistrationDocumentAPI API.
// Include any external packages or services that will be required by this service.
type SMSFNon3GPPRegistrationDocumentAPIService struct {
}

// NewSMSFNon3GPPRegistrationDocumentAPIService creates a default api service
func NewSMSFNon3GPPRegistrationDocumentAPIService() SMSFNon3GPPRegistrationDocumentAPIServicer {
	return &SMSFNon3GPPRegistrationDocumentAPIService{}
}

// CreateSmsfContextNon3gpp - Create the SMSF context data of a UE via non-3GPP access
func (s *SMSFNon3GPPRegistrationDocumentAPIService) CreateSmsfContextNon3gpp(ctx context.Context, ueId string, smsfRegistration SmsfRegistration) (ImplResponse, error) {
	// TODO - update CreateSmsfContextNon3gpp with the required logic for this service method.
	// Add api_smsf_non3_gpp_registration_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(201, SmsfRegistration{}) or use other options such as http.Ok ...
	// return Response(201, SmsfRegistration{}), nil

	// TODO: Uncomment the next line to return response Response(200, SmsfRegistration{}) or use other options such as http.Ok ...
	// return Response(200, SmsfRegistration{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateSmsfContextNon3gpp method not implemented")
}

// DeleteSmsfContextNon3gpp - To remove the SMSF context data of a UE via non-3GPP access
func (s *SMSFNon3GPPRegistrationDocumentAPIService) DeleteSmsfContextNon3gpp(ctx context.Context, ueId string) (ImplResponse, error) {
	// TODO - update DeleteSmsfContextNon3gpp with the required logic for this service method.
	// Add api_smsf_non3_gpp_registration_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteSmsfContextNon3gpp method not implemented")
}

// QuerySmsfContextNon3gpp - Retrieves the SMSF context data of a UE using non-3gpp access
func (s *SMSFNon3GPPRegistrationDocumentAPIService) QuerySmsfContextNon3gpp(ctx context.Context, ueId string, fields []string, supportedFeatures string) (ImplResponse, error) {
	// TODO - update QuerySmsfContextNon3gpp with the required logic for this service method.
	// Add api_smsf_non3_gpp_registration_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, SmsfRegistration{}) or use other options such as http.Ok ...
	// return Response(200, SmsfRegistration{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("QuerySmsfContextNon3gpp method not implemented")
}