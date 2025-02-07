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

// checks if the PostSwitchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSwitchRequest{}

// PostSwitchRequest struct for PostSwitchRequest
type PostSwitchRequest struct {
	// 名前
	Name string `json:"name"`
	// 説明
	Description string `json:"description"`
	// ゾーンコード * tk2 東京第2 * tk3 東京第3 * os3 大阪第3 * is1 石狩第1
	ZoneCode string `json:"zone_code"`
}

type _PostSwitchRequest PostSwitchRequest

// NewPostSwitchRequest instantiates a new PostSwitchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSwitchRequest(name string, description string, zoneCode string) *PostSwitchRequest {
	this := PostSwitchRequest{}
	this.Name = name
	this.Description = description
	this.ZoneCode = zoneCode
	return &this
}

// NewPostSwitchRequestWithDefaults instantiates a new PostSwitchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSwitchRequestWithDefaults() *PostSwitchRequest {
	this := PostSwitchRequest{}
	return &this
}

// GetName returns the Name field value
func (o *PostSwitchRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostSwitchRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PostSwitchRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *PostSwitchRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PostSwitchRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PostSwitchRequest) SetDescription(v string) {
	o.Description = v
}

// GetZoneCode returns the ZoneCode field value
func (o *PostSwitchRequest) GetZoneCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZoneCode
}

// GetZoneCodeOk returns a tuple with the ZoneCode field value
// and a boolean to check if the value has been set.
func (o *PostSwitchRequest) GetZoneCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZoneCode, true
}

// SetZoneCode sets field value
func (o *PostSwitchRequest) SetZoneCode(v string) {
	o.ZoneCode = v
}

func (o PostSwitchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSwitchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["zone_code"] = o.ZoneCode
	return toSerialize, nil
}

func (o *PostSwitchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"zone_code",
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

	varPostSwitchRequest := _PostSwitchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostSwitchRequest)

	if err != nil {
		return err
	}

	*o = PostSwitchRequest(varPostSwitchRequest)

	return err
}

type NullablePostSwitchRequest struct {
	value *PostSwitchRequest
	isSet bool
}

func (v NullablePostSwitchRequest) Get() *PostSwitchRequest {
	return v.value
}

func (v *NullablePostSwitchRequest) Set(val *PostSwitchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSwitchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSwitchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSwitchRequest(val *PostSwitchRequest) *NullablePostSwitchRequest {
	return &NullablePostSwitchRequest{value: val, isSet: true}
}

func (v NullablePostSwitchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSwitchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


