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


// ServiceSpecificAuthorizationInfoDocumentAPIService ServiceSpecificAuthorizationInfoDocumentAPI service
type ServiceSpecificAuthorizationInfoDocumentAPIService service

type ApiCreateServiceSpecificAuthorizationInfoRequest struct {
	ctx context.Context
	ApiService *ServiceSpecificAuthorizationInfoDocumentAPIService
	ueId string
	serviceType ServiceType
	serviceSpecificAuthorizationInfo *ServiceSpecificAuthorizationInfo
}

func (r ApiCreateServiceSpecificAuthorizationInfoRequest) ServiceSpecificAuthorizationInfo(serviceSpecificAuthorizationInfo ServiceSpecificAuthorizationInfo) ApiCreateServiceSpecificAuthorizationInfoRequest {
	r.serviceSpecificAuthorizationInfo = &serviceSpecificAuthorizationInfo
	return r
}

func (r ApiCreateServiceSpecificAuthorizationInfoRequest) Execute() (*ServiceSpecificAuthorizationInfo, *http.Response, error) {
	return r.ApiService.CreateServiceSpecificAuthorizationInfoExecute(r)
}

/*
CreateServiceSpecificAuthorizationInfo Create Service Specific Authorization Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId
 @param serviceType Service Type
 @return ApiCreateServiceSpecificAuthorizationInfoRequest
*/
func (a *ServiceSpecificAuthorizationInfoDocumentAPIService) CreateServiceSpecificAuthorizationInfo(ctx context.Context, ueId string, serviceType ServiceType) ApiCreateServiceSpecificAuthorizationInfoRequest {
	return ApiCreateServiceSpecificAuthorizationInfoRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		serviceType: serviceType,
	}
}

// Execute executes the request
//  @return ServiceSpecificAuthorizationInfo
func (a *ServiceSpecificAuthorizationInfoDocumentAPIService) CreateServiceSpecificAuthorizationInfoExecute(r ApiCreateServiceSpecificAuthorizationInfoRequest) (*ServiceSpecificAuthorizationInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceSpecificAuthorizationInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceSpecificAuthorizationInfoDocumentAPIService.CreateServiceSpecificAuthorizationInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/service-specific-authorizations/{serviceType}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceType"+"}", url.PathEscape(parameterValueToString(r.serviceType, "serviceType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serviceSpecificAuthorizationInfo == nil {
		return localVarReturnValue, nil, reportError("serviceSpecificAuthorizationInfo is required and must be specified")
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
	localVarPostBody = r.serviceSpecificAuthorizationInfo
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

type ApiGetServiceSpecificAuthorizationInfoRequest struct {
	ctx context.Context
	ApiService *ServiceSpecificAuthorizationInfoDocumentAPIService
	ueId string
	serviceType ServiceType
}

func (r ApiGetServiceSpecificAuthorizationInfoRequest) Execute() (*ServiceSpecificAuthorizationInfo, *http.Response, error) {
	return r.ApiService.GetServiceSpecificAuthorizationInfoExecute(r)
}

/*
GetServiceSpecificAuthorizationInfo Retrieve Service Specific Authorization Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId
 @param serviceType Service Type
 @return ApiGetServiceSpecificAuthorizationInfoRequest
*/
func (a *ServiceSpecificAuthorizationInfoDocumentAPIService) GetServiceSpecificAuthorizationInfo(ctx context.Context, ueId string, serviceType ServiceType) ApiGetServiceSpecificAuthorizationInfoRequest {
	return ApiGetServiceSpecificAuthorizationInfoRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		serviceType: serviceType,
	}
}

// Execute executes the request
//  @return ServiceSpecificAuthorizationInfo
func (a *ServiceSpecificAuthorizationInfoDocumentAPIService) GetServiceSpecificAuthorizationInfoExecute(r ApiGetServiceSpecificAuthorizationInfoRequest) (*ServiceSpecificAuthorizationInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceSpecificAuthorizationInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceSpecificAuthorizationInfoDocumentAPIService.GetServiceSpecificAuthorizationInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/service-specific-authorizations/{serviceType}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceType"+"}", url.PathEscape(parameterValueToString(r.serviceType, "serviceType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiModifyServiceSpecificAuthorizationInfoRequest struct {
	ctx context.Context
	ApiService *ServiceSpecificAuthorizationInfoDocumentAPIService
	ueId string
	serviceType ServiceType
	patchItem *[]PatchItem
	supportedFeatures *string
}

func (r ApiModifyServiceSpecificAuthorizationInfoRequest) PatchItem(patchItem []PatchItem) ApiModifyServiceSpecificAuthorizationInfoRequest {
	r.patchItem = &patchItem
	return r
}

// Features required to be supported by the target NF
func (r ApiModifyServiceSpecificAuthorizationInfoRequest) SupportedFeatures(supportedFeatures string) ApiModifyServiceSpecificAuthorizationInfoRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiModifyServiceSpecificAuthorizationInfoRequest) Execute() (*PatchResult, *http.Response, error) {
	return r.ApiService.ModifyServiceSpecificAuthorizationInfoExecute(r)
}

/*
ModifyServiceSpecificAuthorizationInfo Modify Service Specific Authorization Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId
 @param serviceType Service Type
 @return ApiModifyServiceSpecificAuthorizationInfoRequest
*/
func (a *ServiceSpecificAuthorizationInfoDocumentAPIService) ModifyServiceSpecificAuthorizationInfo(ctx context.Context, ueId string, serviceType ServiceType) ApiModifyServiceSpecificAuthorizationInfoRequest {
	return ApiModifyServiceSpecificAuthorizationInfoRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		serviceType: serviceType,
	}
}

// Execute executes the request
//  @return PatchResult
func (a *ServiceSpecificAuthorizationInfoDocumentAPIService) ModifyServiceSpecificAuthorizationInfoExecute(r ApiModifyServiceSpecificAuthorizationInfoRequest) (*PatchResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PatchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceSpecificAuthorizationInfoDocumentAPIService.ModifyServiceSpecificAuthorizationInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/service-specific-authorizations/{serviceType}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceType"+"}", url.PathEscape(parameterValueToString(r.serviceType, "serviceType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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

type ApiRemoveServiceSpecificAuthorizationInfoRequest struct {
	ctx context.Context
	ApiService *ServiceSpecificAuthorizationInfoDocumentAPIService
	ueId string
	serviceType ServiceType
}

func (r ApiRemoveServiceSpecificAuthorizationInfoRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveServiceSpecificAuthorizationInfoExecute(r)
}

/*
RemoveServiceSpecificAuthorizationInfo Delete Service Specific Authorization Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId
 @param serviceType Service Type
 @return ApiRemoveServiceSpecificAuthorizationInfoRequest
*/
func (a *ServiceSpecificAuthorizationInfoDocumentAPIService) RemoveServiceSpecificAuthorizationInfo(ctx context.Context, ueId string, serviceType ServiceType) ApiRemoveServiceSpecificAuthorizationInfoRequest {
	return ApiRemoveServiceSpecificAuthorizationInfoRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		serviceType: serviceType,
	}
}

// Execute executes the request
func (a *ServiceSpecificAuthorizationInfoDocumentAPIService) RemoveServiceSpecificAuthorizationInfoExecute(r ApiRemoveServiceSpecificAuthorizationInfoRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceSpecificAuthorizationInfoDocumentAPIService.RemoveServiceSpecificAuthorizationInfo")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/service-specific-authorizations/{serviceType}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceType"+"}", url.PathEscape(parameterValueToString(r.serviceType, "serviceType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
