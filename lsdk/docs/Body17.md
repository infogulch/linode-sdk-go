# Body17

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | The ID of the Region to create this NodeBalancer in.  | 
**Label** | [**Label**](label.md) |  | [optional] 
**ClientConnThrottle** | [**ClientConnThrottle**](client_conn_throttle.md) |  | [optional] 
**Configs** | [**[]NodeBalancerConfig**](NodeBalancerConfig.md) | The ports to configure this NodeBalancer with on creation. Each config must have a unique port and at least one Node.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


