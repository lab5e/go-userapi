/*
 * The User API
 *
 * API to manage teams, members and tokens
 *
 * API version: 1.3.10 constant-champ
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package userapi

import (
	"encoding/json"
)

// TokenList struct for TokenList
type TokenList struct {
	Tokens *[]Token `json:"tokens,omitempty"`
}

// NewTokenList instantiates a new TokenList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenList() *TokenList {
	this := TokenList{}
	return &this
}

// NewTokenListWithDefaults instantiates a new TokenList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenListWithDefaults() *TokenList {
	this := TokenList{}
	return &this
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *TokenList) GetTokens() []Token {
	if o == nil || o.Tokens == nil {
		var ret []Token
		return ret
	}
	return *o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenList) GetTokensOk() (*[]Token, bool) {
	if o == nil || o.Tokens == nil {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *TokenList) HasTokens() bool {
	if o != nil && o.Tokens != nil {
		return true
	}

	return false
}

// SetTokens gets a reference to the given []Token and assigns it to the Tokens field.
func (o *TokenList) SetTokens(v []Token) {
	o.Tokens = &v
}

func (o TokenList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tokens != nil {
		toSerialize["tokens"] = o.Tokens
	}
	return json.Marshal(toSerialize)
}

type NullableTokenList struct {
	value *TokenList
	isSet bool
}

func (v NullableTokenList) Get() *TokenList {
	return v.value
}

func (v *NullableTokenList) Set(val *TokenList) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenList) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenList(val *TokenList) *NullableTokenList {
	return &NullableTokenList{value: val, isSet: true}
}

func (v NullableTokenList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
