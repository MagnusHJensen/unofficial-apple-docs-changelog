# Sync the List of Devices

Get updates about the list of devices the server manages.

**Platforms:** Device Assignment Services , VPP License Management 

## Discussion

The sync service depends on a cursor returned by the fetch device service. It returns a list of all modifications (additions or deletions) since the specified cursor. The cursor passed to this endpoint should not be older than 7 days.

This service may return the same device more than once. You must resolve duplicates by matching on the device serial number and the `op_type` and `op_date` fields. The record with the latest `op_date` indicates the last known state of the device in DEP.

## Topics

### Request

- [SyncDeviceRequest](/documentation/devicemanagement/syncdevicerequest) - The request to sync the list of devices.

