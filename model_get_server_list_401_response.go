/*
さくらのVPS APIドキュメント

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sakuravps

import (
	"encoding/json"
)

// checks if the GetServerList401Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetServerList401Response{}

// GetServerList401Response struct for GetServerList401Response
type GetServerList401Response struct {
	// エラー内容を示す簡潔な識別子
	Code *string `json:"code,omitempty"`
	// エラーの内容
	Message *string `json:"message,omitempty"`
}

// NewGetServerList401Response instantiates a new GetServerList401Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServerList401Response() *GetServerList401Response {
	this := GetServerList401Response{}
	return &this
}

// NewGetServerList401ResponseWithDefaults instantiates a new GetServerList401Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServerList401ResponseWithDefaults() *GetServerList401Response {
	this := GetServerList401Response{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetServerList401Response) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerList401Response) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetServerList401Response) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetServerList401Response) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetServerList401Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerList401Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetServerList401Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetServerList401Response) SetMessage(v string) {
	o.Message = &v
}

func (o GetServerList401Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetServerList401Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableGetServerList401Response struct {
	value *GetServerList401Response
	isSet bool
}

func (v NullableGetServerList401Response) Get() *GetServerList401Response {
	return v.value
}

func (v *NullableGetServerList401Response) Set(val *GetServerList401Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServerList401Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServerList401Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServerList401Response(val *GetServerList401Response) *NullableGetServerList401Response {
	return &NullableGetServerList401Response{value: val, isSet: true}
}

func (v NullableGetServerList401Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServerList401Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


