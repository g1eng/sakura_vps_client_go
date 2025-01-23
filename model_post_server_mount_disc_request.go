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

// checks if the PostServerMountDiscRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostServerMountDiscRequest{}

// PostServerMountDiscRequest struct for PostServerMountDiscRequest
type PostServerMountDiscRequest struct {
	// ディスクID
	DiscId int32 `json:"disc_id"`
}

type _PostServerMountDiscRequest PostServerMountDiscRequest

// NewPostServerMountDiscRequest instantiates a new PostServerMountDiscRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostServerMountDiscRequest(discId int32) *PostServerMountDiscRequest {
	this := PostServerMountDiscRequest{}
	this.DiscId = discId
	return &this
}

// NewPostServerMountDiscRequestWithDefaults instantiates a new PostServerMountDiscRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostServerMountDiscRequestWithDefaults() *PostServerMountDiscRequest {
	this := PostServerMountDiscRequest{}
	return &this
}

// GetDiscId returns the DiscId field value
func (o *PostServerMountDiscRequest) GetDiscId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DiscId
}

// GetDiscIdOk returns a tuple with the DiscId field value
// and a boolean to check if the value has been set.
func (o *PostServerMountDiscRequest) GetDiscIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscId, true
}

// SetDiscId sets field value
func (o *PostServerMountDiscRequest) SetDiscId(v int32) {
	o.DiscId = v
}

func (o PostServerMountDiscRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostServerMountDiscRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["disc_id"] = o.DiscId
	return toSerialize, nil
}

func (o *PostServerMountDiscRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"disc_id",
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

	varPostServerMountDiscRequest := _PostServerMountDiscRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostServerMountDiscRequest)

	if err != nil {
		return err
	}

	*o = PostServerMountDiscRequest(varPostServerMountDiscRequest)

	return err
}

type NullablePostServerMountDiscRequest struct {
	value *PostServerMountDiscRequest
	isSet bool
}

func (v NullablePostServerMountDiscRequest) Get() *PostServerMountDiscRequest {
	return v.value
}

func (v *NullablePostServerMountDiscRequest) Set(val *PostServerMountDiscRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostServerMountDiscRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostServerMountDiscRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostServerMountDiscRequest(val *PostServerMountDiscRequest) *NullablePostServerMountDiscRequest {
	return &NullablePostServerMountDiscRequest{value: val, isSet: true}
}

func (v NullablePostServerMountDiscRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostServerMountDiscRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


