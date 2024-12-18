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
)

// Class5GVNGroupsInternalDocumentAPIController binds http requests to an api service and writes the service results to the http response
type Class5GVNGroupsInternalDocumentAPIController struct {
	service      Class5GVNGroupsInternalDocumentAPIServicer
	errorHandler ErrorHandler
}

// Class5GVNGroupsInternalDocumentAPIOption for how the controller is set up.
type Class5GVNGroupsInternalDocumentAPIOption func(*Class5GVNGroupsInternalDocumentAPIController)

// WithClass5GVNGroupsInternalDocumentAPIErrorHandler inject ErrorHandler into controller
func WithClass5GVNGroupsInternalDocumentAPIErrorHandler(h ErrorHandler) Class5GVNGroupsInternalDocumentAPIOption {
	return func(c *Class5GVNGroupsInternalDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewClass5GVNGroupsInternalDocumentAPIController creates a default api controller
func NewClass5GVNGroupsInternalDocumentAPIController(s Class5GVNGroupsInternalDocumentAPIServicer, opts ...Class5GVNGroupsInternalDocumentAPIOption) Router {
	controller := &Class5GVNGroupsInternalDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the Class5GVNGroupsInternalDocumentAPIController
func (c *Class5GVNGroupsInternalDocumentAPIController) Routes() Routes {
	return Routes{
		"Query5GVnGroupInternal": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/group-data/5g-vn-groups/internal",
			c.Query5GVnGroupInternal,
		},
	}
}

// Query5GVnGroupInternal - Retrieves the data of 5G VN Group
func (c *Class5GVNGroupsInternalDocumentAPIController) Query5GVnGroupInternal(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var internalGroupIdsParam []string
	if query.Has("internal-group-ids") {
		internalGroupIdsParam = strings.Split(query.Get("internal-group-ids"), ",")
	}
	result, err := c.service.Query5GVnGroupInternal(r.Context(), internalGroupIdsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
