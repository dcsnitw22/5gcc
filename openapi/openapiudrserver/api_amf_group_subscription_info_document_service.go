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

// AMFGroupSubscriptionInfoDocumentAPIService is a service that implements the logic for the AMFGroupSubscriptionInfoDocumentAPIServicer
// This service should implement the business logic for every endpoint for the AMFGroupSubscriptionInfoDocumentAPI API.
// Include any external packages or services that will be required by this service.
type AMFGroupSubscriptionInfoDocumentAPIService struct {
}

// NewAMFGroupSubscriptionInfoDocumentAPIService creates a default api service
func NewAMFGroupSubscriptionInfoDocumentAPIService() AMFGroupSubscriptionInfoDocumentAPIServicer {
	return &AMFGroupSubscriptionInfoDocumentAPIService{}
}

// CreateAmfGroupSubscriptions - Create AmfSubscriptions for a group of UEs or any UE
func (s *AMFGroupSubscriptionInfoDocumentAPIService) CreateAmfGroupSubscriptions(ctx context.Context, ueGroupId string, subsId string, amfSubscriptionInfo []AmfSubscriptionInfo) (ImplResponse, error) {
	// TODO - update CreateAmfGroupSubscriptions with the required logic for this service method.
	// Add api_amf_group_subscription_info_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(201, []AmfSubscriptionInfo{}) or use other options such as http.Ok ...
	// return Response(201, []AmfSubscriptionInfo{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateAmfGroupSubscriptions method not implemented")
}