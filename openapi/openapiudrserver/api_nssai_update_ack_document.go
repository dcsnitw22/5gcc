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

// NSSAIUpdateAckDocumentAPIController binds http requests to an api service and writes the service results to the http response
type NSSAIUpdateAckDocumentAPIController struct {
	service      NSSAIUpdateAckDocumentAPIServicer
	errorHandler ErrorHandler
}

// NSSAIUpdateAckDocumentAPIOption for how the controller is set up.
type NSSAIUpdateAckDocumentAPIOption func(*NSSAIUpdateAckDocumentAPIController)

// WithNSSAIUpdateAckDocumentAPIErrorHandler inject ErrorHandler into controller
func WithNSSAIUpdateAckDocumentAPIErrorHandler(h ErrorHandler) NSSAIUpdateAckDocumentAPIOption {
	return func(c *NSSAIUpdateAckDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewNSSAIUpdateAckDocumentAPIController creates a default api controller
func NewNSSAIUpdateAckDocumentAPIController(s NSSAIUpdateAckDocumentAPIServicer, opts ...NSSAIUpdateAckDocumentAPIOption) Router {
	controller := &NSSAIUpdateAckDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the NSSAIUpdateAckDocumentAPIController
func (c *NSSAIUpdateAckDocumentAPIController) Routes() Routes {
	return Routes{
		"CreateOrUpdateNssaiAck": Route{
			strings.ToUpper("Put"),
			"/nudr-dr/v2/subscription-data/{ueId}/ue-update-confirmation-data/subscribed-snssais",
			c.CreateOrUpdateNssaiAck,
		},
	}
}

// CreateOrUpdateNssaiAck - To store the NSSAI update acknowledgement information of a UE
func (c *NSSAIUpdateAckDocumentAPIController) CreateOrUpdateNssaiAck(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	ueIdParam := params["ueId"]
	if ueIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"ueId"}, nil)
		return
	}
	nssaiAckDataParam := NssaiAckData{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&nssaiAckDataParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertNssaiAckDataRequired(nssaiAckDataParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertNssaiAckDataConstraints(nssaiAckDataParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	var supportedFeaturesParam string
	if query.Has("supported-features") {
		supportedFeaturesParam = query.Get("supported-features")
	}
	result, err := c.service.CreateOrUpdateNssaiAck(r.Context(), ueIdParam, nssaiAckDataParam, supportedFeaturesParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}