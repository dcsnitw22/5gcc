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

// Class5GVNGroupsInternalDocumentAPIService is a service that implements the logic for the Class5GVNGroupsInternalDocumentAPIServicer
// This service should implement the business logic for every endpoint for the Class5GVNGroupsInternalDocumentAPI API.
// Include any external packages or services that will be required by this service.
type Class5GVNGroupsInternalDocumentAPIService struct {
}

// NewClass5GVNGroupsInternalDocumentAPIService creates a default api service
func NewClass5GVNGroupsInternalDocumentAPIService() Class5GVNGroupsInternalDocumentAPIServicer {
	return &Class5GVNGroupsInternalDocumentAPIService{}
}

// Query5GVnGroupInternal - Retrieves the data of 5G VN Group
func (s *Class5GVNGroupsInternalDocumentAPIService) Query5GVnGroupInternal(ctx context.Context, internalGroupIds []string) (ImplResponse, error) {
	// TODO - update Query5GVnGroupInternal with the required logic for this service method.
	// Add api_class5_gvn_groups_internal_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, map[string]Model5GvnGroupConfiguration{}) or use other options such as http.Ok ...
	// return Response(200, map[string]Model5GvnGroupConfiguration{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("Query5GVnGroupInternal method not implemented")
}