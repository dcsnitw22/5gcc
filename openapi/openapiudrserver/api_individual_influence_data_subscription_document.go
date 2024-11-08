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

// IndividualInfluenceDataSubscriptionDocumentAPIController binds http requests to an api service and writes the service results to the http response
type IndividualInfluenceDataSubscriptionDocumentAPIController struct {
	service      IndividualInfluenceDataSubscriptionDocumentAPIServicer
	errorHandler ErrorHandler
}

// IndividualInfluenceDataSubscriptionDocumentAPIOption for how the controller is set up.
type IndividualInfluenceDataSubscriptionDocumentAPIOption func(*IndividualInfluenceDataSubscriptionDocumentAPIController)

// WithIndividualInfluenceDataSubscriptionDocumentAPIErrorHandler inject ErrorHandler into controller
func WithIndividualInfluenceDataSubscriptionDocumentAPIErrorHandler(h ErrorHandler) IndividualInfluenceDataSubscriptionDocumentAPIOption {
	return func(c *IndividualInfluenceDataSubscriptionDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewIndividualInfluenceDataSubscriptionDocumentAPIController creates a default api controller
func NewIndividualInfluenceDataSubscriptionDocumentAPIController(s IndividualInfluenceDataSubscriptionDocumentAPIServicer, opts ...IndividualInfluenceDataSubscriptionDocumentAPIOption) Router {
	controller := &IndividualInfluenceDataSubscriptionDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the IndividualInfluenceDataSubscriptionDocumentAPIController
func (c *IndividualInfluenceDataSubscriptionDocumentAPIController) Routes() Routes {
	return Routes{
		"DeleteIndividualInfluenceDataSubscription": Route{
			strings.ToUpper("Delete"),
			"/nudr-dr/v2/application-data/influenceData/subs-to-notify/{subscriptionId}",
			c.DeleteIndividualInfluenceDataSubscription,
		},
		"ReadIndividualInfluenceDataSubscription": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/application-data/influenceData/subs-to-notify/{subscriptionId}",
			c.ReadIndividualInfluenceDataSubscription,
		},
		"ReplaceIndividualInfluenceDataSubscription": Route{
			strings.ToUpper("Put"),
			"/nudr-dr/v2/application-data/influenceData/subs-to-notify/{subscriptionId}",
			c.ReplaceIndividualInfluenceDataSubscription,
		},
	}
}

// DeleteIndividualInfluenceDataSubscription - Delete an individual Influence Data Subscription resource
func (c *IndividualInfluenceDataSubscriptionDocumentAPIController) DeleteIndividualInfluenceDataSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscriptionId"}, nil)
		return
	}
	result, err := c.service.DeleteIndividualInfluenceDataSubscription(r.Context(), subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ReadIndividualInfluenceDataSubscription - Get an existing individual Influence Data Subscription resource
func (c *IndividualInfluenceDataSubscriptionDocumentAPIController) ReadIndividualInfluenceDataSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscriptionId"}, nil)
		return
	}
	result, err := c.service.ReadIndividualInfluenceDataSubscription(r.Context(), subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ReplaceIndividualInfluenceDataSubscription - Modify an existing individual Influence Data Subscription resource
func (c *IndividualInfluenceDataSubscriptionDocumentAPIController) ReplaceIndividualInfluenceDataSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscriptionId"}, nil)
		return
	}
	trafficInfluSubParam := TrafficInfluSub{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&trafficInfluSubParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTrafficInfluSubRequired(trafficInfluSubParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertTrafficInfluSubConstraints(trafficInfluSubParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ReplaceIndividualInfluenceDataSubscription(r.Context(), subscriptionIdParam, trafficInfluSubParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
