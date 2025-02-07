# \NfsServerAPI

All URIs are relative to *https://secure.sakura.ad.jp/vps/api/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNfsServer**](NfsServerAPI.md#GetNfsServer) | **Get** /nfs-servers/{nfs_server_id} | NFS情報を取得する
[**GetNfsServerInterface**](NfsServerAPI.md#GetNfsServerInterface) | **Get** /nfs-servers/{nfs_server_id}/interface | NFSのインターフェースを取得する
[**GetNfsServerList**](NfsServerAPI.md#GetNfsServerList) | **Get** /nfs-servers | NFS情報一覧を取得する
[**GetNfsServerPowerStatus**](NfsServerAPI.md#GetNfsServerPowerStatus) | **Get** /nfs-servers/{nfs_server_id}/power-status | NFSの電源状態を取得する
[**GetNfsServerStorage**](NfsServerAPI.md#GetNfsServerStorage) | **Get** /nfs-servers/{nfs_server_id}/storage | NFSのストレージ容量情報を取得する
[**PostNfsServerChangeIpv4**](NfsServerAPI.md#PostNfsServerChangeIpv4) | **Post** /nfs-servers/{nfs_server_id}/change-ipv4 | NFSのipv4を更新する
[**PutNfsServer**](NfsServerAPI.md#PutNfsServer) | **Put** /nfs-servers/{nfs_server_id} | NFS情報を更新する
[**PutNfsServerInterface**](NfsServerAPI.md#PutNfsServerInterface) | **Put** /nfs-servers/{nfs_server_id}/interface | NFSのインターフェース情報を更新する



## GetNfsServer

> NfsServer GetNfsServer(ctx, nfsServerId).Execute()

NFS情報を取得する

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/sakura_vps_client_go"
)

func main() {
	nfsServerId := int32(56) // int32 | NFSのID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsServerAPI.GetNfsServer(context.Background(), nfsServerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsServerAPI.GetNfsServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNfsServer`: NfsServer
	fmt.Fprintf(os.Stdout, "Response from `NfsServerAPI.GetNfsServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfsServerId** | **int32** | NFSのID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNfsServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NfsServer**](NfsServer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNfsServerInterface

> NfsServerInterface GetNfsServerInterface(ctx, nfsServerId).Execute()

NFSのインターフェースを取得する

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/sakura_vps_client_go"
)

func main() {
	nfsServerId := int32(56) // int32 | NFSのID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsServerAPI.GetNfsServerInterface(context.Background(), nfsServerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsServerAPI.GetNfsServerInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNfsServerInterface`: NfsServerInterface
	fmt.Fprintf(os.Stdout, "Response from `NfsServerAPI.GetNfsServerInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfsServerId** | **int32** | NFSのID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNfsServerInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NfsServerInterface**](NfsServerInterface.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNfsServerList

> GetNfsServerList200Response GetNfsServerList(ctx).Page(page).PerPage(perPage).Id(id).Switch_(switch_).ZoneCode(zoneCode).Sort(sort).Execute()

NFS情報一覧を取得する

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/sakura_vps_client_go"
)

func main() {
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional) (default to 10)
	id := "id_example" // string | idで絞り込む。カンマ区切りで100件まで指定可能。 (optional)
	switch_ := int32(56) // int32 | switchのid 対象のスイッチに接続されているサーバーに絞り込む (optional)
	zoneCode := "zoneCode_example" // string | ゾーンコードで絞り込む * tk1 東京第1 * tk2 東京第2 * tk3 東京第3 * os1 大阪第1 * os2 大阪第2 * os3 大阪第3 * is1 石狩第1 (optional)
	sort := "sort_example" // string | データの並び順。項目名の頭に`-`をつけると降順で取得する   例: * サービスコード昇順: sort=service_code * サービスコード降順: sort=-service_code (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsServerAPI.GetNfsServerList(context.Background()).Page(page).PerPage(perPage).Id(id).Switch_(switch_).ZoneCode(zoneCode).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsServerAPI.GetNfsServerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNfsServerList`: GetNfsServerList200Response
	fmt.Fprintf(os.Stdout, "Response from `NfsServerAPI.GetNfsServerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNfsServerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **perPage** | **int32** |  | [default to 10]
 **id** | **string** | idで絞り込む。カンマ区切りで100件まで指定可能。 | 
 **switch_** | **int32** | switchのid 対象のスイッチに接続されているサーバーに絞り込む | 
 **zoneCode** | **string** | ゾーンコードで絞り込む * tk1 東京第1 * tk2 東京第2 * tk3 東京第3 * os1 大阪第1 * os2 大阪第2 * os3 大阪第3 * is1 石狩第1 | 
 **sort** | **string** | データの並び順。項目名の頭に&#x60;-&#x60;をつけると降順で取得する   例: * サービスコード昇順: sort&#x3D;service_code * サービスコード降順: sort&#x3D;-service_code | 

### Return type

[**GetNfsServerList200Response**](GetNfsServerList200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNfsServerPowerStatus

> NfsServerPowerStatus GetNfsServerPowerStatus(ctx, nfsServerId).Execute()

NFSの電源状態を取得する

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/sakura_vps_client_go"
)

func main() {
	nfsServerId := int32(56) // int32 | NFSのID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsServerAPI.GetNfsServerPowerStatus(context.Background(), nfsServerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsServerAPI.GetNfsServerPowerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNfsServerPowerStatus`: NfsServerPowerStatus
	fmt.Fprintf(os.Stdout, "Response from `NfsServerAPI.GetNfsServerPowerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfsServerId** | **int32** | NFSのID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNfsServerPowerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NfsServerPowerStatus**](NfsServerPowerStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNfsServerStorage

> NfsStorageInfo GetNfsServerStorage(ctx, nfsServerId).Execute()

NFSのストレージ容量情報を取得する

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/sakura_vps_client_go"
)

func main() {
	nfsServerId := int32(56) // int32 | NFSのID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsServerAPI.GetNfsServerStorage(context.Background(), nfsServerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsServerAPI.GetNfsServerStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNfsServerStorage`: NfsStorageInfo
	fmt.Fprintf(os.Stdout, "Response from `NfsServerAPI.GetNfsServerStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfsServerId** | **int32** | NFSのID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNfsServerStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NfsStorageInfo**](NfsStorageInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNfsServerChangeIpv4

> PostNfsServerChangeIpv4(ctx, nfsServerId).PostNfsServerChangeIpv4Request(postNfsServerChangeIpv4Request).Execute()

NFSのipv4を更新する

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/sakura_vps_client_go"
)

func main() {
	nfsServerId := int32(56) // int32 | NFSのID
	postNfsServerChangeIpv4Request := *openapiclient.NewPostNfsServerChangeIpv4Request("Address_example", "Netmask_example") // PostNfsServerChangeIpv4Request |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NfsServerAPI.PostNfsServerChangeIpv4(context.Background(), nfsServerId).PostNfsServerChangeIpv4Request(postNfsServerChangeIpv4Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsServerAPI.PostNfsServerChangeIpv4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfsServerId** | **int32** | NFSのID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostNfsServerChangeIpv4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postNfsServerChangeIpv4Request** | [**PostNfsServerChangeIpv4Request**](PostNfsServerChangeIpv4Request.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutNfsServer

> NfsServer PutNfsServer(ctx, nfsServerId).PutServerRequest(putServerRequest).Execute()

NFS情報を更新する

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/sakura_vps_client_go"
)

func main() {
	nfsServerId := int32(56) // int32 | NFSのID
	putServerRequest := *openapiclient.NewPutServerRequest("Name_example", "Description_example") // PutServerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsServerAPI.PutNfsServer(context.Background(), nfsServerId).PutServerRequest(putServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsServerAPI.PutNfsServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutNfsServer`: NfsServer
	fmt.Fprintf(os.Stdout, "Response from `NfsServerAPI.PutNfsServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfsServerId** | **int32** | NFSのID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutNfsServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putServerRequest** | [**PutServerRequest**](PutServerRequest.md) |  | 

### Return type

[**NfsServer**](NfsServer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutNfsServerInterface

> NfsServerInterface PutNfsServerInterface(ctx, nfsServerId).NfsServerInterface(nfsServerInterface).Execute()

NFSのインターフェース情報を更新する

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/sakura_vps_client_go"
)

func main() {
	nfsServerId := int32(56) // int32 | NFSのID
	nfsServerInterface := *openapiclient.NewNfsServerInterface(int32(123), "eth0", "switch", "9C:A3:BA:00:00:00", NullableInt32(1)) // NfsServerInterface |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsServerAPI.PutNfsServerInterface(context.Background(), nfsServerId).NfsServerInterface(nfsServerInterface).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsServerAPI.PutNfsServerInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutNfsServerInterface`: NfsServerInterface
	fmt.Fprintf(os.Stdout, "Response from `NfsServerAPI.PutNfsServerInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfsServerId** | **int32** | NFSのID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutNfsServerInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nfsServerInterface** | [**NfsServerInterface**](NfsServerInterface.md) |  | 

### Return type

[**NfsServerInterface**](NfsServerInterface.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

