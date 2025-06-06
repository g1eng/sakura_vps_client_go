/*
さくらのVPS APIドキュメント

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sakuravps

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the HealthCheckHttp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthCheckHttp{}

// HealthCheckHttp struct for HealthCheckHttp
type HealthCheckHttp struct {
	// ポート番号
	Port int32 `json:"port"`
	// 監視用HTTPリクエストのHostヘッダ   RFCの定義に基づいて下記の制限をかけています * ラベルは2つ以上必要 * 各ラベルについて   * 利用できる文字     * 半角数字 0～9     * 半角英小文字 a～z     * 半角記号 -   * 両端が-でないこと   * -が2つ以上続かないこと。ただしpunycodeの接頭辞`xn--`を除く   * 下記パターン(RFCなどで予約または禁止されているパターン)と一致しないこと     * isatap     * wpad     * example     * example0～example9 * 最後のラベルについて   * 利用できる文字     * 半角英小文字 a～z   * 下記のパターン(RFC予約済みのDNS名)と一致しないこと     * test     * localhost
	Host NullableString `json:"host"`
	// 監視対象のパス * 利用できる文字    * 半角数字 0～9   * 半角英字 A～Z、a～z   * 半角記号 _./~%?=-&
	Path string `json:"path"`
	// ベーシック認証のユーザー名 * 利用できる文字    * 半角数字 0～9   * 半角英字 A～Z、a～z   * 半角記号 _.-+!@
	BasicAuthUsername NullableString `json:"basic_auth_username"`
	// ベーシック認証のパスワード * 利用できる文字    * 半角数字 0～9   * 半角英字 A～Z、a～z   * 半角記号 !#$%&()*+,-./:<=>?@[]^_`{|}~
	BasicAuthPassword NullableString `json:"basic_auth_password"`
	// 正常と見なすHTTPステータスコード
	Status int32 `json:"status"`
	// チェック間隔(分)
	IntervalMinutes int32 `json:"interval_minutes"`
	// 監視方法
	Protocol string `json:"protocol"`
}

type _HealthCheckHttp HealthCheckHttp

// NewHealthCheckHttp instantiates a new HealthCheckHttp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheckHttp(port int32, host NullableString, path string, basicAuthUsername NullableString, basicAuthPassword NullableString, status int32, intervalMinutes int32, protocol string) *HealthCheckHttp {
	this := HealthCheckHttp{}
	this.Port = port
	this.Host = host
	this.Path = path
	this.BasicAuthUsername = basicAuthUsername
	this.BasicAuthPassword = basicAuthPassword
	this.Status = status
	this.IntervalMinutes = intervalMinutes
	this.Protocol = protocol
	return &this
}

// NewHealthCheckHttpWithDefaults instantiates a new HealthCheckHttp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckHttpWithDefaults() *HealthCheckHttp {
	this := HealthCheckHttp{}
	return &this
}

// GetPort returns the Port field value
func (o *HealthCheckHttp) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *HealthCheckHttp) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *HealthCheckHttp) SetPort(v int32) {
	o.Port = v
}

// GetHost returns the Host field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HealthCheckHttp) GetHost() string {
	if o == nil || o.Host.Get() == nil {
		var ret string
		return ret
	}

	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HealthCheckHttp) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// SetHost sets field value
func (o *HealthCheckHttp) SetHost(v string) {
	o.Host.Set(&v)
}

// GetPath returns the Path field value
func (o *HealthCheckHttp) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *HealthCheckHttp) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *HealthCheckHttp) SetPath(v string) {
	o.Path = v
}

// GetBasicAuthUsername returns the BasicAuthUsername field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HealthCheckHttp) GetBasicAuthUsername() string {
	if o == nil || o.BasicAuthUsername.Get() == nil {
		var ret string
		return ret
	}

	return *o.BasicAuthUsername.Get()
}

// GetBasicAuthUsernameOk returns a tuple with the BasicAuthUsername field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HealthCheckHttp) GetBasicAuthUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BasicAuthUsername.Get(), o.BasicAuthUsername.IsSet()
}

// SetBasicAuthUsername sets field value
func (o *HealthCheckHttp) SetBasicAuthUsername(v string) {
	o.BasicAuthUsername.Set(&v)
}

// GetBasicAuthPassword returns the BasicAuthPassword field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HealthCheckHttp) GetBasicAuthPassword() string {
	if o == nil || o.BasicAuthPassword.Get() == nil {
		var ret string
		return ret
	}

	return *o.BasicAuthPassword.Get()
}

// GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HealthCheckHttp) GetBasicAuthPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BasicAuthPassword.Get(), o.BasicAuthPassword.IsSet()
}

// SetBasicAuthPassword sets field value
func (o *HealthCheckHttp) SetBasicAuthPassword(v string) {
	o.BasicAuthPassword.Set(&v)
}

// GetStatus returns the Status field value
func (o *HealthCheckHttp) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *HealthCheckHttp) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *HealthCheckHttp) SetStatus(v int32) {
	o.Status = v
}

// GetIntervalMinutes returns the IntervalMinutes field value
func (o *HealthCheckHttp) GetIntervalMinutes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IntervalMinutes
}

// GetIntervalMinutesOk returns a tuple with the IntervalMinutes field value
// and a boolean to check if the value has been set.
func (o *HealthCheckHttp) GetIntervalMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntervalMinutes, true
}

// SetIntervalMinutes sets field value
func (o *HealthCheckHttp) SetIntervalMinutes(v int32) {
	o.IntervalMinutes = v
}

// GetProtocol returns the Protocol field value
func (o *HealthCheckHttp) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *HealthCheckHttp) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *HealthCheckHttp) SetProtocol(v string) {
	o.Protocol = v
}

func (o HealthCheckHttp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthCheckHttp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["port"] = o.Port
	toSerialize["host"] = o.Host.Get()
	toSerialize["path"] = o.Path
	toSerialize["basic_auth_username"] = o.BasicAuthUsername.Get()
	toSerialize["basic_auth_password"] = o.BasicAuthPassword.Get()
	toSerialize["status"] = o.Status
	toSerialize["interval_minutes"] = o.IntervalMinutes
	toSerialize["protocol"] = o.Protocol
	return toSerialize, nil
}

func (o *HealthCheckHttp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"port",
		"host",
		"path",
		"basic_auth_username",
		"basic_auth_password",
		"status",
		"interval_minutes",
		"protocol",
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

	varHealthCheckHttp := _HealthCheckHttp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHealthCheckHttp)

	if err != nil {
		return err
	}

	*o = HealthCheckHttp(varHealthCheckHttp)

	return err
}

type NullableHealthCheckHttp struct {
	value *HealthCheckHttp
	isSet bool
}

func (v NullableHealthCheckHttp) Get() *HealthCheckHttp {
	return v.value
}

func (v *NullableHealthCheckHttp) Set(val *HealthCheckHttp) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckHttp) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckHttp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckHttp(val *HealthCheckHttp) *NullableHealthCheckHttp {
	return &NullableHealthCheckHttp{value: val, isSet: true}
}

func (v NullableHealthCheckHttp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckHttp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


