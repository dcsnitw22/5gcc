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

// MulticastMbsGroupMembDocumentAPIService is a service that implements the logic for the MulticastMbsGroupMembDocumentAPIServicer
// This service should implement the business logic for every endpoint for the MulticastMbsGroupMembDocumentAPI API.
// Include any external packages or services that will be required by this service.
type MulticastMbsGroupMembDocumentAPIService struct {
}

// NewMulticastMbsGroupMembDocumentAPIService creates a default api service
func NewMulticastMbsGroupMembDocumentAPIService() MulticastMbsGroupMembDocumentAPIServicer {
	return &MulticastMbsGroupMembDocumentAPIService{}
}

// Create5GmbsGroup - Create an individual 5G MBS Grouop
func (s *MulticastMbsGroupMembDocumentAPIService) Create5GmbsGroup(ctx context.Context, externalGroupId string, multicastMbsGroupMemb MulticastMbsGroupMemb) (ImplResponse, error) {
	// TODO - update Create5GmbsGroup with the required logic for this service method.
	// Add api_multicast_mbs_group_memb_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, MulticastMbsGroupMemb{}) or use other options such as http.Ok ...
	// return Response(201, MulticastMbsGroupMemb{}), nil

	// TODO: Uncomment the next line to return response Response(400, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(400, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(401, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(401, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(403, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(403, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(404, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(404, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(411, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(411, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(413, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(413, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(415, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(415, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(429, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(429, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(500, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(500, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(502, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(502, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(503, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(503, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("Create5GmbsGroup method not implemented")
}