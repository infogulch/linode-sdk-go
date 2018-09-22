# User

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | This User&#39;s username. This is used for logging in, and may also be displayed alongside actions the User performs (for example, in Events or public StackScripts).  | [optional] 
**Email** | **string** | The email address for this User, for account management communications, and may be used for other communications as configured.  | [optional] 
**Restricted** | **bool** | If true, this User must be granted access to perform actions or access entities on this Account. See [/account/users/{username}/grants](/#operation/getUserGrants) for details on how to configure grants for a restricted User.  | [optional] 
**SshKeys** | **[]string** | A list of SSH Key labels added by this User. These are the keys that will be deployed if this User is included in the &#x60;authorized_users&#x60; field of a [create Linode](/#operation/createLinodeInstance), [rebuild Linode](/#operation/rebuildLinodeInstance), or [create Disk](/#operation/addLinodeDisk) request.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


