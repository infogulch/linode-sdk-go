# LinodeSpecs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disk** | **int32** | The amount of storage space, in GB. this Linode has access to. A typical Linode will divide this space between a primary disk with an &#x60;image&#x60; deployed to it, and a swap disk, usually 512 MB. This is the default configuration created when deploying a Linode with an &#x60;image&#x60; through [POST /linode/instances](/#operation/createLinodeInstance). While this configuration is suitable for 99% of use cases, if you need finer control over your Linode&#39;s disks, see the [/linode/instances/{linodeId}/disks](/#operation/getLinodeDisks) endpoints.  | [optional] 
**Memory** | **int32** | The amount of RAM, in MB, this Linode has access to. Typically a Linode will choose to boot with all of its available RAM, but this can be configured in a Config profile, see the [/linode/instances/{linodeId}/configs](/#operation/getLinodeConfigs) endpoints and the LinodeConfig object for more information.  | [optional] 
**Vcpus** | **int32** | The number of vcpus this Linode has access to.  Typically a Linode will choose to boot with all of its available vcpus, but this can be configured in a Config Profile, see the [/linode/instances/{linodeId}/configs](/#operation/getLinodeConfigs) endpoints and the LinodeConfig object for more information.  | [optional] 
**Transfer** | **int32** | The amount of network transfer this Linode is allotted each month. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


