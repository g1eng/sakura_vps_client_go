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

// checks if the ServerMonitoringSettingsNotificationEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerMonitoringSettingsNotificationEmail{}

// ServerMonitoringSettingsNotificationEmail emailでの通知の設定。会員情報に登録されているメールアドレス宛に送信されます。
type ServerMonitoringSettingsNotificationEmail struct {
	// 通知のON/OFF * true 通知ON * false 通知OFF
	Enabled bool `json:"enabled"`
}

type _ServerMonitoringSettingsNotificationEmail ServerMonitoringSettingsNotificationEmail

// NewServerMonitoringSettingsNotificationEmail instantiates a new ServerMonitoringSettingsNotificationEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerMonitoringSettingsNotificationEmail(enabled bool) *ServerMonitoringSettingsNotificationEmail {
	this := ServerMonitoringSettingsNotificationEmail{}
	this.Enabled = enabled
	return &this
}

// NewServerMonitoringSettingsNotificationEmailWithDefaults instantiates a new ServerMonitoringSettingsNotificationEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerMonitoringSettingsNotificationEmailWithDefaults() *ServerMonitoringSettingsNotificationEmail {
	this := ServerMonitoringSettingsNotificationEmail{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *ServerMonitoringSettingsNotificationEmail) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ServerMonitoringSettingsNotificationEmail) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ServerMonitoringSettingsNotificationEmail) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ServerMonitoringSettingsNotificationEmail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerMonitoringSettingsNotificationEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

func (o *ServerMonitoringSettingsNotificationEmail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
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

	varServerMonitoringSettingsNotificationEmail := _ServerMonitoringSettingsNotificationEmail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServerMonitoringSettingsNotificationEmail)

	if err != nil {
		return err
	}

	*o = ServerMonitoringSettingsNotificationEmail(varServerMonitoringSettingsNotificationEmail)

	return err
}

type NullableServerMonitoringSettingsNotificationEmail struct {
	value *ServerMonitoringSettingsNotificationEmail
	isSet bool
}

func (v NullableServerMonitoringSettingsNotificationEmail) Get() *ServerMonitoringSettingsNotificationEmail {
	return v.value
}

func (v *NullableServerMonitoringSettingsNotificationEmail) Set(val *ServerMonitoringSettingsNotificationEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableServerMonitoringSettingsNotificationEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableServerMonitoringSettingsNotificationEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerMonitoringSettingsNotificationEmail(val *ServerMonitoringSettingsNotificationEmail) *NullableServerMonitoringSettingsNotificationEmail {
	return &NullableServerMonitoringSettingsNotificationEmail{value: val, isSet: true}
}

func (v NullableServerMonitoringSettingsNotificationEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerMonitoringSettingsNotificationEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


