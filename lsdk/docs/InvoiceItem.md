# InvoiceItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The price, in US dollars, of unit price multiplied by quantity. | [optional] 
**From** | [**time.Time**](time.Time.md) | The date the Invoice Item started, based on month. | [optional] 
**Label** | **string** | The Invoice Item&#39;s display label. | [optional] 
**Quantity** | **int32** | The quantity of this Item for the specified Invoice. | [optional] 
**To** | [**time.Time**](time.Time.md) | The date the Invoice Item ended, based on month. | [optional] 
**Type** | **string** | The type of service, ether &#x60;prepay&#x60; or &#x60;misc&#x60;. | [optional] 
**Unitprice** | **int32** | The monthly service fee in US Dollars for this Item. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


