/*
さくらのVPS APIドキュメント

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sakura_vps_client_go

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// ServerMonitoringSettingsHealthCheck - struct for ServerMonitoringSettingsHealthCheck
type ServerMonitoringSettingsHealthCheck struct {
	HealthCheckHttp *HealthCheckHttp
	HealthCheckHttps *HealthCheckHttps
	HealthCheckPing *HealthCheckPing
	HealthCheckPop3 *HealthCheckPop3
	HealthCheckSmtp *HealthCheckSmtp
	HealthCheckSsh *HealthCheckSsh
	HealthCheckTcp *HealthCheckTcp
}

// HealthCheckHttpAsServerMonitoringSettingsHealthCheck is a convenience function that returns HealthCheckHttp wrapped in ServerMonitoringSettingsHealthCheck
func HealthCheckHttpAsServerMonitoringSettingsHealthCheck(v *HealthCheckHttp) ServerMonitoringSettingsHealthCheck {
	return ServerMonitoringSettingsHealthCheck{
		HealthCheckHttp: v,
	}
}

// HealthCheckHttpsAsServerMonitoringSettingsHealthCheck is a convenience function that returns HealthCheckHttps wrapped in ServerMonitoringSettingsHealthCheck
func HealthCheckHttpsAsServerMonitoringSettingsHealthCheck(v *HealthCheckHttps) ServerMonitoringSettingsHealthCheck {
	return ServerMonitoringSettingsHealthCheck{
		HealthCheckHttps: v,
	}
}

// HealthCheckPingAsServerMonitoringSettingsHealthCheck is a convenience function that returns HealthCheckPing wrapped in ServerMonitoringSettingsHealthCheck
func HealthCheckPingAsServerMonitoringSettingsHealthCheck(v *HealthCheckPing) ServerMonitoringSettingsHealthCheck {
	return ServerMonitoringSettingsHealthCheck{
		HealthCheckPing: v,
	}
}

// HealthCheckPop3AsServerMonitoringSettingsHealthCheck is a convenience function that returns HealthCheckPop3 wrapped in ServerMonitoringSettingsHealthCheck
func HealthCheckPop3AsServerMonitoringSettingsHealthCheck(v *HealthCheckPop3) ServerMonitoringSettingsHealthCheck {
	return ServerMonitoringSettingsHealthCheck{
		HealthCheckPop3: v,
	}
}

// HealthCheckSmtpAsServerMonitoringSettingsHealthCheck is a convenience function that returns HealthCheckSmtp wrapped in ServerMonitoringSettingsHealthCheck
func HealthCheckSmtpAsServerMonitoringSettingsHealthCheck(v *HealthCheckSmtp) ServerMonitoringSettingsHealthCheck {
	return ServerMonitoringSettingsHealthCheck{
		HealthCheckSmtp: v,
	}
}

// HealthCheckSshAsServerMonitoringSettingsHealthCheck is a convenience function that returns HealthCheckSsh wrapped in ServerMonitoringSettingsHealthCheck
func HealthCheckSshAsServerMonitoringSettingsHealthCheck(v *HealthCheckSsh) ServerMonitoringSettingsHealthCheck {
	return ServerMonitoringSettingsHealthCheck{
		HealthCheckSsh: v,
	}
}

// HealthCheckTcpAsServerMonitoringSettingsHealthCheck is a convenience function that returns HealthCheckTcp wrapped in ServerMonitoringSettingsHealthCheck
func HealthCheckTcpAsServerMonitoringSettingsHealthCheck(v *HealthCheckTcp) ServerMonitoringSettingsHealthCheck {
	return ServerMonitoringSettingsHealthCheck{
		HealthCheckTcp: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ServerMonitoringSettingsHealthCheck) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HealthCheckHttp
	err = newStrictDecoder(data).Decode(&dst.HealthCheckHttp)
	if err == nil {
		jsonHealthCheckHttp, _ := json.Marshal(dst.HealthCheckHttp)
		if string(jsonHealthCheckHttp) == "{}" { // empty struct
			dst.HealthCheckHttp = nil
		} else {
			if err = validator.Validate(dst.HealthCheckHttp); err != nil {
				dst.HealthCheckHttp = nil
			} else {
				match++
			}
		}
	} else {
		dst.HealthCheckHttp = nil
	}

	// try to unmarshal data into HealthCheckHttps
	err = newStrictDecoder(data).Decode(&dst.HealthCheckHttps)
	if err == nil {
		jsonHealthCheckHttps, _ := json.Marshal(dst.HealthCheckHttps)
		if string(jsonHealthCheckHttps) == "{}" { // empty struct
			dst.HealthCheckHttps = nil
		} else {
			if err = validator.Validate(dst.HealthCheckHttps); err != nil {
				dst.HealthCheckHttps = nil
			} else {
				match++
			}
		}
	} else {
		dst.HealthCheckHttps = nil
	}

	// try to unmarshal data into HealthCheckPing
	err = newStrictDecoder(data).Decode(&dst.HealthCheckPing)
	if err == nil {
		jsonHealthCheckPing, _ := json.Marshal(dst.HealthCheckPing)
		if string(jsonHealthCheckPing) == "{}" { // empty struct
			dst.HealthCheckPing = nil
		} else {
			if err = validator.Validate(dst.HealthCheckPing); err != nil {
				dst.HealthCheckPing = nil
			} else {
				match++
			}
		}
	} else {
		dst.HealthCheckPing = nil
	}

	// try to unmarshal data into HealthCheckPop3
	err = newStrictDecoder(data).Decode(&dst.HealthCheckPop3)
	if err == nil {
		jsonHealthCheckPop3, _ := json.Marshal(dst.HealthCheckPop3)
		if string(jsonHealthCheckPop3) == "{}" { // empty struct
			dst.HealthCheckPop3 = nil
		} else {
			if err = validator.Validate(dst.HealthCheckPop3); err != nil {
				dst.HealthCheckPop3 = nil
			} else {
				match++
			}
		}
	} else {
		dst.HealthCheckPop3 = nil
	}

	// try to unmarshal data into HealthCheckSmtp
	err = newStrictDecoder(data).Decode(&dst.HealthCheckSmtp)
	if err == nil {
		jsonHealthCheckSmtp, _ := json.Marshal(dst.HealthCheckSmtp)
		if string(jsonHealthCheckSmtp) == "{}" { // empty struct
			dst.HealthCheckSmtp = nil
		} else {
			if err = validator.Validate(dst.HealthCheckSmtp); err != nil {
				dst.HealthCheckSmtp = nil
			} else {
				match++
			}
		}
	} else {
		dst.HealthCheckSmtp = nil
	}

	// try to unmarshal data into HealthCheckSsh
	err = newStrictDecoder(data).Decode(&dst.HealthCheckSsh)
	if err == nil {
		jsonHealthCheckSsh, _ := json.Marshal(dst.HealthCheckSsh)
		if string(jsonHealthCheckSsh) == "{}" { // empty struct
			dst.HealthCheckSsh = nil
		} else {
			if err = validator.Validate(dst.HealthCheckSsh); err != nil {
				dst.HealthCheckSsh = nil
			} else {
				match++
			}
		}
	} else {
		dst.HealthCheckSsh = nil
	}

	// try to unmarshal data into HealthCheckTcp
	err = newStrictDecoder(data).Decode(&dst.HealthCheckTcp)
	if err == nil {
		jsonHealthCheckTcp, _ := json.Marshal(dst.HealthCheckTcp)
		if string(jsonHealthCheckTcp) == "{}" { // empty struct
			dst.HealthCheckTcp = nil
		} else {
			if err = validator.Validate(dst.HealthCheckTcp); err != nil {
				dst.HealthCheckTcp = nil
			} else {
				match++
			}
		}
	} else {
		dst.HealthCheckTcp = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.HealthCheckHttp = nil
		dst.HealthCheckHttps = nil
		dst.HealthCheckPing = nil
		dst.HealthCheckPop3 = nil
		dst.HealthCheckSmtp = nil
		dst.HealthCheckSsh = nil
		dst.HealthCheckTcp = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ServerMonitoringSettingsHealthCheck)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ServerMonitoringSettingsHealthCheck)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ServerMonitoringSettingsHealthCheck) MarshalJSON() ([]byte, error) {
	if src.HealthCheckHttp != nil {
		return json.Marshal(&src.HealthCheckHttp)
	}

	if src.HealthCheckHttps != nil {
		return json.Marshal(&src.HealthCheckHttps)
	}

	if src.HealthCheckPing != nil {
		return json.Marshal(&src.HealthCheckPing)
	}

	if src.HealthCheckPop3 != nil {
		return json.Marshal(&src.HealthCheckPop3)
	}

	if src.HealthCheckSmtp != nil {
		return json.Marshal(&src.HealthCheckSmtp)
	}

	if src.HealthCheckSsh != nil {
		return json.Marshal(&src.HealthCheckSsh)
	}

	if src.HealthCheckTcp != nil {
		return json.Marshal(&src.HealthCheckTcp)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ServerMonitoringSettingsHealthCheck) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.HealthCheckHttp != nil {
		return obj.HealthCheckHttp
	}

	if obj.HealthCheckHttps != nil {
		return obj.HealthCheckHttps
	}

	if obj.HealthCheckPing != nil {
		return obj.HealthCheckPing
	}

	if obj.HealthCheckPop3 != nil {
		return obj.HealthCheckPop3
	}

	if obj.HealthCheckSmtp != nil {
		return obj.HealthCheckSmtp
	}

	if obj.HealthCheckSsh != nil {
		return obj.HealthCheckSsh
	}

	if obj.HealthCheckTcp != nil {
		return obj.HealthCheckTcp
	}

	// all schemas are nil
	return nil
}

type NullableServerMonitoringSettingsHealthCheck struct {
	value *ServerMonitoringSettingsHealthCheck
	isSet bool
}

func (v NullableServerMonitoringSettingsHealthCheck) Get() *ServerMonitoringSettingsHealthCheck {
	return v.value
}

func (v *NullableServerMonitoringSettingsHealthCheck) Set(val *ServerMonitoringSettingsHealthCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableServerMonitoringSettingsHealthCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableServerMonitoringSettingsHealthCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerMonitoringSettingsHealthCheck(val *ServerMonitoringSettingsHealthCheck) *NullableServerMonitoringSettingsHealthCheck {
	return &NullableServerMonitoringSettingsHealthCheck{value: val, isSet: true}
}

func (v NullableServerMonitoringSettingsHealthCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerMonitoringSettingsHealthCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


