/*
Hyperledger Cactus Plugin - Connector Besu

Can perform basic tasks on a Besu ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-besu

import (
	"encoding/json"
)

// checks if the Web3BlockHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3BlockHeader{}

// Web3BlockHeader struct for Web3BlockHeader
type Web3BlockHeader struct {
	Number float32 `json:"number"`
	Hash string `json:"hash"`
	ParentHash string `json:"parentHash"`
	Nonce string `json:"nonce"`
	Sha3Uncles string `json:"sha3Uncles"`
	LogsBloom string `json:"logsBloom"`
	TransactionRoot string `json:"transactionRoot"`
	StateRoot string `json:"stateRoot"`
	ReceiptRoot string `json:"receiptRoot"`
	Miner string `json:"miner"`
	ExtraData string `json:"extraData"`
	GasLimit int32 `json:"gasLimit"`
	GasUsed int32 `json:"gasUsed"`
	Timestamp Web3BlockHeaderTimestamp `json:"timestamp"`
}

// NewWeb3BlockHeader instantiates a new Web3BlockHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3BlockHeader(number float32, hash string, parentHash string, nonce string, sha3Uncles string, logsBloom string, transactionRoot string, stateRoot string, receiptRoot string, miner string, extraData string, gasLimit int32, gasUsed int32, timestamp Web3BlockHeaderTimestamp) *Web3BlockHeader {
	this := Web3BlockHeader{}
	this.Number = number
	this.Hash = hash
	this.ParentHash = parentHash
	this.Nonce = nonce
	this.Sha3Uncles = sha3Uncles
	this.LogsBloom = logsBloom
	this.TransactionRoot = transactionRoot
	this.StateRoot = stateRoot
	this.ReceiptRoot = receiptRoot
	this.Miner = miner
	this.ExtraData = extraData
	this.GasLimit = gasLimit
	this.GasUsed = gasUsed
	this.Timestamp = timestamp
	return &this
}

// NewWeb3BlockHeaderWithDefaults instantiates a new Web3BlockHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3BlockHeaderWithDefaults() *Web3BlockHeader {
	this := Web3BlockHeader{}
	return &this
}

// GetNumber returns the Number field value
func (o *Web3BlockHeader) GetNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *Web3BlockHeader) SetNumber(v float32) {
	o.Number = v
}

// GetHash returns the Hash field value
func (o *Web3BlockHeader) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *Web3BlockHeader) SetHash(v string) {
	o.Hash = v
}

// GetParentHash returns the ParentHash field value
func (o *Web3BlockHeader) GetParentHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentHash
}

// GetParentHashOk returns a tuple with the ParentHash field value
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetParentHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentHash, true
}

// SetParentHash sets field value
func (o *Web3BlockHeader) SetParentHash(v string) {
	o.ParentHash = v
}

// GetNonce returns the Nonce field value
func (o *Web3BlockHeader) GetNonce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetNonceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *Web3BlockHeader) SetNonce(v string) {
	o.Nonce = v
}

// GetSha3Uncles returns the Sha3Uncles field value
func (o *Web3BlockHeader) GetSha3Uncles() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha3Uncles
}

// GetSha3UnclesOk returns a tuple with the Sha3Uncles field value
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetSha3UnclesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha3Uncles, true
}

// SetSha3Uncles sets field value
func (o *Web3BlockHeader) SetSha3Uncles(v string) {
	o.Sha3Uncles = v
}

// GetLogsBloom returns the LogsBloom field value
func (o *Web3BlockHeader) GetLogsBloom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogsBloom
}

// GetLogsBloomOk returns a tuple with the LogsBloom field value
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetLogsBloomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogsBloom, true
}

// SetLogsBloom sets field value
func (o *Web3BlockHeader) SetLogsBloom(v string) {
	o.LogsBloom = v
}

// GetTransactionRoot returns the TransactionRoot field value
func (o *Web3BlockHeader) GetTransactionRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionRoot
}

// GetTransactionRootOk returns a tuple with the TransactionRoot field value
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetTransactionRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionRoot, true
}

// SetTransactionRoot sets field value
func (o *Web3BlockHeader) SetTransactionRoot(v string) {
	o.TransactionRoot = v
}

// GetStateRoot returns the StateRoot field value
func (o *Web3BlockHeader) GetStateRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateRoot
}

// GetStateRootOk returns a tuple with the StateRoot field value
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetStateRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateRoot, true
}

// SetStateRoot sets field value
func (o *Web3BlockHeader) SetStateRoot(v string) {
	o.StateRoot = v
}

// GetReceiptRoot returns the ReceiptRoot field value
func (o *Web3BlockHeader) GetReceiptRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceiptRoot
}

// GetReceiptRootOk returns a tuple with the ReceiptRoot field value
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetReceiptRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiptRoot, true
}

// SetReceiptRoot sets field value
func (o *Web3BlockHeader) SetReceiptRoot(v string) {
	o.ReceiptRoot = v
}

// GetMiner returns the Miner field value
func (o *Web3BlockHeader) GetMiner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Miner
}

// GetMinerOk returns a tuple with the Miner field value
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetMinerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Miner, true
}

// SetMiner sets field value
func (o *Web3BlockHeader) SetMiner(v string) {
	o.Miner = v
}

// GetExtraData returns the ExtraData field value
func (o *Web3BlockHeader) GetExtraData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetExtraDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtraData, true
}

// SetExtraData sets field value
func (o *Web3BlockHeader) SetExtraData(v string) {
	o.ExtraData = v
}

// GetGasLimit returns the GasLimit field value
func (o *Web3BlockHeader) GetGasLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetGasLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasLimit, true
}

// SetGasLimit sets field value
func (o *Web3BlockHeader) SetGasLimit(v int32) {
	o.GasLimit = v
}

// GetGasUsed returns the GasUsed field value
func (o *Web3BlockHeader) GetGasUsed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetGasUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *Web3BlockHeader) SetGasUsed(v int32) {
	o.GasUsed = v
}

// GetTimestamp returns the Timestamp field value
func (o *Web3BlockHeader) GetTimestamp() Web3BlockHeaderTimestamp {
	if o == nil {
		var ret Web3BlockHeaderTimestamp
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Web3BlockHeader) GetTimestampOk() (*Web3BlockHeaderTimestamp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *Web3BlockHeader) SetTimestamp(v Web3BlockHeaderTimestamp) {
	o.Timestamp = v
}

func (o Web3BlockHeader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3BlockHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["number"] = o.Number
	toSerialize["hash"] = o.Hash
	toSerialize["parentHash"] = o.ParentHash
	toSerialize["nonce"] = o.Nonce
	toSerialize["sha3Uncles"] = o.Sha3Uncles
	toSerialize["logsBloom"] = o.LogsBloom
	toSerialize["transactionRoot"] = o.TransactionRoot
	toSerialize["stateRoot"] = o.StateRoot
	toSerialize["receiptRoot"] = o.ReceiptRoot
	toSerialize["miner"] = o.Miner
	toSerialize["extraData"] = o.ExtraData
	toSerialize["gasLimit"] = o.GasLimit
	toSerialize["gasUsed"] = o.GasUsed
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

type NullableWeb3BlockHeader struct {
	value *Web3BlockHeader
	isSet bool
}

func (v NullableWeb3BlockHeader) Get() *Web3BlockHeader {
	return v.value
}

func (v *NullableWeb3BlockHeader) Set(val *Web3BlockHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3BlockHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3BlockHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3BlockHeader(val *Web3BlockHeader) *NullableWeb3BlockHeader {
	return &NullableWeb3BlockHeader{value: val, isSet: true}
}

func (v NullableWeb3BlockHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3BlockHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


