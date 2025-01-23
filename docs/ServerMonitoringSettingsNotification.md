# ServerMonitoringSettingsNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | [**ServerMonitoringSettingsNotificationEmail**](ServerMonitoringSettingsNotificationEmail.md) |  | 
**IncomingWebhook** | [**ServerMonitoringSettingsNotificationIncomingWebhook**](ServerMonitoringSettingsNotificationIncomingWebhook.md) |  | 
**IntervalHours** | **int32** | 再通知間隔(時間) | 

## Methods

### NewServerMonitoringSettingsNotification

`func NewServerMonitoringSettingsNotification(email ServerMonitoringSettingsNotificationEmail, incomingWebhook ServerMonitoringSettingsNotificationIncomingWebhook, intervalHours int32, ) *ServerMonitoringSettingsNotification`

NewServerMonitoringSettingsNotification instantiates a new ServerMonitoringSettingsNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerMonitoringSettingsNotificationWithDefaults

`func NewServerMonitoringSettingsNotificationWithDefaults() *ServerMonitoringSettingsNotification`

NewServerMonitoringSettingsNotificationWithDefaults instantiates a new ServerMonitoringSettingsNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ServerMonitoringSettingsNotification) GetEmail() ServerMonitoringSettingsNotificationEmail`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ServerMonitoringSettingsNotification) GetEmailOk() (*ServerMonitoringSettingsNotificationEmail, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ServerMonitoringSettingsNotification) SetEmail(v ServerMonitoringSettingsNotificationEmail)`

SetEmail sets Email field to given value.


### GetIncomingWebhook

`func (o *ServerMonitoringSettingsNotification) GetIncomingWebhook() ServerMonitoringSettingsNotificationIncomingWebhook`

GetIncomingWebhook returns the IncomingWebhook field if non-nil, zero value otherwise.

### GetIncomingWebhookOk

`func (o *ServerMonitoringSettingsNotification) GetIncomingWebhookOk() (*ServerMonitoringSettingsNotificationIncomingWebhook, bool)`

GetIncomingWebhookOk returns a tuple with the IncomingWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingWebhook

`func (o *ServerMonitoringSettingsNotification) SetIncomingWebhook(v ServerMonitoringSettingsNotificationIncomingWebhook)`

SetIncomingWebhook sets IncomingWebhook field to given value.


### GetIntervalHours

`func (o *ServerMonitoringSettingsNotification) GetIntervalHours() int32`

GetIntervalHours returns the IntervalHours field if non-nil, zero value otherwise.

### GetIntervalHoursOk

`func (o *ServerMonitoringSettingsNotification) GetIntervalHoursOk() (*int32, bool)`

GetIntervalHoursOk returns a tuple with the IntervalHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalHours

`func (o *ServerMonitoringSettingsNotification) SetIntervalHours(v int32)`

SetIntervalHours sets IntervalHours field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


