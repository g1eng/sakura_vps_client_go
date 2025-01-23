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

// checks if the NfsServerStorageInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NfsServerStorageInner{}

// NfsServerStorageInner struct for NfsServerStorageInner
type NfsServerStorageInner struct {
	// 種別
	Type string `json:"type"`
	// ストレージ容量(GiB)
	SizeGibibytes int32 `json:"size_gibibytes"`
}

type _NfsServerStorageInner NfsServerStorageInner

// NewNfsServerStorageInner instantiates a new NfsServerStorageInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfsServerStorageInner(type_ string, sizeGibibytes int32) *NfsServerStorageInner {
	this := NfsServerStorageInner{}
	this.Type = type_
	this.SizeGibibytes = sizeGibibytes
	return &this
}

// NewNfsServerStorageInnerWithDefaults instantiates a new NfsServerStorageInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfsServerStorageInnerWithDefaults() *NfsServerStorageInner {
	this := NfsServerStorageInner{}
	return &this
}

// GetType returns the Type field value
func (o *NfsServerStorageInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NfsServerStorageInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NfsServerStorageInner) SetType(v string) {
	o.Type = v
}

// GetSizeGibibytes returns the SizeGibibytes field value
func (o *NfsServerStorageInner) GetSizeGibibytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SizeGibibytes
}

// GetSizeGibibytesOk returns a tuple with the SizeGibibytes field value
// and a boolean to check if the value has been set.
func (o *NfsServerStorageInner) GetSizeGibibytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeGibibytes, true
}

// SetSizeGibibytes sets field value
func (o *NfsServerStorageInner) SetSizeGibibytes(v int32) {
	o.SizeGibibytes = v
}

func (o NfsServerStorageInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NfsServerStorageInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["size_gibibytes"] = o.SizeGibibytes
	return toSerialize, nil
}

func (o *NfsServerStorageInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varNfsServerStorageInner := _NfsServerStorageInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNfsServerStorageInner)

	if err != nil {
		return err
	}

	*o = NfsServerStorageInner(varNfsServerStorageInner)

	return err
}

type NullableNfsServerStorageInner struct {
	value *NfsServerStorageInner
	isSet bool
}

func (v NullableNfsServerStorageInner) Get() *NfsServerStorageInner {
	return v.value
}

func (v *NullableNfsServerStorageInner) Set(val *NfsServerStorageInner) {
	v.value = val
	v.isSet = true
}

func (v NullableNfsServerStorageInner) IsSet() bool {
	return v.isSet
}

func (v *NullableNfsServerStorageInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfsServerStorageInner(val *NfsServerStorageInner) *NullableNfsServerStorageInner {
	return &NullableNfsServerStorageInner{value: val, isSet: true}
}

func (v NullableNfsServerStorageInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfsServerStorageInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


