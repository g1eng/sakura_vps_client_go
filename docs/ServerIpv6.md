# ServerIpv6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **NullableString** | アドレス | 
**Prefixlen** | **NullableInt32** | プレフィックス長 | 
**Gateway** | **NullableString** | ゲートウェイのアドレス | 
**Nameservers** | **[]string** | ネームサーバーのアドレスリスト | 
**Hostname** | **NullableString** | 標準ホスト名 | 
**Ptr** | **NullableString** | 逆引きホスト名 | 

## Methods

### NewServerIpv6

`func NewServerIpv6(address NullableString, prefixlen NullableInt32, gateway NullableString, nameservers []string, hostname NullableString, ptr NullableString, ) *ServerIpv6`

NewServerIpv6 instantiates a new ServerIpv6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerIpv6WithDefaults

`func NewServerIpv6WithDefaults() *ServerIpv6`

NewServerIpv6WithDefaults instantiates a new ServerIpv6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ServerIpv6) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ServerIpv6) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ServerIpv6) SetAddress(v string)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *ServerIpv6) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ServerIpv6) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetPrefixlen

`func (o *ServerIpv6) GetPrefixlen() int32`

GetPrefixlen returns the Prefixlen field if non-nil, zero value otherwise.

### GetPrefixlenOk

`func (o *ServerIpv6) GetPrefixlenOk() (*int32, bool)`

GetPrefixlenOk returns a tuple with the Prefixlen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixlen

`func (o *ServerIpv6) SetPrefixlen(v int32)`

SetPrefixlen sets Prefixlen field to given value.


### SetPrefixlenNil

`func (o *ServerIpv6) SetPrefixlenNil(b bool)`

 SetPrefixlenNil sets the value for Prefixlen to be an explicit nil

### UnsetPrefixlen
`func (o *ServerIpv6) UnsetPrefixlen()`

UnsetPrefixlen ensures that no value is present for Prefixlen, not even an explicit nil
### GetGateway

`func (o *ServerIpv6) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ServerIpv6) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ServerIpv6) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### SetGatewayNil

`func (o *ServerIpv6) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *ServerIpv6) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetNameservers

`func (o *ServerIpv6) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *ServerIpv6) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *ServerIpv6) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.


### GetHostname

`func (o *ServerIpv6) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerIpv6) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerIpv6) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### SetHostnameNil

`func (o *ServerIpv6) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *ServerIpv6) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetPtr

`func (o *ServerIpv6) GetPtr() string`

GetPtr returns the Ptr field if non-nil, zero value otherwise.

### GetPtrOk

`func (o *ServerIpv6) GetPtrOk() (*string, bool)`

GetPtrOk returns a tuple with the Ptr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtr

`func (o *ServerIpv6) SetPtr(v string)`

SetPtr sets Ptr field to given value.


### SetPtrNil

`func (o *ServerIpv6) SetPtrNil(b bool)`

 SetPtrNil sets the value for Ptr to be an explicit nil

### UnsetPtr
`func (o *ServerIpv6) UnsetPtr()`

UnsetPtr ensures that no value is present for Ptr, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


