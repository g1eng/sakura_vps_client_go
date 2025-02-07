/*
さくらのVPS APIドキュメント

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sakura_vps

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the NfsServerPowerStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NfsServerPowerStatus{}

// NfsServerPowerStatus struct for NfsServerPowerStatus
type NfsServerPowerStatus struct {
	// 電源ステータス * power_on 電源ON * in_shutdown シャットダウン中 * power_off 電源OFF * unknown 不明（電源状態を取得できない）
	Status string `json:"status"`
}

type _NfsServerPowerStatus NfsServerPowerStatus

// NewNfsServerPowerStatus instantiates a new NfsServerPowerStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfsServerPowerStatus(status string) *NfsServerPowerStatus {
	this := NfsServerPowerStatus{}
	this.Status = status
	return &this
}

// NewNfsServerPowerStatusWithDefaults instantiates a new NfsServerPowerStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfsServerPowerStatusWithDefaults() *NfsServerPowerStatus {
	this := NfsServerPowerStatus{}
	return &this
}

// GetStatus returns the Status field value
func (o *NfsServerPowerStatus) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *NfsServerPowerStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *NfsServerPowerStatus) SetStatus(v string) {
	o.Status = v
}

func (o NfsServerPowerStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NfsServerPowerStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *NfsServerPowerStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varNfsServerPowerStatus := _NfsServerPowerStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNfsServerPowerStatus)

	if err != nil {
		return err
	}

	*o = NfsServerPowerStatus(varNfsServerPowerStatus)

	return err
}

type NullableNfsServerPowerStatus struct {
	value *NfsServerPowerStatus
	isSet bool
}

func (v NullableNfsServerPowerStatus) Get() *NfsServerPowerStatus {
	return v.value
}

func (v *NullableNfsServerPowerStatus) Set(val *NfsServerPowerStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNfsServerPowerStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNfsServerPowerStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfsServerPowerStatus(val *NfsServerPowerStatus) *NullableNfsServerPowerStatus {
	return &NullableNfsServerPowerStatus{value: val, isSet: true}
}

func (v NullableNfsServerPowerStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfsServerPowerStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


