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
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// NSSAIACKDocumentAPIController binds http requests to an api service and writes the service results to the http response
type NSSAIACKDocumentAPIController struct {
	service      NSSAIACKDocumentAPIServicer
	errorHandler ErrorHandler
}

// NSSAIACKDocumentAPIOption for how the controller is set up.
type NSSAIACKDocumentAPIOption func(*NSSAIACKDocumentAPIController)

// WithNSSAIACKDocumentAPIErrorHandler inject ErrorHandler into controller
func WithNSSAIACKDocumentAPIErrorHandler(h ErrorHandler) NSSAIACKDocumentAPIOption {
	return func(c *NSSAIACKDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewNSSAIACKDocumentAPIController creates a default api controller
func NewNSSAIACKDocumentAPIController(s NSSAIACKDocumentAPIServicer, opts ...NSSAIACKDocumentAPIOption) Router {
	controller := &NSSAIACKDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the NSSAIACKDocumentAPIController
func (c *NSSAIACKDocumentAPIController) Routes() Routes {
	return Routes{
		"QueryNssaiAck": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/{ueId}/ue-update-confirmation-data/subscribed-snssais",
			c.QueryNssaiAck,
		},
	}
}

// QueryNssaiAck - Retrieves the UPU acknowledgement information of a UE
func (c *NSSAIACKDocumentAPIController) QueryNssaiAck(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	ueIdParam := params["ueId"]
	if ueIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"ueId"}, nil)
		return
	}
	var supportedFeaturesParam string
	if query.Has("supported-features") {
		supportedFeaturesParam = query.Get("supported-features")
	}
	result, err := c.service.QueryNssaiAck(r.Context(), ueIdParam, supportedFeaturesParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}