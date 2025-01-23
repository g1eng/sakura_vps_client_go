# ServerMonitoringSettingsNotificationIncomingWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | 通知のON/OFF * true 通知ON * false 通知OFF | 
**WebhooksUrl** | **NullableString** | 通知先のWebhooksURL。 Slack、Discord、Microsoft TeamsのIncoming WebHooksにのみ対応しています。 指定できるURLは各サービスのWebhook URL(https://hooks.slack.com/services/_* など)の形式に制限されています。 Discordの場合は[Slack互換のWebhook URL](https://discord.com/developers/docs/resources/webhook#execute-slackcompatible-webhook)を指定してください。 | 
**SlackTeamName** | **string** | slackのteam name。VPSコントロールパネルの「Slackと自動で連携をする」を利用した場合に設定されます。 | [readonly] 
**SlackChannelName** | **string** | slackのchannel name。VPSコントロールパネルの「Slackと自動で連携をする」を利用した場合に設定されます。 | [readonly] 

## Methods

### NewServerMonitoringSettingsNotificationIncomingWebhook

`func NewServerMonitoringSettingsNotificationIncomingWebhook(enabled bool, webhooksUrl NullableString, slackTeamName string, slackChannelName string, ) *ServerMonitoringSettingsNotificationIncomingWebhook`

NewServerMonitoringSettingsNotificationIncomingWebhook instantiates a new ServerMonitoringSettingsNotificationIncomingWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerMonitoringSettingsNotificationIncomingWebhookWithDefaults

`func NewServerMonitoringSettingsNotificationIncomingWebhookWithDefaults() *ServerMonitoringSettingsNotificationIncomingWebhook`

NewServerMonitoringSettingsNotificationIncomingWebhookWithDefaults instantiates a new ServerMonitoringSettingsNotificationIncomingWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ServerMonitoringSettingsNotificationIncomingWebhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServerMonitoringSettingsNotificationIncomingWebhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServerMonitoringSettingsNotificationIncomingWebhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetWebhooksUrl

`func (o *ServerMonitoringSettingsNotificationIncomingWebhook) GetWebhooksUrl() string`

GetWebhooksUrl returns the WebhooksUrl field if non-nil, zero value otherwise.

### GetWebhooksUrlOk

`func (o *ServerMonitoringSettingsNotificationIncomingWebhook) GetWebhooksUrlOk() (*string, bool)`

GetWebhooksUrlOk returns a tuple with the WebhooksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooksUrl

`func (o *ServerMonitoringSettingsNotificationIncomingWebhook) SetWebhooksUrl(v string)`

SetWebhooksUrl sets WebhooksUrl field to given value.


### SetWebhooksUrlNil

`func (o *ServerMonitoringSettingsNotificationIncomingWebhook) SetWebhooksUrlNil(b bool)`

 SetWebhooksUrlNil sets the value for WebhooksUrl to be an explicit nil

### UnsetWebhooksUrl
`func (o *ServerMonitoringSettingsNotificationIncomingWebhook) UnsetWebhooksUrl()`

UnsetWebhooksUrl ensures that no value is present for WebhooksUrl, not even an explicit nil
### GetSlackTeamName

`func (o *ServerMonitoringSettingsNotificationIncomingWebhook) GetSlackTeamName() string`

GetSlackTeamName returns the SlackTeamName field if non-nil, zero value otherwise.

### GetSlackTeamNameOk

`func (o *ServerMonitoringSettingsNotificationIncomingWebhook) GetSlackTeamNameOk() (*string, bool)`

GetSlackTeamNameOk returns a tuple with the SlackTeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackTeamName

`func (o *ServerMonitoringSettingsNotificationIncomingWebhook) SetSlackTeamName(v string)`

SetSlackTeamName sets SlackTeamName field to given value.


### GetSlackChannelName

`func (o *ServerMonitoringSettingsNotificationIncomingWebhook) GetSlackChannelName() string`

GetSlackChannelName returns the SlackChannelName field if non-nil, zero value otherwise.

### GetSlackChannelNameOk

`func (o *ServerMonitoringSettingsNotificationIncomingWebhook) GetSlackChannelNameOk() (*string, bool)`

GetSlackChannelNameOk returns a tuple with the SlackChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackChannelName

`func (o *ServerMonitoringSettingsNotificationIncomingWebhook) SetSlackChannelName(v string)`

SetSlackChannelName sets SlackChannelName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


