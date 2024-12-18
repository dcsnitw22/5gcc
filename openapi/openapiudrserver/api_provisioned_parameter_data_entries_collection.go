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

// ProvisionedParameterDataEntriesCollectionAPIController binds http requests to an api service and writes the service results to the http response
type ProvisionedParameterDataEntriesCollectionAPIController struct {
	service      ProvisionedParameterDataEntriesCollectionAPIServicer
	errorHandler ErrorHandler
}

// ProvisionedParameterDataEntriesCollectionAPIOption for how the controller is set up.
type ProvisionedParameterDataEntriesCollectionAPIOption func(*ProvisionedParameterDataEntriesCollectionAPIController)

// WithProvisionedParameterDataEntriesCollectionAPIErrorHandler inject ErrorHandler into controller
func WithProvisionedParameterDataEntriesCollectionAPIErrorHandler(h ErrorHandler) ProvisionedParameterDataEntriesCollectionAPIOption {
	return func(c *ProvisionedParameterDataEntriesCollectionAPIController) {
		c.errorHandler = h
	}
}

// NewProvisionedParameterDataEntriesCollectionAPIController creates a default api controller
func NewProvisionedParameterDataEntriesCollectionAPIController(s ProvisionedParameterDataEntriesCollectionAPIServicer, opts ...ProvisionedParameterDataEntriesCollectionAPIOption) Router {
	controller := &ProvisionedParameterDataEntriesCollectionAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ProvisionedParameterDataEntriesCollectionAPIController
func (c *ProvisionedParameterDataEntriesCollectionAPIController) Routes() Routes {
	return Routes{
		"GetMultiplePPDataEntries": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/{ueId}/pp-data-store",
			c.GetMultiplePPDataEntries,
		},
	}
}

// GetMultiplePPDataEntries - get a list of Parameter Provisioning Data Entries
func (c *ProvisionedParameterDataEntriesCollectionAPIController) GetMultiplePPDataEntries(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.GetMultiplePPDataEntries(r.Context(), GetMultiplePpDataEntriesUeIdParameter{}, supportedFeaturesParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
