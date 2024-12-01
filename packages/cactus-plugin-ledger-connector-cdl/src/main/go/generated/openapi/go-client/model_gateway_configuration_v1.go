/*
Hyperledger Cacti Plugin - Connector CDL

Can perform basic tasks on Fujitsu CDL service.

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-cdl

import (
	"encoding/json"
)

// checks if the GatewayConfigurationV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayConfigurationV1{}

// GatewayConfigurationV1 struct for GatewayConfigurationV1
type GatewayConfigurationV1 struct {
	// Gateway URL
	Url string `json:"url"`
	// Value of User-Agent header sent to CDL (to identify this client)
	UserAgent *string `json:"userAgent,omitempty"`
	// Set to true to ignore self-signed and other rejected certificates
	SkipCertCheck *bool `json:"skipCertCheck,omitempty"`
	// CA of CDL API gateway server in PEM format to use
	CaPath *string `json:"caPath,omitempty"`
	// Overwrite server name from cdlApiGateway.url to match one specified in CA
	ServerName *string `json:"serverName,omitempty"`
}

// NewGatewayConfigurationV1 instantiates a new GatewayConfigurationV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayConfigurationV1(url string) *GatewayConfigurationV1 {
	this := GatewayConfigurationV1{}
	this.Url = url
	var skipCertCheck bool = false
	this.SkipCertCheck = &skipCertCheck
	return &this
}

// NewGatewayConfigurationV1WithDefaults instantiates a new GatewayConfigurationV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayConfigurationV1WithDefaults() *GatewayConfigurationV1 {
	this := GatewayConfigurationV1{}
	var skipCertCheck bool = false
	this.SkipCertCheck = &skipCertCheck
	return &this
}

// GetUrl returns the Url field value
func (o *GatewayConfigurationV1) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *GatewayConfigurationV1) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *GatewayConfigurationV1) SetUrl(v string) {
	o.Url = v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *GatewayConfigurationV1) GetUserAgent() string {
	if o == nil || IsNil(o.UserAgent) {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayConfigurationV1) GetUserAgentOk() (*string, bool) {
	if o == nil || IsNil(o.UserAgent) {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *GatewayConfigurationV1) HasUserAgent() bool {
	if o != nil && !IsNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *GatewayConfigurationV1) SetUserAgent(v string) {
	o.UserAgent = &v
}

// GetSkipCertCheck returns the SkipCertCheck field value if set, zero value otherwise.
func (o *GatewayConfigurationV1) GetSkipCertCheck() bool {
	if o == nil || IsNil(o.SkipCertCheck) {
		var ret bool
		return ret
	}
	return *o.SkipCertCheck
}

// GetSkipCertCheckOk returns a tuple with the SkipCertCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayConfigurationV1) GetSkipCertCheckOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipCertCheck) {
		return nil, false
	}
	return o.SkipCertCheck, true
}

// HasSkipCertCheck returns a boolean if a field has been set.
func (o *GatewayConfigurationV1) HasSkipCertCheck() bool {
	if o != nil && !IsNil(o.SkipCertCheck) {
		return true
	}

	return false
}

// SetSkipCertCheck gets a reference to the given bool and assigns it to the SkipCertCheck field.
func (o *GatewayConfigurationV1) SetSkipCertCheck(v bool) {
	o.SkipCertCheck = &v
}

// GetCaPath returns the CaPath field value if set, zero value otherwise.
func (o *GatewayConfigurationV1) GetCaPath() string {
	if o == nil || IsNil(o.CaPath) {
		var ret string
		return ret
	}
	return *o.CaPath
}

// GetCaPathOk returns a tuple with the CaPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayConfigurationV1) GetCaPathOk() (*string, bool) {
	if o == nil || IsNil(o.CaPath) {
		return nil, false
	}
	return o.CaPath, true
}

// HasCaPath returns a boolean if a field has been set.
func (o *GatewayConfigurationV1) HasCaPath() bool {
	if o != nil && !IsNil(o.CaPath) {
		return true
	}

	return false
}

// SetCaPath gets a reference to the given string and assigns it to the CaPath field.
func (o *GatewayConfigurationV1) SetCaPath(v string) {
	o.CaPath = &v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *GatewayConfigurationV1) GetServerName() string {
	if o == nil || IsNil(o.ServerName) {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayConfigurationV1) GetServerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServerName) {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *GatewayConfigurationV1) HasServerName() bool {
	if o != nil && !IsNil(o.ServerName) {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *GatewayConfigurationV1) SetServerName(v string) {
	o.ServerName = &v
}

func (o GatewayConfigurationV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayConfigurationV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	if !IsNil(o.UserAgent) {
		toSerialize["userAgent"] = o.UserAgent
	}
	if !IsNil(o.SkipCertCheck) {
		toSerialize["skipCertCheck"] = o.SkipCertCheck
	}
	if !IsNil(o.CaPath) {
		toSerialize["caPath"] = o.CaPath
	}
	if !IsNil(o.ServerName) {
		toSerialize["serverName"] = o.ServerName
	}
	return toSerialize, nil
}

type NullableGatewayConfigurationV1 struct {
	value *GatewayConfigurationV1
	isSet bool
}

func (v NullableGatewayConfigurationV1) Get() *GatewayConfigurationV1 {
	return v.value
}

func (v *NullableGatewayConfigurationV1) Set(val *GatewayConfigurationV1) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayConfigurationV1) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayConfigurationV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayConfigurationV1(val *GatewayConfigurationV1) *NullableGatewayConfigurationV1 {
	return &NullableGatewayConfigurationV1{value: val, isSet: true}
}

func (v NullableGatewayConfigurationV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayConfigurationV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


