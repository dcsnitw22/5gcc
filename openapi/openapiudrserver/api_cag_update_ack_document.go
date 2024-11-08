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

// CAGUpdateAckDocumentAPIController binds http requests to an api service and writes the service results to the http response
type CAGUpdateAckDocumentAPIController struct {
	service      CAGUpdateAckDocumentAPIServicer
	errorHandler ErrorHandler
}

// CAGUpdateAckDocumentAPIOption for how the controller is set up.
type CAGUpdateAckDocumentAPIOption func(*CAGUpdateAckDocumentAPIController)

// WithCAGUpdateAckDocumentAPIErrorHandler inject ErrorHandler into controller
func WithCAGUpdateAckDocumentAPIErrorHandler(h ErrorHandler) CAGUpdateAckDocumentAPIOption {
	return func(c *CAGUpdateAckDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewCAGUpdateAckDocumentAPIController creates a default api controller
func NewCAGUpdateAckDocumentAPIController(s CAGUpdateAckDocumentAPIServicer, opts ...CAGUpdateAckDocumentAPIOption) Router {
	controller := &CAGUpdateAckDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the CAGUpdateAckDocumentAPIController
func (c *CAGUpdateAckDocumentAPIController) Routes() Routes {
	return Routes{
		"CreateCagUpdateAck": Route{
			strings.ToUpper("Put"),
			"/nudr-dr/v2/subscription-data/{ueId}/ue-update-confirmation-data/subscribed-cag",
			c.CreateCagUpdateAck,
		},
	}
}

// CreateCagUpdateAck - To store the CAG update acknowledgement information of a UE
func (c *CAGUpdateAckDocumentAPIController) CreateCagUpdateAck(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	ueIdParam := params["ueId"]
	if ueIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"ueId"}, nil)
		return
	}
	cagAckDataParam := CagAckData{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&cagAckDataParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCagAckDataRequired(cagAckDataParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCagAckDataConstraints(cagAckDataParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	var supportedFeaturesParam string
	if query.Has("supported-features") {
		supportedFeaturesParam = query.Get("supported-features")
	}
	result, err := c.service.CreateCagUpdateAck(r.Context(), ueIdParam, cagAckDataParam, supportedFeaturesParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
