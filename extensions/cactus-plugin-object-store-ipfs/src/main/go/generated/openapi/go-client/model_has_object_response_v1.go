/*
Hyperledger Cactus Plugin - Object Store - IPFS 

Contains/describes the Hyperledger Cactus Object Store IPFS plugin.

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-object-store-ipfs

import (
	"encoding/json"
)

// checks if the HasObjectResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HasObjectResponseV1{}

// HasObjectResponseV1 struct for HasObjectResponseV1
type HasObjectResponseV1 struct {
	// The key that was used to check the presence of the value in the object store.
	Key string `json:"key"`
	// Date and time encoded as JSON when the presence check was performed by the plugin backend.
	CheckedAt string `json:"checkedAt"`
	// The boolean true or false indicating the presence or absence of an object under 'key'.
	IsPresent bool `json:"isPresent"`
}

// NewHasObjectResponseV1 instantiates a new HasObjectResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHasObjectResponseV1(key string, checkedAt string, isPresent bool) *HasObjectResponseV1 {
	this := HasObjectResponseV1{}
	this.Key = key
	this.CheckedAt = checkedAt
	this.IsPresent = isPresent
	return &this
}

// NewHasObjectResponseV1WithDefaults instantiates a new HasObjectResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHasObjectResponseV1WithDefaults() *HasObjectResponseV1 {
	this := HasObjectResponseV1{}
	return &this
}

// GetKey returns the Key field value
func (o *HasObjectResponseV1) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *HasObjectResponseV1) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *HasObjectResponseV1) SetKey(v string) {
	o.Key = v
}

// GetCheckedAt returns the CheckedAt field value
func (o *HasObjectResponseV1) GetCheckedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckedAt
}

// GetCheckedAtOk returns a tuple with the CheckedAt field value
// and a boolean to check if the value has been set.
func (o *HasObjectResponseV1) GetCheckedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckedAt, true
}

// SetCheckedAt sets field value
func (o *HasObjectResponseV1) SetCheckedAt(v string) {
	o.CheckedAt = v
}

// GetIsPresent returns the IsPresent field value
func (o *HasObjectResponseV1) GetIsPresent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPresent
}

// GetIsPresentOk returns a tuple with the IsPresent field value
// and a boolean to check if the value has been set.
func (o *HasObjectResponseV1) GetIsPresentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPresent, true
}

// SetIsPresent sets field value
func (o *HasObjectResponseV1) SetIsPresent(v bool) {
	o.IsPresent = v
}

func (o HasObjectResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HasObjectResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["checkedAt"] = o.CheckedAt
	toSerialize["isPresent"] = o.IsPresent
	return toSerialize, nil
}

type NullableHasObjectResponseV1 struct {
	value *HasObjectResponseV1
	isSet bool
}

func (v NullableHasObjectResponseV1) Get() *HasObjectResponseV1 {
	return v.value
}

func (v *NullableHasObjectResponseV1) Set(val *HasObjectResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableHasObjectResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableHasObjectResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHasObjectResponseV1(val *HasObjectResponseV1) *NullableHasObjectResponseV1 {
	return &NullableHasObjectResponseV1{value: val, isSet: true}
}

func (v NullableHasObjectResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHasObjectResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


