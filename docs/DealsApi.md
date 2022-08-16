# \DealsApi

All URIs are relative to *http://api.estuary.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DealEstimatePost**](DealsApi.md#DealEstimatePost) | **Post** /deal/estimate | Estimate the cost of a deal
[**DealInfoDealidGet**](DealsApi.md#DealInfoDealidGet) | **Get** /deal/info/{dealid} | Get Deal Info
[**DealProposalPropcidGet**](DealsApi.md#DealProposalPropcidGet) | **Get** /deal/proposal/{propcid} | Get Proposal
[**DealQueryMinerGet**](DealsApi.md#DealQueryMinerGet) | **Get** /deal/query/{miner} | Query Ask
[**DealStatusByProposalPropcidGet**](DealsApi.md#DealStatusByProposalPropcidGet) | **Get** /deal/status-by-proposal/{propcid} | Get Deal Status by PropCid
[**DealStatusMinerPropcidGet**](DealsApi.md#DealStatusMinerPropcidGet) | **Get** /deal/status/{miner}/{propcid} | Deal Status
[**DealTransferInProgressGet**](DealsApi.md#DealTransferInProgressGet) | **Get** /deal/transfer/in-progress | Transfer In Progress
[**DealTransferStatusPost**](DealsApi.md#DealTransferStatusPost) | **Post** /deal/transfer/status | Transfer Status
[**DealsFailuresGet**](DealsApi.md#DealsFailuresGet) | **Get** /deals/failures | Get storage failures for user
[**DealsMakeMinerPost**](DealsApi.md#DealsMakeMinerPost) | **Post** /deals/make/{miner} | Make Deal
[**DealsStatusDealGet**](DealsApi.md#DealsStatusDealGet) | **Get** /deals/status/{deal} | Get Deal Status
[**PublicDealsFailuresGet**](DealsApi.md#PublicDealsFailuresGet) | **Get** /public/deals/failures | Get storage failures



## DealEstimatePost

> DealEstimatePost(ctx).Body(body).Execute()

Estimate the cost of a deal



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
    body := *openapiclient.NewMainEstimateDealBody() // MainEstimateDealBody | The size of the deal in bytes, the replication factor, and the duration of the deal in blocks

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DealsApi.DealEstimatePost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DealsApi.DealEstimatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDealEstimatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MainEstimateDealBody**](MainEstimateDealBody.md) | The size of the deal in bytes, the replication factor, and the duration of the deal in blocks | 

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


## DealInfoDealidGet

> DealInfoDealidGet(ctx, dealid).Execute()

Get Deal Info



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
    dealid := int32(56) // int32 | Deal ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DealsApi.DealInfoDealidGet(context.Background(), dealid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DealsApi.DealInfoDealidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealid** | **int32** | Deal ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDealInfoDealidGetRequest struct via the builder pattern


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


## DealProposalPropcidGet

> DealProposalPropcidGet(ctx, propcid).Execute()

Get Proposal



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
    propcid := "propcid_example" // string | Proposal CID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DealsApi.DealProposalPropcidGet(context.Background(), propcid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DealsApi.DealProposalPropcidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propcid** | **string** | Proposal CID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDealProposalPropcidGetRequest struct via the builder pattern


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


## DealQueryMinerGet

> DealQueryMinerGet(ctx, miner).Execute()

Query Ask



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
    miner := "miner_example" // string | CID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DealsApi.DealQueryMinerGet(context.Background(), miner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DealsApi.DealQueryMinerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**miner** | **string** | CID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDealQueryMinerGetRequest struct via the builder pattern


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


## DealStatusByProposalPropcidGet

> DealStatusByProposalPropcidGet(ctx, propcid).Execute()

Get Deal Status by PropCid



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
    propcid := "propcid_example" // string | PropCid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DealsApi.DealStatusByProposalPropcidGet(context.Background(), propcid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DealsApi.DealStatusByProposalPropcidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propcid** | **string** | PropCid | 

### Other Parameters

Other parameters are passed through a pointer to a apiDealStatusByProposalPropcidGetRequest struct via the builder pattern


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


## DealStatusMinerPropcidGet

> DealStatusMinerPropcidGet(ctx, miner, propcid).Execute()

Deal Status



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
    miner := "miner_example" // string | Miner
    propcid := "propcid_example" // string | Proposal CID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DealsApi.DealStatusMinerPropcidGet(context.Background(), miner, propcid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DealsApi.DealStatusMinerPropcidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**miner** | **string** | Miner | 
**propcid** | **string** | Proposal CID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDealStatusMinerPropcidGetRequest struct via the builder pattern


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


## DealTransferInProgressGet

> DealTransferInProgressGet(ctx).Execute()

Transfer In Progress



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
    resp, r, err := apiClient.DealsApi.DealTransferInProgressGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DealsApi.DealTransferInProgressGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDealTransferInProgressGetRequest struct via the builder pattern


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


## DealTransferStatusPost

> DealTransferStatusPost(ctx).Execute()

Transfer Status



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
    resp, r, err := apiClient.DealsApi.DealTransferStatusPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DealsApi.DealTransferStatusPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDealTransferStatusPostRequest struct via the builder pattern


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


## DealsFailuresGet

> DealsFailuresGet(ctx).Execute()

Get storage failures for user



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
    resp, r, err := apiClient.DealsApi.DealsFailuresGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DealsApi.DealsFailuresGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDealsFailuresGetRequest struct via the builder pattern


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


## DealsMakeMinerPost

> DealsMakeMinerPost(ctx, miner).DealRequest(dealRequest).Execute()

Make Deal



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
    miner := "miner_example" // string | Miner
    dealRequest := "dealRequest_example" // string | Deal Request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DealsApi.DealsMakeMinerPost(context.Background(), miner).DealRequest(dealRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DealsApi.DealsMakeMinerPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**miner** | **string** | Miner | 

### Other Parameters

Other parameters are passed through a pointer to a apiDealsMakeMinerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dealRequest** | **string** | Deal Request | 

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


## DealsStatusDealGet

> DealsStatusDealGet(ctx, deal).Execute()

Get Deal Status



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
    deal := int32(56) // int32 | Deal ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DealsApi.DealsStatusDealGet(context.Background(), deal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DealsApi.DealsStatusDealGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deal** | **int32** | Deal ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDealsStatusDealGetRequest struct via the builder pattern


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


## PublicDealsFailuresGet

> PublicDealsFailuresGet(ctx).Execute()

Get storage failures



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
    resp, r, err := apiClient.DealsApi.PublicDealsFailuresGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DealsApi.PublicDealsFailuresGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPublicDealsFailuresGetRequest struct via the builder pattern


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

