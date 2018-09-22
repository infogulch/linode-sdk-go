# StackScript

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique ID of this StackScript. | [optional] 
**Username** | **string** | The User who created the StackScript.  | [optional] 
**UserGravatarId** | **string** | The Gravatar ID for the User who created the StackScript.  | [optional] 
**Label** | **string** | The StackScript&#39;s label is for display purposes only.  | [optional] 
**Description** | **string** | A description for the StackScript.  | [optional] 
**Images** | **[]string** | An array of Image IDs representing the Images that this StackScript is compatible for deploying with.  | [optional] 
**DeploymentsTotal** | **int32** | The total number of times this StackScript has been deployed.  | [optional] 
**DeploymentsActive** | **int32** | Count of currently active, deployed Linodes created from this StackScript.  | [optional] 
**IsPublic** | **bool** | This determines whether other users can use your StackScript. **Once a StackScript is made public, it cannot be made private.**  | [optional] 
**Created** | [**time.Time**](time.Time.md) | The date this StackScript was created.  | [optional] 
**Updated** | [**time.Time**](time.Time.md) | The date this StackScript was last updated.  | [optional] 
**RevNote** | **string** | This field allows you to add notes for the set of revisions made to this StackScript.  | [optional] 
**Script** | **string** | The script to execute when provisioning a new Linode with this StackScript.  | [optional] 
**UserDefinedFields** | [**[]UserDefinedField**](UserDefinedField.md) | This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized parameters during deployment. See [Variables and UDFs](https://www.linode.com/docs/platform/stackscripts/#variables-and-udfs) for more information.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


