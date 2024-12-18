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

// QueryODBDataBySUPIOrGPSIDocumentAPIController binds http requests to an api service and writes the service results to the http response
type QueryODBDataBySUPIOrGPSIDocumentAPIController struct {
	service      QueryODBDataBySUPIOrGPSIDocumentAPIServicer
	errorHandler ErrorHandler
}

// QueryODBDataBySUPIOrGPSIDocumentAPIOption for how the controller is set up.
type QueryODBDataBySUPIOrGPSIDocumentAPIOption func(*QueryODBDataBySUPIOrGPSIDocumentAPIController)

// WithQueryODBDataBySUPIOrGPSIDocumentAPIErrorHandler inject ErrorHandler into controller
func WithQueryODBDataBySUPIOrGPSIDocumentAPIErrorHandler(h ErrorHandler) QueryODBDataBySUPIOrGPSIDocumentAPIOption {
	return func(c *QueryODBDataBySUPIOrGPSIDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewQueryODBDataBySUPIOrGPSIDocumentAPIController creates a default api controller
func NewQueryODBDataBySUPIOrGPSIDocumentAPIController(s QueryODBDataBySUPIOrGPSIDocumentAPIServicer, opts ...QueryODBDataBySUPIOrGPSIDocumentAPIOption) Router {
	controller := &QueryODBDataBySUPIOrGPSIDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the QueryODBDataBySUPIOrGPSIDocumentAPIController
func (c *QueryODBDataBySUPIOrGPSIDocumentAPIController) Routes() Routes {
	return Routes{
		"GetOdbData": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/{ueId}/operator-determined-barring-data",
			c.GetOdbData,
		},
	}
}

// GetOdbData - Retrieve ODB Data data by SUPI or GPSI
func (c *QueryODBDataBySUPIOrGPSIDocumentAPIController) GetOdbData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ueIdParam := params["ueId"]
	if ueIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"ueId"}, nil)
		return
	}
	result, err := c.service.GetOdbData(r.Context(), ueIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
