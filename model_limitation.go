/*
さくらのVPS APIドキュメント

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sakura_vps_client_go

import (
	"encoding/json"
)

// checks if the Limitation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Limitation{}

// Limitation struct for Limitation
type Limitation struct {
	// CPUリソースの制限可否
	CpuPerformanceLimit *string `json:"cpu_performance_limit,omitempty"`
	// ネットワーク帯域の制限可否
	NetworkBandwidthLimit *string `json:"network_bandwidth_limit,omitempty"`
	// OP25Bの可否
	OutboundPort25Blocking *string `json:"outbound_port_25_blocking,omitempty"`
	// ストレージのIOPSの制限可否
	StorageIopsLimit *string `json:"storage_iops_limit,omitempty"`
}

// NewLimitation instantiates a new Limitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLimitation() *Limitation {
	this := Limitation{}
	return &this
}

// NewLimitationWithDefaults instantiates a new Limitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimitationWithDefaults() *Limitation {
	this := Limitation{}
	return &this
}

// GetCpuPerformanceLimit returns the CpuPerformanceLimit field value if set, zero value otherwise.
func (o *Limitation) GetCpuPerformanceLimit() string {
	if o == nil || IsNil(o.CpuPerformanceLimit) {
		var ret string
		return ret
	}
	return *o.CpuPerformanceLimit
}

// GetCpuPerformanceLimitOk returns a tuple with the CpuPerformanceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Limitation) GetCpuPerformanceLimitOk() (*string, bool) {
	if o == nil || IsNil(o.CpuPerformanceLimit) {
		return nil, false
	}
	return o.CpuPerformanceLimit, true
}

// HasCpuPerformanceLimit returns a boolean if a field has been set.
func (o *Limitation) HasCpuPerformanceLimit() bool {
	if o != nil && !IsNil(o.CpuPerformanceLimit) {
		return true
	}

	return false
}

// SetCpuPerformanceLimit gets a reference to the given string and assigns it to the CpuPerformanceLimit field.
func (o *Limitation) SetCpuPerformanceLimit(v string) {
	o.CpuPerformanceLimit = &v
}

// GetNetworkBandwidthLimit returns the NetworkBandwidthLimit field value if set, zero value otherwise.
func (o *Limitation) GetNetworkBandwidthLimit() string {
	if o == nil || IsNil(o.NetworkBandwidthLimit) {
		var ret string
		return ret
	}
	return *o.NetworkBandwidthLimit
}

// GetNetworkBandwidthLimitOk returns a tuple with the NetworkBandwidthLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Limitation) GetNetworkBandwidthLimitOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkBandwidthLimit) {
		return nil, false
	}
	return o.NetworkBandwidthLimit, true
}

// HasNetworkBandwidthLimit returns a boolean if a field has been set.
func (o *Limitation) HasNetworkBandwidthLimit() bool {
	if o != nil && !IsNil(o.NetworkBandwidthLimit) {
		return true
	}

	return false
}

// SetNetworkBandwidthLimit gets a reference to the given string and assigns it to the NetworkBandwidthLimit field.
func (o *Limitation) SetNetworkBandwidthLimit(v string) {
	o.NetworkBandwidthLimit = &v
}

// GetOutboundPort25Blocking returns the OutboundPort25Blocking field value if set, zero value otherwise.
func (o *Limitation) GetOutboundPort25Blocking() string {
	if o == nil || IsNil(o.OutboundPort25Blocking) {
		var ret string
		return ret
	}
	return *o.OutboundPort25Blocking
}

// GetOutboundPort25BlockingOk returns a tuple with the OutboundPort25Blocking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Limitation) GetOutboundPort25BlockingOk() (*string, bool) {
	if o == nil || IsNil(o.OutboundPort25Blocking) {
		return nil, false
	}
	return o.OutboundPort25Blocking, true
}

// HasOutboundPort25Blocking returns a boolean if a field has been set.
func (o *Limitation) HasOutboundPort25Blocking() bool {
	if o != nil && !IsNil(o.OutboundPort25Blocking) {
		return true
	}

	return false
}

// SetOutboundPort25Blocking gets a reference to the given string and assigns it to the OutboundPort25Blocking field.
func (o *Limitation) SetOutboundPort25Blocking(v string) {
	o.OutboundPort25Blocking = &v
}

// GetStorageIopsLimit returns the StorageIopsLimit field value if set, zero value otherwise.
func (o *Limitation) GetStorageIopsLimit() string {
	if o == nil || IsNil(o.StorageIopsLimit) {
		var ret string
		return ret
	}
	return *o.StorageIopsLimit
}

// GetStorageIopsLimitOk returns a tuple with the StorageIopsLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Limitation) GetStorageIopsLimitOk() (*string, bool) {
	if o == nil || IsNil(o.StorageIopsLimit) {
		return nil, false
	}
	return o.StorageIopsLimit, true
}

// HasStorageIopsLimit returns a boolean if a field has been set.
func (o *Limitation) HasStorageIopsLimit() bool {
	if o != nil && !IsNil(o.StorageIopsLimit) {
		return true
	}

	return false
}

// SetStorageIopsLimit gets a reference to the given string and assigns it to the StorageIopsLimit field.
func (o *Limitation) SetStorageIopsLimit(v string) {
	o.StorageIopsLimit = &v
}

func (o Limitation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Limitation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CpuPerformanceLimit) {
		toSerialize["cpu_performance_limit"] = o.CpuPerformanceLimit
	}
	if !IsNil(o.NetworkBandwidthLimit) {
		toSerialize["network_bandwidth_limit"] = o.NetworkBandwidthLimit
	}
	if !IsNil(o.OutboundPort25Blocking) {
		toSerialize["outbound_port_25_blocking"] = o.OutboundPort25Blocking
	}
	if !IsNil(o.StorageIopsLimit) {
		toSerialize["storage_iops_limit"] = o.StorageIopsLimit
	}
	return toSerialize, nil
}

type NullableLimitation struct {
	value *Limitation
	isSet bool
}

func (v NullableLimitation) Get() *Limitation {
	return v.value
}

func (v *NullableLimitation) Set(val *Limitation) {
	v.value = val
	v.isSet = true
}

func (v NullableLimitation) IsSet() bool {
	return v.isSet
}

func (v *NullableLimitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLimitation(val *Limitation) *NullableLimitation {
	return &NullableLimitation{value: val, isSet: true}
}

func (v NullableLimitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLimitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


