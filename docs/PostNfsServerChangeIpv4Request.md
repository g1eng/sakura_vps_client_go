# PostNfsServerChangeIpv4Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | ipv4アドレス | 
**Netmask** | **string** | ネットマスク | 

## Methods

### NewPostNfsServerChangeIpv4Request

`func NewPostNfsServerChangeIpv4Request(address string, netmask string, ) *PostNfsServerChangeIpv4Request`

NewPostNfsServerChangeIpv4Request instantiates a new PostNfsServerChangeIpv4Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostNfsServerChangeIpv4RequestWithDefaults

`func NewPostNfsServerChangeIpv4RequestWithDefaults() *PostNfsServerChangeIpv4Request`

NewPostNfsServerChangeIpv4RequestWithDefaults instantiates a new PostNfsServerChangeIpv4Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PostNfsServerChangeIpv4Request) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PostNfsServerChangeIpv4Request) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PostNfsServerChangeIpv4Request) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetNetmask

`func (o *PostNfsServerChangeIpv4Request) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *PostNfsServerChangeIpv4Request) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *PostNfsServerChangeIpv4Request) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


