/*
 * Namf_Communication
 *
 * AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_commn_server

import (
	"context"
)

// N1N2MessageCollectionDocumentAPIService is a service that implements the logic for the N1N2MessageCollectionDocumentAPIServicer
// This service should implement the business logic for every endpoint for the N1N2MessageCollectionDocumentAPI API.
// Include any external packages or services that will be required by this service.
type N1N2MessageCollectionDocumentAPIService struct {
}

// NewN1N2MessageCollectionDocumentAPIService creates a default api service
func NewN1N2MessageCollectionDocumentAPIService() N1N2MessageCollectionDocumentAPIServicer {
	return &N1N2MessageCollectionDocumentAPIService{}
}

// N1N2MessageTransfer - Namf_Communication N1N2 Message Transfer (UE Specific) service Operation
func (s *N1N2MessageCollectionDocumentAPIService) N1N2MessageTransfer(ctx context.Context, ueContextId string, n1N2MessageTransferReqData N1N2MessageTransferReqData) (ImplResponse, error) {
	// TODO - update N1N2MessageTransfer with the required logic for this service method.
	// Add api_n1_n2_message_collection_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(202, N1N2MessageTransferRspData{}) or use other options such as http.Ok ...
	// return Response(202, N1N2MessageTransferRspData{}), nil

	// TODO: Uncomment the next line to return response Response(200, N1N2MessageTransferRspData{}) or use other options such as http.Ok ...
	return Response(200, N1N2MessageTransferRspData{}), nil

	// TODO: Uncomment the next line to return response Response(307, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(307, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(400, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(400, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(403, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(403, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(404, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(404, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(409, N1N2MessageTransferError{}) or use other options such as http.Ok ...
	// return Response(409, N1N2MessageTransferError{}), nil

	// TODO: Uncomment the next line to return response Response(411, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(411, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(413, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(413, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(415, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(415, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(429, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(429, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(500, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(500, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(503, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(503, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(504, N1N2MessageTransferError{}) or use other options such as http.Ok ...
	// return Response(504, N1N2MessageTransferError{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	// return Response(http.StatusNotImplemented, nil), errors.New("N1N2MessageTransfer method not implemented")
}