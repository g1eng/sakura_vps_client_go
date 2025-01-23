/*
さくらのVPS APIドキュメント

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sakura_vps_client_go

import (
	"encoding/json"
)

// checks if the ProblemDetails429 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProblemDetails429{}

// ProblemDetails429 struct for ProblemDetails429
type ProblemDetails429 struct {
	Code *string `json:"code,omitempty"`
	// エラーの内容
	Message *string `json:"message,omitempty"`
}

// NewProblemDetails429 instantiates a new ProblemDetails429 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblemDetails429() *ProblemDetails429 {
	this := ProblemDetails429{}
	return &this
}

// NewProblemDetails429WithDefaults instantiates a new ProblemDetails429 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemDetails429WithDefaults() *ProblemDetails429 {
	this := ProblemDetails429{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProblemDetails429) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails429) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProblemDetails429) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProblemDetails429) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ProblemDetails429) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails429) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ProblemDetails429) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ProblemDetails429) SetMessage(v string) {
	o.Message = &v
}

func (o ProblemDetails429) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProblemDetails429) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableProblemDetails429 struct {
	value *ProblemDetails429
	isSet bool
}

func (v NullableProblemDetails429) Get() *ProblemDetails429 {
	return v.value
}

func (v *NullableProblemDetails429) Set(val *ProblemDetails429) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemDetails429) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemDetails429) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemDetails429(val *ProblemDetails429) *NullableProblemDetails429 {
	return &NullableProblemDetails429{value: val, isSet: true}
}

func (v NullableProblemDetails429) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemDetails429) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


