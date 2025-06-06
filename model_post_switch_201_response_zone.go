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

// checks if the PostSwitch201ResponseZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSwitch201ResponseZone{}

// PostSwitch201ResponseZone ゾーン情報
type PostSwitch201ResponseZone struct {
	// ゾーンコード * tk2 東京第2 * tk3 東京第3 * os3 大阪第3 * is1 石狩第1
	Code string `json:"code"`
	// ゾーン名称
	Name string `json:"name"`
}

type _PostSwitch201ResponseZone PostSwitch201ResponseZone

// NewPostSwitch201ResponseZone instantiates a new PostSwitch201ResponseZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSwitch201ResponseZone(code string, name string) *PostSwitch201ResponseZone {
	this := PostSwitch201ResponseZone{}
	this.Code = code
	this.Name = name
	return &this
}

// NewPostSwitch201ResponseZoneWithDefaults instantiates a new PostSwitch201ResponseZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSwitch201ResponseZoneWithDefaults() *PostSwitch201ResponseZone {
	this := PostSwitch201ResponseZone{}
	return &this
}

// GetCode returns the Code field value
func (o *PostSwitch201ResponseZone) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *PostSwitch201ResponseZone) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *PostSwitch201ResponseZone) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value
func (o *PostSwitch201ResponseZone) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostSwitch201ResponseZone) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PostSwitch201ResponseZone) SetName(v string) {
	o.Name = v
}

func (o PostSwitch201ResponseZone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSwitch201ResponseZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *PostSwitch201ResponseZone) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"name",
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

	varPostSwitch201ResponseZone := _PostSwitch201ResponseZone{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostSwitch201ResponseZone)

	if err != nil {
		return err
	}

	*o = PostSwitch201ResponseZone(varPostSwitch201ResponseZone)

	return err
}

type NullablePostSwitch201ResponseZone struct {
	value *PostSwitch201ResponseZone
	isSet bool
}

func (v NullablePostSwitch201ResponseZone) Get() *PostSwitch201ResponseZone {
	return v.value
}

func (v *NullablePostSwitch201ResponseZone) Set(val *PostSwitch201ResponseZone) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSwitch201ResponseZone) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSwitch201ResponseZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSwitch201ResponseZone(val *PostSwitch201ResponseZone) *NullablePostSwitch201ResponseZone {
	return &NullablePostSwitch201ResponseZone{value: val, isSet: true}
}

func (v NullablePostSwitch201ResponseZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSwitch201ResponseZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


