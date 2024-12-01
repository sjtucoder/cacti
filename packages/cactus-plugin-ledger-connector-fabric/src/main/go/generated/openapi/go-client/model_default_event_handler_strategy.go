/*
Hyperledger Cactus Plugin - Connector Fabric

Can perform basic tasks on a fabric ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-fabric

import (
	"encoding/json"
	"fmt"
)

// DefaultEventHandlerStrategy the model 'DefaultEventHandlerStrategy'
type DefaultEventHandlerStrategy string

// List of DefaultEventHandlerStrategy
const (
	MSPID_SCOPE_ALLFORTX DefaultEventHandlerStrategy = "MSPID_SCOPE_ALLFORTX"
	MSPID_SCOPE_ANYFORTX DefaultEventHandlerStrategy = "MSPID_SCOPE_ANYFORTX"
	NETWORK_SCOPE_ALLFORTX DefaultEventHandlerStrategy = "NETWORK_SCOPE_ALLFORTX"
	NETWORK_SCOPE_ANYFORTX DefaultEventHandlerStrategy = "NETWORK_SCOPE_ANYFORTX"
)

// All allowed values of DefaultEventHandlerStrategy enum
var AllowedDefaultEventHandlerStrategyEnumValues = []DefaultEventHandlerStrategy{
	"MSPID_SCOPE_ALLFORTX",
	"MSPID_SCOPE_ANYFORTX",
	"NETWORK_SCOPE_ALLFORTX",
	"NETWORK_SCOPE_ANYFORTX",
}

func (v *DefaultEventHandlerStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DefaultEventHandlerStrategy(value)
	for _, existing := range AllowedDefaultEventHandlerStrategyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DefaultEventHandlerStrategy", value)
}

// NewDefaultEventHandlerStrategyFromValue returns a pointer to a valid DefaultEventHandlerStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDefaultEventHandlerStrategyFromValue(v string) (*DefaultEventHandlerStrategy, error) {
	ev := DefaultEventHandlerStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DefaultEventHandlerStrategy: valid values are %v", v, AllowedDefaultEventHandlerStrategyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DefaultEventHandlerStrategy) IsValid() bool {
	for _, existing := range AllowedDefaultEventHandlerStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DefaultEventHandlerStrategy value
func (v DefaultEventHandlerStrategy) Ptr() *DefaultEventHandlerStrategy {
	return &v
}

type NullableDefaultEventHandlerStrategy struct {
	value *DefaultEventHandlerStrategy
	isSet bool
}

func (v NullableDefaultEventHandlerStrategy) Get() *DefaultEventHandlerStrategy {
	return v.value
}

func (v *NullableDefaultEventHandlerStrategy) Set(val *DefaultEventHandlerStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultEventHandlerStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultEventHandlerStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultEventHandlerStrategy(val *DefaultEventHandlerStrategy) *NullableDefaultEventHandlerStrategy {
	return &NullableDefaultEventHandlerStrategy{value: val, isSet: true}
}

func (v NullableDefaultEventHandlerStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultEventHandlerStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

