# ImagePublic

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of this Image. | [optional] 
**Label** | **string** | A short description of the Image.  | [optional] 
**Created** | [**time.Time**](time.Time.md) | When this Image was created. | [optional] 
**CreatedBy** | **string** | The name of the User who created this Image, or \&quot;linode\&quot; for official Images.  | [optional] 
**Deprecated** | **bool** | Whether or not this Image is deprecated. Will only be true for deprecated public Images.  | [optional] 
**Description** | **string** | A detailed description of this Image. | [optional] 
**IsPublic** | **bool** | True if the Image is public. | [optional] 
**Size** | **int32** | The minimum size this Image needs to deploy. Size is in MB.  | [optional] 
**Type** | **string** | How the Image was created. Manual Images can be created at any time. \&quot;Automatic\&quot; Images are created automatically from a deleted Linode.  | [optional] 
**Expiry** | [**time.Time**](time.Time.md) | Only Images created automatically (from a deleted Linode; type&#x3D;automatic) will expire.  | [optional] 
**Vendor** | **string** | The upstream distribution vendor. &#x60;None&#x60; for private Images.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


