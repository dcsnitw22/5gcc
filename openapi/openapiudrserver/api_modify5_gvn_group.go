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

// Modify5GVnGroupAPIController binds http requests to an api service and writes the service results to the http response
type Modify5GVnGroupAPIController struct {
	service      Modify5GVnGroupAPIServicer
	errorHandler ErrorHandler
}

// Modify5GVnGroupAPIOption for how the controller is set up.
type Modify5GVnGroupAPIOption func(*Modify5GVnGroupAPIController)

// WithModify5GVnGroupAPIErrorHandler inject ErrorHandler into controller
func WithModify5GVnGroupAPIErrorHandler(h ErrorHandler) Modify5GVnGroupAPIOption {
	return func(c *Modify5GVnGroupAPIController) {
		c.errorHandler = h
	}
}

// NewModify5GVnGroupAPIController creates a default api controller
func NewModify5GVnGroupAPIController(s Modify5GVnGroupAPIServicer, opts ...Modify5GVnGroupAPIOption) Router {
	controller := &Modify5GVnGroupAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the Modify5GVnGroupAPIController
func (c *Modify5GVnGroupAPIController) Routes() Routes {
	return Routes{
		"Modify5GVnGroup": Route{
			strings.ToUpper("Patch"),
			"/nudr-dr/v2/subscription-data/group-data/5g-vn-groups/{externalGroupId}",
			c.Modify5GVnGroup,
		},
	}
}

// Modify5GVnGroup - modify the 5GVnGroup
func (c *Modify5GVnGroupAPIController) Modify5GVnGroup(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.Modify5GVnGroup(r.Context(), externalGroupIdParam, patchItemParam, supportedFeaturesParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}