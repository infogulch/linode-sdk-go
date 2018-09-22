# Event

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique ID of this Event. | [optional] 
**Action** | **string** | The action that caused this Event. New actions may be added in the future.  | [optional] 
**Created** | [**time.Time**](time.Time.md) | When this Event was created. | [optional] 
**Entity** | [**EventEntity**](Event_entity.md) |  | [optional] 
**PercentComplete** | **int32** | A percentage estimating the amount of time remaining for an Event. Returns &#x60;null&#x60; for notification events.  | [optional] 
**Rate** | **string** | The rate of completion of the Event. Only some Events will return rate; for example, migration and resize Events.  | [optional] 
**Read** | **bool** | If this Event has been read. | [optional] 
**Seen** | **bool** | If this Event has been seen. | [optional] 
**Status** | **string** | The current status of this Event. | [optional] 
**TimeRemaining** | **int32** | The estimated time remaining until the completion of this Event. This value is only returned for in-progress events.  | [optional] 
**Username** | **string** | The username of the User who caused the Event.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


