# \ProfileApi

All URIs are relative to *https://api.linode.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSSHKey**](ProfileApi.md#AddSSHKey) | **Post** /profile/sshkeys | Add SSH Key
[**CreatePersonalAccessToken**](ProfileApi.md#CreatePersonalAccessToken) | **Post** /profile/tokens | Create Personal Access Token
[**DeletePersonalAccessToken**](ProfileApi.md#DeletePersonalAccessToken) | **Delete** /profile/tokens/{tokenId} | Revoke Personal Access Token
[**DeleteProfileApp**](ProfileApi.md#DeleteProfileApp) | **Delete** /profile/apps/{appId} | Revoke App Access
[**DeleteSSHKey**](ProfileApi.md#DeleteSSHKey) | **Delete** /profile/sshkeys/{sshKeyId} | Delete SSH Key
[**GetPersonalAccessToken**](ProfileApi.md#GetPersonalAccessToken) | **Get** /profile/tokens/{tokenId} | View Personal Access Token
[**GetPersonalAccessTokens**](ProfileApi.md#GetPersonalAccessTokens) | **Get** /profile/tokens | List Personal Access Tokens
[**GetProfile**](ProfileApi.md#GetProfile) | **Get** /profile | View Profile
[**GetProfileApp**](ProfileApi.md#GetProfileApp) | **Get** /profile/apps/{appId} | View Authorized App
[**GetProfileApps**](ProfileApi.md#GetProfileApps) | **Get** /profile/apps | List Authorized Apps
[**GetProfileGrants**](ProfileApi.md#GetProfileGrants) | **Get** /profile/grants | List Grants
[**GetSSHKey**](ProfileApi.md#GetSSHKey) | **Get** /profile/sshkeys/{sshKeyId} | View SSH Key
[**GetSSHKeys**](ProfileApi.md#GetSSHKeys) | **Get** /profile/sshkeys | List SSH Keys
[**TfaConfirm**](ProfileApi.md#TfaConfirm) | **Post** /profile/tfa-enable-confirm | Confirm/Enable Two Factor Authentication
[**TfaDisable**](ProfileApi.md#TfaDisable) | **Post** /profile/tfa-disable | Disable Two Factor Authentication
[**TfaEnable**](ProfileApi.md#TfaEnable) | **Post** /profile/tfa-enable | Create Two Factor Secret
[**UpdatePersonalAccessToken**](ProfileApi.md#UpdatePersonalAccessToken) | **Put** /profile/tokens/{tokenId} | Update Personal Access Token
[**UpdateProfile**](ProfileApi.md#UpdateProfile) | **Put** /profile | Update Profile
[**UpdateSSHKey**](ProfileApi.md#UpdateSSHKey) | **Put** /profile/sshkeys/{sshKeyId} | Update SSH Key


# **AddSSHKey**
> SshKey AddSSHKey(ctx, optional)
Add SSH Key

Adds an SSH Key to your Account profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddSSHKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddSSHKeyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshKeyRequest** | [**optional.Interface of SshKeyRequest**](SshKeyRequest.md)| Add SSH Key | 

### Return type

[**SshKey**](SSHKey.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePersonalAccessToken**
> PersonalAccessToken CreatePersonalAccessToken(ctx, uNKNOWNBASETYPE)
Create Personal Access Token

Creates a Personal Access Token for your User. The raw token will be returned in the response, but will never be returned again afterward so be sure to take note of it. You may create a token with _at most_ the scopes of your current token. The created token will be able to access your Account until the given expiry, or until it is revoked. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| Information about the requested token. | 

### Return type

[**PersonalAccessToken**](PersonalAccessToken.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePersonalAccessToken**
> map[string]interface{} DeletePersonalAccessToken(ctx, tokenId)
Revoke Personal Access Token

Revokes a Personal Access Token. The token will be invalidated immediately, and requests using that token will fail with a 401. It is possible to revoke access to the token making the request to revoke a token, but keep in mind that doing so could lose you access to the api and require you to create a new token through some other means. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenId** | **int32**| The ID of the token to access. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProfileApp**
> map[string]interface{} DeleteProfileApp(ctx, appId)
Revoke App Access

Expires this app token. This token may no longer be used to access your Account. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| The authorized app ID to manage. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSSHKey**
> map[string]interface{} DeleteSSHKey(ctx, sshKeyId)
Delete SSH Key

Deletes an SSH Key you have access to.  **Note:** deleting an SSH Key will *not* remove it from any Linode or Disk that was deployed with `authorized_keys`. In those cases, the keys must be manually deleted on the Linode or Disk. This endpoint will only delete the key's association from your Profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sshKeyId** | **int32**| The ID of the SSHKey | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPersonalAccessToken**
> PersonalAccessToken GetPersonalAccessToken(ctx, tokenId)
View Personal Access Token

Returns a single Personal Access Token. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenId** | **int32**| The ID of the token to access. | 

### Return type

[**PersonalAccessToken**](PersonalAccessToken.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPersonalAccessTokens**
> InlineResponse20033 GetPersonalAccessTokens(ctx, )
List Personal Access Tokens

Returns a paginated list of Personal Access Tokens currently active for your User. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProfile**
> Profile GetProfile(ctx, )
View Profile

Returns information about the current User. This can be used to see who is acting in applications where more than one token is managed. For example, in third-party OAuth applications.  This endpoint is always accessible, no matter what OAuth scopes the acting token has. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Profile**](Profile.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProfileApp**
> AuthorizedApp GetProfileApp(ctx, appId)
View Authorized App

Returns information about a single app you've authorized to access your Account. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int32**| The authorized app ID to manage. | 

### Return type

[**AuthorizedApp**](AuthorizedApp.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProfileApps**
> InlineResponse20032 GetProfileApps(ctx, optional)
List Authorized Apps

This is a collection of OAuth apps that you've given access to your Account, and includes the level of access granted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetProfileAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetProfileAppsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| The page of a collection to return. | [default to 1]
 **pageSize** | **optional.Int32**| The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProfileGrants**
> GrantsResponse GetProfileGrants(ctx, )
List Grants

This returns a GrantsResponse describing what the acting User has been granted access to.  For unrestricted users, this will return a  204 and no body because unrestricted users have access to everything without grants.  This will not return information about entities you do not have access to.  This endpoint is useful when writing third-party OAuth applications to see what options you should present to the acting User.  For example, if they do not have `global.add_linodes`, you might not display a button to deploy a new Linode.  Any client may access this endpoint; no OAuth scopes are required. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GrantsResponse**](GrantsResponse.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSSHKey**
> SshKey GetSSHKey(ctx, sshKeyId)
View SSH Key

Returns a single SSH Key object identified by `id` that you have access to view. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sshKeyId** | **int32**| The ID of the SSHKey | 

### Return type

[**SshKey**](SSHKey.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSSHKeys**
> InlineResponse20034 GetSSHKeys(ctx, optional)
List SSH Keys

Returns a collection of SSH Keys you've added to your Profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSSHKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSSHKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| The page of a collection to return. | [default to 1]
 **pageSize** | **optional.Int32**| The number of items to return per page. | [default to 100]

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TfaConfirm**
> map[string]interface{} TfaConfirm(ctx, uNKNOWNBASETYPE)
Confirm/Enable Two Factor Authentication

Confirms that you can successfully generate Two Factor codes and enables TFA on your Account. Once this is complete, login attempts from untrusted computers will be required to provide a Two Factor code before they are successful. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| The Two Factor code you generated with your Two Factor secret. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TfaDisable**
> map[string]interface{} TfaDisable(ctx, )
Disable Two Factor Authentication

Disables Two Factor Authentication for your User. Once successful, login attempts from untrusted computers will only require a password before being successful. This is less secure, and is discouraged. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TfaEnable**
> map[string]interface{} TfaEnable(ctx, )
Create Two Factor Secret

Generates a Two Factor secret for your User. TFA will not be enabled until you have successfully confirmed the code you were given with [tfa-enable-confirm](/#operation/tfaConfirm) (see below). Once enabled, logins from untrusted computers will be required to provide a TFA code before they are successful. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePersonalAccessToken**
> PersonalAccessToken UpdatePersonalAccessToken(ctx, tokenId, personalAccessToken)
Update Personal Access Token

Updates a Personal Access Token. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenId** | **int32**| The ID of the token to access. | 
  **personalAccessToken** | [**PersonalAccessToken**](PersonalAccessToken.md)| The fields to update. | 

### Return type

[**PersonalAccessToken**](PersonalAccessToken.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProfile**
> Profile UpdateProfile(ctx, profile)
Update Profile

Update information in your Profile.  This endpoint requires the \"account:read_write\" OAuth Scope. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profile** | [**Profile**](Profile.md)| The fields to update. | 

### Return type

[**Profile**](Profile.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSSHKey**
> SshKey UpdateSSHKey(ctx, sshKeyId, sshKey)
Update SSH Key

Updates an SSH Key that you have permission to `read_write`. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sshKeyId** | **int32**| The ID of the SSHKey | 
  **sshKey** | [**SshKey**](SshKey.md)| The fields to update.  | 

### Return type

[**SshKey**](SSHKey.md)

### Authorization

[oauth](../README.md#oauth), [personalAccessToken](../README.md#personalAccessToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

