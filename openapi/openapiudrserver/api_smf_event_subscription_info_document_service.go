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

// SMFEventSubscriptionInfoDocumentAPIService is a service that implements the logic for the SMFEventSubscriptionInfoDocumentAPIServicer
// This service should implement the business logic for every endpoint for the SMFEventSubscriptionInfoDocumentAPI API.
// Include any external packages or services that will be required by this service.
type SMFEventSubscriptionInfoDocumentAPIService struct {
}

// NewSMFEventSubscriptionInfoDocumentAPIService creates a default api service
func NewSMFEventSubscriptionInfoDocumentAPIService() SMFEventSubscriptionInfoDocumentAPIServicer {
	return &SMFEventSubscriptionInfoDocumentAPIService{}
}

// CreateSMFSubscriptions - Create SMF Subscription Info
func (s *SMFEventSubscriptionInfoDocumentAPIService) CreateSMFSubscriptions(ctx context.Context, ueId string, subsId string, smfSubscriptionInfo SmfSubscriptionInfo) (ImplResponse, error) {
	// TODO - update CreateSMFSubscriptions with the required logic for this service method.
	// Add api_smf_event_subscription_info_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(201, SmfSubscriptionInfo{}) or use other options such as http.Ok ...
	// return Response(201, SmfSubscriptionInfo{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateSMFSubscriptions method not implemented")
}

// GetSmfGroupSubscriptions - Retrieve SMF Subscription Info for a group of UEs or any UE
func (s *SMFEventSubscriptionInfoDocumentAPIService) GetSmfGroupSubscriptions(ctx context.Context, ueGroupId string, subsId string) (ImplResponse, error) {
	// TODO - update GetSmfGroupSubscriptions with the required logic for this service method.
	// Add api_smf_event_subscription_info_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, SmfSubscriptionInfo{}) or use other options such as http.Ok ...
	// return Response(200, SmfSubscriptionInfo{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetSmfGroupSubscriptions method not implemented")
}

// GetSmfSubscriptionInfo - Retrieve SMF Subscription Info
func (s *SMFEventSubscriptionInfoDocumentAPIService) GetSmfSubscriptionInfo(ctx context.Context, ueId string, subsId string) (ImplResponse, error) {
	// TODO - update GetSmfSubscriptionInfo with the required logic for this service method.
	// Add api_smf_event_subscription_info_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, SmfSubscriptionInfo{}) or use other options such as http.Ok ...
	// return Response(200, SmfSubscriptionInfo{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetSmfSubscriptionInfo method not implemented")
}

// ModifySmfGroupSubscriptions - Modify SMF Subscription Info for a group of UEs or any UE
func (s *SMFEventSubscriptionInfoDocumentAPIService) ModifySmfGroupSubscriptions(ctx context.Context, ueGroupId string, subsId string, patchItem []PatchItem, supportedFeatures string) (ImplResponse, error) {
	// TODO - update ModifySmfGroupSubscriptions with the required logic for this service method.
	// Add api_smf_event_subscription_info_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(200, PatchResult{}) or use other options such as http.Ok ...
	// return Response(200, PatchResult{}), nil

	// TODO: Uncomment the next line to return response Response(403, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(403, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ModifySmfGroupSubscriptions method not implemented")
}

// ModifySmfSubscriptionInfo - Modify SMF Subscription Info
func (s *SMFEventSubscriptionInfoDocumentAPIService) ModifySmfSubscriptionInfo(ctx context.Context, ueId string, subsId string, patchItem []PatchItem, supportedFeatures string) (ImplResponse, error) {
	// TODO - update ModifySmfSubscriptionInfo with the required logic for this service method.
	// Add api_smf_event_subscription_info_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(200, PatchResult{}) or use other options such as http.Ok ...
	// return Response(200, PatchResult{}), nil

	// TODO: Uncomment the next line to return response Response(403, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(403, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ModifySmfSubscriptionInfo method not implemented")
}

// RemoveSmfGroupSubscriptions - Delete SMF Subscription Info for a group of UEs or any UE
func (s *SMFEventSubscriptionInfoDocumentAPIService) RemoveSmfGroupSubscriptions(ctx context.Context, ueGroupId string, subsId string) (ImplResponse, error) {
	// TODO - update RemoveSmfGroupSubscriptions with the required logic for this service method.
	// Add api_smf_event_subscription_info_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("RemoveSmfGroupSubscriptions method not implemented")
}

// RemoveSmfSubscriptionsInfo - Delete SMF Subscription Info
func (s *SMFEventSubscriptionInfoDocumentAPIService) RemoveSmfSubscriptionsInfo(ctx context.Context, ueId string, subsId string) (ImplResponse, error) {
	// TODO - update RemoveSmfSubscriptionsInfo with the required logic for this service method.
	// Add api_smf_event_subscription_info_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("RemoveSmfSubscriptionsInfo method not implemented")
}