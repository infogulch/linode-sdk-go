# Profile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **int32** | Your unique ID in our system. This value will never change, and can safely be used to identify your User.  | [optional] 
**Username** | **string** | Your username, used for logging in to our system.  | [optional] 
**Email** | **string** | Your email address.  This address will be used for communication with Linode as necessary.  | [optional] 
**Timezone** | **string** | The timezone you prefer to see times in.  This is not used by the API, and is for the benefit of clients only.  All times the API returns are in UTC.  | [optional] 
**EmailNotifications** | **bool** | If true, you will receive email notifications about account activity.  If false, you may still receive business-critical communications through email.  | [optional] 
**Referrals** | [**ProfileReferrals**](Profile_referrals.md) |  | [optional] 
**IpWhitelistEnabled** | **bool** | If true, logins for your User will only be allowed from whitelisted IPs. This setting is currently deprecated, and cannot be enabled.  If you disable this setting, you will not be able to re-enable it.  | [optional] 
**LishAuthMethod** | **string** | What methods of authentication are allowed when connecting via Lish.  \&quot;keys_only\&quot; is the most secure if you intend to use Lish, and \&quot;disabled\&quot; is recommended if you do not intend to use Lish at all.  | [optional] 
**AuthorizedKeys** | **[]string** | The list of SSH Keys authorized to use Lish for your User. This value is ignored if &#x60;lish_auth_method&#x60; is \&quot;disabled.\&quot;  | [optional] 
**TwoFactorAuth** | **bool** | If true, logins from untrusted computers will require Two Factor Authentication.  See [/profile/tfa-enable](/#operation/tfaEnable) to enable Two Factor Authentication.  | [optional] 
**Restricted** | **bool** | If true, your User has restrictions on what can be accessed on your Account. To get details on what entities/actions you can access/perform, see [/profile/grants](/#operation/getProfileGrants).  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


