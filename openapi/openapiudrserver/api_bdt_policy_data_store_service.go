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

// BdtPolicyDataStoreAPIService is a service that implements the logic for the BdtPolicyDataStoreAPIServicer
// This service should implement the business logic for every endpoint for the BdtPolicyDataStoreAPI API.
// Include any external packages or services that will be required by this service.
type BdtPolicyDataStoreAPIService struct {
}

// NewBdtPolicyDataStoreAPIService creates a default api service
func NewBdtPolicyDataStoreAPIService() BdtPolicyDataStoreAPIServicer {
	return &BdtPolicyDataStoreAPIService{}
}

// ReadBdtPolicyData - Retrieve applied BDT Policy Data
func (s *BdtPolicyDataStoreAPIService) ReadBdtPolicyData(ctx context.Context, bdtPolicyIds []string, internalGroupIds []string, supis []string) (ImplResponse, error) {
	// TODO - update ReadBdtPolicyData with the required logic for this service method.
	// Add api_bdt_policy_data_store_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []BdtPolicyData{}) or use other options such as http.Ok ...
	// return Response(200, []BdtPolicyData{}), nil

	// TODO: Uncomment the next line to return response Response(400, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(400, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(401, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(401, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(403, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(403, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(404, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(404, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(406, {}) or use other options such as http.Ok ...
	// return Response(406, nil),nil

	// TODO: Uncomment the next line to return response Response(414, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(414, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(429, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(429, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(500, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(500, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(503, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(503, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ReadBdtPolicyData method not implemented")
}
