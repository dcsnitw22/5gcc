# \DefaultAPI

All URIs are relative to *https://example.com/npcf-smpolicycontrol/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SmPoliciesPost**](DefaultAPI.md#SmPoliciesPost) | **Post** /sm-policies | 
[**SmPoliciesSmPolicyIdDeletePost**](DefaultAPI.md#SmPoliciesSmPolicyIdDeletePost) | **Post** /sm-policies/{smPolicyId}/delete | 
[**SmPoliciesSmPolicyIdGet**](DefaultAPI.md#SmPoliciesSmPolicyIdGet) | **Get** /sm-policies/{smPolicyId} | 
[**SmPoliciesSmPolicyIdUpdatePost**](DefaultAPI.md#SmPoliciesSmPolicyIdUpdatePost) | **Post** /sm-policies/{smPolicyId}/update | 



## SmPoliciesPost

> SmPolicyDecision SmPoliciesPost(ctx).SmPolicyContextData(smPolicyContextData).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    smPolicyContextData := *openapiclient.NewSmPolicyContextData("Supi_example", int32(123), *openapiclient.NewPduSessionType(), "Dnn_example", "NotificationUri_example", *openapiclient.NewSnssai(int32(123))) // SmPolicyContextData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.SmPoliciesPost(context.Background()).SmPolicyContextData(smPolicyContextData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SmPoliciesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SmPoliciesPost`: SmPolicyDecision
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SmPoliciesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSmPoliciesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smPolicyContextData** | [**SmPolicyContextData**](SmPolicyContextData.md) |  | 

### Return type

[**SmPolicyDecision**](SmPolicyDecision.md)

### Authorization

[oAuth2Clientcredentials](../README.md#oAuth2Clientcredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmPoliciesSmPolicyIdDeletePost

> SmPoliciesSmPolicyIdDeletePost(ctx, smPolicyId).SmPolicyDeleteData(smPolicyDeleteData).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    smPolicyId := "smPolicyId_example" // string | Identifier of a policy association
    smPolicyDeleteData := *openapiclient.NewSmPolicyDeleteData() // SmPolicyDeleteData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.SmPoliciesSmPolicyIdDeletePost(context.Background(), smPolicyId).SmPolicyDeleteData(smPolicyDeleteData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SmPoliciesSmPolicyIdDeletePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smPolicyId** | **string** | Identifier of a policy association | 

### Other Parameters

Other parameters are passed through a pointer to a apiSmPoliciesSmPolicyIdDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smPolicyDeleteData** | [**SmPolicyDeleteData**](SmPolicyDeleteData.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2Clientcredentials](../README.md#oAuth2Clientcredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmPoliciesSmPolicyIdGet

> SmPolicyControl SmPoliciesSmPolicyIdGet(ctx, smPolicyId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    smPolicyId := "smPolicyId_example" // string | Identifier of a policy association

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.SmPoliciesSmPolicyIdGet(context.Background(), smPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SmPoliciesSmPolicyIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SmPoliciesSmPolicyIdGet`: SmPolicyControl
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SmPoliciesSmPolicyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smPolicyId** | **string** | Identifier of a policy association | 

### Other Parameters

Other parameters are passed through a pointer to a apiSmPoliciesSmPolicyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmPolicyControl**](SmPolicyControl.md)

### Authorization

[oAuth2Clientcredentials](../README.md#oAuth2Clientcredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmPoliciesSmPolicyIdUpdatePost

> SmPolicyDecision SmPoliciesSmPolicyIdUpdatePost(ctx, smPolicyId).SmPolicyUpdateContextData(smPolicyUpdateContextData).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    smPolicyId := "smPolicyId_example" // string | Identifier of a policy association
    smPolicyUpdateContextData := *openapiclient.NewSmPolicyUpdateContextData() // SmPolicyUpdateContextData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.SmPoliciesSmPolicyIdUpdatePost(context.Background(), smPolicyId).SmPolicyUpdateContextData(smPolicyUpdateContextData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SmPoliciesSmPolicyIdUpdatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SmPoliciesSmPolicyIdUpdatePost`: SmPolicyDecision
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SmPoliciesSmPolicyIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smPolicyId** | **string** | Identifier of a policy association | 

### Other Parameters

Other parameters are passed through a pointer to a apiSmPoliciesSmPolicyIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smPolicyUpdateContextData** | [**SmPolicyUpdateContextData**](SmPolicyUpdateContextData.md) |  | 

### Return type

[**SmPolicyDecision**](SmPolicyDecision.md)

### Authorization

[oAuth2Clientcredentials](../README.md#oAuth2Clientcredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

