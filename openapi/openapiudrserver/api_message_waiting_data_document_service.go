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

// MessageWaitingDataDocumentAPIService is a service that implements the logic for the MessageWaitingDataDocumentAPIServicer
// This service should implement the business logic for every endpoint for the MessageWaitingDataDocumentAPI API.
// Include any external packages or services that will be required by this service.
type MessageWaitingDataDocumentAPIService struct {
}

// NewMessageWaitingDataDocumentAPIService creates a default api service
func NewMessageWaitingDataDocumentAPIService() MessageWaitingDataDocumentAPIServicer {
	return &MessageWaitingDataDocumentAPIService{}
}

// CreateMessageWaitingData - Create the Message Waiting Data of the UE
func (s *MessageWaitingDataDocumentAPIService) CreateMessageWaitingData(ctx context.Context, ueId string, messageWaitingData MessageWaitingData) (ImplResponse, error) {
	// TODO - update CreateMessageWaitingData with the required logic for this service method.
	// Add api_message_waiting_data_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, MessageWaitingData{}) or use other options such as http.Ok ...
	// return Response(201, MessageWaitingData{}), nil

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateMessageWaitingData method not implemented")
}

// DeleteMessageWaitingData - To remove the Message Waiting Data of the UE
func (s *MessageWaitingDataDocumentAPIService) DeleteMessageWaitingData(ctx context.Context, ueId string) (ImplResponse, error) {
	// TODO - update DeleteMessageWaitingData with the required logic for this service method.
	// Add api_message_waiting_data_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteMessageWaitingData method not implemented")
}

// ModifyMessageWaitingData - Modify the Message Waiting Data of the UE
func (s *MessageWaitingDataDocumentAPIService) ModifyMessageWaitingData(ctx context.Context, ueId string, patchItem []PatchItem) (ImplResponse, error) {
	// TODO - update ModifyMessageWaitingData with the required logic for this service method.
	// Add api_message_waiting_data_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(403, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(403, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(404, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(404, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ModifyMessageWaitingData method not implemented")
}

// QueryMessageWaitingData - Retrieves the Message Waiting Data of the UE
func (s *MessageWaitingDataDocumentAPIService) QueryMessageWaitingData(ctx context.Context, ueId string, fields []string, supportedFeatures string) (ImplResponse, error) {
	// TODO - update QueryMessageWaitingData with the required logic for this service method.
	// Add api_message_waiting_data_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, MessageWaitingData{}) or use other options such as http.Ok ...
	// return Response(200, MessageWaitingData{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("QueryMessageWaitingData method not implemented")
}