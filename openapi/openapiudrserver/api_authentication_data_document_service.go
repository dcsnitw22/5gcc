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

// AuthenticationDataDocumentAPIService is a service that implements the logic for the AuthenticationDataDocumentAPIServicer
// This service should implement the business logic for every endpoint for the AuthenticationDataDocumentAPI API.
// Include any external packages or services that will be required by this service.
type AuthenticationDataDocumentAPIService struct {
}

// NewAuthenticationDataDocumentAPIService creates a default api service
func NewAuthenticationDataDocumentAPIService() AuthenticationDataDocumentAPIServicer {
	return &AuthenticationDataDocumentAPIService{}
}

// QueryAuthSubsData - Retrieves the authentication subscription data of a UE
func (s *AuthenticationDataDocumentAPIService) QueryAuthSubsData(ctx context.Context, ueId string, supportedFeatures string) (ImplResponse, error) {
	// TODO - update QueryAuthSubsData with the required logic for this service method.
	// Add api_authentication_data_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, AuthenticationSubscription{}) or use other options such as http.Ok ...
	// return Response(200, AuthenticationSubscription{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("QueryAuthSubsData method not implemented")
}
