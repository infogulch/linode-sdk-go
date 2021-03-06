# Backup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique ID of this Backup. | [optional] 
**Type** | **string** | This indicates whether the Backup is an automatic Backup or manual snapshot taken by the User at a specific point in time.  | [optional] 
**Status** | **string** | The current state of a specific Backup. | [optional] 
**Created** | [**time.Time**](time.Time.md) | The date the Backup was taken. | [optional] 
**Updated** | [**time.Time**](time.Time.md) | The date the Backup was most recently updated. | [optional] 
**Finished** | [**time.Time**](time.Time.md) | The date the Backup completed. | [optional] 
**Label** | **string** | A label for Backups that are of type &#x60;snapshot&#x60;. | [optional] 
**Configs** | **[]string** | A list of the labels of the Configuration profiles that are part of the Backup.  | [optional] 
**Disks** | [**[]BackupDisks**](Backup_disks.md) | A list of the disks that are part of the Backup.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


