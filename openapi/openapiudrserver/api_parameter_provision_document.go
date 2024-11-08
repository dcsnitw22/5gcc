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

// ParameterProvisionDocumentAPIController binds http requests to an api service and writes the service results to the http response
type ParameterProvisionDocumentAPIController struct {
	service      ParameterProvisionDocumentAPIServicer
	errorHandler ErrorHandler
}

// ParameterProvisionDocumentAPIOption for how the controller is set up.
type ParameterProvisionDocumentAPIOption func(*ParameterProvisionDocumentAPIController)

// WithParameterProvisionDocumentAPIErrorHandler inject ErrorHandler into controller
func WithParameterProvisionDocumentAPIErrorHandler(h ErrorHandler) ParameterProvisionDocumentAPIOption {
	return func(c *ParameterProvisionDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewParameterProvisionDocumentAPIController creates a default api controller
func NewParameterProvisionDocumentAPIController(s ParameterProvisionDocumentAPIServicer, opts ...ParameterProvisionDocumentAPIOption) Router {
	controller := &ParameterProvisionDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ParameterProvisionDocumentAPIController
func (c *ParameterProvisionDocumentAPIController) Routes() Routes {
	return Routes{
		"GetppData": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/{ueId}/pp-data",
			c.GetppData,
		},
	}
}

// GetppData - Read the profile of a given UE
func (c *ParameterProvisionDocumentAPIController) GetppData(w http.ResponseWriter, r *http.Request) {
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
	ifNoneMatchParam := r.Header.Get("If-None-Match")
	ifModifiedSinceParam := r.Header.Get("If-Modified-Since")
	result, err := c.service.GetppData(r.Context(), ueIdParam, supportedFeaturesParam, ifNoneMatchParam, ifModifiedSinceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
