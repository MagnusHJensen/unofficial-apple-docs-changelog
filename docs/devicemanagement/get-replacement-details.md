# Get Replacement Details

Get information about the device that a replacement device replaces.

## Discussion

This endpoint returns the original device serial number and replacement date for a replacement device. Call this endpoint only when the device’s `is_replacement_device` field is `true` in the response from [Get a List of Devices](/documentation/devicemanagement/fetch-devices), [Sync the List of Devices](/documentation/devicemanagement/sync-devices), or [Get Device Details](/documentation/devicemanagement/device-details). Calling it for a device whose `is_replacement_device` value is `false` returns HTTP 404 `DEVICE_NOT_FOUND`.

This request requires `X-Server-Protocol-Version` 10 or later.

## Topics

### Response

- [GetReplacementDetailsResponse](/documentation/devicemanagement/getreplacementdetailsresponse) - Information about a replacement device, including the original device it replaces and the date the replacement occurred.

