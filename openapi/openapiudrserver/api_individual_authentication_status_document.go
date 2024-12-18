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

// IndividualAuthenticationStatusDocumentAPIController binds http requests to an api service and writes the service results to the http response
type IndividualAuthenticationStatusDocumentAPIController struct {
	service      IndividualAuthenticationStatusDocumentAPIServicer
	errorHandler ErrorHandler
}

// IndividualAuthenticationStatusDocumentAPIOption for how the controller is set up.
type IndividualAuthenticationStatusDocumentAPIOption func(*IndividualAuthenticationStatusDocumentAPIController)

// WithIndividualAuthenticationStatusDocumentAPIErrorHandler inject ErrorHandler into controller
func WithIndividualAuthenticationStatusDocumentAPIErrorHandler(h ErrorHandler) IndividualAuthenticationStatusDocumentAPIOption {
	return func(c *IndividualAuthenticationStatusDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewIndividualAuthenticationStatusDocumentAPIController creates a default api controller
func NewIndividualAuthenticationStatusDocumentAPIController(s IndividualAuthenticationStatusDocumentAPIServicer, opts ...IndividualAuthenticationStatusDocumentAPIOption) Router {
	controller := &IndividualAuthenticationStatusDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the IndividualAuthenticationStatusDocumentAPIController
func (c *IndividualAuthenticationStatusDocumentAPIController) Routes() Routes {
	return Routes{
		"CreateIndividualAuthenticationStatus": Route{
			strings.ToUpper("Put"),
			"/nudr-dr/v2/subscription-data/{ueId}/authentication-data/authentication-status/{servingNetworkName}",
			c.CreateIndividualAuthenticationStatus,
		},
	}
}

// CreateIndividualAuthenticationStatus - To store the individual Authentication Status data of a UE
func (c *IndividualAuthenticationStatusDocumentAPIController) CreateIndividualAuthenticationStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ueIdParam := params["ueId"]
	if ueIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"ueId"}, nil)
		return
	}
	servingNetworkNameParam := params["servingNetworkName"]
	if servingNetworkNameParam == "" {
		c.errorHandler(w, r, &RequiredError{"servingNetworkName"}, nil)
		return
	}
	authEventParam := AuthEvent{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&authEventParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertAuthEventRequired(authEventParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertAuthEventConstraints(authEventParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateIndividualAuthenticationStatus(r.Context(), ueIdParam, servingNetworkNameParam, authEventParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
