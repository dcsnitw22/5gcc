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

// LCSPrivacySubscriptionDataAPIController binds http requests to an api service and writes the service results to the http response
type LCSPrivacySubscriptionDataAPIController struct {
	service      LCSPrivacySubscriptionDataAPIServicer
	errorHandler ErrorHandler
}

// LCSPrivacySubscriptionDataAPIOption for how the controller is set up.
type LCSPrivacySubscriptionDataAPIOption func(*LCSPrivacySubscriptionDataAPIController)

// WithLCSPrivacySubscriptionDataAPIErrorHandler inject ErrorHandler into controller
func WithLCSPrivacySubscriptionDataAPIErrorHandler(h ErrorHandler) LCSPrivacySubscriptionDataAPIOption {
	return func(c *LCSPrivacySubscriptionDataAPIController) {
		c.errorHandler = h
	}
}

// NewLCSPrivacySubscriptionDataAPIController creates a default api controller
func NewLCSPrivacySubscriptionDataAPIController(s LCSPrivacySubscriptionDataAPIServicer, opts ...LCSPrivacySubscriptionDataAPIOption) Router {
	controller := &LCSPrivacySubscriptionDataAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the LCSPrivacySubscriptionDataAPIController
func (c *LCSPrivacySubscriptionDataAPIController) Routes() Routes {
	return Routes{
		"QueryLcsPrivacyData": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/{ueId}/lcs-privacy-data",
			c.QueryLcsPrivacyData,
		},
	}
}

// QueryLcsPrivacyData - Retrieves the LCS Privacy subscription data of a UE
func (c *LCSPrivacySubscriptionDataAPIController) QueryLcsPrivacyData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	ueIdParam := params["ueId"]
	if ueIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"ueId"}, nil)
		return
	}
	var fieldsParam []string
	if query.Has("fields") {
		fieldsParam = strings.Split(query.Get("fields"), ",")
	}
	var supportedFeaturesParam string
	if query.Has("supported-features") {
		supportedFeaturesParam = query.Get("supported-features")
	}
	ifNoneMatchParam := r.Header.Get("If-None-Match")
	ifModifiedSinceParam := r.Header.Get("If-Modified-Since")
	result, err := c.service.QueryLcsPrivacyData(r.Context(), ueIdParam, fieldsParam, supportedFeaturesParam, ifNoneMatchParam, ifModifiedSinceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}