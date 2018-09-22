# LinodeType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID representing the Linode Type. | [optional] 
**Label** | **string** | The Linode Type&#39;s label is for display purposes only.  | [optional] 
**Disk** | **int32** | The Disk size, in MB, of the Linode Type.  | [optional] 
**Class** | **string** | The class of the Linode Type. We currently offer three classes of Linodes:    * nanode - Our entry-level Linode Type.   * standard - Our flagship Linode Type.   * highmem - A Linode Type featuring high memory availability.  | [optional] 
**Price** | [**LinodeTypePrice**](LinodeType_price.md) |  | [optional] 
**Addons** | [**LinodeTypeAddons**](LinodeType_addons.md) |  | [optional] 
**NetworkOut** | **int32** | The Mbits outbound bandwidth allocation.  | [optional] 
**Memory** | **int32** | Amount of RAM included in this Linode Type.  | [optional] 
**Successor** | **string** | The Linode Type that a [mutate](#operation/mutateLinodeInstance) will upgrade to for a Linode of this type.  If \&quot;null\&quot;, a Linode of this type may not mutate.  | [optional] 
**Transfer** | **int32** | The monthly outbound transfer amount, in MB.  | [optional] 
**Vcpus** | **int32** | The number of VCPU cores this Linode Type offers.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


