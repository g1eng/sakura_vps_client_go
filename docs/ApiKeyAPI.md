# \ApiKeyAPI

All URIs are relative to *https://secure.sakura.ad.jp/vps/api/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiKey**](ApiKeyAPI.md#DeleteApiKey) | **Delete** /api-keys/{api_key_id} | APIキーを削除する
[**DeleteRole**](ApiKeyAPI.md#DeleteRole) | **Delete** /roles/{role_id} | ロールを削除する
[**GetApiKey**](ApiKeyAPI.md#GetApiKey) | **Get** /api-keys/{api_key_id} | APIキーを取得する
[**GetApiKeyList**](ApiKeyAPI.md#GetApiKeyList) | **Get** /api-keys | APIキーの一覧を取得する
[**GetPermissionList**](ApiKeyAPI.md#GetPermissionList) | **Get** /permissions | 権限の一覧を取得する
[**GetRole**](ApiKeyAPI.md#GetRole) | **Get** /roles/{role_id} | ロールを取得する
[**GetRoleList**](ApiKeyAPI.md#GetRoleList) | **Get** /roles | ロール一覧を取得する
[**PostApiKey**](ApiKeyAPI.md#PostApiKey) | **Post** /api-keys | APIキーを作成する
[**PostApiKeyRotate**](ApiKeyAPI.md#PostApiKeyRotate) | **Post** /api-keys/{api_key_id}/rotate | APIキーのトークンのローテーションを行う
[**PostRole**](ApiKeyAPI.md#PostRole) | **Post** /roles | ロールを作成する
[**PutApiKey**](ApiKeyAPI.md#PutApiKey) | **Put** /api-keys/{api_key_id} | APIキーを更新する
[**PutRole**](ApiKeyAPI.md#PutRole) | **Put** /roles/{role_id} | ロールを更新する



## DeleteApiKey

> DeleteApiKey(ctx, apiKeyId).Execute()

APIキーを削除する

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
	apiKeyId := int32(56) // int32 | APIキーID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApiKeyAPI.DeleteApiKey(context.Background(), apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.DeleteApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **int32** | APIキーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiKeyRequest struct via the builder pattern


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


## DeleteRole

> DeleteRole(ctx, roleId).Execute()

ロールを削除する

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
	roleId := int32(56) // int32 | ロールID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApiKeyAPI.DeleteRole(context.Background(), roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.DeleteRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** | ロールID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


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


## GetApiKey

> ApiKey GetApiKey(ctx, apiKeyId).Execute()

APIキーを取得する

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
	apiKeyId := int32(56) // int32 | APIキーID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.GetApiKey(context.Background(), apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.GetApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiKey`: ApiKey
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.GetApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **int32** | APIキーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKeyList

> GetApiKeyList200Response GetApiKeyList(ctx).Page(page).PerPage(perPage).Order(order).Role(role).Search(search).Execute()

APIキーの一覧を取得する

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
	order := "order_example" // string | データの並び順。項目名の頭に`-`をつけると降順で取得する   例: * 名称昇順: sort=name * 名称降順: sort=-name (optional)
	role := int32(56) // int32 | roleのid。ロールで絞り込む (optional)
	search := "search_example" // string | 名前から部分一致検索 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.GetApiKeyList(context.Background()).Page(page).PerPage(perPage).Order(order).Role(role).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.GetApiKeyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiKeyList`: GetApiKeyList200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.GetApiKeyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **perPage** | **int32** |  | [default to 10]
 **order** | **string** | データの並び順。項目名の頭に&#x60;-&#x60;をつけると降順で取得する   例: * 名称昇順: sort&#x3D;name * 名称降順: sort&#x3D;-name | 
 **role** | **int32** | roleのid。ロールで絞り込む | 
 **search** | **string** | 名前から部分一致検索 | 

### Return type

[**GetApiKeyList200Response**](GetApiKeyList200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPermissionList

> GetPermissionList200Response GetPermissionList(ctx).Page(page).PerPage(perPage).Category(category).Code(code).Execute()

権限の一覧を取得する

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
	category := "server" // string | カテゴリーで絞り込む (optional)
	code := "get-server-list,get-server,put-server" // string | コードで絞り込む。カンマ区切りで複数指定可能。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.GetPermissionList(context.Background()).Page(page).PerPage(perPage).Category(category).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.GetPermissionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPermissionList`: GetPermissionList200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.GetPermissionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **perPage** | **int32** |  | [default to 10]
 **category** | **string** | カテゴリーで絞り込む | 
 **code** | **string** | コードで絞り込む。カンマ区切りで複数指定可能。 | 

### Return type

[**GetPermissionList200Response**](GetPermissionList200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> Role GetRole(ctx, roleId).Execute()

ロールを取得する

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
	roleId := int32(56) // int32 | ロールID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.GetRole(context.Background(), roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.GetRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRole`: Role
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** | ロールID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Role**](Role.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleList

> GetRoleList200Response GetRoleList(ctx).Page(page).PerPage(perPage).Order(order).Search(search).Execute()

ロール一覧を取得する

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
	order := "order_example" // string | データの並び順。項目名の頭に`-`をつけると降順で取得する   例: * 名称昇順: sort=name * 名称降順: sort=-name (optional)
	search := "search_example" // string | 名前と説明文から部分一致検索 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.GetRoleList(context.Background()).Page(page).PerPage(perPage).Order(order).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.GetRoleList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleList`: GetRoleList200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.GetRoleList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **perPage** | **int32** |  | [default to 10]
 **order** | **string** | データの並び順。項目名の頭に&#x60;-&#x60;をつけると降順で取得する   例: * 名称昇順: sort&#x3D;name * 名称降順: sort&#x3D;-name | 
 **search** | **string** | 名前と説明文から部分一致検索 | 

### Return type

[**GetRoleList200Response**](GetRoleList200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiKey

> ApiKey PostApiKey(ctx).ApiKey(apiKey).Execute()

APIキーを作成する

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
	apiKey := *openapiclient.NewApiKey(int32(123), "Name_example", int32(123), "Token_example") // ApiKey |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.PostApiKey(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.PostApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiKey`: ApiKey
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.PostApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | [**ApiKey**](ApiKey.md) |  | 

### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiKeyRotate

> ApiKey PostApiKeyRotate(ctx, apiKeyId).Execute()

APIキーのトークンのローテーションを行う

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
	apiKeyId := int32(56) // int32 | APIキーID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.PostApiKeyRotate(context.Background(), apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.PostApiKeyRotate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiKeyRotate`: ApiKey
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.PostApiKeyRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **int32** | APIキーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiKeyRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRole

> Role PostRole(ctx).Role(role).Execute()

ロールを作成する

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
	role := *openapiclient.NewRole(int32(123), "Name_example", "Description_example", "PermissionFiltering_example", []string{"AllowedPermissions_example"}, "ResourceFiltering_example", "TODO") // Role |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.PostRole(context.Background()).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.PostRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRole`: Role
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.PostRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | [**Role**](Role.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutApiKey

> ApiKey PutApiKey(ctx, apiKeyId).ApiKey(apiKey).Execute()

APIキーを更新する

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
	apiKeyId := int32(56) // int32 | APIキーID
	apiKey := *openapiclient.NewApiKey(int32(123), "Name_example", int32(123), "Token_example") // ApiKey |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.PutApiKey(context.Background(), apiKeyId).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.PutApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutApiKey`: ApiKey
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.PutApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **int32** | APIキーID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | [**ApiKey**](ApiKey.md) |  | 

### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRole

> Role PutRole(ctx, roleId).Role(role).Execute()

ロールを更新する

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
	roleId := int32(56) // int32 | ロールID
	role := *openapiclient.NewRole(int32(123), "Name_example", "Description_example", "PermissionFiltering_example", []string{"AllowedPermissions_example"}, "ResourceFiltering_example", "TODO") // Role |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.PutRole(context.Background(), roleId).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.PutRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRole`: Role
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.PutRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** | ロールID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | [**Role**](Role.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

