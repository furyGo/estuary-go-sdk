# \ContentApi

All URIs are relative to *http://api.estuary.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentAddCarPost**](ContentApi.md#ContentAddCarPost) | **Post** /content/add-car | Add Car object
[**ContentAddIpfsPost**](ContentApi.md#ContentAddIpfsPost) | **Post** /content/add-ipfs | Add IPFS object
[**ContentAddPost**](ContentApi.md#ContentAddPost) | **Post** /content/add | Add new content
[**ContentAggregatedContentGet**](ContentApi.md#ContentAggregatedContentGet) | **Get** /content/aggregated/{content} | Get aggregated content stats
[**ContentAllDealsGet**](ContentApi.md#ContentAllDealsGet) | **Get** /content/all-deals | Get all deals for a user
[**ContentBwUsageContentGet**](ContentApi.md#ContentBwUsageContentGet) | **Get** /content/bw-usage/{content} | Get content bandwidth
[**ContentCreatePost**](ContentApi.md#ContentCreatePost) | **Post** /content/create | Add a new content
[**ContentDealsGet**](ContentApi.md#ContentDealsGet) | **Get** /content/deals | Content with deals
[**ContentEnsureReplicationDatacidGet**](ContentApi.md#ContentEnsureReplicationDatacidGet) | **Get** /content/ensure-replication/{datacid} | Ensure Replication
[**ContentFailuresContentGet**](ContentApi.md#ContentFailuresContentGet) | **Get** /content/failures/{content} | List all failures for a content
[**ContentImportdealPost**](ContentApi.md#ContentImportdealPost) | **Post** /content/importdeal | Import a deal
[**ContentListGet**](ContentApi.md#ContentListGet) | **Get** /content/list | List all pinned content
[**ContentReadContGet**](ContentApi.md#ContentReadContGet) | **Get** /content/read/{cont} | Read content
[**ContentStagingZonesGet**](ContentApi.md#ContentStagingZonesGet) | **Get** /content/staging-zones | Get staging zone for user
[**ContentStatsGet**](ContentApi.md#ContentStatsGet) | **Get** /content/stats | Get content statistics
[**ContentStatusIdGet**](ContentApi.md#ContentStatusIdGet) | **Get** /content/status/{id} | Content Status



## ContentAddCarPost

> ContentAddCarPost(ctx).Body(body).Filename(filename).Commp(commp).Size(size).Execute()

Add Car object



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
    body := "body_example" // string | Car
    filename := "filename_example" // string | Filename (optional)
    commp := "commp_example" // string | Commp (optional)
    size := "size_example" // string | Size (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentAddCarPost(context.Background()).Body(body).Filename(filename).Commp(commp).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentAddCarPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentAddCarPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | Car | 
 **filename** | **string** | Filename | 
 **commp** | **string** | Commp | 
 **size** | **string** | Size | 

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


## ContentAddIpfsPost

> ContentAddIpfsPost(ctx).Body(body).Execute()

Add IPFS object



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
    body := *openapiclient.NewUtilContentAddIpfsBody() // UtilContentAddIpfsBody | IPFS Body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentAddIpfsPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentAddIpfsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentAddIpfsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UtilContentAddIpfsBody**](UtilContentAddIpfsBody.md) | IPFS Body | 

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


## ContentAddPost

> ContentAddPost(ctx, coluuid, dir).File(file).Execute()

Add new content



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
    coluuid := "coluuid_example" // string | Collection UUID
    dir := "dir_example" // string | Directory
    file := os.NewFile(1234, "some_file") // *os.File | File to upload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentAddPost(context.Background(), coluuid, dir).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentAddPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**coluuid** | **string** | Collection UUID | 
**dir** | **string** | Directory | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentAddPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | File to upload | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentAggregatedContentGet

> string ContentAggregatedContentGet(ctx, content).Execute()

Get aggregated content stats



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
    content := "content_example" // string | Content ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentAggregatedContentGet(context.Background(), content).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentAggregatedContentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentAggregatedContentGet`: string
    fmt.Fprintf(os.Stdout, "Response from `ContentApi.ContentAggregatedContentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**content** | **string** | Content ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentAggregatedContentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentAllDealsGet

> ContentAllDealsGet(ctx).Begin(begin).Duration(duration).All(all).Execute()

Get all deals for a user



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
    begin := "begin_example" // string | Begin
    duration := "duration_example" // string | Duration
    all := "all_example" // string | All

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentAllDealsGet(context.Background()).Begin(begin).Duration(duration).All(all).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentAllDealsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentAllDealsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **begin** | **string** | Begin | 
 **duration** | **string** | Duration | 
 **all** | **string** | All | 

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


## ContentBwUsageContentGet

> ContentBwUsageContentGet(ctx, content).Execute()

Get content bandwidth



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
    content := "content_example" // string | Content ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentBwUsageContentGet(context.Background(), content).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentBwUsageContentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**content** | **string** | Content ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentBwUsageContentGetRequest struct via the builder pattern


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


## ContentCreatePost

> ContentCreatePost(ctx).Body(body).Execute()

Add a new content



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
    body := "body_example" // string | Content

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentCreatePost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentCreatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | Content | 

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


## ContentDealsGet

> ContentDealsGet(ctx).Limit(limit).Offset(offset).Execute()

Content with deals



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
    limit := int32(56) // int32 | Limit (optional)
    offset := int32(56) // int32 | Offset (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentDealsGet(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentDealsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDealsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit | 
 **offset** | **int32** | Offset | 

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


## ContentEnsureReplicationDatacidGet

> ContentEnsureReplicationDatacidGet(ctx, datacid).Execute()

Ensure Replication



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
    datacid := "datacid_example" // string | Data CID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentEnsureReplicationDatacidGet(context.Background(), datacid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentEnsureReplicationDatacidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacid** | **string** | Data CID | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentEnsureReplicationDatacidGetRequest struct via the builder pattern


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


## ContentFailuresContentGet

> string ContentFailuresContentGet(ctx, content).Execute()

List all failures for a content



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
    content := "content_example" // string | Content ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentFailuresContentGet(context.Background(), content).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentFailuresContentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentFailuresContentGet`: string
    fmt.Fprintf(os.Stdout, "Response from `ContentApi.ContentFailuresContentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**content** | **string** | Content ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentFailuresContentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentImportdealPost

> ContentImportdealPost(ctx).Body(body).Execute()

Import a deal



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
    body := *openapiclient.NewMainImportDealBody() // MainImportDealBody | Import a deal

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentImportdealPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentImportdealPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentImportdealPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MainImportDealBody**](MainImportDealBody.md) | Import a deal | 

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


## ContentListGet

> []string ContentListGet(ctx).Execute()

List all pinned content



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
    resp, r, err := apiClient.ContentApi.ContentListGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentListGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `ContentApi.ContentListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiContentListGetRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentReadContGet

> ContentReadContGet(ctx, cont).Execute()

Read content



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
    cont := "cont_example" // string | CID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentReadContGet(context.Background(), cont).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentReadContGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cont** | **string** | CID | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentReadContGetRequest struct via the builder pattern


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


## ContentStagingZonesGet

> ContentStagingZonesGet(ctx).Execute()

Get staging zone for user



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
    resp, r, err := apiClient.ContentApi.ContentStagingZonesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentStagingZonesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiContentStagingZonesGetRequest struct via the builder pattern


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


## ContentStatsGet

> ContentStatsGet(ctx, limit).Execute()

Get content statistics



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
    limit := "limit_example" // string | limit

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentStatsGet(context.Background(), limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentStatsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**limit** | **string** | limit | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentStatsGetRequest struct via the builder pattern


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


## ContentStatusIdGet

> ContentStatusIdGet(ctx, id).Execute()

Content Status



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
    id := int32(56) // int32 | Content ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentApi.ContentStatusIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.ContentStatusIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Content ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentStatusIdGetRequest struct via the builder pattern


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

