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

// UpdateTheRoamingInformationOfTheEPCDomainDocumentAPIController binds http requests to an api service and writes the service results to the http response
type UpdateTheRoamingInformationOfTheEPCDomainDocumentAPIController struct {
	service      UpdateTheRoamingInformationOfTheEPCDomainDocumentAPIServicer
	errorHandler ErrorHandler
}

// UpdateTheRoamingInformationOfTheEPCDomainDocumentAPIOption for how the controller is set up.
type UpdateTheRoamingInformationOfTheEPCDomainDocumentAPIOption func(*UpdateTheRoamingInformationOfTheEPCDomainDocumentAPIController)

// WithUpdateTheRoamingInformationOfTheEPCDomainDocumentAPIErrorHandler inject ErrorHandler into controller
func WithUpdateTheRoamingInformationOfTheEPCDomainDocumentAPIErrorHandler(h ErrorHandler) UpdateTheRoamingInformationOfTheEPCDomainDocumentAPIOption {
	return func(c *UpdateTheRoamingInformationOfTheEPCDomainDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewUpdateTheRoamingInformationOfTheEPCDomainDocumentAPIController creates a default api controller
func NewUpdateTheRoamingInformationOfTheEPCDomainDocumentAPIController(s UpdateTheRoamingInformationOfTheEPCDomainDocumentAPIServicer, opts ...UpdateTheRoamingInformationOfTheEPCDomainDocumentAPIOption) Router {
	controller := &UpdateTheRoamingInformationOfTheEPCDomainDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the UpdateTheRoamingInformationOfTheEPCDomainDocumentAPIController
func (c *UpdateTheRoamingInformationOfTheEPCDomainDocumentAPIController) Routes() Routes {
	return Routes{
		"UpdateRoamingInformation": Route{
			strings.ToUpper("Put"),
			"/nudr-dr/v2/subscription-data/{ueId}/context-data/roaming-information",
			c.UpdateRoamingInformation,
		},
	}
}

// UpdateRoamingInformation - Update the Roaming Information of the EPC domain
func (c *UpdateTheRoamingInformationOfTheEPCDomainDocumentAPIController) UpdateRoamingInformation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ueIdParam := params["ueId"]
	if ueIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"ueId"}, nil)
		return
	}
	roamingInfoUpdateParam := RoamingInfoUpdate{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&roamingInfoUpdateParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertRoamingInfoUpdateRequired(roamingInfoUpdateParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertRoamingInfoUpdateConstraints(roamingInfoUpdateParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateRoamingInformation(r.Context(), ueIdParam, roamingInfoUpdateParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}