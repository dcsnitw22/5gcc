/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.0.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiserver

import (
	"context"
	"net/http"
	"os"
)

// IndividualPDUSessionHSMFAPIRouter defines the required methods for binding the api requests to a responses for the IndividualPDUSessionHSMFAPI
// The IndividualPDUSessionHSMFAPIRouter implementation should parse necessary information from the http request,
// pass the data to a IndividualPDUSessionHSMFAPIServicer to perform the required actions, then write the service results to the http response.
type IndividualPDUSessionHSMFAPIRouter interface {
	ReleasePduSession(http.ResponseWriter, *http.Request)
	UpdatePduSession(http.ResponseWriter, *http.Request)
}

// IndividualSMContextAPIRouter defines the required methods for binding the api requests to a responses for the IndividualSMContextAPI
// The IndividualSMContextAPIRouter implementation should parse necessary information from the http request,
// pass the data to a IndividualSMContextAPIServicer to perform the required actions, then write the service results to the http response.
type IndividualSMContextAPIRouter interface {
	ReleaseSmContext(http.ResponseWriter, *http.Request)
	RetrieveSmContext(http.ResponseWriter, *http.Request)
	UpdateSmContext(http.ResponseWriter, *http.Request)
}

// PDUSessionsCollectionAPIRouter defines the required methods for binding the api requests to a responses for the PDUSessionsCollectionAPI
// The PDUSessionsCollectionAPIRouter implementation should parse necessary information from the http request,
// pass the data to a PDUSessionsCollectionAPIServicer to perform the required actions, then write the service results to the http response.
type PDUSessionsCollectionAPIRouter interface {
	PostPduSessions(http.ResponseWriter, *http.Request)
}

// SMContextsCollectionAPIRouter defines the required methods for binding the api requests to a responses for the SMContextsCollectionAPI
// The SMContextsCollectionAPIRouter implementation should parse necessary information from the http request,
// pass the data to a SMContextsCollectionAPIServicer to perform the required actions, then write the service results to the http response.
type SMContextsCollectionAPIRouter interface {
	PostSmContexts(http.ResponseWriter, *http.Request)
}

// IndividualPDUSessionHSMFAPIServicer defines the api actions for the IndividualPDUSessionHSMFAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type IndividualPDUSessionHSMFAPIServicer interface {
	ReleasePduSession(context.Context, string, ReleaseData) (ImplResponse, error)
	UpdatePduSession(context.Context, string, HsmfUpdateData) (ImplResponse, error)
}

// IndividualSMContextAPIServicer defines the api actions for the IndividualSMContextAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type IndividualSMContextAPIServicer interface {
	ReleaseSmContext(context.Context, string, SmContextReleaseData) (ImplResponse, error)
	RetrieveSmContext(context.Context, string, SmContextRetrieveData) (ImplResponse, error)
	UpdateSmContext(context.Context, string, SmContextUpdateData) (ImplResponse, error)
}

// PDUSessionsCollectionAPIServicer defines the api actions for the PDUSessionsCollectionAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type PDUSessionsCollectionAPIServicer interface {
	PostPduSessions(context.Context, PduSessionCreateData) (ImplResponse, error)
}

// SMContextsCollectionAPIServicer defines the api actions for the SMContextsCollectionAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type SMContextsCollectionAPIServicer interface {
	PostSmContexts(context.Context, SmContextCreateData, *os.File) (ImplResponse, error)
}