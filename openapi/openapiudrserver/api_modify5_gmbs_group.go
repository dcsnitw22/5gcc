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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Modify5GmbsGroupAPIController binds http requests to an api service and writes the service results to the http response
type Modify5GmbsGroupAPIController struct {
	service      Modify5GmbsGroupAPIServicer
	errorHandler ErrorHandler
}

// Modify5GmbsGroupAPIOption for how the controller is set up.
type Modify5GmbsGroupAPIOption func(*Modify5GmbsGroupAPIController)

// WithModify5GmbsGroupAPIErrorHandler inject ErrorHandler into controller
func WithModify5GmbsGroupAPIErrorHandler(h ErrorHandler) Modify5GmbsGroupAPIOption {
	return func(c *Modify5GmbsGroupAPIController) {
		c.errorHandler = h
	}
}

// NewModify5GmbsGroupAPIController creates a default api controller
func NewModify5GmbsGroupAPIController(s Modify5GmbsGroupAPIServicer, opts ...Modify5GmbsGroupAPIOption) Router {
	controller := &Modify5GmbsGroupAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the Modify5GmbsGroupAPIController
func (c *Modify5GmbsGroupAPIController) Routes() Routes {
	return Routes{
		"Modify5GmbsGroup": Route{
			strings.ToUpper("Patch"),
			"/nudr-dr/v2/subscription-data/group-data/mbs-group-membership/{externalGroupId}",
			c.Modify5GmbsGroup,
		},
	}
}

// Modify5GmbsGroup - modify the 5GmbsGroup
func (c *Modify5GmbsGroupAPIController) Modify5GmbsGroup(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	externalGroupIdParam := params["externalGroupId"]
	if externalGroupIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"externalGroupId"}, nil)
		return
	}
	patchItemParam := []PatchItem{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&patchItemParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	for _, el := range patchItemParam {
		if err := AssertPatchItemRequired(el); err != nil {
			c.errorHandler(w, r, err, nil)
			return
		}
	}
	var supportedFeaturesParam string
	if query.Has("supported-features") {
		supportedFeaturesParam = query.Get("supported-features")
	}
	result, err := c.service.Modify5GmbsGroup(r.Context(), externalGroupIdParam, patchItemParam, supportedFeaturesParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}