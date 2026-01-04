# PacketFilterRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** |  | 
**Name** | **string** |  | 
**Ports** | **[]string** | アクセスを許可するポートのリスト | 
**SourceIp** | **string** | アクセスを許可するipアドレス。ipアドレスで制限しない場合は\&quot;0.0.0.0\&quot; | 
**SourceIpPrefixLength** | **int32** | source_ipのプレフィックス長。ipアドレスで制限しない場合は0 | 

## Methods

### NewPacketFilterRule

`func NewPacketFilterRule(protocol string, name string, ports []string, sourceIp string, sourceIpPrefixLength int32, ) *PacketFilterRule`

NewPacketFilterRule instantiates a new PacketFilterRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPacketFilterRuleWithDefaults

`func NewPacketFilterRuleWithDefaults() *PacketFilterRule`

NewPacketFilterRuleWithDefaults instantiates a new PacketFilterRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *PacketFilterRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PacketFilterRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PacketFilterRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetName

`func (o *PacketFilterRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PacketFilterRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PacketFilterRule) SetName(v string)`

SetName sets Name field to given value.


### GetPorts

`func (o *PacketFilterRule) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *PacketFilterRule) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *PacketFilterRule) SetPorts(v []string)`

SetPorts sets Ports field to given value.


### GetSourceIp

`func (o *PacketFilterRule) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *PacketFilterRule) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *PacketFilterRule) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.


### GetSourceIpPrefixLength

`func (o *PacketFilterRule) GetSourceIpPrefixLength() int32`

GetSourceIpPrefixLength returns the SourceIpPrefixLength field if non-nil, zero value otherwise.

### GetSourceIpPrefixLengthOk

`func (o *PacketFilterRule) GetSourceIpPrefixLengthOk() (*int32, bool)`

GetSourceIpPrefixLengthOk returns a tuple with the SourceIpPrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpPrefixLength

`func (o *PacketFilterRule) SetSourceIpPrefixLength(v int32)`

SetSourceIpPrefixLength sets SourceIpPrefixLength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


