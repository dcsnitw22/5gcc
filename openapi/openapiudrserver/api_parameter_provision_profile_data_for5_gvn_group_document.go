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

// ParameterProvisionProfileDataFor5GVNGroupDocumentAPIController binds http requests to an api service and writes the service results to the http response
type ParameterProvisionProfileDataFor5GVNGroupDocumentAPIController struct {
	service      ParameterProvisionProfileDataFor5GVNGroupDocumentAPIServicer
	errorHandler ErrorHandler
}

// ParameterProvisionProfileDataFor5GVNGroupDocumentAPIOption for how the controller is set up.
type ParameterProvisionProfileDataFor5GVNGroupDocumentAPIOption func(*ParameterProvisionProfileDataFor5GVNGroupDocumentAPIController)

// WithParameterProvisionProfileDataFor5GVNGroupDocumentAPIErrorHandler inject ErrorHandler into controller
func WithParameterProvisionProfileDataFor5GVNGroupDocumentAPIErrorHandler(h ErrorHandler) ParameterProvisionProfileDataFor5GVNGroupDocumentAPIOption {
	return func(c *ParameterProvisionProfileDataFor5GVNGroupDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewParameterProvisionProfileDataFor5GVNGroupDocumentAPIController creates a default api controller
func NewParameterProvisionProfileDataFor5GVNGroupDocumentAPIController(s ParameterProvisionProfileDataFor5GVNGroupDocumentAPIServicer, opts ...ParameterProvisionProfileDataFor5GVNGroupDocumentAPIOption) Router {
	controller := &ParameterProvisionProfileDataFor5GVNGroupDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ParameterProvisionProfileDataFor5GVNGroupDocumentAPIController
func (c *ParameterProvisionProfileDataFor5GVNGroupDocumentAPIController) Routes() Routes {
	return Routes{
		"Query5GVNGroupPPData": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/group-data/5g-vn-groups/pp-profile-data",
			c.Query5GVNGroupPPData,
		},
	}
}

// Query5GVNGroupPPData - Retrieves the parameter provision profile data for 5G VN Group
func (c *ParameterProvisionProfileDataFor5GVNGroupDocumentAPIController) Query5GVNGroupPPData(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var extGroupIdsParam []string
	if query.Has("ext-group-ids") {
		extGroupIdsParam = strings.Split(query.Get("ext-group-ids"), ",")
	}
	var supportedFeaturesParam string
	if query.Has("supported-features") {
		supportedFeaturesParam = query.Get("supported-features")
	}
	result, err := c.service.Query5GVNGroupPPData(r.Context(), extGroupIdsParam, supportedFeaturesParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
