# \AutoretrieveApi

All URIs are relative to *http://api.estuary.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminAutoretrieveInitPost**](AutoretrieveApi.md#AdminAutoretrieveInitPost) | **Post** /admin/autoretrieve/init | Register autoretrieve server
[**AdminAutoretrieveListGet**](AutoretrieveApi.md#AdminAutoretrieveListGet) | **Get** /admin/autoretrieve/list | List autoretrieve servers
[**AutoretrieveHeartbeatPost**](AutoretrieveApi.md#AutoretrieveHeartbeatPost) | **Post** /autoretrieve/heartbeat | Marks autoretrieve server as up



## AdminAutoretrieveInitPost

> AdminAutoretrieveInitPost(ctx).PubKey(pubKey).Execute()

Register autoretrieve server



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
    pubKey := "pubKey_example" // string | Autoretrieve's public key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoretrieveApi.AdminAutoretrieveInitPost(context.Background()).PubKey(pubKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoretrieveApi.AdminAutoretrieveInitPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminAutoretrieveInitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pubKey** | **string** | Autoretrieve&#39;s public key | 

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


## AdminAutoretrieveListGet

> AdminAutoretrieveListGet(ctx).Execute()

List autoretrieve servers



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
    resp, r, err := apiClient.AutoretrieveApi.AdminAutoretrieveListGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoretrieveApi.AdminAutoretrieveListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAutoretrieveListGetRequest struct via the builder pattern


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


## AutoretrieveHeartbeatPost

> AutoretrieveHeartbeatPost(ctx).Token(token).Execute()

Marks autoretrieve server as up



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
    token := "token_example" // string | Autoretrieve's auth token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoretrieveApi.AutoretrieveHeartbeatPost(context.Background()).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoretrieveApi.AutoretrieveHeartbeatPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoretrieveHeartbeatPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Autoretrieve&#39;s auth token | 

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

