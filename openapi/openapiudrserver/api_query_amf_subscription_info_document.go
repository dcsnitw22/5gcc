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

// QueryAMFSubscriptionInfoDocumentAPIController binds http requests to an api service and writes the service results to the http response
type QueryAMFSubscriptionInfoDocumentAPIController struct {
	service      QueryAMFSubscriptionInfoDocumentAPIServicer
	errorHandler ErrorHandler
}

// QueryAMFSubscriptionInfoDocumentAPIOption for how the controller is set up.
type QueryAMFSubscriptionInfoDocumentAPIOption func(*QueryAMFSubscriptionInfoDocumentAPIController)

// WithQueryAMFSubscriptionInfoDocumentAPIErrorHandler inject ErrorHandler into controller
func WithQueryAMFSubscriptionInfoDocumentAPIErrorHandler(h ErrorHandler) QueryAMFSubscriptionInfoDocumentAPIOption {
	return func(c *QueryAMFSubscriptionInfoDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewQueryAMFSubscriptionInfoDocumentAPIController creates a default api controller
func NewQueryAMFSubscriptionInfoDocumentAPIController(s QueryAMFSubscriptionInfoDocumentAPIServicer, opts ...QueryAMFSubscriptionInfoDocumentAPIOption) Router {
	controller := &QueryAMFSubscriptionInfoDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the QueryAMFSubscriptionInfoDocumentAPIController
func (c *QueryAMFSubscriptionInfoDocumentAPIController) Routes() Routes {
	return Routes{
		"GetAmfGroupSubscriptions": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/amf-subscriptions",
			c.GetAmfGroupSubscriptions,
		},
		"GetAmfSubscriptionInfo": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/amf-subscriptions",
			c.GetAmfSubscriptionInfo,
		},
	}
}

// GetAmfGroupSubscriptions - Retrieve AMF subscription Info for a group of UEs or any UE
func (c *QueryAMFSubscriptionInfoDocumentAPIController) GetAmfGroupSubscriptions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ueGroupIdParam := params["ueGroupId"]
	if ueGroupIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"ueGroupId"}, nil)
		return
	}
	subsIdParam := params["subsId"]
	if subsIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subsId"}, nil)
		return
	}
	result, err := c.service.GetAmfGroupSubscriptions(r.Context(), ueGroupIdParam, subsIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetAmfSubscriptionInfo - Retrieve AMF subscription Info
func (c *QueryAMFSubscriptionInfoDocumentAPIController) GetAmfSubscriptionInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ueIdParam := params["ueId"]
	if ueIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"ueId"}, nil)
		return
	}
	subsIdParam := params["subsId"]
	if subsIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subsId"}, nil)
		return
	}
	result, err := c.service.GetAmfSubscriptionInfo(r.Context(), ueIdParam, subsIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
