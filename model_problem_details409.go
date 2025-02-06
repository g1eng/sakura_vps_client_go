/*
さくらのVPS APIドキュメント

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sakura_vps_client_go

import (
	"encoding/json"
)

// checks if the ProblemDetails409 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProblemDetails409{}

// ProblemDetails409 struct for ProblemDetails409
type ProblemDetails409 struct {
	Code *string `json:"code,omitempty"`
	// エラーの内容
	Message *string `json:"message,omitempty"`
}

// NewProblemDetails409 instantiates a new ProblemDetails409 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblemDetails409() *ProblemDetails409 {
	this := ProblemDetails409{}
	return &this
}

// NewProblemDetails409WithDefaults instantiates a new ProblemDetails409 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemDetails409WithDefaults() *ProblemDetails409 {
	this := ProblemDetails409{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProblemDetails409) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails409) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProblemDetails409) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProblemDetails409) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ProblemDetails409) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails409) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ProblemDetails409) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ProblemDetails409) SetMessage(v string) {
	o.Message = &v
}

func (o ProblemDetails409) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProblemDetails409) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableProblemDetails409 struct {
	value *ProblemDetails409
	isSet bool
}

func (v NullableProblemDetails409) Get() *ProblemDetails409 {
	return v.value
}

func (v *NullableProblemDetails409) Set(val *ProblemDetails409) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemDetails409) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemDetails409) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemDetails409(val *ProblemDetails409) *NullableProblemDetails409 {
	return &NullableProblemDetails409{value: val, isSet: true}
}

func (v NullableProblemDetails409) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemDetails409) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


