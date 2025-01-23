/*
さくらのVPS APIドキュメント

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sakura_vps_client_go

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ServerMonitoringSettingsNotificationIncomingWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerMonitoringSettingsNotificationIncomingWebhook{}

// ServerMonitoringSettingsNotificationIncomingWebhook incoming webhookでの通知の設定
type ServerMonitoringSettingsNotificationIncomingWebhook struct {
	// 通知のON/OFF * true 通知ON * false 通知OFF
	Enabled bool `json:"enabled"`
	// 通知先のWebhooksURL。 Slack、Discord、Microsoft TeamsのIncoming WebHooksにのみ対応しています。 指定できるURLは各サービスのWebhook URL(https://hooks.slack.com/services/_* など)の形式に制限されています。 Discordの場合は[Slack互換のWebhook URL](https://discord.com/developers/docs/resources/webhook#execute-slackcompatible-webhook)を指定してください。
	WebhooksUrl NullableString `json:"webhooks_url"`
	// slackのteam name。VPSコントロールパネルの「Slackと自動で連携をする」を利用した場合に設定されます。
	SlackTeamName string `json:"slack_team_name"`
	// slackのchannel name。VPSコントロールパネルの「Slackと自動で連携をする」を利用した場合に設定されます。
	SlackChannelName string `json:"slack_channel_name"`
}

type _ServerMonitoringSettingsNotificationIncomingWebhook ServerMonitoringSettingsNotificationIncomingWebhook

// NewServerMonitoringSettingsNotificationIncomingWebhook instantiates a new ServerMonitoringSettingsNotificationIncomingWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerMonitoringSettingsNotificationIncomingWebhook(enabled bool, webhooksUrl NullableString, slackTeamName string, slackChannelName string) *ServerMonitoringSettingsNotificationIncomingWebhook {
	this := ServerMonitoringSettingsNotificationIncomingWebhook{}
	this.Enabled = enabled
	this.WebhooksUrl = webhooksUrl
	this.SlackTeamName = slackTeamName
	this.SlackChannelName = slackChannelName
	return &this
}

// NewServerMonitoringSettingsNotificationIncomingWebhookWithDefaults instantiates a new ServerMonitoringSettingsNotificationIncomingWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerMonitoringSettingsNotificationIncomingWebhookWithDefaults() *ServerMonitoringSettingsNotificationIncomingWebhook {
	this := ServerMonitoringSettingsNotificationIncomingWebhook{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *ServerMonitoringSettingsNotificationIncomingWebhook) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ServerMonitoringSettingsNotificationIncomingWebhook) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ServerMonitoringSettingsNotificationIncomingWebhook) SetEnabled(v bool) {
	o.Enabled = v
}

// GetWebhooksUrl returns the WebhooksUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServerMonitoringSettingsNotificationIncomingWebhook) GetWebhooksUrl() string {
	if o == nil || o.WebhooksUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.WebhooksUrl.Get()
}

// GetWebhooksUrlOk returns a tuple with the WebhooksUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerMonitoringSettingsNotificationIncomingWebhook) GetWebhooksUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebhooksUrl.Get(), o.WebhooksUrl.IsSet()
}

// SetWebhooksUrl sets field value
func (o *ServerMonitoringSettingsNotificationIncomingWebhook) SetWebhooksUrl(v string) {
	o.WebhooksUrl.Set(&v)
}

// GetSlackTeamName returns the SlackTeamName field value
func (o *ServerMonitoringSettingsNotificationIncomingWebhook) GetSlackTeamName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SlackTeamName
}

// GetSlackTeamNameOk returns a tuple with the SlackTeamName field value
// and a boolean to check if the value has been set.
func (o *ServerMonitoringSettingsNotificationIncomingWebhook) GetSlackTeamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlackTeamName, true
}

// SetSlackTeamName sets field value
func (o *ServerMonitoringSettingsNotificationIncomingWebhook) SetSlackTeamName(v string) {
	o.SlackTeamName = v
}

// GetSlackChannelName returns the SlackChannelName field value
func (o *ServerMonitoringSettingsNotificationIncomingWebhook) GetSlackChannelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SlackChannelName
}

// GetSlackChannelNameOk returns a tuple with the SlackChannelName field value
// and a boolean to check if the value has been set.
func (o *ServerMonitoringSettingsNotificationIncomingWebhook) GetSlackChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlackChannelName, true
}

// SetSlackChannelName sets field value
func (o *ServerMonitoringSettingsNotificationIncomingWebhook) SetSlackChannelName(v string) {
	o.SlackChannelName = v
}

func (o ServerMonitoringSettingsNotificationIncomingWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerMonitoringSettingsNotificationIncomingWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["webhooks_url"] = o.WebhooksUrl.Get()
	toSerialize["slack_team_name"] = o.SlackTeamName
	toSerialize["slack_channel_name"] = o.SlackChannelName
	return toSerialize, nil
}

func (o *ServerMonitoringSettingsNotificationIncomingWebhook) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"webhooks_url",
		"slack_team_name",
		"slack_channel_name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varServerMonitoringSettingsNotificationIncomingWebhook := _ServerMonitoringSettingsNotificationIncomingWebhook{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServerMonitoringSettingsNotificationIncomingWebhook)

	if err != nil {
		return err
	}

	*o = ServerMonitoringSettingsNotificationIncomingWebhook(varServerMonitoringSettingsNotificationIncomingWebhook)

	return err
}

type NullableServerMonitoringSettingsNotificationIncomingWebhook struct {
	value *ServerMonitoringSettingsNotificationIncomingWebhook
	isSet bool
}

func (v NullableServerMonitoringSettingsNotificationIncomingWebhook) Get() *ServerMonitoringSettingsNotificationIncomingWebhook {
	return v.value
}

func (v *NullableServerMonitoringSettingsNotificationIncomingWebhook) Set(val *ServerMonitoringSettingsNotificationIncomingWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableServerMonitoringSettingsNotificationIncomingWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableServerMonitoringSettingsNotificationIncomingWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerMonitoringSettingsNotificationIncomingWebhook(val *ServerMonitoringSettingsNotificationIncomingWebhook) *NullableServerMonitoringSettingsNotificationIncomingWebhook {
	return &NullableServerMonitoringSettingsNotificationIncomingWebhook{value: val, isSet: true}
}

func (v NullableServerMonitoringSettingsNotificationIncomingWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerMonitoringSettingsNotificationIncomingWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


