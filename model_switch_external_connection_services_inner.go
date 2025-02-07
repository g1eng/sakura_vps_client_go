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

// checks if the SwitchExternalConnectionServicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwitchExternalConnectionServicesInner{}

// SwitchExternalConnectionServicesInner struct for SwitchExternalConnectionServicesInner
type SwitchExternalConnectionServicesInner struct {
	// サービスカテゴリー
	ServiceCategory string `json:"service_category"`
	// サービス名
	ServiceName string `json:"service_name"`
	// スイッチコード
	SwitchCode string `json:"switch_code"`
}

type _SwitchExternalConnectionServicesInner SwitchExternalConnectionServicesInner

// NewSwitchExternalConnectionServicesInner instantiates a new SwitchExternalConnectionServicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwitchExternalConnectionServicesInner(serviceCategory string, serviceName string, switchCode string) *SwitchExternalConnectionServicesInner {
	this := SwitchExternalConnectionServicesInner{}
	this.ServiceCategory = serviceCategory
	this.ServiceName = serviceName
	this.SwitchCode = switchCode
	return &this
}

// NewSwitchExternalConnectionServicesInnerWithDefaults instantiates a new SwitchExternalConnectionServicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwitchExternalConnectionServicesInnerWithDefaults() *SwitchExternalConnectionServicesInner {
	this := SwitchExternalConnectionServicesInner{}
	return &this
}

// GetServiceCategory returns the ServiceCategory field value
func (o *SwitchExternalConnectionServicesInner) GetServiceCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceCategory
}

// GetServiceCategoryOk returns a tuple with the ServiceCategory field value
// and a boolean to check if the value has been set.
func (o *SwitchExternalConnectionServicesInner) GetServiceCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceCategory, true
}

// SetServiceCategory sets field value
func (o *SwitchExternalConnectionServicesInner) SetServiceCategory(v string) {
	o.ServiceCategory = v
}

// GetServiceName returns the ServiceName field value
func (o *SwitchExternalConnectionServicesInner) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *SwitchExternalConnectionServicesInner) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *SwitchExternalConnectionServicesInner) SetServiceName(v string) {
	o.ServiceName = v
}

// GetSwitchCode returns the SwitchCode field value
func (o *SwitchExternalConnectionServicesInner) GetSwitchCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SwitchCode
}

// GetSwitchCodeOk returns a tuple with the SwitchCode field value
// and a boolean to check if the value has been set.
func (o *SwitchExternalConnectionServicesInner) GetSwitchCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SwitchCode, true
}

// SetSwitchCode sets field value
func (o *SwitchExternalConnectionServicesInner) SetSwitchCode(v string) {
	o.SwitchCode = v
}

func (o SwitchExternalConnectionServicesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwitchExternalConnectionServicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["service_category"] = o.ServiceCategory
	toSerialize["service_name"] = o.ServiceName
	toSerialize["switch_code"] = o.SwitchCode
	return toSerialize, nil
}

func (o *SwitchExternalConnectionServicesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"service_category",
		"service_name",
		"switch_code",
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

	varSwitchExternalConnectionServicesInner := _SwitchExternalConnectionServicesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSwitchExternalConnectionServicesInner)

	if err != nil {
		return err
	}

	*o = SwitchExternalConnectionServicesInner(varSwitchExternalConnectionServicesInner)

	return err
}

type NullableSwitchExternalConnectionServicesInner struct {
	value *SwitchExternalConnectionServicesInner
	isSet bool
}

func (v NullableSwitchExternalConnectionServicesInner) Get() *SwitchExternalConnectionServicesInner {
	return v.value
}

func (v *NullableSwitchExternalConnectionServicesInner) Set(val *SwitchExternalConnectionServicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchExternalConnectionServicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchExternalConnectionServicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchExternalConnectionServicesInner(val *SwitchExternalConnectionServicesInner) *NullableSwitchExternalConnectionServicesInner {
	return &NullableSwitchExternalConnectionServicesInner{value: val, isSet: true}
}

func (v NullableSwitchExternalConnectionServicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchExternalConnectionServicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


