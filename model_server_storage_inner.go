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

// checks if the ServerStorageInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerStorageInner{}

// ServerStorageInner struct for ServerStorageInner
type ServerStorageInner struct {
	// ポート番号
	Port int32 `json:"port"`
	// 種別
	Type string `json:"type"`
	// ストレージ容量(GiB)
	SizeGibibytes int32 `json:"size_gibibytes"`
}

type _ServerStorageInner ServerStorageInner

// NewServerStorageInner instantiates a new ServerStorageInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerStorageInner(port int32, type_ string, sizeGibibytes int32) *ServerStorageInner {
	this := ServerStorageInner{}
	this.Port = port
	this.Type = type_
	this.SizeGibibytes = sizeGibibytes
	return &this
}

// NewServerStorageInnerWithDefaults instantiates a new ServerStorageInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerStorageInnerWithDefaults() *ServerStorageInner {
	this := ServerStorageInner{}
	return &this
}

// GetPort returns the Port field value
func (o *ServerStorageInner) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ServerStorageInner) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ServerStorageInner) SetPort(v int32) {
	o.Port = v
}

// GetType returns the Type field value
func (o *ServerStorageInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServerStorageInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServerStorageInner) SetType(v string) {
	o.Type = v
}

// GetSizeGibibytes returns the SizeGibibytes field value
func (o *ServerStorageInner) GetSizeGibibytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SizeGibibytes
}

// GetSizeGibibytesOk returns a tuple with the SizeGibibytes field value
// and a boolean to check if the value has been set.
func (o *ServerStorageInner) GetSizeGibibytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeGibibytes, true
}

// SetSizeGibibytes sets field value
func (o *ServerStorageInner) SetSizeGibibytes(v int32) {
	o.SizeGibibytes = v
}

func (o ServerStorageInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerStorageInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["port"] = o.Port
	toSerialize["type"] = o.Type
	toSerialize["size_gibibytes"] = o.SizeGibibytes
	return toSerialize, nil
}

func (o *ServerStorageInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"port",
		"type",
		"size_gibibytes",
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

	varServerStorageInner := _ServerStorageInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServerStorageInner)

	if err != nil {
		return err
	}

	*o = ServerStorageInner(varServerStorageInner)

	return err
}

type NullableServerStorageInner struct {
	value *ServerStorageInner
	isSet bool
}

func (v NullableServerStorageInner) Get() *ServerStorageInner {
	return v.value
}

func (v *NullableServerStorageInner) Set(val *ServerStorageInner) {
	v.value = val
	v.isSet = true
}

func (v NullableServerStorageInner) IsSet() bool {
	return v.isSet
}

func (v *NullableServerStorageInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerStorageInner(val *ServerStorageInner) *NullableServerStorageInner {
	return &NullableServerStorageInner{value: val, isSet: true}
}

func (v NullableServerStorageInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerStorageInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


