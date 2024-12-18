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

// PFDDataStoreAPIService is a service that implements the logic for the PFDDataStoreAPIServicer
// This service should implement the business logic for every endpoint for the PFDDataStoreAPI API.
// Include any external packages or services that will be required by this service.
type PFDDataStoreAPIService struct {
}

// NewPFDDataStoreAPIService creates a default api service
func NewPFDDataStoreAPIService() PFDDataStoreAPIServicer {
	return &PFDDataStoreAPIService{}
}

// ReadPFDData - Retrieve PFDs for application identifier(s)
func (s *PFDDataStoreAPIService) ReadPFDData(ctx context.Context, appId []string, suppFeat string) (ImplResponse, error) {
	// TODO - update ReadPFDData with the required logic for this service method.
	// Add api_pfd_data_store_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []PfdDataForAppExt{}) or use other options such as http.Ok ...
	// return Response(200, []PfdDataForAppExt{}), nil

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

	return Response(http.StatusNotImplemented, nil), errors.New("ReadPFDData method not implemented")
}
