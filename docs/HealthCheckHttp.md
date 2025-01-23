# HealthCheckHttp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **int32** | ポート番号 | 
**Host** | **NullableString** | 監視用HTTPリクエストのHostヘッダ   RFCの定義に基づいて下記の制限をかけています * ラベルは2つ以上必要 * 各ラベルについて   * 利用できる文字     * 半角数字 0～9     * 半角英小文字 a～z     * 半角記号 -   * 両端が-でないこと   * -が2つ以上続かないこと。ただしpunycodeの接頭辞&#x60;xn--&#x60;を除く   * 下記パターン(RFCなどで予約または禁止されているパターン)と一致しないこと     * isatap     * wpad     * example     * example0～example9 * 最後のラベルについて   * 利用できる文字     * 半角英小文字 a～z   * 下記のパターン(RFC予約済みのDNS名)と一致しないこと     * test     * localhost | 
**Path** | **string** | 監視対象のパス * 利用できる文字    * 半角数字 0～9   * 半角英字 A～Z、a～z   * 半角記号 _./~%?&#x3D;-&amp; | 
**BasicAuthUsername** | **NullableString** | ベーシック認証のユーザー名 * 利用できる文字    * 半角数字 0～9   * 半角英字 A～Z、a～z   * 半角記号 _.-+!@ | 
**BasicAuthPassword** | **NullableString** | ベーシック認証のパスワード * 利用できる文字    * 半角数字 0～9   * 半角英字 A～Z、a～z   * 半角記号 !#$%&amp;()*+,-./:&lt;&#x3D;&gt;?@[]^_&#x60;{|}~ | 
**Status** | **int32** | 正常と見なすHTTPステータスコード | 
**IntervalMinutes** | **int32** | チェック間隔(分) | 
**Protocol** | **string** | 監視方法 | 

## Methods

### NewHealthCheckHttp

`func NewHealthCheckHttp(port int32, host NullableString, path string, basicAuthUsername NullableString, basicAuthPassword NullableString, status int32, intervalMinutes int32, protocol string, ) *HealthCheckHttp`

NewHealthCheckHttp instantiates a new HealthCheckHttp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckHttpWithDefaults

`func NewHealthCheckHttpWithDefaults() *HealthCheckHttp`

NewHealthCheckHttpWithDefaults instantiates a new HealthCheckHttp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *HealthCheckHttp) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HealthCheckHttp) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HealthCheckHttp) SetPort(v int32)`

SetPort sets Port field to given value.


### GetHost

`func (o *HealthCheckHttp) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HealthCheckHttp) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HealthCheckHttp) SetHost(v string)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *HealthCheckHttp) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *HealthCheckHttp) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPath

`func (o *HealthCheckHttp) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HealthCheckHttp) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HealthCheckHttp) SetPath(v string)`

SetPath sets Path field to given value.


### GetBasicAuthUsername

`func (o *HealthCheckHttp) GetBasicAuthUsername() string`

GetBasicAuthUsername returns the BasicAuthUsername field if non-nil, zero value otherwise.

### GetBasicAuthUsernameOk

`func (o *HealthCheckHttp) GetBasicAuthUsernameOk() (*string, bool)`

GetBasicAuthUsernameOk returns a tuple with the BasicAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUsername

`func (o *HealthCheckHttp) SetBasicAuthUsername(v string)`

SetBasicAuthUsername sets BasicAuthUsername field to given value.


### SetBasicAuthUsernameNil

`func (o *HealthCheckHttp) SetBasicAuthUsernameNil(b bool)`

 SetBasicAuthUsernameNil sets the value for BasicAuthUsername to be an explicit nil

### UnsetBasicAuthUsername
`func (o *HealthCheckHttp) UnsetBasicAuthUsername()`

UnsetBasicAuthUsername ensures that no value is present for BasicAuthUsername, not even an explicit nil
### GetBasicAuthPassword

`func (o *HealthCheckHttp) GetBasicAuthPassword() string`

GetBasicAuthPassword returns the BasicAuthPassword field if non-nil, zero value otherwise.

### GetBasicAuthPasswordOk

`func (o *HealthCheckHttp) GetBasicAuthPasswordOk() (*string, bool)`

GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPassword

`func (o *HealthCheckHttp) SetBasicAuthPassword(v string)`

SetBasicAuthPassword sets BasicAuthPassword field to given value.


### SetBasicAuthPasswordNil

`func (o *HealthCheckHttp) SetBasicAuthPasswordNil(b bool)`

 SetBasicAuthPasswordNil sets the value for BasicAuthPassword to be an explicit nil

### UnsetBasicAuthPassword
`func (o *HealthCheckHttp) UnsetBasicAuthPassword()`

UnsetBasicAuthPassword ensures that no value is present for BasicAuthPassword, not even an explicit nil
### GetStatus

`func (o *HealthCheckHttp) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthCheckHttp) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthCheckHttp) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetIntervalMinutes

`func (o *HealthCheckHttp) GetIntervalMinutes() int32`

GetIntervalMinutes returns the IntervalMinutes field if non-nil, zero value otherwise.

### GetIntervalMinutesOk

`func (o *HealthCheckHttp) GetIntervalMinutesOk() (*int32, bool)`

GetIntervalMinutesOk returns a tuple with the IntervalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMinutes

`func (o *HealthCheckHttp) SetIntervalMinutes(v int32)`

SetIntervalMinutes sets IntervalMinutes field to given value.


### GetProtocol

`func (o *HealthCheckHttp) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HealthCheckHttp) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HealthCheckHttp) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


