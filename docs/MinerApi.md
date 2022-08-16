# \MinerApi

All URIs are relative to *http://api.estuary.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublicMinersDealsMinerGet**](MinerApi.md#PublicMinersDealsMinerGet) | **Get** /public/miners/deals/{miner} | Get all miners deals
[**PublicMinersStatsMinerGet**](MinerApi.md#PublicMinersStatsMinerGet) | **Get** /public/miners/stats/{miner} | Get miner stats



## PublicMinersDealsMinerGet

> PublicMinersDealsMinerGet(ctx, miner).Execute()

Get all miners deals



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
    miner := "miner_example" // string | Filter by miner

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MinerApi.PublicMinersDealsMinerGet(context.Background(), miner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MinerApi.PublicMinersDealsMinerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**miner** | **string** | Filter by miner | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicMinersDealsMinerGetRequest struct via the builder pattern


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


## PublicMinersStatsMinerGet

> PublicMinersStatsMinerGet(ctx, miner).Execute()

Get miner stats



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
    miner := "miner_example" // string | Filter by miner

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MinerApi.PublicMinersStatsMinerGet(context.Background(), miner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MinerApi.PublicMinersStatsMinerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**miner** | **string** | Filter by miner | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicMinersStatsMinerGetRequest struct via the builder pattern


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

