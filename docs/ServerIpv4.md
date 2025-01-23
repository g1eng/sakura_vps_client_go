# ServerIpv4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | アドレス | 
**Netmask** | **string** | サブネットマスク | 
**Gateway** | **string** | ゲートウェイのアドレス | 
**Nameservers** | **[]string** | ネームサーバーのアドレスリスト | 
**Hostname** | **string** | 標準ホスト名 | 
**Ptr** | **string** | 逆引きホスト名 | 

## Methods

### NewServerIpv4

`func NewServerIpv4(address string, netmask string, gateway string, nameservers []string, hostname string, ptr string, ) *ServerIpv4`

NewServerIpv4 instantiates a new ServerIpv4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerIpv4WithDefaults

`func NewServerIpv4WithDefaults() *ServerIpv4`

NewServerIpv4WithDefaults instantiates a new ServerIpv4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ServerIpv4) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ServerIpv4) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ServerIpv4) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetNetmask

`func (o *ServerIpv4) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *ServerIpv4) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *ServerIpv4) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.


### GetGateway

`func (o *ServerIpv4) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ServerIpv4) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ServerIpv4) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetNameservers

`func (o *ServerIpv4) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *ServerIpv4) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *ServerIpv4) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.


### GetHostname

`func (o *ServerIpv4) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerIpv4) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerIpv4) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPtr

`func (o *ServerIpv4) GetPtr() string`

GetPtr returns the Ptr field if non-nil, zero value otherwise.

### GetPtrOk

`func (o *ServerIpv4) GetPtrOk() (*string, bool)`

GetPtrOk returns a tuple with the Ptr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtr

`func (o *ServerIpv4) SetPtr(v string)`

SetPtr sets Ptr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


