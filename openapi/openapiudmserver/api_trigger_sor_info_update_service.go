/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// TriggerSORInfoUpdateAPIService is a service that implements the logic for the TriggerSORInfoUpdateAPIServicer
// This service should implement the business logic for every endpoint for the TriggerSORInfoUpdateAPI API.
// Include any external packages or services that will be required by this service.
type TriggerSORInfoUpdateAPIService struct {
}

// NewTriggerSORInfoUpdateAPIService creates a default api service
func NewTriggerSORInfoUpdateAPIService() TriggerSORInfoUpdateAPIServicer {
	return &TriggerSORInfoUpdateAPIService{}
}

// UpdateSORInfo - Nudm_Sdm custom operation to trigger SOR info update
func (s *TriggerSORInfoUpdateAPIService) UpdateSORInfo(ctx context.Context, supi string, sorUpdateInfo SorUpdateInfo) (ImplResponse, error) {
	// TODO - update UpdateSORInfo with the required logic for this service method.
	// Add api_trigger_sor_info_update_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, SorInfo{}) or use other options such as http.Ok ...
	// return Response(200, SorInfo{}), nil

	// TODO: Uncomment the next line to return response Response(400, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(400, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(404, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(404, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(500, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(500, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(503, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(503, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateSORInfo method not implemented")
}