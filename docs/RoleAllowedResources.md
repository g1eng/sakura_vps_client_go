# RoleAllowedResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to **[]int32** | 利用できるサーバーのid | [optional] 
**Switches** | Pointer to **[]int32** | 利用できるスイッチのid | [optional] 
**NfsServers** | Pointer to **[]int32** | 利用できるNFSのid | [optional] 

## Methods

### NewRoleAllowedResources

`func NewRoleAllowedResources() *RoleAllowedResources`

NewRoleAllowedResources instantiates a new RoleAllowedResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAllowedResourcesWithDefaults

`func NewRoleAllowedResourcesWithDefaults() *RoleAllowedResources`

NewRoleAllowedResourcesWithDefaults instantiates a new RoleAllowedResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *RoleAllowedResources) GetServers() []int32`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *RoleAllowedResources) GetServersOk() (*[]int32, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *RoleAllowedResources) SetServers(v []int32)`

SetServers sets Servers field to given value.

### HasServers

`func (o *RoleAllowedResources) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetSwitches

`func (o *RoleAllowedResources) GetSwitches() []int32`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *RoleAllowedResources) GetSwitchesOk() (*[]int32, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *RoleAllowedResources) SetSwitches(v []int32)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *RoleAllowedResources) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.

### GetNfsServers

`func (o *RoleAllowedResources) GetNfsServers() []int32`

GetNfsServers returns the NfsServers field if non-nil, zero value otherwise.

### GetNfsServersOk

`func (o *RoleAllowedResources) GetNfsServersOk() (*[]int32, bool)`

GetNfsServersOk returns a tuple with the NfsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsServers

`func (o *RoleAllowedResources) SetNfsServers(v []int32)`

SetNfsServers sets NfsServers field to given value.

### HasNfsServers

`func (o *RoleAllowedResources) HasNfsServers() bool`

HasNfsServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


