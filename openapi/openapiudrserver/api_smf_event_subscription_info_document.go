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

// SMFEventSubscriptionInfoDocumentAPIController binds http requests to an api service and writes the service results to the http response
type SMFEventSubscriptionInfoDocumentAPIController struct {
	service      SMFEventSubscriptionInfoDocumentAPIServicer
	errorHandler ErrorHandler
}

// SMFEventSubscriptionInfoDocumentAPIOption for how the controller is set up.
type SMFEventSubscriptionInfoDocumentAPIOption func(*SMFEventSubscriptionInfoDocumentAPIController)

// WithSMFEventSubscriptionInfoDocumentAPIErrorHandler inject ErrorHandler into controller
func WithSMFEventSubscriptionInfoDocumentAPIErrorHandler(h ErrorHandler) SMFEventSubscriptionInfoDocumentAPIOption {
	return func(c *SMFEventSubscriptionInfoDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewSMFEventSubscriptionInfoDocumentAPIController creates a default api controller
func NewSMFEventSubscriptionInfoDocumentAPIController(s SMFEventSubscriptionInfoDocumentAPIServicer, opts ...SMFEventSubscriptionInfoDocumentAPIOption) Router {
	controller := &SMFEventSubscriptionInfoDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the SMFEventSubscriptionInfoDocumentAPIController
func (c *SMFEventSubscriptionInfoDocumentAPIController) Routes() Routes {
	return Routes{
		"CreateSMFSubscriptions": Route{
			strings.ToUpper("Put"),
			"/nudr-dr/v2/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/smf-subscriptions",
			c.CreateSMFSubscriptions,
		},
		"GetSmfGroupSubscriptions": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/smf-subscriptions",
			c.GetSmfGroupSubscriptions,
		},
		"GetSmfSubscriptionInfo": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/smf-subscriptions",
			c.GetSmfSubscriptionInfo,
		},
		"ModifySmfGroupSubscriptions": Route{
			strings.ToUpper("Patch"),
			"/nudr-dr/v2/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/smf-subscriptions",
			c.ModifySmfGroupSubscriptions,
		},
		"ModifySmfSubscriptionInfo": Route{
			strings.ToUpper("Patch"),
			"/nudr-dr/v2/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/smf-subscriptions",
			c.ModifySmfSubscriptionInfo,
		},
		"RemoveSmfGroupSubscriptions": Route{
			strings.ToUpper("Delete"),
			"/nudr-dr/v2/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/smf-subscriptions",
			c.RemoveSmfGroupSubscriptions,
		},
		"RemoveSmfSubscriptionsInfo": Route{
			strings.ToUpper("Delete"),
			"/nudr-dr/v2/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/smf-subscriptions",
			c.RemoveSmfSubscriptionsInfo,
		},
	}
}

// CreateSMFSubscriptions - Create SMF Subscription Info
func (c *SMFEventSubscriptionInfoDocumentAPIController) CreateSMFSubscriptions(w http.ResponseWriter, r *http.Request) {
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
	smfSubscriptionInfoParam := SmfSubscriptionInfo{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&smfSubscriptionInfoParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSmfSubscriptionInfoRequired(smfSubscriptionInfoParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertSmfSubscriptionInfoConstraints(smfSubscriptionInfoParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateSMFSubscriptions(r.Context(), ueIdParam, subsIdParam, smfSubscriptionInfoParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetSmfGroupSubscriptions - Retrieve SMF Subscription Info for a group of UEs or any UE
func (c *SMFEventSubscriptionInfoDocumentAPIController) GetSmfGroupSubscriptions(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.GetSmfGroupSubscriptions(r.Context(), ueGroupIdParam, subsIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetSmfSubscriptionInfo - Retrieve SMF Subscription Info
func (c *SMFEventSubscriptionInfoDocumentAPIController) GetSmfSubscriptionInfo(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.GetSmfSubscriptionInfo(r.Context(), ueIdParam, subsIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ModifySmfGroupSubscriptions - Modify SMF Subscription Info for a group of UEs or any UE
func (c *SMFEventSubscriptionInfoDocumentAPIController) ModifySmfGroupSubscriptions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
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
	result, err := c.service.ModifySmfGroupSubscriptions(r.Context(), ueGroupIdParam, subsIdParam, patchItemParam, supportedFeaturesParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ModifySmfSubscriptionInfo - Modify SMF Subscription Info
func (c *SMFEventSubscriptionInfoDocumentAPIController) ModifySmfSubscriptionInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
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
	result, err := c.service.ModifySmfSubscriptionInfo(r.Context(), ueIdParam, subsIdParam, patchItemParam, supportedFeaturesParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// RemoveSmfGroupSubscriptions - Delete SMF Subscription Info for a group of UEs or any UE
func (c *SMFEventSubscriptionInfoDocumentAPIController) RemoveSmfGroupSubscriptions(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.RemoveSmfGroupSubscriptions(r.Context(), ueGroupIdParam, subsIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// RemoveSmfSubscriptionsInfo - Delete SMF Subscription Info
func (c *SMFEventSubscriptionInfoDocumentAPIController) RemoveSmfSubscriptionsInfo(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.RemoveSmfSubscriptionsInfo(r.Context(), ueIdParam, subsIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
