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
)

// ApplicationDataSubscriptionsCollectionAPIController binds http requests to an api service and writes the service results to the http response
type ApplicationDataSubscriptionsCollectionAPIController struct {
	service      ApplicationDataSubscriptionsCollectionAPIServicer
	errorHandler ErrorHandler
}

// ApplicationDataSubscriptionsCollectionAPIOption for how the controller is set up.
type ApplicationDataSubscriptionsCollectionAPIOption func(*ApplicationDataSubscriptionsCollectionAPIController)

// WithApplicationDataSubscriptionsCollectionAPIErrorHandler inject ErrorHandler into controller
func WithApplicationDataSubscriptionsCollectionAPIErrorHandler(h ErrorHandler) ApplicationDataSubscriptionsCollectionAPIOption {
	return func(c *ApplicationDataSubscriptionsCollectionAPIController) {
		c.errorHandler = h
	}
}

// NewApplicationDataSubscriptionsCollectionAPIController creates a default api controller
func NewApplicationDataSubscriptionsCollectionAPIController(s ApplicationDataSubscriptionsCollectionAPIServicer, opts ...ApplicationDataSubscriptionsCollectionAPIOption) Router {
	controller := &ApplicationDataSubscriptionsCollectionAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ApplicationDataSubscriptionsCollectionAPIController
func (c *ApplicationDataSubscriptionsCollectionAPIController) Routes() Routes {
	return Routes{
		"CreateIndividualApplicationDataSubscription": Route{
			strings.ToUpper("Post"),
			"/nudr-dr/v2/application-data/subs-to-notify",
			c.CreateIndividualApplicationDataSubscription,
		},
		"ReadApplicationDataChangeSubscriptions": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/application-data/subs-to-notify",
			c.ReadApplicationDataChangeSubscriptions,
		},
	}
}

// CreateIndividualApplicationDataSubscription - Create a subscription to receive notification of application data changes
func (c *ApplicationDataSubscriptionsCollectionAPIController) CreateIndividualApplicationDataSubscription(w http.ResponseWriter, r *http.Request) {
	applicationDataSubsParam := ApplicationDataSubs{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&applicationDataSubsParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertApplicationDataSubsRequired(applicationDataSubsParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertApplicationDataSubsConstraints(applicationDataSubsParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateIndividualApplicationDataSubscription(r.Context(), applicationDataSubsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ReadApplicationDataChangeSubscriptions - Read Application Data change Subscriptions
func (c *ApplicationDataSubscriptionsCollectionAPIController) ReadApplicationDataChangeSubscriptions(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var dataFilterParam DataFilter
	if query.Has("data-filter") {
		// dataFilterParam = query.Get("data-filter")
	}
	result, err := c.service.ReadApplicationDataChangeSubscriptions(r.Context(), dataFilterParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}