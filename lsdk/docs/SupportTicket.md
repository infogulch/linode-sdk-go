# SupportTicket

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the Support Ticket.  | [optional] 
**Attachments** | **[]string** | A list of filenames representing attached files associated with this Ticket.  | [optional] 
**Closed** | [**time.Time**](time.Time.md) | The date and time this Ticket was closed.  | [optional] 
**Closable** | **bool** | Whether the Support Ticket may be closed.  | [optional] 
**Description** | **string** | The full details of the issue or question.  | [optional] 
**Entity** | [**SupportTicketEntity**](SupportTicket_entity.md) |  | [optional] 
**GravatarId** | **string** | The Gravatar ID of the User who opened this Ticket.  | [optional] 
**Opened** | [**time.Time**](time.Time.md) | The date and time this Ticket was created.  | [optional] 
**OpenedBy** | **string** | The User who opened this Ticket.  | [optional] 
**Status** | **string** | The current status of this Ticket. | [optional] 
**Summary** | **string** | The summary or title for this Ticket.  | [optional] 
**Updated** | [**time.Time**](time.Time.md) | The date and time this Ticket was last updated.  | [optional] 
**UpdatedBy** | **string** | The User who last updated this Ticket.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


