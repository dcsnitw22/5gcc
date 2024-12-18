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

// HSSEventSubscriptionInfoDocumentAPIService is a service that implements the logic for the HSSEventSubscriptionInfoDocumentAPIServicer
// This service should implement the business logic for every endpoint for the HSSEventSubscriptionInfoDocumentAPI API.
// Include any external packages or services that will be required by this service.
type HSSEventSubscriptionInfoDocumentAPIService struct {
}

// NewHSSEventSubscriptionInfoDocumentAPIService creates a default api service
func NewHSSEventSubscriptionInfoDocumentAPIService() HSSEventSubscriptionInfoDocumentAPIServicer {
	return &HSSEventSubscriptionInfoDocumentAPIService{}
}

// CreateHSSSubscriptions - Create HSS Subscription Info
func (s *HSSEventSubscriptionInfoDocumentAPIService) CreateHSSSubscriptions(ctx context.Context, ueId string, subsId string, hssSubscriptionInfo HssSubscriptionInfo) (ImplResponse, error) {
	// TODO - update CreateHSSSubscriptions with the required logic for this service method.
	// Add api_hss_event_subscription_info_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(201, HssSubscriptionInfo{}) or use other options such as http.Ok ...
	// return Response(201, HssSubscriptionInfo{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateHSSSubscriptions method not implemented")
}

// GetHssGroupSubscriptions - Retrieve HSS Subscription Info
func (s *HSSEventSubscriptionInfoDocumentAPIService) GetHssGroupSubscriptions(ctx context.Context, externalGroupId string, subsId string) (ImplResponse, error) {
	// TODO - update GetHssGroupSubscriptions with the required logic for this service method.
	// Add api_hss_event_subscription_info_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, HssSubscriptionInfo{}) or use other options such as http.Ok ...
	// return Response(200, HssSubscriptionInfo{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetHssGroupSubscriptions method not implemented")
}

// GetHssSubscriptionInfo - Retrieve HSS Subscription Info
func (s *HSSEventSubscriptionInfoDocumentAPIService) GetHssSubscriptionInfo(ctx context.Context, ueId string, subsId string) (ImplResponse, error) {
	// TODO - update GetHssSubscriptionInfo with the required logic for this service method.
	// Add api_hss_event_subscription_info_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, SmfSubscriptionInfo{}) or use other options such as http.Ok ...
	// return Response(200, SmfSubscriptionInfo{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetHssSubscriptionInfo method not implemented")
}

// ModifyHssGroupSubscriptions - Modify HSS Subscription Info
func (s *HSSEventSubscriptionInfoDocumentAPIService) ModifyHssGroupSubscriptions(ctx context.Context, externalGroupId string, subsId string, patchItem []PatchItem, supportedFeatures string) (ImplResponse, error) {
	// TODO - update ModifyHssGroupSubscriptions with the required logic for this service method.
	// Add api_hss_event_subscription_info_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(200, PatchResult{}) or use other options such as http.Ok ...
	// return Response(200, PatchResult{}), nil

	// TODO: Uncomment the next line to return response Response(403, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(403, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ModifyHssGroupSubscriptions method not implemented")
}

// ModifyHssSubscriptionInfo - Modify HSS Subscription Info
func (s *HSSEventSubscriptionInfoDocumentAPIService) ModifyHssSubscriptionInfo(ctx context.Context, ueId string, subsId string, patchItem []PatchItem, supportedFeatures string) (ImplResponse, error) {
	// TODO - update ModifyHssSubscriptionInfo with the required logic for this service method.
	// Add api_hss_event_subscription_info_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(200, PatchResult{}) or use other options such as http.Ok ...
	// return Response(200, PatchResult{}), nil

	// TODO: Uncomment the next line to return response Response(403, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(403, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ModifyHssSubscriptionInfo method not implemented")
}

// RemoveHssGroupSubscriptions - Delete HSS Subscription Info
func (s *HSSEventSubscriptionInfoDocumentAPIService) RemoveHssGroupSubscriptions(ctx context.Context, externalGroupId string, subsId string) (ImplResponse, error) {
	// TODO - update RemoveHssGroupSubscriptions with the required logic for this service method.
	// Add api_hss_event_subscription_info_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("RemoveHssGroupSubscriptions method not implemented")
}

// RemoveHssSubscriptionsInfo - Delete HSS Subscription Info
func (s *HSSEventSubscriptionInfoDocumentAPIService) RemoveHssSubscriptionsInfo(ctx context.Context, ueId string, subsId string) (ImplResponse, error) {
	// TODO - update RemoveHssSubscriptionsInfo with the required logic for this service method.
	// Add api_hss_event_subscription_info_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("RemoveHssSubscriptionsInfo method not implemented")
}
