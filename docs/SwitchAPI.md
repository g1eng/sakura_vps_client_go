# \SwitchAPI

All URIs are relative to *https://secure.sakura.ad.jp/vps/api/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSwitch**](SwitchAPI.md#DeleteSwitch) | **Delete** /switches/{switch_id} | スイッチを削除する
[**GetSwitch**](SwitchAPI.md#GetSwitch) | **Get** /switches/{switch_id} | スイッチ情報を取得する
[**GetSwitchList**](SwitchAPI.md#GetSwitchList) | **Get** /switches | スイッチ情報一覧を取得する
[**PostSwitch**](SwitchAPI.md#PostSwitch) | **Post** /switches | スイッチを作成する
[**PutSwitch**](SwitchAPI.md#PutSwitch) | **Put** /switches/{switch_id} | スイッチ情報を更新する



## DeleteSwitch

> DeleteSwitch(ctx, switchId).Execute()

スイッチを削除する

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
	switchId := int32(56) // int32 | スイッチID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SwitchAPI.DeleteSwitch(context.Background(), switchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.DeleteSwitch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**switchId** | **int32** | スイッチID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSwitchRequest struct via the builder pattern


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


## GetSwitch

> Switch GetSwitch(ctx, switchId).Execute()

スイッチ情報を取得する

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
	switchId := int32(56) // int32 | スイッチID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetSwitch(context.Background(), switchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetSwitch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSwitch`: Switch
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetSwitch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**switchId** | **int32** | スイッチID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSwitchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Switch**](Switch.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSwitchList

> GetSwitchList200Response GetSwitchList(ctx).Page(page).PerPage(perPage).Id(id).ZoneCode(zoneCode).Interface_(interface_).Sort(sort).Search(search).Execute()

スイッチ情報一覧を取得する

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
	zoneCode := "zoneCode_example" // string | ゾーンコードで絞り込む * tk1 東京第1 * tk2 東京第2 * tk3 東京第3 * os1 大阪第1 * os2 大阪第2 * os3 大阪第3 * is1 石狩第1 (optional)
	interface_ := int32(56) // int32 | interfaceのid 対象のインターフェースに接続されているスイッチに絞り込む (optional)
	sort := "sort_example" // string | データの並び順。項目名の頭に`-`をつけると降順で取得する   例: * 接続インターフェース数昇順: sort=interfaces_count * 接続インターフェース数降順: sort=-interfaces_count (optional)
	search := "search_example" // string | 名前と説明文から部分一致検索 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetSwitchList(context.Background()).Page(page).PerPage(perPage).Id(id).ZoneCode(zoneCode).Interface_(interface_).Sort(sort).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetSwitchList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSwitchList`: GetSwitchList200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetSwitchList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSwitchListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **perPage** | **int32** |  | [default to 10]
 **id** | **string** | idで絞り込む。カンマ区切りで100件まで指定可能。 | 
 **zoneCode** | **string** | ゾーンコードで絞り込む * tk1 東京第1 * tk2 東京第2 * tk3 東京第3 * os1 大阪第1 * os2 大阪第2 * os3 大阪第3 * is1 石狩第1 | 
 **interface_** | **int32** | interfaceのid 対象のインターフェースに接続されているスイッチに絞り込む | 
 **sort** | **string** | データの並び順。項目名の頭に&#x60;-&#x60;をつけると降順で取得する   例: * 接続インターフェース数昇順: sort&#x3D;interfaces_count * 接続インターフェース数降順: sort&#x3D;-interfaces_count | 
 **search** | **string** | 名前と説明文から部分一致検索 | 

### Return type

[**GetSwitchList200Response**](GetSwitchList200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSwitch

> PostSwitch201Response PostSwitch(ctx).PostSwitchRequest(postSwitchRequest).Execute()

スイッチを作成する



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
	postSwitchRequest := *openapiclient.NewPostSwitchRequest("Name_example", "Description_example", "ZoneCode_example") // PostSwitchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.PostSwitch(context.Background()).PostSwitchRequest(postSwitchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.PostSwitch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSwitch`: PostSwitch201Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.PostSwitch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSwitchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postSwitchRequest** | [**PostSwitchRequest**](PostSwitchRequest.md) |  | 

### Return type

[**PostSwitch201Response**](PostSwitch201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSwitch

> Switch PutSwitch(ctx, switchId).Switch_(switch_).Execute()

スイッチ情報を更新する

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
	switchId := int32(56) // int32 | スイッチID
	switch_ := *openapiclient.NewSwitch(int32(123), "Name_example", "Description_example", "SwitchCode_example", *openapiclient.NewSwitchZone("Code_example", "石狩第1"), []int32{int32(123)}, []int32{int32(123)}, "TODO") // Switch |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.PutSwitch(context.Background(), switchId).Switch_(switch_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.PutSwitch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSwitch`: Switch
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.PutSwitch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**switchId** | **int32** | スイッチID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSwitchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **switch_** | [**Switch**](Switch.md) |  | 

### Return type

[**Switch**](Switch.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

