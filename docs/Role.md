# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | id | [readonly] 
**Name** | **string** | 名前 | 
**Description** | **string** | 説明 | 
**PermissionFiltering** | **string** | 利用できる権限を制限するか | 
**AllowedPermissions** | **[]string** | 利用できる権限。permission_filteringがenabledの場合のみ指定可能。**権限の一覧を取得する**&#x60;/permissions&#x60;のcode値を指定します。 | 
**ResourceFiltering** | **string** | 利用できるリソースを制限するか | 
**AllowedResources** | [**NullableRoleAllowedResources**](RoleAllowedResources.md) |  | 

## Methods

### NewRole

`func NewRole(id int32, name string, description string, permissionFiltering string, allowedPermissions []string, resourceFiltering string, allowedResources NullableRoleAllowedResources, ) *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Role) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Role) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Role) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Role) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Role) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Role) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Role) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Role) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Role) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPermissionFiltering

`func (o *Role) GetPermissionFiltering() string`

GetPermissionFiltering returns the PermissionFiltering field if non-nil, zero value otherwise.

### GetPermissionFilteringOk

`func (o *Role) GetPermissionFilteringOk() (*string, bool)`

GetPermissionFilteringOk returns a tuple with the PermissionFiltering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionFiltering

`func (o *Role) SetPermissionFiltering(v string)`

SetPermissionFiltering sets PermissionFiltering field to given value.


### GetAllowedPermissions

`func (o *Role) GetAllowedPermissions() []string`

GetAllowedPermissions returns the AllowedPermissions field if non-nil, zero value otherwise.

### GetAllowedPermissionsOk

`func (o *Role) GetAllowedPermissionsOk() (*[]string, bool)`

GetAllowedPermissionsOk returns a tuple with the AllowedPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPermissions

`func (o *Role) SetAllowedPermissions(v []string)`

SetAllowedPermissions sets AllowedPermissions field to given value.


### GetResourceFiltering

`func (o *Role) GetResourceFiltering() string`

GetResourceFiltering returns the ResourceFiltering field if non-nil, zero value otherwise.

### GetResourceFilteringOk

`func (o *Role) GetResourceFilteringOk() (*string, bool)`

GetResourceFilteringOk returns a tuple with the ResourceFiltering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceFiltering

`func (o *Role) SetResourceFiltering(v string)`

SetResourceFiltering sets ResourceFiltering field to given value.


### GetAllowedResources

`func (o *Role) GetAllowedResources() RoleAllowedResources`

GetAllowedResources returns the AllowedResources field if non-nil, zero value otherwise.

### GetAllowedResourcesOk

`func (o *Role) GetAllowedResourcesOk() (*RoleAllowedResources, bool)`

GetAllowedResourcesOk returns a tuple with the AllowedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedResources

`func (o *Role) SetAllowedResources(v RoleAllowedResources)`

SetAllowedResources sets AllowedResources field to given value.


### SetAllowedResourcesNil

`func (o *Role) SetAllowedResourcesNil(b bool)`

 SetAllowedResourcesNil sets the value for AllowedResources to be an explicit nil

### UnsetAllowedResources
`func (o *Role) UnsetAllowedResources()`

UnsetAllowedResources ensures that no value is present for AllowedResources, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


