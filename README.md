# これは何

さくらのVPS APIをGoで利用するための非公式クライアントライブラリです。

このライブラリは[上流の仕様書](https://manual.sakura.ad.jp/vps/api/api-doc/)が [CC-BY 4.0国際ライセンス](https://creativecommons.org/licenses/by/4.0/deed.ja)で許諾する部分を切り出して作成した[仕様書](/spec/spec.json)を元に、OpenAPI Generator の Go generator を用いてコード生成を行ったものです。

従って、このライブラリは[さくらインターネット](https://www.sakura.ad.jp)が著作権を有する仕様書(バージョン: 4.30)を元として作成されています。

# 利用方法

本リポジトリでは、クライアントライブラリの包括的な利用方法のサポートを提供しません。
その代わり、仕様書に含まれるAPIの利用方法をご自身のエディタで参照できるよう、/docs配下にAPIの利用方法のドキュメンテーションが格納されています。

goplsのようなLSPサーバーを用いれば、それぞれの関数の利用方法について有益なサジェストを得ることができるでしょう。


# ライセンス

Apache 2.0


# Go API client for sakura_vps_client_go

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 4.3.0
- Package version: 1.0.0
- Generator version: 7.11.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import sakura_vps_client_go "github.com/g1eng/sakura_vps_client_go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sakura_vps_client_go.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), sakura_vps_client_go.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sakura_vps_client_go.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), sakura_vps_client_go.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sakura_vps_client_go.ContextOperationServerIndices` and `sakura_vps_client_go.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), sakura_vps_client_go.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sakura_vps_client_go.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://secure.sakura.ad.jp/vps/api/v7*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApiKeyAPI* | [**DeleteApiKey**](docs/ApiKeyAPI.md#deleteapikey) | **Delete** /api-keys/{api_key_id} | APIキーを削除する
*ApiKeyAPI* | [**DeleteRole**](docs/ApiKeyAPI.md#deleterole) | **Delete** /roles/{role_id} | ロールを削除する
*ApiKeyAPI* | [**GetApiKey**](docs/ApiKeyAPI.md#getapikey) | **Get** /api-keys/{api_key_id} | APIキーを取得する
*ApiKeyAPI* | [**GetApiKeyList**](docs/ApiKeyAPI.md#getapikeylist) | **Get** /api-keys | APIキーの一覧を取得する
*ApiKeyAPI* | [**GetPermissionList**](docs/ApiKeyAPI.md#getpermissionlist) | **Get** /permissions | 権限の一覧を取得する
*ApiKeyAPI* | [**GetRole**](docs/ApiKeyAPI.md#getrole) | **Get** /roles/{role_id} | ロールを取得する
*ApiKeyAPI* | [**GetRoleList**](docs/ApiKeyAPI.md#getrolelist) | **Get** /roles | ロール一覧を取得する
*ApiKeyAPI* | [**PostApiKey**](docs/ApiKeyAPI.md#postapikey) | **Post** /api-keys | APIキーを作成する
*ApiKeyAPI* | [**PostApiKeyRotate**](docs/ApiKeyAPI.md#postapikeyrotate) | **Post** /api-keys/{api_key_id}/rotate | APIキーのトークンのローテーションを行う
*ApiKeyAPI* | [**PostRole**](docs/ApiKeyAPI.md#postrole) | **Post** /roles | ロールを作成する
*ApiKeyAPI* | [**PutApiKey**](docs/ApiKeyAPI.md#putapikey) | **Put** /api-keys/{api_key_id} | APIキーを更新する
*ApiKeyAPI* | [**PutRole**](docs/ApiKeyAPI.md#putrole) | **Put** /roles/{role_id} | ロールを更新する
*DiscAPI* | [**GetDiscList**](docs/DiscAPI.md#getdisclist) | **Get** /discs | ディスク情報一覧を取得する
*NfsServerAPI* | [**GetNfsServer**](docs/NfsServerAPI.md#getnfsserver) | **Get** /nfs-servers/{nfs_server_id} | NFS情報を取得する
*NfsServerAPI* | [**GetNfsServerInterface**](docs/NfsServerAPI.md#getnfsserverinterface) | **Get** /nfs-servers/{nfs_server_id}/interface | NFSのインターフェースを取得する
*NfsServerAPI* | [**GetNfsServerList**](docs/NfsServerAPI.md#getnfsserverlist) | **Get** /nfs-servers | NFS情報一覧を取得する
*NfsServerAPI* | [**GetNfsServerPowerStatus**](docs/NfsServerAPI.md#getnfsserverpowerstatus) | **Get** /nfs-servers/{nfs_server_id}/power-status | NFSの電源状態を取得する
*NfsServerAPI* | [**PostNfsServerChangeIpv4**](docs/NfsServerAPI.md#postnfsserverchangeipv4) | **Post** /nfs-servers/{nfs_server_id}/change-ipv4 | NFSのipv4を更新する
*NfsServerAPI* | [**PutNfsServer**](docs/NfsServerAPI.md#putnfsserver) | **Put** /nfs-servers/{nfs_server_id} | NFS情報を更新する
*NfsServerAPI* | [**PutNfsServerInterface**](docs/NfsServerAPI.md#putnfsserverinterface) | **Put** /nfs-servers/{nfs_server_id}/interface | NFSのインターフェース情報を更新する
*ServerAPI* | [**DeleteServerMonitorings**](docs/ServerAPI.md#deleteservermonitorings) | **Delete** /servers/{server_id}/monitorings/{server_monitoring_id} | サーバーのサーバー監視情報を削除する
*ServerAPI* | [**GetServer**](docs/ServerAPI.md#getserver) | **Get** /servers/{server_id} | サーバー情報を取得する
*ServerAPI* | [**GetServerInterface**](docs/ServerAPI.md#getserverinterface) | **Get** /servers/{server_id}/interfaces/{interface_id} | サーバーのインターフェース情報を取得する
*ServerAPI* | [**GetServerInterfaceList**](docs/ServerAPI.md#getserverinterfacelist) | **Get** /servers/{server_id}/interfaces | サーバーのインターフェース情報一覧を取得する
*ServerAPI* | [**GetServerLimitation**](docs/ServerAPI.md#getserverlimitation) | **Get** /servers/{server_id}/limitation | サーバーの制限情報を取得する
*ServerAPI* | [**GetServerList**](docs/ServerAPI.md#getserverlist) | **Get** /servers | サーバー情報一覧を取得する
*ServerAPI* | [**GetServerMonitoring**](docs/ServerAPI.md#getservermonitoring) | **Get** /servers/{server_id}/monitorings/{server_monitoring_id} | サーバーのサーバー監視情報を取得する
*ServerAPI* | [**GetServerMonitoringHealth**](docs/ServerAPI.md#getservermonitoringhealth) | **Get** /servers/{server_id}/monitorings/{server_monitoring_id}/health | サーバー監視の監視状態を取得する
*ServerAPI* | [**GetServerMonitoringList**](docs/ServerAPI.md#getservermonitoringlist) | **Get** /servers/{server_id}/monitorings | サーバーのサーバー監視情報一覧を取得する
*ServerAPI* | [**GetServerPowerStatus**](docs/ServerAPI.md#getserverpowerstatus) | **Get** /servers/{server_id}/power-status | サーバーの電源状態を取得する
*ServerAPI* | [**GetServerVideoDevice**](docs/ServerAPI.md#getservervideodevice) | **Get** /servers/{server_id}/video-device | サーバーのビデオデバイスの設定を取得する
*ServerAPI* | [**GetServerVncConsoleKeymap**](docs/ServerAPI.md#getservervncconsolekeymap) | **Get** /servers/{server_id}/vnc-console-keymap | コンパネのVNCコンソールのキーマップ情報を取得する
*ServerAPI* | [**PostServerForceReboot**](docs/ServerAPI.md#postserverforcereboot) | **Post** /servers/{server_id}/force-reboot | サーバーを強制再起動する
*ServerAPI* | [**PostServerMonitoring**](docs/ServerAPI.md#postservermonitoring) | **Post** /servers/{server_id}/monitorings | サーバーのサーバー監視を作成する
*ServerAPI* | [**PostServerMountDisc**](docs/ServerAPI.md#postservermountdisc) | **Post** /servers/{server_id}/mount-disc | サーバーにディスクをマウントする
*ServerAPI* | [**PostServerPowerOn**](docs/ServerAPI.md#postserverpoweron) | **Post** /servers/{server_id}/power-on | サーバーを起動する
*ServerAPI* | [**PostServerShutdown**](docs/ServerAPI.md#postservershutdown) | **Post** /servers/{server_id}/shutdown | サーバーをシャットダウンする
*ServerAPI* | [**PutServer**](docs/ServerAPI.md#putserver) | **Put** /servers/{server_id} | サーバー情報を更新する
*ServerAPI* | [**PutServerInterface**](docs/ServerAPI.md#putserverinterface) | **Put** /servers/{server_id}/interfaces/{interface_id} | サーバーのインターフェース情報を更新する
*ServerAPI* | [**PutServerIpv4Ptr**](docs/ServerAPI.md#putserveripv4ptr) | **Put** /servers/{server_id}/ipv4-ptr | サーバーのipv4の逆引きホスト名を更新する
*ServerAPI* | [**PutServerIpv6Ptr**](docs/ServerAPI.md#putserveripv6ptr) | **Put** /servers/{server_id}/ipv6-ptr | サーバーのipv6の逆引きホスト名を更新する
*ServerAPI* | [**PutServerMonitoring**](docs/ServerAPI.md#putservermonitoring) | **Put** /servers/{server_id}/monitorings/{server_monitoring_id} | サーバーのサーバー監視情報を更新する
*ServerAPI* | [**PutServerVideoDevice**](docs/ServerAPI.md#putservervideodevice) | **Put** /servers/{server_id}/video-device | サーバーのビデオデバイスの設定を更新する
*ServerAPI* | [**PutServerVncConsoleKeymap**](docs/ServerAPI.md#putservervncconsolekeymap) | **Put** /servers/{server_id}/vnc-console-keymap | コンパネのVNCコンソールのキーマップ情報を変更する
*SwitchAPI* | [**DeleteSwitch**](docs/SwitchAPI.md#deleteswitch) | **Delete** /switches/{switch_id} | スイッチを削除する
*SwitchAPI* | [**GetSwitch**](docs/SwitchAPI.md#getswitch) | **Get** /switches/{switch_id} | スイッチ情報を取得する
*SwitchAPI* | [**GetSwitchList**](docs/SwitchAPI.md#getswitchlist) | **Get** /switches | スイッチ情報一覧を取得する
*SwitchAPI* | [**PostSwitch**](docs/SwitchAPI.md#postswitch) | **Post** /switches | スイッチを作成する
*SwitchAPI* | [**PutSwitch**](docs/SwitchAPI.md#putswitch) | **Put** /switches/{switch_id} | スイッチ情報を更新する
*ZoneAPI* | [**GetZoneList**](docs/ZoneAPI.md#getzonelist) | **Get** /zones | ゾーン情報一覧を取得する


## Documentation For Models

 - [ApiKey](docs/ApiKey.md)
 - [Disc](docs/Disc.md)
 - [GetApiKeyList200Response](docs/GetApiKeyList200Response.md)
 - [GetDiscList200Response](docs/GetDiscList200Response.md)
 - [GetNfsServerList200Response](docs/GetNfsServerList200Response.md)
 - [GetPermissionList200Response](docs/GetPermissionList200Response.md)
 - [GetRoleList200Response](docs/GetRoleList200Response.md)
 - [GetServerInterfaceList200Response](docs/GetServerInterfaceList200Response.md)
 - [GetServerList200Response](docs/GetServerList200Response.md)
 - [GetServerList401Response](docs/GetServerList401Response.md)
 - [GetServerList403Response](docs/GetServerList403Response.md)
 - [GetServerMonitoringList200Response](docs/GetServerMonitoringList200Response.md)
 - [GetServerVideoDevice200Response](docs/GetServerVideoDevice200Response.md)
 - [GetSwitchList200Response](docs/GetSwitchList200Response.md)
 - [GetZoneList200Response](docs/GetZoneList200Response.md)
 - [HealthCheckHttp](docs/HealthCheckHttp.md)
 - [HealthCheckHttpBase](docs/HealthCheckHttpBase.md)
 - [HealthCheckHttps](docs/HealthCheckHttps.md)
 - [HealthCheckPing](docs/HealthCheckPing.md)
 - [HealthCheckPop3](docs/HealthCheckPop3.md)
 - [HealthCheckSmtp](docs/HealthCheckSmtp.md)
 - [HealthCheckSsh](docs/HealthCheckSsh.md)
 - [HealthCheckTcp](docs/HealthCheckTcp.md)
 - [InvalidParameterDetailInner](docs/InvalidParameterDetailInner.md)
 - [Keymap](docs/Keymap.md)
 - [Limitation](docs/Limitation.md)
 - [NfsServer](docs/NfsServer.md)
 - [NfsServerContract](docs/NfsServerContract.md)
 - [NfsServerInterface](docs/NfsServerInterface.md)
 - [NfsServerIpv4](docs/NfsServerIpv4.md)
 - [NfsServerPowerStatus](docs/NfsServerPowerStatus.md)
 - [NfsServerStorageInner](docs/NfsServerStorageInner.md)
 - [Pagination](docs/Pagination.md)
 - [Permission](docs/Permission.md)
 - [PostNfsServerChangeIpv4Request](docs/PostNfsServerChangeIpv4Request.md)
 - [PostServerMountDiscRequest](docs/PostServerMountDiscRequest.md)
 - [PostServerShutdownRequest](docs/PostServerShutdownRequest.md)
 - [PostSwitch201Response](docs/PostSwitch201Response.md)
 - [PostSwitch201ResponseZone](docs/PostSwitch201ResponseZone.md)
 - [PostSwitchRequest](docs/PostSwitchRequest.md)
 - [ProblemDetails400](docs/ProblemDetails400.md)
 - [ProblemDetails400Errors](docs/ProblemDetails400Errors.md)
 - [ProblemDetails404](docs/ProblemDetails404.md)
 - [ProblemDetails409](docs/ProblemDetails409.md)
 - [ProblemDetails429](docs/ProblemDetails429.md)
 - [ProblemDetails503](docs/ProblemDetails503.md)
 - [PutServerIpv4Ptr200Response](docs/PutServerIpv4Ptr200Response.md)
 - [PutServerIpv4PtrRequest](docs/PutServerIpv4PtrRequest.md)
 - [PutServerRequest](docs/PutServerRequest.md)
 - [PutServerVideoDeviceRequest](docs/PutServerVideoDeviceRequest.md)
 - [Role](docs/Role.md)
 - [RoleAllowedResources](docs/RoleAllowedResources.md)
 - [Server](docs/Server.md)
 - [ServerContract](docs/ServerContract.md)
 - [ServerInterface](docs/ServerInterface.md)
 - [ServerIpv4](docs/ServerIpv4.md)
 - [ServerIpv6](docs/ServerIpv6.md)
 - [ServerMonitoring](docs/ServerMonitoring.md)
 - [ServerMonitoringHealth](docs/ServerMonitoringHealth.md)
 - [ServerMonitoringSettings](docs/ServerMonitoringSettings.md)
 - [ServerMonitoringSettingsHealthCheck](docs/ServerMonitoringSettingsHealthCheck.md)
 - [ServerMonitoringSettingsNotification](docs/ServerMonitoringSettingsNotification.md)
 - [ServerMonitoringSettingsNotificationEmail](docs/ServerMonitoringSettingsNotificationEmail.md)
 - [ServerMonitoringSettingsNotificationIncomingWebhook](docs/ServerMonitoringSettingsNotificationIncomingWebhook.md)
 - [ServerPowerStatus](docs/ServerPowerStatus.md)
 - [ServerStorageInner](docs/ServerStorageInner.md)
 - [ServerZone](docs/ServerZone.md)
 - [Switch](docs/Switch.md)
 - [SwitchExternalConnection](docs/SwitchExternalConnection.md)
 - [SwitchExternalConnectionServicesInner](docs/SwitchExternalConnectionServicesInner.md)
 - [SwitchZone](docs/SwitchZone.md)
 - [Zone](docs/Zone.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### ApiKeyAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), sakura_vps_client_go.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



