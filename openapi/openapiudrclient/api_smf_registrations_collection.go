/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_udr_cli

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// SMFRegistrationsCollectionAPIService SMFRegistrationsCollectionAPI service
type SMFRegistrationsCollectionAPIService service

type ApiQuerySmfRegListRequest struct {
	ctx context.Context
	ApiService *SMFRegistrationsCollectionAPIService
	ueId string
	supportedFeatures *string
}

// Supported Features
func (r ApiQuerySmfRegListRequest) SupportedFeatures(supportedFeatures string) ApiQuerySmfRegListRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiQuerySmfRegListRequest) Execute() ([]SmfRegistration, *http.Response, error) {
	return r.ApiService.QuerySmfRegListExecute(r)
}

/*
QuerySmfRegList Retrieves the SMF registration list of a UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @return ApiQuerySmfRegListRequest
*/
func (a *SMFRegistrationsCollectionAPIService) QuerySmfRegList(ctx context.Context, ueId string) ApiQuerySmfRegListRequest {
	return ApiQuerySmfRegListRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
//  @return []SmfRegistration
func (a *SMFRegistrationsCollectionAPIService) QuerySmfRegListExecute(r ApiQuerySmfRegListRequest) ([]SmfRegistration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SmfRegistration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMFRegistrationsCollectionAPIService.QuerySmfRegList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/smf-registrations"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.supportedFeatures != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "supported-features", r.supportedFeatures, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
