/*
さくらのVPS APIドキュメント

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sakura_vps

import (
	"encoding/json"
)

// checks if the InvalidParameterDetailInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvalidParameterDetailInner{}

// InvalidParameterDetailInner struct for InvalidParameterDetailInner
type InvalidParameterDetailInner struct {
	// エラー内容を示す簡潔な識別子
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewInvalidParameterDetailInner instantiates a new InvalidParameterDetailInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidParameterDetailInner() *InvalidParameterDetailInner {
	this := InvalidParameterDetailInner{}
	return &this
}

// NewInvalidParameterDetailInnerWithDefaults instantiates a new InvalidParameterDetailInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidParameterDetailInnerWithDefaults() *InvalidParameterDetailInner {
	this := InvalidParameterDetailInner{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InvalidParameterDetailInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidParameterDetailInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InvalidParameterDetailInner) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InvalidParameterDetailInner) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InvalidParameterDetailInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidParameterDetailInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InvalidParameterDetailInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InvalidParameterDetailInner) SetMessage(v string) {
	o.Message = &v
}

func (o InvalidParameterDetailInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvalidParameterDetailInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableInvalidParameterDetailInner struct {
	value *InvalidParameterDetailInner
	isSet bool
}

func (v NullableInvalidParameterDetailInner) Get() *InvalidParameterDetailInner {
	return v.value
}

func (v *NullableInvalidParameterDetailInner) Set(val *InvalidParameterDetailInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidParameterDetailInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidParameterDetailInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidParameterDetailInner(val *InvalidParameterDetailInner) *NullableInvalidParameterDetailInner {
	return &NullableInvalidParameterDetailInner{value: val, isSet: true}
}

func (v NullableInvalidParameterDetailInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvalidParameterDetailInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


