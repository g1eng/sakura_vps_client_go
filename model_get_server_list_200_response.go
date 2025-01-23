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

// checks if the GetServerList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetServerList200Response{}

// GetServerList200Response struct for GetServerList200Response
type GetServerList200Response struct {
	// データ総数
	Count int32 `json:"count"`
	// 次のページへのURL
	Next NullableString `json:"next"`
	// 前のページへのURL
	Previous NullableString `json:"previous"`
	Results []Server `json:"results"`
}

type _GetServerList200Response GetServerList200Response

// NewGetServerList200Response instantiates a new GetServerList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServerList200Response(count int32, next NullableString, previous NullableString, results []Server) *GetServerList200Response {
	this := GetServerList200Response{}
	this.Count = count
	this.Next = next
	this.Previous = previous
	this.Results = results
	return &this
}

// NewGetServerList200ResponseWithDefaults instantiates a new GetServerList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServerList200ResponseWithDefaults() *GetServerList200Response {
	this := GetServerList200Response{}
	return &this
}

// GetCount returns the Count field value
func (o *GetServerList200Response) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetServerList200Response) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GetServerList200Response) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetServerList200Response) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}

	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetServerList200Response) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// SetNext sets field value
func (o *GetServerList200Response) SetNext(v string) {
	o.Next.Set(&v)
}

// GetPrevious returns the Previous field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetServerList200Response) GetPrevious() string {
	if o == nil || o.Previous.Get() == nil {
		var ret string
		return ret
	}

	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetServerList200Response) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// SetPrevious sets field value
func (o *GetServerList200Response) SetPrevious(v string) {
	o.Previous.Set(&v)
}

// GetResults returns the Results field value
func (o *GetServerList200Response) GetResults() []Server {
	if o == nil {
		var ret []Server
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *GetServerList200Response) GetResultsOk() ([]Server, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *GetServerList200Response) SetResults(v []Server) {
	o.Results = v
}

func (o GetServerList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetServerList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["next"] = o.Next.Get()
	toSerialize["previous"] = o.Previous.Get()
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *GetServerList200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"next",
		"previous",
		"results",
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

	varGetServerList200Response := _GetServerList200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetServerList200Response)

	if err != nil {
		return err
	}

	*o = GetServerList200Response(varGetServerList200Response)

	return err
}

type NullableGetServerList200Response struct {
	value *GetServerList200Response
	isSet bool
}

func (v NullableGetServerList200Response) Get() *GetServerList200Response {
	return v.value
}

func (v *NullableGetServerList200Response) Set(val *GetServerList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServerList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServerList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServerList200Response(val *GetServerList200Response) *NullableGetServerList200Response {
	return &NullableGetServerList200Response{value: val, isSet: true}
}

func (v NullableGetServerList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServerList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


