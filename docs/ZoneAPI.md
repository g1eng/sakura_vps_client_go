# \ZoneAPI

All URIs are relative to *https://secure.sakura.ad.jp/vps/api/v7*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetZoneList**](ZoneAPI.md#GetZoneList) | **Get** /zones | ゾーン情報一覧を取得する



## GetZoneList

> GetZoneList200Response GetZoneList(ctx).Page(page).PerPage(perPage).Execute()

ゾーン情報一覧を取得する

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZoneAPI.GetZoneList(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneAPI.GetZoneList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZoneList`: GetZoneList200Response
	fmt.Fprintf(os.Stdout, "Response from `ZoneAPI.GetZoneList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **perPage** | **int32** |  | [default to 10]

### Return type

[**GetZoneList200Response**](GetZoneList200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

