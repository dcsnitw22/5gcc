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

// EventExposureGroupSubscriptionsCollectionAPIService is a service that implements the logic for the EventExposureGroupSubscriptionsCollectionAPIServicer
// This service should implement the business logic for every endpoint for the EventExposureGroupSubscriptionsCollectionAPI API.
// Include any external packages or services that will be required by this service.
type EventExposureGroupSubscriptionsCollectionAPIService struct {
}

// NewEventExposureGroupSubscriptionsCollectionAPIService creates a default api service
func NewEventExposureGroupSubscriptionsCollectionAPIService() EventExposureGroupSubscriptionsCollectionAPIServicer {
	return &EventExposureGroupSubscriptionsCollectionAPIService{}
}

// CreateEeGroupSubscriptions - Create individual EE subscription for a group of UEs or any UE
func (s *EventExposureGroupSubscriptionsCollectionAPIService) CreateEeGroupSubscriptions(ctx context.Context, ueGroupId string, eeSubscription EeSubscription) (ImplResponse, error) {
	// TODO - update CreateEeGroupSubscriptions with the required logic for this service method.
	// Add api_event_exposure_group_subscriptions_collection_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, EeSubscription{}) or use other options such as http.Ok ...
	// return Response(201, EeSubscription{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateEeGroupSubscriptions method not implemented")
}

// QueryEeGroupSubscriptions - Retrieves the ee subscriptions of a group of UEs or any UE
func (s *EventExposureGroupSubscriptionsCollectionAPIService) QueryEeGroupSubscriptions(ctx context.Context, ueGroupId string, supportedFeatures string) (ImplResponse, error) {
	// TODO - update QueryEeGroupSubscriptions with the required logic for this service method.
	// Add api_event_exposure_group_subscriptions_collection_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []EeSubscription{}) or use other options such as http.Ok ...
	// return Response(200, []EeSubscription{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("QueryEeGroupSubscriptions method not implemented")
}