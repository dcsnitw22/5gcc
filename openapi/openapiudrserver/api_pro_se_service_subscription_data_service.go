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

// ProSeServiceSubscriptionDataAPIService is a service that implements the logic for the ProSeServiceSubscriptionDataAPIServicer
// This service should implement the business logic for every endpoint for the ProSeServiceSubscriptionDataAPI API.
// Include any external packages or services that will be required by this service.
type ProSeServiceSubscriptionDataAPIService struct {
}

// NewProSeServiceSubscriptionDataAPIService creates a default api service
func NewProSeServiceSubscriptionDataAPIService() ProSeServiceSubscriptionDataAPIServicer {
	return &ProSeServiceSubscriptionDataAPIService{}
}

// QueryPorseData - Retrieves the subscribed ProSe service Data of a UE
func (s *ProSeServiceSubscriptionDataAPIService) QueryPorseData(ctx context.Context, ueId string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (ImplResponse, error) {
	// TODO - update QueryPorseData with the required logic for this service method.
	// Add api_pro_se_service_subscription_data_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ProseSubscriptionData{}) or use other options such as http.Ok ...
	// return Response(200, ProseSubscriptionData{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("QueryPorseData method not implemented")
}