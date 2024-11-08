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


// SMFRegistrationDocumentAPIService SMFRegistrationDocumentAPI service
type SMFRegistrationDocumentAPIService service

type ApiCreateOrUpdateSmfRegistrationRequest struct {
	ctx context.Context
	ApiService *SMFRegistrationDocumentAPIService
	ueId string
	pduSessionId int32
	smfRegistration *SmfRegistration
}

func (r ApiCreateOrUpdateSmfRegistrationRequest) SmfRegistration(smfRegistration SmfRegistration) ApiCreateOrUpdateSmfRegistrationRequest {
	r.smfRegistration = &smfRegistration
	return r
}

func (r ApiCreateOrUpdateSmfRegistrationRequest) Execute() (*SmfRegistration, *http.Response, error) {
	return r.ApiService.CreateOrUpdateSmfRegistrationExecute(r)
}

/*
CreateOrUpdateSmfRegistration To create an individual SMF context data of a UE in the UDR

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @param pduSessionId PDU session id
 @return ApiCreateOrUpdateSmfRegistrationRequest
*/
func (a *SMFRegistrationDocumentAPIService) CreateOrUpdateSmfRegistration(ctx context.Context, ueId string, pduSessionId int32) ApiCreateOrUpdateSmfRegistrationRequest {
	return ApiCreateOrUpdateSmfRegistrationRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		pduSessionId: pduSessionId,
	}
}

// Execute executes the request
//  @return SmfRegistration
func (a *SMFRegistrationDocumentAPIService) CreateOrUpdateSmfRegistrationExecute(r ApiCreateOrUpdateSmfRegistrationRequest) (*SmfRegistration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmfRegistration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMFRegistrationDocumentAPIService.CreateOrUpdateSmfRegistration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/smf-registrations/{pduSessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pduSessionId"+"}", url.PathEscape(parameterValueToString(r.pduSessionId, "pduSessionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pduSessionId < 0 {
		return localVarReturnValue, nil, reportError("pduSessionId must be greater than 0")
	}
	if r.pduSessionId > 255 {
		return localVarReturnValue, nil, reportError("pduSessionId must be less than 255")
	}
	if r.smfRegistration == nil {
		return localVarReturnValue, nil, reportError("smfRegistration is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.smfRegistration
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

type ApiDeleteSmfRegistrationRequest struct {
	ctx context.Context
	ApiService *SMFRegistrationDocumentAPIService
	ueId string
	pduSessionId int32
}

func (r ApiDeleteSmfRegistrationRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSmfRegistrationExecute(r)
}

/*
DeleteSmfRegistration To remove an individual SMF context data of a UE the UDR

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @param pduSessionId PDU session id
 @return ApiDeleteSmfRegistrationRequest
*/
func (a *SMFRegistrationDocumentAPIService) DeleteSmfRegistration(ctx context.Context, ueId string, pduSessionId int32) ApiDeleteSmfRegistrationRequest {
	return ApiDeleteSmfRegistrationRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		pduSessionId: pduSessionId,
	}
}

// Execute executes the request
func (a *SMFRegistrationDocumentAPIService) DeleteSmfRegistrationExecute(r ApiDeleteSmfRegistrationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMFRegistrationDocumentAPIService.DeleteSmfRegistration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/smf-registrations/{pduSessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pduSessionId"+"}", url.PathEscape(parameterValueToString(r.pduSessionId, "pduSessionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pduSessionId < 0 {
		return nil, reportError("pduSessionId must be greater than 0")
	}
	if r.pduSessionId > 255 {
		return nil, reportError("pduSessionId must be less than 255")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiQuerySmfRegistrationRequest struct {
	ctx context.Context
	ApiService *SMFRegistrationDocumentAPIService
	ueId string
	pduSessionId int32
	fields *[]string
	supportedFeatures *string
}

// attributes to be retrieved
func (r ApiQuerySmfRegistrationRequest) Fields(fields []string) ApiQuerySmfRegistrationRequest {
	r.fields = &fields
	return r
}

// Supported Features
func (r ApiQuerySmfRegistrationRequest) SupportedFeatures(supportedFeatures string) ApiQuerySmfRegistrationRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiQuerySmfRegistrationRequest) Execute() (*SmfRegistration, *http.Response, error) {
	return r.ApiService.QuerySmfRegistrationExecute(r)
}

/*
QuerySmfRegistration Retrieves the individual SMF registration of a UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @param pduSessionId PDU session id
 @return ApiQuerySmfRegistrationRequest
*/
func (a *SMFRegistrationDocumentAPIService) QuerySmfRegistration(ctx context.Context, ueId string, pduSessionId int32) ApiQuerySmfRegistrationRequest {
	return ApiQuerySmfRegistrationRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		pduSessionId: pduSessionId,
	}
}

// Execute executes the request
//  @return SmfRegistration
func (a *SMFRegistrationDocumentAPIService) QuerySmfRegistrationExecute(r ApiQuerySmfRegistrationRequest) (*SmfRegistration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmfRegistration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMFRegistrationDocumentAPIService.QuerySmfRegistration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/smf-registrations/{pduSessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pduSessionId"+"}", url.PathEscape(parameterValueToString(r.pduSessionId, "pduSessionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pduSessionId < 0 {
		return localVarReturnValue, nil, reportError("pduSessionId must be greater than 0")
	}
	if r.pduSessionId > 255 {
		return localVarReturnValue, nil, reportError("pduSessionId must be less than 255")
	}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "csv")
	}
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

type ApiUpdateSmfContextRequest struct {
	ctx context.Context
	ApiService *SMFRegistrationDocumentAPIService
	ueId string
	pduSessionId int32
	patchItem *[]PatchItem
	supportedFeatures *string
}

func (r ApiUpdateSmfContextRequest) PatchItem(patchItem []PatchItem) ApiUpdateSmfContextRequest {
	r.patchItem = &patchItem
	return r
}

// Features required to be supported by the target NF
func (r ApiUpdateSmfContextRequest) SupportedFeatures(supportedFeatures string) ApiUpdateSmfContextRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiUpdateSmfContextRequest) Execute() (*PatchResult, *http.Response, error) {
	return r.ApiService.UpdateSmfContextExecute(r)
}

/*
UpdateSmfContext To modify the SMF context data of a UE in the UDR

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @param pduSessionId PDU session id
 @return ApiUpdateSmfContextRequest
*/
func (a *SMFRegistrationDocumentAPIService) UpdateSmfContext(ctx context.Context, ueId string, pduSessionId int32) ApiUpdateSmfContextRequest {
	return ApiUpdateSmfContextRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		pduSessionId: pduSessionId,
	}
}

// Execute executes the request
//  @return PatchResult
func (a *SMFRegistrationDocumentAPIService) UpdateSmfContextExecute(r ApiUpdateSmfContextRequest) (*PatchResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PatchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMFRegistrationDocumentAPIService.UpdateSmfContext")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/smf-registrations/{pduSessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pduSessionId"+"}", url.PathEscape(parameterValueToString(r.pduSessionId, "pduSessionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pduSessionId < 0 {
		return localVarReturnValue, nil, reportError("pduSessionId must be greater than 0")
	}
	if r.pduSessionId > 255 {
		return localVarReturnValue, nil, reportError("pduSessionId must be less than 255")
	}
	if r.patchItem == nil {
		return localVarReturnValue, nil, reportError("patchItem is required and must be specified")
	}

	if r.supportedFeatures != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "supported-features", r.supportedFeatures, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.patchItem
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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
