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

// QueryMulticastMbsGroupMembDocumentAPIController binds http requests to an api service and writes the service results to the http response
type QueryMulticastMbsGroupMembDocumentAPIController struct {
	service      QueryMulticastMbsGroupMembDocumentAPIServicer
	errorHandler ErrorHandler
}

// QueryMulticastMbsGroupMembDocumentAPIOption for how the controller is set up.
type QueryMulticastMbsGroupMembDocumentAPIOption func(*QueryMulticastMbsGroupMembDocumentAPIController)

// WithQueryMulticastMbsGroupMembDocumentAPIErrorHandler inject ErrorHandler into controller
func WithQueryMulticastMbsGroupMembDocumentAPIErrorHandler(h ErrorHandler) QueryMulticastMbsGroupMembDocumentAPIOption {
	return func(c *QueryMulticastMbsGroupMembDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewQueryMulticastMbsGroupMembDocumentAPIController creates a default api controller
func NewQueryMulticastMbsGroupMembDocumentAPIController(s QueryMulticastMbsGroupMembDocumentAPIServicer, opts ...QueryMulticastMbsGroupMembDocumentAPIOption) Router {
	controller := &QueryMulticastMbsGroupMembDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the QueryMulticastMbsGroupMembDocumentAPIController
func (c *QueryMulticastMbsGroupMembDocumentAPIController) Routes() Routes {
	return Routes{
		"GetMulticastMbsGroupMemb": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/group-data/mbs-group-membership/{externalGroupId}",
			c.GetMulticastMbsGroupMemb,
		},
	}
}

// GetMulticastMbsGroupMemb - Retrieve a 5GmbsGroup
func (c *QueryMulticastMbsGroupMembDocumentAPIController) GetMulticastMbsGroupMemb(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	externalGroupIdParam := params["externalGroupId"]
	if externalGroupIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"externalGroupId"}, nil)
		return
	}
	result, err := c.service.GetMulticastMbsGroupMemb(r.Context(), externalGroupIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}