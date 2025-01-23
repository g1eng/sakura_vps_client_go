# Server

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | id | 
**Name** | **string** | 名前 | 
**Description** | **string** | 説明 | 
**ServiceType** | **string** | サービスタイプ | 
**ServiceStatus** | **string** | サービス状況 * on_trial お試し中 * link_down_on_trial お試し中（一時停止） * in_use 利用中 * link_down 一時停止中 | 
**CpuCores** | **int32** | CPUコア数 | 
**MemoryMebibytes** | **int32** | メモリ容量(MiB) | 
**Storage** | [**[]ServerStorageInner**](ServerStorageInner.md) | ストレージ情報 | 
**Zone** | [**ServerZone**](ServerZone.md) |  | 
**Options** | **[]string** | オプション（追加ソフトウェア） | 
**Version** | **string** | プランの世代 | 
**Ipv4** | [**ServerIpv4**](ServerIpv4.md) |  | 
**Ipv6** | [**ServerIpv6**](ServerIpv6.md) |  | 
**Contract** | [**ServerContract**](ServerContract.md) |  | 
**PowerStatus** | **string** | 電源ステータス * power_on 電源ON * in_shutdown シャットダウン中 * power_off 電源OFF * installing OSインストール中 * in_scaleup スケールアップ中 * migration サーバー移行作業中 * unknown 不明（電源状態を取得できない） このエンドポイントが返す電源ステータスはキャッシュされた情報のため、最新の正確な電源ステータスが必要な場合は **サーバーの電源状態を取得する**&#x60;/servers/{server_id}/power-status&#x60;をご利用ください | 

## Methods

### NewServer

`func NewServer(id int32, name string, description string, serviceType string, serviceStatus string, cpuCores int32, memoryMebibytes int32, storage []ServerStorageInner, zone ServerZone, options []string, version string, ipv4 ServerIpv4, ipv6 ServerIpv6, contract ServerContract, powerStatus string, ) *Server`

NewServer instantiates a new Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWithDefaults

`func NewServerWithDefaults() *Server`

NewServerWithDefaults instantiates a new Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Server) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Server) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Server) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Server) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Server) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Server) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Server) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Server) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Server) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetServiceType

`func (o *Server) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *Server) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *Server) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.


### GetServiceStatus

`func (o *Server) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *Server) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *Server) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetCpuCores

`func (o *Server) GetCpuCores() int32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *Server) GetCpuCoresOk() (*int32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *Server) SetCpuCores(v int32)`

SetCpuCores sets CpuCores field to given value.


### GetMemoryMebibytes

`func (o *Server) GetMemoryMebibytes() int32`

GetMemoryMebibytes returns the MemoryMebibytes field if non-nil, zero value otherwise.

### GetMemoryMebibytesOk

`func (o *Server) GetMemoryMebibytesOk() (*int32, bool)`

GetMemoryMebibytesOk returns a tuple with the MemoryMebibytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMebibytes

`func (o *Server) SetMemoryMebibytes(v int32)`

SetMemoryMebibytes sets MemoryMebibytes field to given value.


### GetStorage

`func (o *Server) GetStorage() []ServerStorageInner`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *Server) GetStorageOk() (*[]ServerStorageInner, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *Server) SetStorage(v []ServerStorageInner)`

SetStorage sets Storage field to given value.


### GetZone

`func (o *Server) GetZone() ServerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Server) GetZoneOk() (*ServerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Server) SetZone(v ServerZone)`

SetZone sets Zone field to given value.


### GetOptions

`func (o *Server) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Server) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Server) SetOptions(v []string)`

SetOptions sets Options field to given value.


### GetVersion

`func (o *Server) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Server) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Server) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetIpv4

`func (o *Server) GetIpv4() ServerIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *Server) GetIpv4Ok() (*ServerIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *Server) SetIpv4(v ServerIpv4)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *Server) GetIpv6() ServerIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *Server) GetIpv6Ok() (*ServerIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *Server) SetIpv6(v ServerIpv6)`

SetIpv6 sets Ipv6 field to given value.


### GetContract

`func (o *Server) GetContract() ServerContract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *Server) GetContractOk() (*ServerContract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *Server) SetContract(v ServerContract)`

SetContract sets Contract field to given value.


### GetPowerStatus

`func (o *Server) GetPowerStatus() string`

GetPowerStatus returns the PowerStatus field if non-nil, zero value otherwise.

### GetPowerStatusOk

`func (o *Server) GetPowerStatusOk() (*string, bool)`

GetPowerStatusOk returns a tuple with the PowerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerStatus

`func (o *Server) SetPowerStatus(v string)`

SetPowerStatus sets PowerStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


