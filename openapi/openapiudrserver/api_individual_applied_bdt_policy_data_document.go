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

// IndividualAppliedBDTPolicyDataDocumentAPIController binds http requests to an api service and writes the service results to the http response
type IndividualAppliedBDTPolicyDataDocumentAPIController struct {
	service      IndividualAppliedBDTPolicyDataDocumentAPIServicer
	errorHandler ErrorHandler
}

// IndividualAppliedBDTPolicyDataDocumentAPIOption for how the controller is set up.
type IndividualAppliedBDTPolicyDataDocumentAPIOption func(*IndividualAppliedBDTPolicyDataDocumentAPIController)

// WithIndividualAppliedBDTPolicyDataDocumentAPIErrorHandler inject ErrorHandler into controller
func WithIndividualAppliedBDTPolicyDataDocumentAPIErrorHandler(h ErrorHandler) IndividualAppliedBDTPolicyDataDocumentAPIOption {
	return func(c *IndividualAppliedBDTPolicyDataDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewIndividualAppliedBDTPolicyDataDocumentAPIController creates a default api controller
func NewIndividualAppliedBDTPolicyDataDocumentAPIController(s IndividualAppliedBDTPolicyDataDocumentAPIServicer, opts ...IndividualAppliedBDTPolicyDataDocumentAPIOption) Router {
	controller := &IndividualAppliedBDTPolicyDataDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the IndividualAppliedBDTPolicyDataDocumentAPIController
func (c *IndividualAppliedBDTPolicyDataDocumentAPIController) Routes() Routes {
	return Routes{
		"CreateIndividualAppliedBdtPolicyData": Route{
			strings.ToUpper("Put"),
			"/nudr-dr/v2/application-data/bdtPolicyData/{bdtPolicyId}",
			c.CreateIndividualAppliedBdtPolicyData,
		},
		"DeleteIndividualAppliedBdtPolicyData": Route{
			strings.ToUpper("Delete"),
			"/nudr-dr/v2/application-data/bdtPolicyData/{bdtPolicyId}",
			c.DeleteIndividualAppliedBdtPolicyData,
		},
		"UpdateIndividualAppliedBdtPolicyData": Route{
			strings.ToUpper("Patch"),
			"/nudr-dr/v2/application-data/bdtPolicyData/{bdtPolicyId}",
			c.UpdateIndividualAppliedBdtPolicyData,
		},
	}
}

// CreateIndividualAppliedBdtPolicyData - Create an individual applied BDT Policy Data resource
func (c *IndividualAppliedBDTPolicyDataDocumentAPIController) CreateIndividualAppliedBdtPolicyData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bdtPolicyIdParam := params["bdtPolicyId"]
	if bdtPolicyIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"bdtPolicyId"}, nil)
		return
	}
	bdtPolicyDataParam := BdtPolicyData{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bdtPolicyDataParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertBdtPolicyDataRequired(bdtPolicyDataParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertBdtPolicyDataConstraints(bdtPolicyDataParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateIndividualAppliedBdtPolicyData(r.Context(), bdtPolicyIdParam, bdtPolicyDataParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteIndividualAppliedBdtPolicyData - Delete an individual Applied BDT Policy Data resource
func (c *IndividualAppliedBDTPolicyDataDocumentAPIController) DeleteIndividualAppliedBdtPolicyData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bdtPolicyIdParam := params["bdtPolicyId"]
	if bdtPolicyIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"bdtPolicyId"}, nil)
		return
	}
	result, err := c.service.DeleteIndividualAppliedBdtPolicyData(r.Context(), bdtPolicyIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateIndividualAppliedBdtPolicyData - Modify part of the properties of an individual Applied BDT Policy Data resource
func (c *IndividualAppliedBDTPolicyDataDocumentAPIController) UpdateIndividualAppliedBdtPolicyData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bdtPolicyIdParam := params["bdtPolicyId"]
	if bdtPolicyIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"bdtPolicyId"}, nil)
		return
	}
	bdtPolicyDataPatchParam := BdtPolicyDataPatch{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bdtPolicyDataPatchParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertBdtPolicyDataPatchRequired(bdtPolicyDataPatchParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertBdtPolicyDataPatchConstraints(bdtPolicyDataPatchParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateIndividualAppliedBdtPolicyData(r.Context(), bdtPolicyIdParam, bdtPolicyDataPatchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}