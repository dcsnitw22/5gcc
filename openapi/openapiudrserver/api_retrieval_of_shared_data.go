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

// RetrievalOfSharedDataAPIController binds http requests to an api service and writes the service results to the http response
type RetrievalOfSharedDataAPIController struct {
	service      RetrievalOfSharedDataAPIServicer
	errorHandler ErrorHandler
}

// RetrievalOfSharedDataAPIOption for how the controller is set up.
type RetrievalOfSharedDataAPIOption func(*RetrievalOfSharedDataAPIController)

// WithRetrievalOfSharedDataAPIErrorHandler inject ErrorHandler into controller
func WithRetrievalOfSharedDataAPIErrorHandler(h ErrorHandler) RetrievalOfSharedDataAPIOption {
	return func(c *RetrievalOfSharedDataAPIController) {
		c.errorHandler = h
	}
}

// NewRetrievalOfSharedDataAPIController creates a default api controller
func NewRetrievalOfSharedDataAPIController(s RetrievalOfSharedDataAPIServicer, opts ...RetrievalOfSharedDataAPIOption) Router {
	controller := &RetrievalOfSharedDataAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the RetrievalOfSharedDataAPIController
func (c *RetrievalOfSharedDataAPIController) Routes() Routes {
	return Routes{
		"GetSharedData": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/shared-data",
			c.GetSharedData,
		},
	}
}

// GetSharedData - retrieve shared data
func (c *RetrievalOfSharedDataAPIController) GetSharedData(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var sharedDataIdsParam []string
	if query.Has("shared-data-ids") {
		sharedDataIdsParam = strings.Split(query.Get("shared-data-ids"), ",")
	}
	var supportedFeaturesParam string
	if query.Has("supported-features") {
		supportedFeaturesParam = query.Get("supported-features")
	}
	result, err := c.service.GetSharedData(r.Context(), sharedDataIdsParam, supportedFeaturesParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}