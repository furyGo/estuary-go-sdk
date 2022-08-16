# \CollectionsApi

All URIs are relative to *http://api.estuary.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CollectionsAddContentPost**](CollectionsApi.md#CollectionsAddContentPost) | **Post** /collections/add-content | Add contents to a collection
[**CollectionsColuuidCommitPost**](CollectionsApi.md#CollectionsColuuidCommitPost) | **Post** /collections/{coluuid}/commit | Produce a CID of the collection contents
[**CollectionsColuuidDelete**](CollectionsApi.md#CollectionsColuuidDelete) | **Delete** /collections/{coluuid} | Deletes a collection
[**CollectionsContentGet**](CollectionsApi.md#CollectionsContentGet) | **Get** /collections/content | Get contents in a collection
[**CollectionsCreatePost**](CollectionsApi.md#CollectionsCreatePost) | **Post** /collections/create | Create a new collection
[**CollectionsFsAddPost**](CollectionsApi.md#CollectionsFsAddPost) | **Post** /collections/fs/add | Add a file to a collection
[**CollectionsListGet**](CollectionsApi.md#CollectionsListGet) | **Get** /collections/list | List all collections



## CollectionsAddContentPost

> map[string]string CollectionsAddContentPost(ctx).Body(body).Execute()

Add contents to a collection



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
    body := *openapiclient.NewMainAddContentToCollectionParams() // MainAddContentToCollectionParams | Contents to add to collection

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.CollectionsAddContentPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.CollectionsAddContentPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CollectionsAddContentPost`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.CollectionsAddContentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsAddContentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MainAddContentToCollectionParams**](MainAddContentToCollectionParams.md) | Contents to add to collection | 

### Return type

**map[string]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectionsColuuidCommitPost

> string CollectionsColuuidCommitPost(ctx, coluuid).Execute()

Produce a CID of the collection contents



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
    coluuid := "coluuid_example" // string | coluuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.CollectionsColuuidCommitPost(context.Background(), coluuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.CollectionsColuuidCommitPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CollectionsColuuidCommitPost`: string
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.CollectionsColuuidCommitPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**coluuid** | **string** | coluuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsColuuidCommitPostRequest struct via the builder pattern


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


## CollectionsColuuidDelete

> CollectionsColuuidDelete(ctx, coluuid).Execute()

Deletes a collection



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
    coluuid := "coluuid_example" // string | Collection ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.CollectionsColuuidDelete(context.Background(), coluuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.CollectionsColuuidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**coluuid** | **string** | Collection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsColuuidDeleteRequest struct via the builder pattern


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


## CollectionsContentGet

> string CollectionsContentGet(ctx).Coluuid(coluuid).Dir(dir).Execute()

Get contents in a collection



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
    dir := "dir_example" // string | Directory (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.CollectionsContentGet(context.Background()).Coluuid(coluuid).Dir(dir).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.CollectionsContentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CollectionsContentGet`: string
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.CollectionsContentGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsContentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coluuid** | **string** | Collection UUID | 
 **dir** | **string** | Directory | 

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


## CollectionsCreatePost

> MainCollection CollectionsCreatePost(ctx).Body(body).Execute()

Create a new collection



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
    body := *openapiclient.NewMainCreateCollectionBody() // MainCreateCollectionBody | Collection name and description

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.CollectionsCreatePost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.CollectionsCreatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CollectionsCreatePost`: MainCollection
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.CollectionsCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MainCreateCollectionBody**](MainCreateCollectionBody.md) | Collection name and description | 

### Return type

[**MainCollection**](MainCollection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectionsFsAddPost

> CollectionsFsAddPost(ctx).Coluuid(coluuid).Content(content).Path(path).Execute()

Add a file to a collection



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
    coluuid := "coluuid_example" // string | Collection ID
    content := "content_example" // string | Content
    path := "path_example" // string | Path to file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.CollectionsFsAddPost(context.Background()).Coluuid(coluuid).Content(content).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.CollectionsFsAddPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsFsAddPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coluuid** | **string** | Collection ID | 
 **content** | **string** | Content | 
 **path** | **string** | Path to file | 

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


## CollectionsListGet

> []MainCollection CollectionsListGet(ctx, id).Execute()

List all collections



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
    id := int32(56) // int32 | User ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.CollectionsListGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.CollectionsListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CollectionsListGet`: []MainCollection
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.CollectionsListGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MainCollection**](MainCollection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

