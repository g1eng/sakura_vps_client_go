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

// checks if the ServerIpv6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerIpv6{}

// ServerIpv6 struct for ServerIpv6
type ServerIpv6 struct {
	// アドレス
	Address NullableString `json:"address"`
	// プレフィックス長
	Prefixlen NullableInt32 `json:"prefixlen"`
	// ゲートウェイのアドレス
	Gateway NullableString `json:"gateway"`
	// ネームサーバーのアドレスリスト
	Nameservers []string `json:"nameservers"`
	// 標準ホスト名
	Hostname NullableString `json:"hostname"`
	// 逆引きホスト名
	Ptr NullableString `json:"ptr"`
}

type _ServerIpv6 ServerIpv6

// NewServerIpv6 instantiates a new ServerIpv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerIpv6(address NullableString, prefixlen NullableInt32, gateway NullableString, nameservers []string, hostname NullableString, ptr NullableString) *ServerIpv6 {
	this := ServerIpv6{}
	this.Address = address
	this.Prefixlen = prefixlen
	this.Gateway = gateway
	this.Nameservers = nameservers
	this.Hostname = hostname
	this.Ptr = ptr
	return &this
}

// NewServerIpv6WithDefaults instantiates a new ServerIpv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerIpv6WithDefaults() *ServerIpv6 {
	this := ServerIpv6{}
	return &this
}

// GetAddress returns the Address field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServerIpv6) GetAddress() string {
	if o == nil || o.Address.Get() == nil {
		var ret string
		return ret
	}

	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerIpv6) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// SetAddress sets field value
func (o *ServerIpv6) SetAddress(v string) {
	o.Address.Set(&v)
}

// GetPrefixlen returns the Prefixlen field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ServerIpv6) GetPrefixlen() int32 {
	if o == nil || o.Prefixlen.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Prefixlen.Get()
}

// GetPrefixlenOk returns a tuple with the Prefixlen field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerIpv6) GetPrefixlenOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefixlen.Get(), o.Prefixlen.IsSet()
}

// SetPrefixlen sets field value
func (o *ServerIpv6) SetPrefixlen(v int32) {
	o.Prefixlen.Set(&v)
}

// GetGateway returns the Gateway field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServerIpv6) GetGateway() string {
	if o == nil || o.Gateway.Get() == nil {
		var ret string
		return ret
	}

	return *o.Gateway.Get()
}

// GetGatewayOk returns a tuple with the Gateway field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerIpv6) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gateway.Get(), o.Gateway.IsSet()
}

// SetGateway sets field value
func (o *ServerIpv6) SetGateway(v string) {
	o.Gateway.Set(&v)
}

// GetNameservers returns the Nameservers field value
func (o *ServerIpv6) GetNameservers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Nameservers
}

// GetNameserversOk returns a tuple with the Nameservers field value
// and a boolean to check if the value has been set.
func (o *ServerIpv6) GetNameserversOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nameservers, true
}

// SetNameservers sets field value
func (o *ServerIpv6) SetNameservers(v []string) {
	o.Nameservers = v
}

// GetHostname returns the Hostname field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServerIpv6) GetHostname() string {
	if o == nil || o.Hostname.Get() == nil {
		var ret string
		return ret
	}

	return *o.Hostname.Get()
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerIpv6) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hostname.Get(), o.Hostname.IsSet()
}

// SetHostname sets field value
func (o *ServerIpv6) SetHostname(v string) {
	o.Hostname.Set(&v)
}

// GetPtr returns the Ptr field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServerIpv6) GetPtr() string {
	if o == nil || o.Ptr.Get() == nil {
		var ret string
		return ret
	}

	return *o.Ptr.Get()
}

// GetPtrOk returns a tuple with the Ptr field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerIpv6) GetPtrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ptr.Get(), o.Ptr.IsSet()
}

// SetPtr sets field value
func (o *ServerIpv6) SetPtr(v string) {
	o.Ptr.Set(&v)
}

func (o ServerIpv6) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerIpv6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address.Get()
	toSerialize["prefixlen"] = o.Prefixlen.Get()
	toSerialize["gateway"] = o.Gateway.Get()
	toSerialize["nameservers"] = o.Nameservers
	toSerialize["hostname"] = o.Hostname.Get()
	toSerialize["ptr"] = o.Ptr.Get()
	return toSerialize, nil
}

func (o *ServerIpv6) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"prefixlen",
		"gateway",
		"nameservers",
		"hostname",
		"ptr",
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

	varServerIpv6 := _ServerIpv6{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServerIpv6)

	if err != nil {
		return err
	}

	*o = ServerIpv6(varServerIpv6)

	return err
}

type NullableServerIpv6 struct {
	value *ServerIpv6
	isSet bool
}

func (v NullableServerIpv6) Get() *ServerIpv6 {
	return v.value
}

func (v *NullableServerIpv6) Set(val *ServerIpv6) {
	v.value = val
	v.isSet = true
}

func (v NullableServerIpv6) IsSet() bool {
	return v.isSet
}

func (v *NullableServerIpv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerIpv6(val *ServerIpv6) *NullableServerIpv6 {
	return &NullableServerIpv6{value: val, isSet: true}
}

func (v NullableServerIpv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerIpv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


