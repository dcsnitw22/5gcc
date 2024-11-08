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

// ParameterProvisionProfileDataFor5GMBSGroupDocumentAPIController binds http requests to an api service and writes the service results to the http response
type ParameterProvisionProfileDataFor5GMBSGroupDocumentAPIController struct {
	service      ParameterProvisionProfileDataFor5GMBSGroupDocumentAPIServicer
	errorHandler ErrorHandler
}

// ParameterProvisionProfileDataFor5GMBSGroupDocumentAPIOption for how the controller is set up.
type ParameterProvisionProfileDataFor5GMBSGroupDocumentAPIOption func(*ParameterProvisionProfileDataFor5GMBSGroupDocumentAPIController)

// WithParameterProvisionProfileDataFor5GMBSGroupDocumentAPIErrorHandler inject ErrorHandler into controller
func WithParameterProvisionProfileDataFor5GMBSGroupDocumentAPIErrorHandler(h ErrorHandler) ParameterProvisionProfileDataFor5GMBSGroupDocumentAPIOption {
	return func(c *ParameterProvisionProfileDataFor5GMBSGroupDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewParameterProvisionProfileDataFor5GMBSGroupDocumentAPIController creates a default api controller
func NewParameterProvisionProfileDataFor5GMBSGroupDocumentAPIController(s ParameterProvisionProfileDataFor5GMBSGroupDocumentAPIServicer, opts ...ParameterProvisionProfileDataFor5GMBSGroupDocumentAPIOption) Router {
	controller := &ParameterProvisionProfileDataFor5GMBSGroupDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ParameterProvisionProfileDataFor5GMBSGroupDocumentAPIController
func (c *ParameterProvisionProfileDataFor5GMBSGroupDocumentAPIController) Routes() Routes {
	return Routes{
		"Query5GMbsGroupPPData": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/group-data/mbs-group-membership/pp-profile-data",
			c.Query5GMbsGroupPPData,
		},
	}
}

// Query5GMbsGroupPPData - Retrieves the parameter provision profile data for 5G MBS Group
func (c *ParameterProvisionProfileDataFor5GMBSGroupDocumentAPIController) Query5GMbsGroupPPData(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var extGroupIdsParam []string
	if query.Has("ext-group-ids") {
		extGroupIdsParam = strings.Split(query.Get("ext-group-ids"), ",")
	}
	var supportedFeaturesParam string
	if query.Has("supported-features") {
		supportedFeaturesParam = query.Get("supported-features")
	}
	result, err := c.service.Query5GMbsGroupPPData(r.Context(), extGroupIdsParam, supportedFeaturesParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
