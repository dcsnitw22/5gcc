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

// SubsToNotifyCollectionAPIController binds http requests to an api service and writes the service results to the http response
type SubsToNotifyCollectionAPIController struct {
	service      SubsToNotifyCollectionAPIServicer
	errorHandler ErrorHandler
}

// SubsToNotifyCollectionAPIOption for how the controller is set up.
type SubsToNotifyCollectionAPIOption func(*SubsToNotifyCollectionAPIController)

// WithSubsToNotifyCollectionAPIErrorHandler inject ErrorHandler into controller
func WithSubsToNotifyCollectionAPIErrorHandler(h ErrorHandler) SubsToNotifyCollectionAPIOption {
	return func(c *SubsToNotifyCollectionAPIController) {
		c.errorHandler = h
	}
}

// NewSubsToNotifyCollectionAPIController creates a default api controller
func NewSubsToNotifyCollectionAPIController(s SubsToNotifyCollectionAPIServicer, opts ...SubsToNotifyCollectionAPIOption) Router {
	controller := &SubsToNotifyCollectionAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the SubsToNotifyCollectionAPIController
func (c *SubsToNotifyCollectionAPIController) Routes() Routes {
	return Routes{
		"QuerySubsToNotify": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/subs-to-notify",
			c.QuerySubsToNotify,
		},
		"RemoveMultipleSubscriptionDataSubscriptions": Route{
			strings.ToUpper("Delete"),
			"/nudr-dr/v2/subscription-data/subs-to-notify",
			c.RemoveMultipleSubscriptionDataSubscriptions,
		},
		"SubscriptionDataSubscriptions": Route{
			strings.ToUpper("Post"),
			"/nudr-dr/v2/subscription-data/subs-to-notify",
			c.SubscriptionDataSubscriptions,
		},
	}
}

// QuerySubsToNotify - Retrieves the list of subscriptions
func (c *SubsToNotifyCollectionAPIController) QuerySubsToNotify(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	ueIdParam := query.Get("ue-id")
	var supportedFeaturesParam string
	if query.Has("supported-features") {
		supportedFeaturesParam = query.Get("supported-features")
	}
	result, err := c.service.QuerySubsToNotify(r.Context(), ueIdParam, supportedFeaturesParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// RemoveMultipleSubscriptionDataSubscriptions - Deletes subscriptions identified by a given ue-id parameter
func (c *SubsToNotifyCollectionAPIController) RemoveMultipleSubscriptionDataSubscriptions(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	ueIdParam := query.Get("ue-id")
	var nfInstanceIdParam string
	if query.Has("nf-instance-id") {
		nfInstanceIdParam = query.Get("nf-instance-id")
	}
	deleteAllNfsParam, err := parseBoolParameter(
		query.Get("delete-all-nfs"),
		WithParse[bool](parseBool),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	implicitUnsubscribeIndicationParam, err := parseBoolParameter(
		query.Get("implicit-unsubscribe-indication"),
		WithParse[bool](parseBool),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.RemoveMultipleSubscriptionDataSubscriptions(r.Context(), ueIdParam, nfInstanceIdParam, deleteAllNfsParam, implicitUnsubscribeIndicationParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// SubscriptionDataSubscriptions - Subscription data subscriptions
func (c *SubsToNotifyCollectionAPIController) SubscriptionDataSubscriptions(w http.ResponseWriter, r *http.Request) {
	subscriptionDataSubscriptionsParam := SubscriptionDataSubscriptions{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&subscriptionDataSubscriptionsParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSubscriptionDataSubscriptionsRequired(subscriptionDataSubscriptionsParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertSubscriptionDataSubscriptionsConstraints(subscriptionDataSubscriptionsParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.SubscriptionDataSubscription(r.Context(), subscriptionDataSubscriptionsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
