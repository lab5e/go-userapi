/*
 * The User API
 *
 * API to manage teams, members and tokens
 *
 * API version: 1.3.6 crooked-daija
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package userapi

import (
	"encoding/json"
)

// UserProfile Your user profile. You can change the contents of the user profile via the log in service you are using.
type UserProfile struct {
	UserId      *string `json:"userId,omitempty"`
	Email       *string `json:"email,omitempty"`
	AvatarUrl   *string `json:"avatarUrl,omitempty"`
	Name        *string `json:"name,omitempty"`
	ProfileUrl  *string `json:"profileUrl,omitempty"`
	GithubLogin *string `json:"githubLogin,omitempty"`
	Provider    *string `json:"provider,omitempty"`
	LogoutUrl   *string `json:"logoutUrl,omitempty"`
}

// NewUserProfile instantiates a new UserProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProfile() *UserProfile {
	this := UserProfile{}
	return &this
}

// NewUserProfileWithDefaults instantiates a new UserProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProfileWithDefaults() *UserProfile {
	this := UserProfile{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UserProfile) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UserProfile) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *UserProfile) SetUserId(v string) {
	o.UserId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserProfile) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserProfile) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserProfile) SetEmail(v string) {
	o.Email = &v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *UserProfile) GetAvatarUrl() string {
	if o == nil || o.AvatarUrl == nil {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetAvatarUrlOk() (*string, bool) {
	if o == nil || o.AvatarUrl == nil {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *UserProfile) HasAvatarUrl() bool {
	if o != nil && o.AvatarUrl != nil {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *UserProfile) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserProfile) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserProfile) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserProfile) SetName(v string) {
	o.Name = &v
}

// GetProfileUrl returns the ProfileUrl field value if set, zero value otherwise.
func (o *UserProfile) GetProfileUrl() string {
	if o == nil || o.ProfileUrl == nil {
		var ret string
		return ret
	}
	return *o.ProfileUrl
}

// GetProfileUrlOk returns a tuple with the ProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetProfileUrlOk() (*string, bool) {
	if o == nil || o.ProfileUrl == nil {
		return nil, false
	}
	return o.ProfileUrl, true
}

// HasProfileUrl returns a boolean if a field has been set.
func (o *UserProfile) HasProfileUrl() bool {
	if o != nil && o.ProfileUrl != nil {
		return true
	}

	return false
}

// SetProfileUrl gets a reference to the given string and assigns it to the ProfileUrl field.
func (o *UserProfile) SetProfileUrl(v string) {
	o.ProfileUrl = &v
}

// GetGithubLogin returns the GithubLogin field value if set, zero value otherwise.
func (o *UserProfile) GetGithubLogin() string {
	if o == nil || o.GithubLogin == nil {
		var ret string
		return ret
	}
	return *o.GithubLogin
}

// GetGithubLoginOk returns a tuple with the GithubLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetGithubLoginOk() (*string, bool) {
	if o == nil || o.GithubLogin == nil {
		return nil, false
	}
	return o.GithubLogin, true
}

// HasGithubLogin returns a boolean if a field has been set.
func (o *UserProfile) HasGithubLogin() bool {
	if o != nil && o.GithubLogin != nil {
		return true
	}

	return false
}

// SetGithubLogin gets a reference to the given string and assigns it to the GithubLogin field.
func (o *UserProfile) SetGithubLogin(v string) {
	o.GithubLogin = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *UserProfile) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *UserProfile) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *UserProfile) SetProvider(v string) {
	o.Provider = &v
}

// GetLogoutUrl returns the LogoutUrl field value if set, zero value otherwise.
func (o *UserProfile) GetLogoutUrl() string {
	if o == nil || o.LogoutUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoutUrl
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetLogoutUrlOk() (*string, bool) {
	if o == nil || o.LogoutUrl == nil {
		return nil, false
	}
	return o.LogoutUrl, true
}

// HasLogoutUrl returns a boolean if a field has been set.
func (o *UserProfile) HasLogoutUrl() bool {
	if o != nil && o.LogoutUrl != nil {
		return true
	}

	return false
}

// SetLogoutUrl gets a reference to the given string and assigns it to the LogoutUrl field.
func (o *UserProfile) SetLogoutUrl(v string) {
	o.LogoutUrl = &v
}

func (o UserProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.AvatarUrl != nil {
		toSerialize["avatarUrl"] = o.AvatarUrl
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProfileUrl != nil {
		toSerialize["profileUrl"] = o.ProfileUrl
	}
	if o.GithubLogin != nil {
		toSerialize["githubLogin"] = o.GithubLogin
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.LogoutUrl != nil {
		toSerialize["logoutUrl"] = o.LogoutUrl
	}
	return json.Marshal(toSerialize)
}

type NullableUserProfile struct {
	value *UserProfile
	isSet bool
}

func (v NullableUserProfile) Get() *UserProfile {
	return v.value
}

func (v *NullableUserProfile) Set(val *UserProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProfile(val *UserProfile) *NullableUserProfile {
	return &NullableUserProfile{value: val, isSet: true}
}

func (v NullableUserProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
