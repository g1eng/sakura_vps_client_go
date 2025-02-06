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

// checks if the PostSwitch201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSwitch201Response{}

// PostSwitch201Response struct for PostSwitch201Response
type PostSwitch201Response struct {
	// id
	Id int32 `json:"id"`
	// 名前
	Name string `json:"name"`
	// 説明
	Description string `json:"description"`
	// スイッチコード
	SwitchCode string `json:"switch_code"`
	Zone PostSwitch201ResponseZone `json:"zone"`
	// 接続されているサーバーのインターフェースid
	ServerInterfaces []int32 `json:"server_interfaces"`
	// 接続されている追加ストレージ（NFS）のインターフェースid
	NfsServerInterfaces []int32 `json:"nfs_server_interfaces"`
	// 接続されている外部接続の情報
	ExternalConnection map[string]interface{} `json:"external_connection"`
}

type _PostSwitch201Response PostSwitch201Response

// NewPostSwitch201Response instantiates a new PostSwitch201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSwitch201Response(id int32, name string, description string, switchCode string, zone PostSwitch201ResponseZone, serverInterfaces []int32, nfsServerInterfaces []int32, externalConnection map[string]interface{}) *PostSwitch201Response {
	this := PostSwitch201Response{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.SwitchCode = switchCode
	this.Zone = zone
	this.ServerInterfaces = serverInterfaces
	this.NfsServerInterfaces = nfsServerInterfaces
	this.ExternalConnection = externalConnection
	return &this
}

// NewPostSwitch201ResponseWithDefaults instantiates a new PostSwitch201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSwitch201ResponseWithDefaults() *PostSwitch201Response {
	this := PostSwitch201Response{}
	return &this
}

// GetId returns the Id field value
func (o *PostSwitch201Response) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostSwitch201Response) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostSwitch201Response) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PostSwitch201Response) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostSwitch201Response) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PostSwitch201Response) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *PostSwitch201Response) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PostSwitch201Response) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PostSwitch201Response) SetDescription(v string) {
	o.Description = v
}

// GetSwitchCode returns the SwitchCode field value
func (o *PostSwitch201Response) GetSwitchCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SwitchCode
}

// GetSwitchCodeOk returns a tuple with the SwitchCode field value
// and a boolean to check if the value has been set.
func (o *PostSwitch201Response) GetSwitchCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SwitchCode, true
}

// SetSwitchCode sets field value
func (o *PostSwitch201Response) SetSwitchCode(v string) {
	o.SwitchCode = v
}

// GetZone returns the Zone field value
func (o *PostSwitch201Response) GetZone() PostSwitch201ResponseZone {
	if o == nil {
		var ret PostSwitch201ResponseZone
		return ret
	}

	return o.Zone
}

// GetZoneOk returns a tuple with the Zone field value
// and a boolean to check if the value has been set.
func (o *PostSwitch201Response) GetZoneOk() (*PostSwitch201ResponseZone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zone, true
}

// SetZone sets field value
func (o *PostSwitch201Response) SetZone(v PostSwitch201ResponseZone) {
	o.Zone = v
}

// GetServerInterfaces returns the ServerInterfaces field value
func (o *PostSwitch201Response) GetServerInterfaces() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ServerInterfaces
}

// GetServerInterfacesOk returns a tuple with the ServerInterfaces field value
// and a boolean to check if the value has been set.
func (o *PostSwitch201Response) GetServerInterfacesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerInterfaces, true
}

// SetServerInterfaces sets field value
func (o *PostSwitch201Response) SetServerInterfaces(v []int32) {
	o.ServerInterfaces = v
}

// GetNfsServerInterfaces returns the NfsServerInterfaces field value
func (o *PostSwitch201Response) GetNfsServerInterfaces() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.NfsServerInterfaces
}

// GetNfsServerInterfacesOk returns a tuple with the NfsServerInterfaces field value
// and a boolean to check if the value has been set.
func (o *PostSwitch201Response) GetNfsServerInterfacesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NfsServerInterfaces, true
}

// SetNfsServerInterfaces sets field value
func (o *PostSwitch201Response) SetNfsServerInterfaces(v []int32) {
	o.NfsServerInterfaces = v
}

// GetExternalConnection returns the ExternalConnection field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *PostSwitch201Response) GetExternalConnection() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ExternalConnection
}

// GetExternalConnectionOk returns a tuple with the ExternalConnection field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostSwitch201Response) GetExternalConnectionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExternalConnection) {
		return map[string]interface{}{}, false
	}
	return o.ExternalConnection, true
}

// SetExternalConnection sets field value
func (o *PostSwitch201Response) SetExternalConnection(v map[string]interface{}) {
	o.ExternalConnection = v
}

func (o PostSwitch201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSwitch201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["switch_code"] = o.SwitchCode
	toSerialize["zone"] = o.Zone
	toSerialize["server_interfaces"] = o.ServerInterfaces
	toSerialize["nfs_server_interfaces"] = o.NfsServerInterfaces
	if o.ExternalConnection != nil {
		toSerialize["external_connection"] = o.ExternalConnection
	}
	return toSerialize, nil
}

func (o *PostSwitch201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
		"switch_code",
		"zone",
		"server_interfaces",
		"nfs_server_interfaces",
		"external_connection",
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

	varPostSwitch201Response := _PostSwitch201Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostSwitch201Response)

	if err != nil {
		return err
	}

	*o = PostSwitch201Response(varPostSwitch201Response)

	return err
}

type NullablePostSwitch201Response struct {
	value *PostSwitch201Response
	isSet bool
}

func (v NullablePostSwitch201Response) Get() *PostSwitch201Response {
	return v.value
}

func (v *NullablePostSwitch201Response) Set(val *PostSwitch201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSwitch201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSwitch201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSwitch201Response(val *PostSwitch201Response) *NullablePostSwitch201Response {
	return &NullablePostSwitch201Response{value: val, isSet: true}
}

func (v NullablePostSwitch201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSwitch201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


