/*
さくらのVPS APIドキュメント

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sakura_vps_client_go

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ServerMonitoring type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerMonitoring{}

// ServerMonitoring struct for ServerMonitoring
type ServerMonitoring struct {
	// id
	Id int32 `json:"id"`
	// 名前
	Name string `json:"name"`
	// 説明
	Description string `json:"description"`
	// 監視リソースID
	MonitoringResourceId string `json:"monitoring_resource_id"`
	// 更新ステータス * waiting 更新待ち * updating 更新中 * completed 更新完了 * error 更新エラー
	UpdateStatus string `json:"update_status"`
	Settings ServerMonitoringSettings `json:"settings"`
}

type _ServerMonitoring ServerMonitoring

// NewServerMonitoring instantiates a new ServerMonitoring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerMonitoring(id int32, name string, description string, monitoringResourceId string, updateStatus string, settings ServerMonitoringSettings) *ServerMonitoring {
	this := ServerMonitoring{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.MonitoringResourceId = monitoringResourceId
	this.UpdateStatus = updateStatus
	this.Settings = settings
	return &this
}

// NewServerMonitoringWithDefaults instantiates a new ServerMonitoring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerMonitoringWithDefaults() *ServerMonitoring {
	this := ServerMonitoring{}
	return &this
}

// GetId returns the Id field value
func (o *ServerMonitoring) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServerMonitoring) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServerMonitoring) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ServerMonitoring) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServerMonitoring) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServerMonitoring) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ServerMonitoring) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServerMonitoring) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServerMonitoring) SetDescription(v string) {
	o.Description = v
}

// GetMonitoringResourceId returns the MonitoringResourceId field value
func (o *ServerMonitoring) GetMonitoringResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitoringResourceId
}

// GetMonitoringResourceIdOk returns a tuple with the MonitoringResourceId field value
// and a boolean to check if the value has been set.
func (o *ServerMonitoring) GetMonitoringResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitoringResourceId, true
}

// SetMonitoringResourceId sets field value
func (o *ServerMonitoring) SetMonitoringResourceId(v string) {
	o.MonitoringResourceId = v
}

// GetUpdateStatus returns the UpdateStatus field value
func (o *ServerMonitoring) GetUpdateStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdateStatus
}

// GetUpdateStatusOk returns a tuple with the UpdateStatus field value
// and a boolean to check if the value has been set.
func (o *ServerMonitoring) GetUpdateStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateStatus, true
}

// SetUpdateStatus sets field value
func (o *ServerMonitoring) SetUpdateStatus(v string) {
	o.UpdateStatus = v
}

// GetSettings returns the Settings field value
func (o *ServerMonitoring) GetSettings() ServerMonitoringSettings {
	if o == nil {
		var ret ServerMonitoringSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *ServerMonitoring) GetSettingsOk() (*ServerMonitoringSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *ServerMonitoring) SetSettings(v ServerMonitoringSettings) {
	o.Settings = v
}

func (o ServerMonitoring) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerMonitoring) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["monitoring_resource_id"] = o.MonitoringResourceId
	toSerialize["update_status"] = o.UpdateStatus
	toSerialize["settings"] = o.Settings
	return toSerialize, nil
}

func (o *ServerMonitoring) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
		"monitoring_resource_id",
		"update_status",
		"settings",
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

	varServerMonitoring := _ServerMonitoring{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServerMonitoring)

	if err != nil {
		return err
	}

	*o = ServerMonitoring(varServerMonitoring)

	return err
}

type NullableServerMonitoring struct {
	value *ServerMonitoring
	isSet bool
}

func (v NullableServerMonitoring) Get() *ServerMonitoring {
	return v.value
}

func (v *NullableServerMonitoring) Set(val *ServerMonitoring) {
	v.value = val
	v.isSet = true
}

func (v NullableServerMonitoring) IsSet() bool {
	return v.isSet
}

func (v *NullableServerMonitoring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerMonitoring(val *ServerMonitoring) *NullableServerMonitoring {
	return &NullableServerMonitoring{value: val, isSet: true}
}

func (v NullableServerMonitoring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerMonitoring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


