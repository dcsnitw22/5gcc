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

// EventExposureDataDocumentAPIService is a service that implements the logic for the EventExposureDataDocumentAPIServicer
// This service should implement the business logic for every endpoint for the EventExposureDataDocumentAPI API.
// Include any external packages or services that will be required by this service.
type EventExposureDataDocumentAPIService struct {
}

// NewEventExposureDataDocumentAPIService creates a default api service
func NewEventExposureDataDocumentAPIService() EventExposureDataDocumentAPIServicer {
	return &EventExposureDataDocumentAPIService{}
}

// QueryEEData - Retrieves the ee profile data of a UE
func (s *EventExposureDataDocumentAPIService) QueryEEData(ctx context.Context, ueId string, fields []string, supportedFeatures string) (ImplResponse, error) {
	// TODO - update QueryEEData with the required logic for this service method.
	// Add api_event_exposure_data_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, EeProfileData{}) or use other options such as http.Ok ...
	// return Response(200, EeProfileData{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("QueryEEData method not implemented")
}
