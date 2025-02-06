/*
さくらのVPS APIドキュメント

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sakura_vps_client_go

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetSwitchList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSwitchList200Response{}

// GetSwitchList200Response struct for GetSwitchList200Response
type GetSwitchList200Response struct {
	// データ総数
	Count int32 `json:"count"`
	// 次のページへのURL
	Next NullableString `json:"next"`
	// 前のページへのURL
	Previous NullableString `json:"previous"`
	Results []Switch `json:"results"`
}

type _GetSwitchList200Response GetSwitchList200Response

// NewGetSwitchList200Response instantiates a new GetSwitchList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSwitchList200Response(count int32, next NullableString, previous NullableString, results []Switch) *GetSwitchList200Response {
	this := GetSwitchList200Response{}
	this.Count = count
	this.Next = next
	this.Previous = previous
	this.Results = results
	return &this
}

// NewGetSwitchList200ResponseWithDefaults instantiates a new GetSwitchList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSwitchList200ResponseWithDefaults() *GetSwitchList200Response {
	this := GetSwitchList200Response{}
	return &this
}

// GetCount returns the Count field value
func (o *GetSwitchList200Response) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetSwitchList200Response) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GetSwitchList200Response) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetSwitchList200Response) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}

	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetSwitchList200Response) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// SetNext sets field value
func (o *GetSwitchList200Response) SetNext(v string) {
	o.Next.Set(&v)
}

// GetPrevious returns the Previous field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetSwitchList200Response) GetPrevious() string {
	if o == nil || o.Previous.Get() == nil {
		var ret string
		return ret
	}

	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetSwitchList200Response) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// SetPrevious sets field value
func (o *GetSwitchList200Response) SetPrevious(v string) {
	o.Previous.Set(&v)
}

// GetResults returns the Results field value
func (o *GetSwitchList200Response) GetResults() []Switch {
	if o == nil {
		var ret []Switch
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *GetSwitchList200Response) GetResultsOk() ([]Switch, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *GetSwitchList200Response) SetResults(v []Switch) {
	o.Results = v
}

func (o GetSwitchList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSwitchList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["next"] = o.Next.Get()
	toSerialize["previous"] = o.Previous.Get()
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *GetSwitchList200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetSwitchList200Response := _GetSwitchList200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSwitchList200Response)

	if err != nil {
		return err
	}

	*o = GetSwitchList200Response(varGetSwitchList200Response)

	return err
}

type NullableGetSwitchList200Response struct {
	value *GetSwitchList200Response
	isSet bool
}

func (v NullableGetSwitchList200Response) Get() *GetSwitchList200Response {
	return v.value
}

func (v *NullableGetSwitchList200Response) Set(val *GetSwitchList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSwitchList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSwitchList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSwitchList200Response(val *GetSwitchList200Response) *NullableGetSwitchList200Response {
	return &NullableGetSwitchList200Response{value: val, isSet: true}
}

func (v NullableGetSwitchList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSwitchList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


