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

// AuthenticationSoRDocumentAPIService is a service that implements the logic for the AuthenticationSoRDocumentAPIServicer
// This service should implement the business logic for every endpoint for the AuthenticationSoRDocumentAPI API.
// Include any external packages or services that will be required by this service.
type AuthenticationSoRDocumentAPIService struct {
}

// NewAuthenticationSoRDocumentAPIService creates a default api service
func NewAuthenticationSoRDocumentAPIService() AuthenticationSoRDocumentAPIServicer {
	return &AuthenticationSoRDocumentAPIService{}
}

// CreateAuthenticationSoR - To store the SoR acknowledgement information of a UE and ME support of SOR CMCI
func (s *AuthenticationSoRDocumentAPIService) CreateAuthenticationSoR(ctx context.Context, ueId string, sorData SorData, supportedFeatures string) (ImplResponse, error) {
	// TODO - update CreateAuthenticationSoR with the required logic for this service method.
	// Add api_authentication_so_r_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateAuthenticationSoR method not implemented")
}

// QueryAuthSoR - Retrieves the SoR acknowledgement information of a UE and ME support of SOR CMCI
func (s *AuthenticationSoRDocumentAPIService) QueryAuthSoR(ctx context.Context, ueId string, supportedFeatures string) (ImplResponse, error) {
	// TODO - update QueryAuthSoR with the required logic for this service method.
	// Add api_authentication_so_r_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, SorData{}) or use other options such as http.Ok ...
	// return Response(200, SorData{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("QueryAuthSoR method not implemented")
}

// UpdateAuthenticationSoR - Updates the ME support of SOR CMCI information of a UE
func (s *AuthenticationSoRDocumentAPIService) UpdateAuthenticationSoR(ctx context.Context, ueId string, patchItem []PatchItem, supportedFeatures string) (ImplResponse, error) {
	// TODO - update UpdateAuthenticationSoR with the required logic for this service method.
	// Add api_authentication_so_r_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(200, PatchResult{}) or use other options such as http.Ok ...
	// return Response(200, PatchResult{}), nil

	// TODO: Uncomment the next line to return response Response(403, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(403, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateAuthenticationSoR method not implemented")
}
