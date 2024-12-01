/*
Hyperledger Cacti Plugin - Connector Ethereum

Can perform basic tasks on a Ethereum ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-ethereum

import (
	"encoding/json"
)

// checks if the InvokeRawWeb3EthMethodV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvokeRawWeb3EthMethodV1Response{}

// InvokeRawWeb3EthMethodV1Response struct for InvokeRawWeb3EthMethodV1Response
type InvokeRawWeb3EthMethodV1Response struct {
	// Status code of the operation
	Status float32 `json:"status"`
	// Output of requested web3.eth method
	Data interface{} `json:"data,omitempty"`
	// Error details
	ErrorDetail *string `json:"errorDetail,omitempty"`
}

// NewInvokeRawWeb3EthMethodV1Response instantiates a new InvokeRawWeb3EthMethodV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvokeRawWeb3EthMethodV1Response(status float32) *InvokeRawWeb3EthMethodV1Response {
	this := InvokeRawWeb3EthMethodV1Response{}
	this.Status = status
	return &this
}

// NewInvokeRawWeb3EthMethodV1ResponseWithDefaults instantiates a new InvokeRawWeb3EthMethodV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvokeRawWeb3EthMethodV1ResponseWithDefaults() *InvokeRawWeb3EthMethodV1Response {
	this := InvokeRawWeb3EthMethodV1Response{}
	return &this
}

// GetStatus returns the Status field value
func (o *InvokeRawWeb3EthMethodV1Response) GetStatus() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InvokeRawWeb3EthMethodV1Response) GetStatusOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InvokeRawWeb3EthMethodV1Response) SetStatus(v float32) {
	o.Status = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvokeRawWeb3EthMethodV1Response) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvokeRawWeb3EthMethodV1Response) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InvokeRawWeb3EthMethodV1Response) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *InvokeRawWeb3EthMethodV1Response) SetData(v interface{}) {
	o.Data = v
}

// GetErrorDetail returns the ErrorDetail field value if set, zero value otherwise.
func (o *InvokeRawWeb3EthMethodV1Response) GetErrorDetail() string {
	if o == nil || IsNil(o.ErrorDetail) {
		var ret string
		return ret
	}
	return *o.ErrorDetail
}

// GetErrorDetailOk returns a tuple with the ErrorDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvokeRawWeb3EthMethodV1Response) GetErrorDetailOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorDetail) {
		return nil, false
	}
	return o.ErrorDetail, true
}

// HasErrorDetail returns a boolean if a field has been set.
func (o *InvokeRawWeb3EthMethodV1Response) HasErrorDetail() bool {
	if o != nil && !IsNil(o.ErrorDetail) {
		return true
	}

	return false
}

// SetErrorDetail gets a reference to the given string and assigns it to the ErrorDetail field.
func (o *InvokeRawWeb3EthMethodV1Response) SetErrorDetail(v string) {
	o.ErrorDetail = &v
}

func (o InvokeRawWeb3EthMethodV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvokeRawWeb3EthMethodV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.ErrorDetail) {
		toSerialize["errorDetail"] = o.ErrorDetail
	}
	return toSerialize, nil
}

type NullableInvokeRawWeb3EthMethodV1Response struct {
	value *InvokeRawWeb3EthMethodV1Response
	isSet bool
}

func (v NullableInvokeRawWeb3EthMethodV1Response) Get() *InvokeRawWeb3EthMethodV1Response {
	return v.value
}

func (v *NullableInvokeRawWeb3EthMethodV1Response) Set(val *InvokeRawWeb3EthMethodV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableInvokeRawWeb3EthMethodV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableInvokeRawWeb3EthMethodV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvokeRawWeb3EthMethodV1Response(val *InvokeRawWeb3EthMethodV1Response) *NullableInvokeRawWeb3EthMethodV1Response {
	return &NullableInvokeRawWeb3EthMethodV1Response{value: val, isSet: true}
}

func (v NullableInvokeRawWeb3EthMethodV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvokeRawWeb3EthMethodV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


