/*
Hyperledger Cactus Plugin - Odap Hermes

Implementation for Odap and Hermes

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-satp-hermes

import (
	"encoding/json"
)

// checks if the RollbackAckV1Message type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RollbackAckV1Message{}

// RollbackAckV1Message struct for RollbackAckV1Message
type RollbackAckV1Message struct {
	SessionID string `json:"sessionID"`
	Success bool `json:"success"`
	Signature string `json:"signature"`
}

// NewRollbackAckV1Message instantiates a new RollbackAckV1Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRollbackAckV1Message(sessionID string, success bool, signature string) *RollbackAckV1Message {
	this := RollbackAckV1Message{}
	this.SessionID = sessionID
	this.Success = success
	this.Signature = signature
	return &this
}

// NewRollbackAckV1MessageWithDefaults instantiates a new RollbackAckV1Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollbackAckV1MessageWithDefaults() *RollbackAckV1Message {
	this := RollbackAckV1Message{}
	return &this
}

// GetSessionID returns the SessionID field value
func (o *RollbackAckV1Message) GetSessionID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionID
}

// GetSessionIDOk returns a tuple with the SessionID field value
// and a boolean to check if the value has been set.
func (o *RollbackAckV1Message) GetSessionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionID, true
}

// SetSessionID sets field value
func (o *RollbackAckV1Message) SetSessionID(v string) {
	o.SessionID = v
}

// GetSuccess returns the Success field value
func (o *RollbackAckV1Message) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *RollbackAckV1Message) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *RollbackAckV1Message) SetSuccess(v bool) {
	o.Success = v
}

// GetSignature returns the Signature field value
func (o *RollbackAckV1Message) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *RollbackAckV1Message) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *RollbackAckV1Message) SetSignature(v string) {
	o.Signature = v
}

func (o RollbackAckV1Message) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RollbackAckV1Message) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionID"] = o.SessionID
	toSerialize["success"] = o.Success
	toSerialize["signature"] = o.Signature
	return toSerialize, nil
}

type NullableRollbackAckV1Message struct {
	value *RollbackAckV1Message
	isSet bool
}

func (v NullableRollbackAckV1Message) Get() *RollbackAckV1Message {
	return v.value
}

func (v *NullableRollbackAckV1Message) Set(val *RollbackAckV1Message) {
	v.value = val
	v.isSet = true
}

func (v NullableRollbackAckV1Message) IsSet() bool {
	return v.isSet
}

func (v *NullableRollbackAckV1Message) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRollbackAckV1Message(val *RollbackAckV1Message) *NullableRollbackAckV1Message {
	return &NullableRollbackAckV1Message{value: val, isSet: true}
}

func (v NullableRollbackAckV1Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRollbackAckV1Message) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


