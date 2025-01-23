# NfsServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | id | 
**Name** | **string** | 名前 | 
**Description** | **string** | 説明 | 
**ServiceStatus** | **string** | サービス状況 * in_preparation 準備中 * on_trial お試し中 * link_down_on_trial お試し中（一時停止） * in_use 利用中 * link_down 一時停止中 | 
**SettingStatus** | **string** | 設定状況 * done 設定完了 * in_update 設定更新中 * failed 設定更新失敗 | 
**Storage** | [**[]NfsServerStorageInner**](NfsServerStorageInner.md) | ストレージ情報 | 
**Zone** | [**ServerZone**](ServerZone.md) |  | 
**Ipv4** | [**NfsServerIpv4**](NfsServerIpv4.md) |  | 
**Contract** | [**NfsServerContract**](NfsServerContract.md) |  | 
**PowerStatus** | **string** | 電源ステータス * power_on 電源ON * in_shutdown シャットダウン中 * power_off 電源OFF * unknown 不明（電源状態を取得できない） このエンドポイントが返す電源ステータスはキャッシュされた情報のため、最新の正確な電源ステータスではない場合があります | 

## Methods

### NewNfsServer

`func NewNfsServer(id int32, name string, description string, serviceStatus string, settingStatus string, storage []NfsServerStorageInner, zone ServerZone, ipv4 NfsServerIpv4, contract NfsServerContract, powerStatus string, ) *NfsServer`

NewNfsServer instantiates a new NfsServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsServerWithDefaults

`func NewNfsServerWithDefaults() *NfsServer`

NewNfsServerWithDefaults instantiates a new NfsServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NfsServer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NfsServer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NfsServer) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *NfsServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NfsServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NfsServer) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NfsServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NfsServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NfsServer) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetServiceStatus

`func (o *NfsServer) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *NfsServer) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *NfsServer) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetSettingStatus

`func (o *NfsServer) GetSettingStatus() string`

GetSettingStatus returns the SettingStatus field if non-nil, zero value otherwise.

### GetSettingStatusOk

`func (o *NfsServer) GetSettingStatusOk() (*string, bool)`

GetSettingStatusOk returns a tuple with the SettingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingStatus

`func (o *NfsServer) SetSettingStatus(v string)`

SetSettingStatus sets SettingStatus field to given value.


### GetStorage

`func (o *NfsServer) GetStorage() []NfsServerStorageInner`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *NfsServer) GetStorageOk() (*[]NfsServerStorageInner, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *NfsServer) SetStorage(v []NfsServerStorageInner)`

SetStorage sets Storage field to given value.


### GetZone

`func (o *NfsServer) GetZone() ServerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *NfsServer) GetZoneOk() (*ServerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *NfsServer) SetZone(v ServerZone)`

SetZone sets Zone field to given value.


### GetIpv4

`func (o *NfsServer) GetIpv4() NfsServerIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *NfsServer) GetIpv4Ok() (*NfsServerIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *NfsServer) SetIpv4(v NfsServerIpv4)`

SetIpv4 sets Ipv4 field to given value.


### GetContract

`func (o *NfsServer) GetContract() NfsServerContract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *NfsServer) GetContractOk() (*NfsServerContract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *NfsServer) SetContract(v NfsServerContract)`

SetContract sets Contract field to given value.


### GetPowerStatus

`func (o *NfsServer) GetPowerStatus() string`

GetPowerStatus returns the PowerStatus field if non-nil, zero value otherwise.

### GetPowerStatusOk

`func (o *NfsServer) GetPowerStatusOk() (*string, bool)`

GetPowerStatusOk returns a tuple with the PowerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerStatus

`func (o *NfsServer) SetPowerStatus(v string)`

SetPowerStatus sets PowerStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


