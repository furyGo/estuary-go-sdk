# \PinningApi

All URIs are relative to *http://api.estuary.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PinningPinsGet**](PinningApi.md#PinningPinsGet) | **Get** /pinning/pins | List all pin status objects
[**PinningPinsPinidDelete**](PinningApi.md#PinningPinsPinidDelete) | **Delete** /pinning/pins/{pinid} | Delete a pinned object
[**PinningPinsPinidGet**](PinningApi.md#PinningPinsPinidGet) | **Get** /pinning/pins/{pinid} | Get a pin status object
[**PinningPinsPinidPost**](PinningApi.md#PinningPinsPinidPost) | **Post** /pinning/pins/{pinid} | Replace a pinned object
[**PinningPinsPost**](PinningApi.md#PinningPinsPost) | **Post** /pinning/pins | Add and pin object



## PinningPinsGet

> PinningPinsGet(ctx).Execute()

List all pin status objects



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PinningApi.PinningPinsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PinningApi.PinningPinsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPinningPinsGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PinningPinsPinidDelete

> PinningPinsPinidDelete(ctx, pinid).Execute()

Delete a pinned object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pinid := "pinid_example" // string | Pin ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PinningApi.PinningPinsPinidDelete(context.Background(), pinid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PinningApi.PinningPinsPinidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pinid** | **string** | Pin ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPinningPinsPinidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PinningPinsPinidGet

> PinningPinsPinidGet(ctx, pinid).Execute()

Get a pin status object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pinid := "pinid_example" // string | cid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PinningApi.PinningPinsPinidGet(context.Background(), pinid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PinningApi.PinningPinsPinidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pinid** | **string** | cid | 

### Other Parameters

Other parameters are passed through a pointer to a apiPinningPinsPinidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PinningPinsPinidPost

> PinningPinsPinidPost(ctx, pinid).Execute()

Replace a pinned object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pinid := "pinid_example" // string | Pin ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PinningApi.PinningPinsPinidPost(context.Background(), pinid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PinningApi.PinningPinsPinidPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pinid** | **string** | Pin ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPinningPinsPinidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PinningPinsPost

> PinningPinsPost(ctx, cid, name).Execute()

Add and pin object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cid := "cid_example" // string | cid
    name := "name_example" // string | name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PinningApi.PinningPinsPost(context.Background(), cid, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PinningApi.PinningPinsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** | cid | 
**name** | **string** | name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPinningPinsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

