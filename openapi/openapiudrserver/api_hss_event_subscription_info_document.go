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

// HSSEventSubscriptionInfoDocumentAPIController binds http requests to an api service and writes the service results to the http response
type HSSEventSubscriptionInfoDocumentAPIController struct {
	service      HSSEventSubscriptionInfoDocumentAPIServicer
	errorHandler ErrorHandler
}

// HSSEventSubscriptionInfoDocumentAPIOption for how the controller is set up.
type HSSEventSubscriptionInfoDocumentAPIOption func(*HSSEventSubscriptionInfoDocumentAPIController)

// WithHSSEventSubscriptionInfoDocumentAPIErrorHandler inject ErrorHandler into controller
func WithHSSEventSubscriptionInfoDocumentAPIErrorHandler(h ErrorHandler) HSSEventSubscriptionInfoDocumentAPIOption {
	return func(c *HSSEventSubscriptionInfoDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewHSSEventSubscriptionInfoDocumentAPIController creates a default api controller
func NewHSSEventSubscriptionInfoDocumentAPIController(s HSSEventSubscriptionInfoDocumentAPIServicer, opts ...HSSEventSubscriptionInfoDocumentAPIOption) Router {
	controller := &HSSEventSubscriptionInfoDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the HSSEventSubscriptionInfoDocumentAPIController
func (c *HSSEventSubscriptionInfoDocumentAPIController) Routes() Routes {
	return Routes{
		"CreateHSSSubscriptions": Route{
			strings.ToUpper("Put"),
			"/nudr-dr/v2/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions",
			c.CreateHSSSubscriptions,
		},
		"GetHssGroupSubscriptions": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/hss-subscriptions",
			c.GetHssGroupSubscriptions,
		},
		"GetHssSubscriptionInfo": Route{
			strings.ToUpper("Get"),
			"/nudr-dr/v2/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions",
			c.GetHssSubscriptionInfo,
		},
		"ModifyHssGroupSubscriptions": Route{
			strings.ToUpper("Patch"),
			"/nudr-dr/v2/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/hss-subscriptions",
			c.ModifyHssGroupSubscriptions,
		},
		"ModifyHssSubscriptionInfo": Route{
			strings.ToUpper("Patch"),
			"/nudr-dr/v2/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions",
			c.ModifyHssSubscriptionInfo,
		},
		"RemoveHssGroupSubscriptions": Route{
			strings.ToUpper("Delete"),
			"/nudr-dr/v2/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/hss-subscriptions",
			c.RemoveHssGroupSubscriptions,
		},
		"RemoveHssSubscriptionsInfo": Route{
			strings.ToUpper("Delete"),
			"/nudr-dr/v2/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions",
			c.RemoveHssSubscriptionsInfo,
		},
	}
}

// CreateHSSSubscriptions - Create HSS Subscription Info
func (c *HSSEventSubscriptionInfoDocumentAPIController) CreateHSSSubscriptions(w http.ResponseWriter, r *http.Request) {
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
	hssSubscriptionInfoParam := HssSubscriptionInfo{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&hssSubscriptionInfoParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertHssSubscriptionInfoRequired(hssSubscriptionInfoParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertHssSubscriptionInfoConstraints(hssSubscriptionInfoParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateHSSSubscriptions(r.Context(), ueIdParam, subsIdParam, hssSubscriptionInfoParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetHssGroupSubscriptions - Retrieve HSS Subscription Info
func (c *HSSEventSubscriptionInfoDocumentAPIController) GetHssGroupSubscriptions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	externalGroupIdParam := params["externalGroupId"]
	if externalGroupIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"externalGroupId"}, nil)
		return
	}
	subsIdParam := params["subsId"]
	if subsIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subsId"}, nil)
		return
	}
	result, err := c.service.GetHssGroupSubscriptions(r.Context(), externalGroupIdParam, subsIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetHssSubscriptionInfo - Retrieve HSS Subscription Info
func (c *HSSEventSubscriptionInfoDocumentAPIController) GetHssSubscriptionInfo(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.GetHssSubscriptionInfo(r.Context(), ueIdParam, subsIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ModifyHssGroupSubscriptions - Modify HSS Subscription Info
func (c *HSSEventSubscriptionInfoDocumentAPIController) ModifyHssGroupSubscriptions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	externalGroupIdParam := params["externalGroupId"]
	if externalGroupIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"externalGroupId"}, nil)
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
	result, err := c.service.ModifyHssGroupSubscriptions(r.Context(), externalGroupIdParam, subsIdParam, patchItemParam, supportedFeaturesParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ModifyHssSubscriptionInfo - Modify HSS Subscription Info
func (c *HSSEventSubscriptionInfoDocumentAPIController) ModifyHssSubscriptionInfo(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.ModifyHssSubscriptionInfo(r.Context(), ueIdParam, subsIdParam, patchItemParam, supportedFeaturesParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// RemoveHssGroupSubscriptions - Delete HSS Subscription Info
func (c *HSSEventSubscriptionInfoDocumentAPIController) RemoveHssGroupSubscriptions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	externalGroupIdParam := params["externalGroupId"]
	if externalGroupIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"externalGroupId"}, nil)
		return
	}
	subsIdParam := params["subsId"]
	if subsIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subsId"}, nil)
		return
	}
	result, err := c.service.RemoveHssGroupSubscriptions(r.Context(), externalGroupIdParam, subsIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// RemoveHssSubscriptionsInfo - Delete HSS Subscription Info
func (c *HSSEventSubscriptionInfoDocumentAPIController) RemoveHssSubscriptionsInfo(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.RemoveHssSubscriptionsInfo(r.Context(), ueIdParam, subsIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}