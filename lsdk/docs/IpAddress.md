# IpAddress

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The IP address.  | [optional] 
**Gateway** | **string** | The default gateway for this address.  | [optional] 
**SubnetMask** | **string** | The mask that separates host bits from network bits for this address.  | [optional] 
**Prefix** | **int32** | The number of bits set in the subnet mask.  | [optional] 
**Type** | **string** | The type of address this is.  | [optional] 
**Public** | **bool** | Whether this is a public or private IP address.  | [optional] 
**Rdns** | **string** | The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if not explicitly set.  | [optional] 
**LinodeId** | **int32** | The ID of the Linode this address currently belongs to. For IPv4 addresses, this is by default the Linode that this address was assigned to on creation, and these addresses my be moved using the [/networking/ipv4/assign](/#operation/assignIPs) endpoint. For SLAAC and link-local addresses, this value may not be changed.  | [optional] 
**Region** | **string** | The Region this IP address resides in.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


