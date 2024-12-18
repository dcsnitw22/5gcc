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

// Query5GVnGroupConfigurationDocumentAPIController binds http requests to an api service and writes the service results to the http response
type Query5GVnGroupConfigurationDocumentAPIController struct {
	service      Query5GVnGroupConfigurationDocumentAPIServicer
	errorHandler ErrorHandler
}

// Query5GVnGroupConfigurationDocumentAPIOption for how the controller is set up.
type Query5GVnGroupConfigurationDocumentAPIOption func(*Query5GVnGroupConfigurationDocumentAPIController)

// WithQuery5GVnGroupConfigurationDocumentAPIErrorHandler inject ErrorHandler into controller
func WithQuery5GVnGroupConfigurationDocumentAPIErrorHandler(h ErrorHandler) Query5GVnGroupConfigurationDocumentAPIOption {
	return func(c *Query5GVnGroupConfigurationDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewQuery5GVnGroupConfigurationDocumentAPIController creates a default api controller
func NewQuery5GVnGroupConfigurationDocumentAPIController(s Query5GVnGroupConfigurationDocumentAPIServicer, opts ...Query5GVnGroupConfigurationDocumentAPIOption) Router {
	controller := &Query5GVnGroupConfigurationDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the Query5GVnGroupConfigurationDocumentAPIController
func (c *Query5GVnGroupConfigurationDocumentAPIController) Routes() Routes {
	return Routes{
		"Get5GVnGroupConfiguration": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/group-data/5g-vn-groups/{externalGroupId}",
			c.Get5GVnGroupConfiguration,
		},
	}
}

// Get5GVnGroupConfiguration - Retrieve a 5GVnGroup configuration
func (c *Query5GVnGroupConfigurationDocumentAPIController) Get5GVnGroupConfiguration(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	externalGroupIdParam := params["externalGroupId"]
	if externalGroupIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"externalGroupId"}, nil)
		return
	}
	result, err := c.service.Get5GVnGroupConfiguration(r.Context(), externalGroupIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
