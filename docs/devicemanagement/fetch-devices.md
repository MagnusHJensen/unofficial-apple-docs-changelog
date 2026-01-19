# Get a List of Devices

Get a list of devices that are managed by the server.

**Platforms:** Device Assignment Services , VPP License Management 

## Discussion

This request fetches a list of all devices that are assigned to this MDM server at the time of the request. This service should be used for loading an initial list of devices into the MDM server’s data store. Once the list of devices is loaded, [Sync the List of Devices](/documentation/devicemanagement/sync-devices) to update the list.

This request provides a limited number of entries per request, using cursors to provide position information across requests.

## Topics

### Request and Response

- [FetchDeviceRequest](/documentation/devicemanagement/fetchdevicerequest) - The request for a list of devices.
- [FetchDeviceResponse](/documentation/devicemanagement/fetchdeviceresponse) - The response that contains a list of devices.

