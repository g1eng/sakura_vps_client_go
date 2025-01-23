# ServerMonitoringSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | 監視のON/OFF * true 監視ON * false 監視OFF | 
**HealthCheck** | [**ServerMonitoringSettingsHealthCheck**](ServerMonitoringSettingsHealthCheck.md) |  | 
**Notification** | [**ServerMonitoringSettingsNotification**](ServerMonitoringSettingsNotification.md) |  | 

## Methods

### NewServerMonitoringSettings

`func NewServerMonitoringSettings(enabled bool, healthCheck ServerMonitoringSettingsHealthCheck, notification ServerMonitoringSettingsNotification, ) *ServerMonitoringSettings`

NewServerMonitoringSettings instantiates a new ServerMonitoringSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerMonitoringSettingsWithDefaults

`func NewServerMonitoringSettingsWithDefaults() *ServerMonitoringSettings`

NewServerMonitoringSettingsWithDefaults instantiates a new ServerMonitoringSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ServerMonitoringSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServerMonitoringSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServerMonitoringSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHealthCheck

`func (o *ServerMonitoringSettings) GetHealthCheck() ServerMonitoringSettingsHealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *ServerMonitoringSettings) GetHealthCheckOk() (*ServerMonitoringSettingsHealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *ServerMonitoringSettings) SetHealthCheck(v ServerMonitoringSettingsHealthCheck)`

SetHealthCheck sets HealthCheck field to given value.


### GetNotification

`func (o *ServerMonitoringSettings) GetNotification() ServerMonitoringSettingsNotification`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *ServerMonitoringSettings) GetNotificationOk() (*ServerMonitoringSettingsNotification, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *ServerMonitoringSettings) SetNotification(v ServerMonitoringSettingsNotification)`

SetNotification sets Notification field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


