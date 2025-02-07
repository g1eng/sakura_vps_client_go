# \ServerAPI

All URIs are relative to *https://secure.sakura.ad.jp/vps/api/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServerMonitorings**](ServerAPI.md#DeleteServerMonitorings) | **Delete** /servers/{server_id}/monitorings/{server_monitoring_id} | サーバーのサーバー監視情報を削除する
[**GetServer**](ServerAPI.md#GetServer) | **Get** /servers/{server_id} | サーバー情報を取得する
[**GetServerInterface**](ServerAPI.md#GetServerInterface) | **Get** /servers/{server_id}/interfaces/{interface_id} | サーバーのインターフェース情報を取得する
[**GetServerInterfaceList**](ServerAPI.md#GetServerInterfaceList) | **Get** /servers/{server_id}/interfaces | サーバーのインターフェース情報一覧を取得する
[**GetServerLimitation**](ServerAPI.md#GetServerLimitation) | **Get** /servers/{server_id}/limitation | サーバーの制限情報を取得する
[**GetServerList**](ServerAPI.md#GetServerList) | **Get** /servers | サーバー情報一覧を取得する
[**GetServerMonitoring**](ServerAPI.md#GetServerMonitoring) | **Get** /servers/{server_id}/monitorings/{server_monitoring_id} | サーバーのサーバー監視情報を取得する
[**GetServerMonitoringHealth**](ServerAPI.md#GetServerMonitoringHealth) | **Get** /servers/{server_id}/monitorings/{server_monitoring_id}/health | サーバー監視の監視状態を取得する
[**GetServerMonitoringList**](ServerAPI.md#GetServerMonitoringList) | **Get** /servers/{server_id}/monitorings | サーバーのサーバー監視情報一覧を取得する
[**GetServerPowerStatus**](ServerAPI.md#GetServerPowerStatus) | **Get** /servers/{server_id}/power-status | サーバーの電源状態を取得する
[**GetServerVideoDevice**](ServerAPI.md#GetServerVideoDevice) | **Get** /servers/{server_id}/video-device | サーバーのビデオデバイスの設定を取得する
[**GetServerVncConsoleKeymap**](ServerAPI.md#GetServerVncConsoleKeymap) | **Get** /servers/{server_id}/vnc-console-keymap | コンパネのVNCコンソールのキーマップ情報を取得する
[**PostServerForceReboot**](ServerAPI.md#PostServerForceReboot) | **Post** /servers/{server_id}/force-reboot | サーバーを強制再起動する
[**PostServerMonitoring**](ServerAPI.md#PostServerMonitoring) | **Post** /servers/{server_id}/monitorings | サーバーのサーバー監視を作成する
[**PostServerMountDisc**](ServerAPI.md#PostServerMountDisc) | **Post** /servers/{server_id}/mount-disc | サーバーにディスクをマウントする
[**PostServerPowerOn**](ServerAPI.md#PostServerPowerOn) | **Post** /servers/{server_id}/power-on | サーバーを起動する
[**PostServerShutdown**](ServerAPI.md#PostServerShutdown) | **Post** /servers/{server_id}/shutdown | サーバーをシャットダウンする
[**PutServer**](ServerAPI.md#PutServer) | **Put** /servers/{server_id} | サーバー情報を更新する
[**PutServerInterface**](ServerAPI.md#PutServerInterface) | **Put** /servers/{server_id}/interfaces/{interface_id} | サーバーのインターフェース情報を更新する
[**PutServerIpv4Ptr**](ServerAPI.md#PutServerIpv4Ptr) | **Put** /servers/{server_id}/ipv4-ptr | サーバーのipv4の逆引きホスト名を更新する
[**PutServerIpv6Ptr**](ServerAPI.md#PutServerIpv6Ptr) | **Put** /servers/{server_id}/ipv6-ptr | サーバーのipv6の逆引きホスト名を更新する
[**PutServerMonitoring**](ServerAPI.md#PutServerMonitoring) | **Put** /servers/{server_id}/monitorings/{server_monitoring_id} | サーバーのサーバー監視情報を更新する
[**PutServerVideoDevice**](ServerAPI.md#PutServerVideoDevice) | **Put** /servers/{server_id}/video-device | サーバーのビデオデバイスの設定を更新する
[**PutServerVncConsoleKeymap**](ServerAPI.md#PutServerVncConsoleKeymap) | **Put** /servers/{server_id}/vnc-console-keymap | コンパネのVNCコンソールのキーマップ情報を変更する



## DeleteServerMonitorings

> DeleteServerMonitorings(ctx, serverId, serverMonitoringId).Execute()

サーバーのサーバー監視情報を削除する

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
	serverId := int32(56) // int32 | サーバーID
	serverMonitoringId := int32(56) // int32 | サーバー監視ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerAPI.DeleteServerMonitorings(context.Background(), serverId, serverMonitoringId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.DeleteServerMonitorings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 
**serverMonitoringId** | **int32** | サーバー監視ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerMonitoringsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServer

> Server GetServer(ctx, serverId).Execute()

サーバー情報を取得する

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
	serverId := int32(56) // int32 | サーバーID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServer`: Server
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Server**](Server.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerInterface

> ServerInterface GetServerInterface(ctx, serverId, interfaceId).Execute()

サーバーのインターフェース情報を取得する

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
	serverId := int32(56) // int32 | サーバーID
	interfaceId := int32(56) // int32 | サーバーインターフェースID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServerInterface(context.Background(), serverId, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServerInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerInterface`: ServerInterface
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServerInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 
**interfaceId** | **int32** | サーバーインターフェースID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServerInterface**](ServerInterface.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerInterfaceList

> GetServerInterfaceList200Response GetServerInterfaceList(ctx, serverId).Page(page).PerPage(perPage).Execute()

サーバーのインターフェース情報一覧を取得する

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
	serverId := int32(56) // int32 | サーバーID
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServerInterfaceList(context.Background(), serverId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServerInterfaceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerInterfaceList`: GetServerInterfaceList200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServerInterfaceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInterfaceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **perPage** | **int32** |  | [default to 10]

### Return type

[**GetServerInterfaceList200Response**](GetServerInterfaceList200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerLimitation

> Limitation GetServerLimitation(ctx, serverId).Execute()

サーバーの制限情報を取得する

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
	serverId := int32(56) // int32 | サーバーID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServerLimitation(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServerLimitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerLimitation`: Limitation
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServerLimitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerLimitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Limitation**](Limitation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerList

> GetServerList200Response GetServerList(ctx).Page(page).PerPage(perPage).Id(id).Switch_(switch_).ZoneCode(zoneCode).ServiceType(serviceType).Ipv4Address(ipv4Address).MonitoringResourceId(monitoringResourceId).Sort(sort).Search(search).Execute()

サーバー情報一覧を取得する

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
	serviceType := "serviceType_example" // string | サービスタイプで絞り込む (optional)
	ipv4Address := "192.0.2.0" // string | IPアドレスで絞り込む (optional)
	monitoringResourceId := "monitoringResourceId_example" // string | サーバー監視リソースIDで絞り込む (optional)
	sort := "sort_example" // string | データの並び順。項目名の頭に`-`をつけると降順で取得する   例: * サービスコード昇順: sort=service_code * サービスコード降順: sort=-service_code (optional)
	search := "search_example" // string | 名前、説明文、ホスト名、ipv4アドレス、ipv6アドレス、サービスコードから部分一致検索 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServerList(context.Background()).Page(page).PerPage(perPage).Id(id).Switch_(switch_).ZoneCode(zoneCode).ServiceType(serviceType).Ipv4Address(ipv4Address).MonitoringResourceId(monitoringResourceId).Sort(sort).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerList`: GetServerList200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **perPage** | **int32** |  | [default to 10]
 **id** | **string** | idで絞り込む。カンマ区切りで100件まで指定可能。 | 
 **switch_** | **int32** | switchのid 対象のスイッチに接続されているサーバーに絞り込む | 
 **zoneCode** | **string** | ゾーンコードで絞り込む * tk1 東京第1 * tk2 東京第2 * tk3 東京第3 * os1 大阪第1 * os2 大阪第2 * os3 大阪第3 * is1 石狩第1 | 
 **serviceType** | **string** | サービスタイプで絞り込む | 
 **ipv4Address** | **string** | IPアドレスで絞り込む | 
 **monitoringResourceId** | **string** | サーバー監視リソースIDで絞り込む | 
 **sort** | **string** | データの並び順。項目名の頭に&#x60;-&#x60;をつけると降順で取得する   例: * サービスコード昇順: sort&#x3D;service_code * サービスコード降順: sort&#x3D;-service_code | 
 **search** | **string** | 名前、説明文、ホスト名、ipv4アドレス、ipv6アドレス、サービスコードから部分一致検索 | 

### Return type

[**GetServerList200Response**](GetServerList200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerMonitoring

> ServerMonitoring GetServerMonitoring(ctx, serverId, serverMonitoringId).Page(page).PerPage(perPage).Execute()

サーバーのサーバー監視情報を取得する

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
	serverId := int32(56) // int32 | サーバーID
	serverMonitoringId := int32(56) // int32 | サーバー監視ID
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServerMonitoring(context.Background(), serverId, serverMonitoringId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServerMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerMonitoring`: ServerMonitoring
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServerMonitoring`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 
**serverMonitoringId** | **int32** | サーバー監視ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** |  | 
 **perPage** | **int32** |  | [default to 10]

### Return type

[**ServerMonitoring**](ServerMonitoring.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerMonitoringHealth

> ServerMonitoringHealth GetServerMonitoringHealth(ctx, serverId, serverMonitoringId).Execute()

サーバー監視の監視状態を取得する

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
	serverId := int32(56) // int32 | サーバーID
	serverMonitoringId := int32(56) // int32 | サーバー監視ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServerMonitoringHealth(context.Background(), serverId, serverMonitoringId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServerMonitoringHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerMonitoringHealth`: ServerMonitoringHealth
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServerMonitoringHealth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 
**serverMonitoringId** | **int32** | サーバー監視ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerMonitoringHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServerMonitoringHealth**](ServerMonitoringHealth.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerMonitoringList

> GetServerMonitoringList200Response GetServerMonitoringList(ctx, serverId).Page(page).PerPage(perPage).MonitoringResourceId(monitoringResourceId).Execute()

サーバーのサーバー監視情報一覧を取得する

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
	serverId := int32(56) // int32 | サーバーID
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional) (default to 10)
	monitoringResourceId := "monitoringResourceId_example" // string | サーバー監視リソースIDで絞り込む (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServerMonitoringList(context.Background(), serverId).Page(page).PerPage(perPage).MonitoringResourceId(monitoringResourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServerMonitoringList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerMonitoringList`: GetServerMonitoringList200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServerMonitoringList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerMonitoringListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **perPage** | **int32** |  | [default to 10]
 **monitoringResourceId** | **string** | サーバー監視リソースIDで絞り込む | 

### Return type

[**GetServerMonitoringList200Response**](GetServerMonitoringList200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerPowerStatus

> ServerPowerStatus GetServerPowerStatus(ctx, serverId).Execute()

サーバーの電源状態を取得する

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
	serverId := int32(56) // int32 | サーバーID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServerPowerStatus(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServerPowerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerPowerStatus`: ServerPowerStatus
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServerPowerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerPowerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerPowerStatus**](ServerPowerStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerVideoDevice

> GetServerVideoDevice200Response GetServerVideoDevice(ctx, serverId).Execute()

サーバーのビデオデバイスの設定を取得する

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
	serverId := int32(56) // int32 | サーバーID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServerVideoDevice(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServerVideoDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerVideoDevice`: GetServerVideoDevice200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServerVideoDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerVideoDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetServerVideoDevice200Response**](GetServerVideoDevice200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerVncConsoleKeymap

> Keymap GetServerVncConsoleKeymap(ctx, serverId).Execute()

コンパネのVNCコンソールのキーマップ情報を取得する

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
	serverId := int32(56) // int32 | サーバーID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.GetServerVncConsoleKeymap(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.GetServerVncConsoleKeymap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerVncConsoleKeymap`: Keymap
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.GetServerVncConsoleKeymap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerVncConsoleKeymapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Keymap**](Keymap.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServerForceReboot

> PostServerForceReboot(ctx, serverId).Execute()

サーバーを強制再起動する

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
	serverId := int32(56) // int32 | サーバーID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerAPI.PostServerForceReboot(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.PostServerForceReboot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServerForceRebootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServerMonitoring

> ServerMonitoring PostServerMonitoring(ctx, serverId).ServerMonitoring(serverMonitoring).Execute()

サーバーのサーバー監視を作成する

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
	serverId := int32(56) // int32 | サーバーID
	serverMonitoring := *openapiclient.NewServerMonitoring(int32(123), "Name_example", "Description_example", "MonitoringResourceId_example", "UpdateStatus_example", *openapiclient.NewServerMonitoringSettings(false, openapiclient.ServerMonitoring_settings_health_check{HealthCheckHttp: openapiclient.NewHealthCheckHttp(int32(80), "Host_example", "/", "BasicAuthUsername_example", "BasicAuthPassword_example", int32(123), int32(123), "Protocol_example")}, *openapiclient.NewServerMonitoringSettingsNotification(*openapiclient.NewServerMonitoringSettingsNotificationEmail(false), *openapiclient.NewServerMonitoringSettingsNotificationIncomingWebhook(false, "WebhooksUrl_example", "SlackTeamName_example", "SlackChannelName_example"), int32(123)))) // ServerMonitoring |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.PostServerMonitoring(context.Background(), serverId).ServerMonitoring(serverMonitoring).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.PostServerMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServerMonitoring`: ServerMonitoring
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.PostServerMonitoring`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServerMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverMonitoring** | [**ServerMonitoring**](ServerMonitoring.md) |  | 

### Return type

[**ServerMonitoring**](ServerMonitoring.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServerMountDisc

> PostServerMountDisc(ctx, serverId).PostServerMountDiscRequest(postServerMountDiscRequest).Execute()

サーバーにディスクをマウントする

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
	serverId := int32(56) // int32 | サーバーID
	postServerMountDiscRequest := *openapiclient.NewPostServerMountDiscRequest(int32(123)) // PostServerMountDiscRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerAPI.PostServerMountDisc(context.Background(), serverId).PostServerMountDiscRequest(postServerMountDiscRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.PostServerMountDisc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServerMountDiscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postServerMountDiscRequest** | [**PostServerMountDiscRequest**](PostServerMountDiscRequest.md) |  | 

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


## PostServerPowerOn

> PostServerPowerOn(ctx, serverId).Execute()

サーバーを起動する

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
	serverId := int32(56) // int32 | サーバーID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerAPI.PostServerPowerOn(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.PostServerPowerOn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServerPowerOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServerShutdown

> PostServerShutdown(ctx, serverId).PostServerShutdownRequest(postServerShutdownRequest).Execute()

サーバーをシャットダウンする

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
	serverId := int32(56) // int32 | サーバーID
	postServerShutdownRequest := *openapiclient.NewPostServerShutdownRequest() // PostServerShutdownRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerAPI.PostServerShutdown(context.Background(), serverId).PostServerShutdownRequest(postServerShutdownRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.PostServerShutdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServerShutdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postServerShutdownRequest** | [**PostServerShutdownRequest**](PostServerShutdownRequest.md) |  | 

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


## PutServer

> Server PutServer(ctx, serverId).PutServerRequest(putServerRequest).Execute()

サーバー情報を更新する

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
	serverId := int32(56) // int32 | サーバーID
	putServerRequest := *openapiclient.NewPutServerRequest("Name_example", "Description_example") // PutServerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.PutServer(context.Background(), serverId).PutServerRequest(putServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.PutServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServer`: Server
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.PutServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putServerRequest** | [**PutServerRequest**](PutServerRequest.md) |  | 

### Return type

[**Server**](Server.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServerInterface

> ServerInterface PutServerInterface(ctx, serverId, interfaceId).ServerInterface(serverInterface).Execute()

サーバーのインターフェース情報を更新する

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
	serverId := int32(56) // int32 | サーバーID
	interfaceId := int32(56) // int32 | サーバーインターフェースID
	serverInterface := *openapiclient.NewServerInterface(int32(123), "eth0", true, "global", "9C:A3:BA:00:00:00", NullableInt32(1)) // ServerInterface |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.PutServerInterface(context.Background(), serverId, interfaceId).ServerInterface(serverInterface).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.PutServerInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServerInterface`: ServerInterface
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.PutServerInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 
**interfaceId** | **int32** | サーバーインターフェースID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServerInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serverInterface** | [**ServerInterface**](ServerInterface.md) |  | 

### Return type

[**ServerInterface**](ServerInterface.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServerIpv4Ptr

> PutServerIpv4Ptr200Response PutServerIpv4Ptr(ctx, serverId).PutServerIpv4PtrRequest(putServerIpv4PtrRequest).Execute()

サーバーのipv4の逆引きホスト名を更新する

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
	serverId := int32(56) // int32 | サーバーID
	putServerIpv4PtrRequest := *openapiclient.NewPutServerIpv4PtrRequest("example.jp") // PutServerIpv4PtrRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.PutServerIpv4Ptr(context.Background(), serverId).PutServerIpv4PtrRequest(putServerIpv4PtrRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.PutServerIpv4Ptr``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServerIpv4Ptr`: PutServerIpv4Ptr200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.PutServerIpv4Ptr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServerIpv4PtrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putServerIpv4PtrRequest** | [**PutServerIpv4PtrRequest**](PutServerIpv4PtrRequest.md) |  | 

### Return type

[**PutServerIpv4Ptr200Response**](PutServerIpv4Ptr200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServerIpv6Ptr

> PutServerIpv4Ptr200Response PutServerIpv6Ptr(ctx, serverId).PutServerIpv4PtrRequest(putServerIpv4PtrRequest).Execute()

サーバーのipv6の逆引きホスト名を更新する

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
	serverId := int32(56) // int32 | サーバーID
	putServerIpv4PtrRequest := *openapiclient.NewPutServerIpv4PtrRequest("example.jp") // PutServerIpv4PtrRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.PutServerIpv6Ptr(context.Background(), serverId).PutServerIpv4PtrRequest(putServerIpv4PtrRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.PutServerIpv6Ptr``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServerIpv6Ptr`: PutServerIpv4Ptr200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.PutServerIpv6Ptr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServerIpv6PtrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putServerIpv4PtrRequest** | [**PutServerIpv4PtrRequest**](PutServerIpv4PtrRequest.md) |  | 

### Return type

[**PutServerIpv4Ptr200Response**](PutServerIpv4Ptr200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServerMonitoring

> ServerMonitoring PutServerMonitoring(ctx, serverId, serverMonitoringId).ServerMonitoring(serverMonitoring).Execute()

サーバーのサーバー監視情報を更新する

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
	serverId := int32(56) // int32 | サーバーID
	serverMonitoringId := int32(56) // int32 | サーバー監視ID
	serverMonitoring := *openapiclient.NewServerMonitoring(int32(123), "Name_example", "Description_example", "MonitoringResourceId_example", "UpdateStatus_example", *openapiclient.NewServerMonitoringSettings(false, openapiclient.ServerMonitoring_settings_health_check{HealthCheckHttp: openapiclient.NewHealthCheckHttp(int32(80), "Host_example", "/", "BasicAuthUsername_example", "BasicAuthPassword_example", int32(123), int32(123), "Protocol_example")}, *openapiclient.NewServerMonitoringSettingsNotification(*openapiclient.NewServerMonitoringSettingsNotificationEmail(false), *openapiclient.NewServerMonitoringSettingsNotificationIncomingWebhook(false, "WebhooksUrl_example", "SlackTeamName_example", "SlackChannelName_example"), int32(123)))) // ServerMonitoring |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.PutServerMonitoring(context.Background(), serverId, serverMonitoringId).ServerMonitoring(serverMonitoring).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.PutServerMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServerMonitoring`: ServerMonitoring
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.PutServerMonitoring`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 
**serverMonitoringId** | **int32** | サーバー監視ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServerMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serverMonitoring** | [**ServerMonitoring**](ServerMonitoring.md) |  | 

### Return type

[**ServerMonitoring**](ServerMonitoring.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServerVideoDevice

> PutServerVideoDeviceRequest PutServerVideoDevice(ctx, serverId).PutServerVideoDeviceRequest(putServerVideoDeviceRequest).Execute()

サーバーのビデオデバイスの設定を更新する

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
	serverId := int32(56) // int32 | サーバーID
	putServerVideoDeviceRequest := *openapiclient.NewPutServerVideoDeviceRequest("Type_example") // PutServerVideoDeviceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.PutServerVideoDevice(context.Background(), serverId).PutServerVideoDeviceRequest(putServerVideoDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.PutServerVideoDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServerVideoDevice`: PutServerVideoDeviceRequest
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.PutServerVideoDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServerVideoDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putServerVideoDeviceRequest** | [**PutServerVideoDeviceRequest**](PutServerVideoDeviceRequest.md) |  | 

### Return type

[**PutServerVideoDeviceRequest**](PutServerVideoDeviceRequest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServerVncConsoleKeymap

> Keymap PutServerVncConsoleKeymap(ctx, serverId).Keymap(keymap).Execute()

コンパネのVNCコンソールのキーマップ情報を変更する

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
	serverId := int32(56) // int32 | サーバーID
	keymap := *openapiclient.NewKeymap("Layout_example") // Keymap |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerAPI.PutServerVncConsoleKeymap(context.Background(), serverId).Keymap(keymap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAPI.PutServerVncConsoleKeymap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServerVncConsoleKeymap`: Keymap
	fmt.Fprintf(os.Stdout, "Response from `ServerAPI.PutServerVncConsoleKeymap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | サーバーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServerVncConsoleKeymapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keymap** | [**Keymap**](Keymap.md) |  | 

### Return type

[**Keymap**](Keymap.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

