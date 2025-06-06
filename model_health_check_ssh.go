/*
さくらのVPS APIドキュメント

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sakuravps

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the HealthCheckSsh type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthCheckSsh{}

// HealthCheckSsh struct for HealthCheckSsh
type HealthCheckSsh struct {
	// 監視方法
	Protocol string `json:"protocol"`
	// ポート番号
	Port int32 `json:"port"`
	// チェック間隔(分)
	IntervalMinutes int32 `json:"interval_minutes"`
}

type _HealthCheckSsh HealthCheckSsh

// NewHealthCheckSsh instantiates a new HealthCheckSsh object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheckSsh(protocol string, port int32, intervalMinutes int32) *HealthCheckSsh {
	this := HealthCheckSsh{}
	this.Protocol = protocol
	this.Port = port
	this.IntervalMinutes = intervalMinutes
	return &this
}

// NewHealthCheckSshWithDefaults instantiates a new HealthCheckSsh object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckSshWithDefaults() *HealthCheckSsh {
	this := HealthCheckSsh{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *HealthCheckSsh) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *HealthCheckSsh) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *HealthCheckSsh) SetProtocol(v string) {
	o.Protocol = v
}

// GetPort returns the Port field value
func (o *HealthCheckSsh) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *HealthCheckSsh) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *HealthCheckSsh) SetPort(v int32) {
	o.Port = v
}

// GetIntervalMinutes returns the IntervalMinutes field value
func (o *HealthCheckSsh) GetIntervalMinutes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IntervalMinutes
}

// GetIntervalMinutesOk returns a tuple with the IntervalMinutes field value
// and a boolean to check if the value has been set.
func (o *HealthCheckSsh) GetIntervalMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntervalMinutes, true
}

// SetIntervalMinutes sets field value
func (o *HealthCheckSsh) SetIntervalMinutes(v int32) {
	o.IntervalMinutes = v
}

func (o HealthCheckSsh) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthCheckSsh) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["protocol"] = o.Protocol
	toSerialize["port"] = o.Port
	toSerialize["interval_minutes"] = o.IntervalMinutes
	return toSerialize, nil
}

func (o *HealthCheckSsh) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"protocol",
		"port",
		"interval_minutes",
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

	varHealthCheckSsh := _HealthCheckSsh{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHealthCheckSsh)

	if err != nil {
		return err
	}

	*o = HealthCheckSsh(varHealthCheckSsh)

	return err
}

type NullableHealthCheckSsh struct {
	value *HealthCheckSsh
	isSet bool
}

func (v NullableHealthCheckSsh) Get() *HealthCheckSsh {
	return v.value
}

func (v *NullableHealthCheckSsh) Set(val *HealthCheckSsh) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckSsh) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckSsh) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckSsh(val *HealthCheckSsh) *NullableHealthCheckSsh {
	return &NullableHealthCheckSsh{value: val, isSet: true}
}

func (v NullableHealthCheckSsh) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckSsh) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


