# Go API client for userapi

API to manage teams, members and tokens

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.3.4 breezy-leafy
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://api.lab5e.com/user](https://api.lab5e.com/user)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./userapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.lab5e.com/user*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ProfileApi* | [**UserGetUserProfile**](docs/ProfileApi.md#usergetuserprofile) | **Get** /user/profile | Logged in profile
*TeamsApi* | [**UserAcceptInvite**](docs/TeamsApi.md#useracceptinvite) | **Post** /user/teams/accept | Accept invite
*TeamsApi* | [**UserCreateTeam**](docs/TeamsApi.md#usercreateteam) | **Post** /user/teams | Create team
*TeamsApi* | [**UserDeleteInvite**](docs/TeamsApi.md#userdeleteinvite) | **Delete** /user/teams/{teamId}/invites/{code} | Delete invite
*TeamsApi* | [**UserDeleteMember**](docs/TeamsApi.md#userdeletemember) | **Delete** /user/teams/{teamId}/members/{userId} | Remove member
*TeamsApi* | [**UserDeleteTeam**](docs/TeamsApi.md#userdeleteteam) | **Delete** /user/teams/{teamId} | Remove team
*TeamsApi* | [**UserGenerateInvite**](docs/TeamsApi.md#usergenerateinvite) | **Post** /user/teams/{teamId}/invites | Generate invite
*TeamsApi* | [**UserListInvites**](docs/TeamsApi.md#userlistinvites) | **Get** /user/teams/{teamId}/invites | List invites
*TeamsApi* | [**UserListTeams**](docs/TeamsApi.md#userlistteams) | **Get** /user/teams | List teams
*TeamsApi* | [**UserRetrieveInvite**](docs/TeamsApi.md#userretrieveinvite) | **Get** /user/teams/{teamId}/invites/{code} | Retrieve invite
*TeamsApi* | [**UserRetrieveMember**](docs/TeamsApi.md#userretrievemember) | **Get** /user/teams/{teamId}/members/{userId} | Retrieve member
*TeamsApi* | [**UserRetrieveTeam**](docs/TeamsApi.md#userretrieveteam) | **Get** /user/teams/{teamId} | Retrieve team
*TeamsApi* | [**UserRetrieveTeamMembers**](docs/TeamsApi.md#userretrieveteammembers) | **Get** /user/teams/{teamId}/members | List members
*TeamsApi* | [**UserUpdateMember**](docs/TeamsApi.md#userupdatemember) | **Patch** /user/teams/{teamId}/members/{userId} | Update member
*TeamsApi* | [**UserUpdateTeam**](docs/TeamsApi.md#userupdateteam) | **Patch** /user/teams/{teamId} | Update team
*TokensApi* | [**UserCreateToken**](docs/TokensApi.md#usercreatetoken) | **Post** /user/tokens | Create token
*TokensApi* | [**UserDeleteToken**](docs/TokensApi.md#userdeletetoken) | **Delete** /user/tokens/{token} | Remove token
*TokensApi* | [**UserListTokens**](docs/TokensApi.md#userlisttokens) | **Get** /user/tokens | List tokens
*TokensApi* | [**UserRetrieveToken**](docs/TokensApi.md#userretrievetoken) | **Get** /user/tokens/{token} | Retrieve token
*TokensApi* | [**UserUpdateToken**](docs/TokensApi.md#userupdatetoken) | **Patch** /user/tokens/{token} | Update token


## Documentation For Models

 - [AcceptInviteRequest](docs/AcceptInviteRequest.md)
 - [Invite](docs/Invite.md)
 - [InviteList](docs/InviteList.md)
 - [InviteRequest](docs/InviteRequest.md)
 - [Member](docs/Member.md)
 - [MemberList](docs/MemberList.md)
 - [ProtobufAny](docs/ProtobufAny.md)
 - [RpcStatus](docs/RpcStatus.md)
 - [Team](docs/Team.md)
 - [TeamList](docs/TeamList.md)
 - [Token](docs/Token.md)
 - [TokenList](docs/TokenList.md)
 - [UserProfile](docs/UserProfile.md)


## Documentation For Authorization



### APIToken

- **Type**: API key
- **API key parameter name**: X-API-Token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-API-Token and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

dev@lab5e.com

