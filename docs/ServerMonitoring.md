# ServerMonitoring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | id | [readonly] 
**Name** | **string** | 名前 | 
**Description** | **string** | 説明 | 
**MonitoringResourceId** | **string** | 監視リソースID | [readonly] 
**UpdateStatus** | **string** | 更新ステータス * waiting 更新待ち * updating 更新中 * completed 更新完了 * error 更新エラー | [readonly] 
**Settings** | [**ServerMonitoringSettings**](ServerMonitoringSettings.md) |  | 

## Methods

### NewServerMonitoring

`func NewServerMonitoring(id int32, name string, description string, monitoringResourceId string, updateStatus string, settings ServerMonitoringSettings, ) *ServerMonitoring`

NewServerMonitoring instantiates a new ServerMonitoring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerMonitoringWithDefaults

`func NewServerMonitoringWithDefaults() *ServerMonitoring`

NewServerMonitoringWithDefaults instantiates a new ServerMonitoring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerMonitoring) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerMonitoring) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerMonitoring) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ServerMonitoring) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerMonitoring) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerMonitoring) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ServerMonitoring) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerMonitoring) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerMonitoring) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMonitoringResourceId

`func (o *ServerMonitoring) GetMonitoringResourceId() string`

GetMonitoringResourceId returns the MonitoringResourceId field if non-nil, zero value otherwise.

### GetMonitoringResourceIdOk

`func (o *ServerMonitoring) GetMonitoringResourceIdOk() (*string, bool)`

GetMonitoringResourceIdOk returns a tuple with the MonitoringResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringResourceId

`func (o *ServerMonitoring) SetMonitoringResourceId(v string)`

SetMonitoringResourceId sets MonitoringResourceId field to given value.


### GetUpdateStatus

`func (o *ServerMonitoring) GetUpdateStatus() string`

GetUpdateStatus returns the UpdateStatus field if non-nil, zero value otherwise.

### GetUpdateStatusOk

`func (o *ServerMonitoring) GetUpdateStatusOk() (*string, bool)`

GetUpdateStatusOk returns a tuple with the UpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStatus

`func (o *ServerMonitoring) SetUpdateStatus(v string)`

SetUpdateStatus sets UpdateStatus field to given value.


### GetSettings

`func (o *ServerMonitoring) GetSettings() ServerMonitoringSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ServerMonitoring) GetSettingsOk() (*ServerMonitoringSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ServerMonitoring) SetSettings(v ServerMonitoringSettings)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


