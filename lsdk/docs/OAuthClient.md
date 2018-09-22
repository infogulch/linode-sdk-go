# OAuthClient

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The OAuth Client ID.  This is used to identify the client, and is a publicly-known value (it is not a secret).  | [optional] 
**RedirectUri** | **string** | The location a successful log in from https://login.linode.com should be redirected to for this client.  The receiver of this redirect should be ready to accept an OAuth exchange code and finish the OAuth exchange.  | [optional] 
**Label** | **string** | The name of this application.  This will be presented to users when they are asked to grant it access to their Account.  | [optional] 
**Status** | **string** | The status of this application.  &#x60;active&#x60; by default.  | [optional] 
**Secret** | **string** | The OAuth Client secret, used in the OAuth exchange.  This is returned as &#x60;&lt;REDACTED&gt;&#x60; except when an OAuth Client is created or its secret is reset.  This is a secret, and should not be shared or disclosed publicly.  | [optional] 
**ThumbnailUrl** | **string** | The URL where this client&#39;s thumbnail may be viewed, or &#x60;null&#x60; if this client does not have a thumbnail set.  | [optional] 
**Public** | **bool** | If this is a public or private OAuth Client.  Public clients have a slightly different authentication workflow than private clients.  See the [OAuth spec](https://oauth.net/2/) for more details.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


