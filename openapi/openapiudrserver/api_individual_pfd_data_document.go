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

// IndividualPFDDataDocumentAPIController binds http requests to an api service and writes the service results to the http response
type IndividualPFDDataDocumentAPIController struct {
	service      IndividualPFDDataDocumentAPIServicer
	errorHandler ErrorHandler
}

// IndividualPFDDataDocumentAPIOption for how the controller is set up.
type IndividualPFDDataDocumentAPIOption func(*IndividualPFDDataDocumentAPIController)

// WithIndividualPFDDataDocumentAPIErrorHandler inject ErrorHandler into controller
func WithIndividualPFDDataDocumentAPIErrorHandler(h ErrorHandler) IndividualPFDDataDocumentAPIOption {
	return func(c *IndividualPFDDataDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewIndividualPFDDataDocumentAPIController creates a default api controller
func NewIndividualPFDDataDocumentAPIController(s IndividualPFDDataDocumentAPIServicer, opts ...IndividualPFDDataDocumentAPIOption) Router {
	controller := &IndividualPFDDataDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the IndividualPFDDataDocumentAPIController
func (c *IndividualPFDDataDocumentAPIController) Routes() Routes {
	return Routes{
		"CreateOrReplaceIndividualPFDData": Route{
			strings.ToUpper("Put"),
			"/nudr-dr/v2/application-data/pfds/{appId}",
			c.CreateOrReplaceIndividualPFDData,
		},
		"DeleteIndividualPFDData": Route{
			strings.ToUpper("Delete"),
			"/nudr-dr/v2/application-data/pfds/{appId}",
			c.DeleteIndividualPFDData,
		},
		"ReadIndividualPFDData": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/application-data/pfds/{appId}",
			c.ReadIndividualPFDData,
		},
	}
}

// CreateOrReplaceIndividualPFDData - Create or update the corresponding PFDs for the specified application identifier
func (c *IndividualPFDDataDocumentAPIController) CreateOrReplaceIndividualPFDData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	appIdParam := params["appId"]
	if appIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"appId"}, nil)
		return
	}
	pfdDataForAppExtParam := PfdDataForAppExt{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&pfdDataForAppExtParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPfdDataForAppExtRequired(pfdDataForAppExtParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertPfdDataForAppExtConstraints(pfdDataForAppExtParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateOrReplaceIndividualPFDData(r.Context(), appIdParam, pfdDataForAppExtParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteIndividualPFDData - Delete the corresponding PFDs of the specified application identifier
func (c *IndividualPFDDataDocumentAPIController) DeleteIndividualPFDData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	appIdParam := params["appId"]
	if appIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"appId"}, nil)
		return
	}
	result, err := c.service.DeleteIndividualPFDData(r.Context(), appIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ReadIndividualPFDData - Retrieve the corresponding PFDs of the specified application identifier
func (c *IndividualPFDDataDocumentAPIController) ReadIndividualPFDData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	appIdParam := params["appId"]
	if appIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"appId"}, nil)
		return
	}
	var suppFeatParam string
	if query.Has("supp-feat") {
		suppFeatParam = query.Get("supp-feat")
	}
	result, err := c.service.ReadIndividualPFDData(r.Context(), appIdParam, suppFeatParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}