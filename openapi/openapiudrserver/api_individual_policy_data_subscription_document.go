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

// IndividualPolicyDataSubscriptionDocumentAPIController binds http requests to an api service and writes the service results to the http response
type IndividualPolicyDataSubscriptionDocumentAPIController struct {
	service      IndividualPolicyDataSubscriptionDocumentAPIServicer
	errorHandler ErrorHandler
}

// IndividualPolicyDataSubscriptionDocumentAPIOption for how the controller is set up.
type IndividualPolicyDataSubscriptionDocumentAPIOption func(*IndividualPolicyDataSubscriptionDocumentAPIController)

// WithIndividualPolicyDataSubscriptionDocumentAPIErrorHandler inject ErrorHandler into controller
func WithIndividualPolicyDataSubscriptionDocumentAPIErrorHandler(h ErrorHandler) IndividualPolicyDataSubscriptionDocumentAPIOption {
	return func(c *IndividualPolicyDataSubscriptionDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewIndividualPolicyDataSubscriptionDocumentAPIController creates a default api controller
func NewIndividualPolicyDataSubscriptionDocumentAPIController(s IndividualPolicyDataSubscriptionDocumentAPIServicer, opts ...IndividualPolicyDataSubscriptionDocumentAPIOption) Router {
	controller := &IndividualPolicyDataSubscriptionDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the IndividualPolicyDataSubscriptionDocumentAPIController
func (c *IndividualPolicyDataSubscriptionDocumentAPIController) Routes() Routes {
	return Routes{
		"DeleteIndividualPolicyDataSubscription": Route{
			strings.ToUpper("Delete"),
			"/nudr-dr/v2/policy-data/subs-to-notify/{subsId}",
			c.DeleteIndividualPolicyDataSubscription,
		},
		"ReplaceIndividualPolicyDataSubscription": Route{
			strings.ToUpper("Put"),
			"/nudr-dr/v2/policy-data/subs-to-notify/{subsId}",
			c.ReplaceIndividualPolicyDataSubscription,
		},
	}
}

// DeleteIndividualPolicyDataSubscription - Delete the individual Policy Data subscription
func (c *IndividualPolicyDataSubscriptionDocumentAPIController) DeleteIndividualPolicyDataSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subsIdParam := params["subsId"]
	if subsIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subsId"}, nil)
		return
	}
	result, err := c.service.DeleteIndividualPolicyDataSubscription(r.Context(), subsIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ReplaceIndividualPolicyDataSubscription - Modify a subscription to receive notification of policy data changes
func (c *IndividualPolicyDataSubscriptionDocumentAPIController) ReplaceIndividualPolicyDataSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subsIdParam := params["subsId"]
	if subsIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subsId"}, nil)
		return
	}
	policyDataSubscriptionParam := PolicyDataSubscription{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&policyDataSubscriptionParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPolicyDataSubscriptionRequired(policyDataSubscriptionParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertPolicyDataSubscriptionConstraints(policyDataSubscriptionParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ReplaceIndividualPolicyDataSubscription(r.Context(), subsIdParam, policyDataSubscriptionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
