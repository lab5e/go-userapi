/*
 * The User API
 *
 * API to manage teams, members and tokens
 *
 * API version: 1.3.11 lucky-fremont
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package userapi

import (
	"encoding/json"
)

// Token struct for Token
type Token struct {
	// The resource of the token.  The token applies to the specified resource and any resources below this so the resource `/` applies to the root resource and any resource below the root resource. In the same manner `/collections` will apply to all collectins while `/collections/{id}` will apply to that particular collection.
	Resource *string `json:"resource,omitempty"`
	// Write flag for token.  If this is set to `true` the token can be used for write operations in the API such as POST, DELETE and PATCH.
	Write *bool `json:"write,omitempty"`
	// Use this in the `X-API-Token` header when using the API.
	Token *string `json:"token,omitempty"`
	// Tags for the token.
	Tags *map[string]string `json:"tags,omitempty"`
	// The last time the token was used. Time in ms since epoch.
	LastUse *string `json:"lastUse,omitempty"`
	Uses    *string `json:"uses,omitempty"`
}

// NewToken instantiates a new Token object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToken() *Token {
	this := Token{}
	return &this
}

// NewTokenWithDefaults instantiates a new Token object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenWithDefaults() *Token {
	this := Token{}
	return &this
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *Token) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *Token) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *Token) SetResource(v string) {
	o.Resource = &v
}

// GetWrite returns the Write field value if set, zero value otherwise.
func (o *Token) GetWrite() bool {
	if o == nil || o.Write == nil {
		var ret bool
		return ret
	}
	return *o.Write
}

// GetWriteOk returns a tuple with the Write field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetWriteOk() (*bool, bool) {
	if o == nil || o.Write == nil {
		return nil, false
	}
	return o.Write, true
}

// HasWrite returns a boolean if a field has been set.
func (o *Token) HasWrite() bool {
	if o != nil && o.Write != nil {
		return true
	}

	return false
}

// SetWrite gets a reference to the given bool and assigns it to the Write field.
func (o *Token) SetWrite(v bool) {
	o.Write = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Token) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Token) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Token) SetToken(v string) {
	o.Token = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Token) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Token) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Token) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetLastUse returns the LastUse field value if set, zero value otherwise.
func (o *Token) GetLastUse() string {
	if o == nil || o.LastUse == nil {
		var ret string
		return ret
	}
	return *o.LastUse
}

// GetLastUseOk returns a tuple with the LastUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetLastUseOk() (*string, bool) {
	if o == nil || o.LastUse == nil {
		return nil, false
	}
	return o.LastUse, true
}

// HasLastUse returns a boolean if a field has been set.
func (o *Token) HasLastUse() bool {
	if o != nil && o.LastUse != nil {
		return true
	}

	return false
}

// SetLastUse gets a reference to the given string and assigns it to the LastUse field.
func (o *Token) SetLastUse(v string) {
	o.LastUse = &v
}

// GetUses returns the Uses field value if set, zero value otherwise.
func (o *Token) GetUses() string {
	if o == nil || o.Uses == nil {
		var ret string
		return ret
	}
	return *o.Uses
}

// GetUsesOk returns a tuple with the Uses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetUsesOk() (*string, bool) {
	if o == nil || o.Uses == nil {
		return nil, false
	}
	return o.Uses, true
}

// HasUses returns a boolean if a field has been set.
func (o *Token) HasUses() bool {
	if o != nil && o.Uses != nil {
		return true
	}

	return false
}

// SetUses gets a reference to the given string and assigns it to the Uses field.
func (o *Token) SetUses(v string) {
	o.Uses = &v
}

func (o Token) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	if o.Write != nil {
		toSerialize["write"] = o.Write
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.LastUse != nil {
		toSerialize["lastUse"] = o.LastUse
	}
	if o.Uses != nil {
		toSerialize["uses"] = o.Uses
	}
	return json.Marshal(toSerialize)
}

type NullableToken struct {
	value *Token
	isSet bool
}

func (v NullableToken) Get() *Token {
	return v.value
}

func (v *NullableToken) Set(val *Token) {
	v.value = val
	v.isSet = true
}

func (v NullableToken) IsSet() bool {
	return v.isSet
}

func (v *NullableToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToken(val *Token) *NullableToken {
	return &NullableToken{value: val, isSet: true}
}

func (v NullableToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
