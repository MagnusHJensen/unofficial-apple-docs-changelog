# FetchDeviceResponse

The response that contains a list of devices.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### cursor

- **Type:** `string`
- **Required:** No

Indicates when this request was processed by the server. The MDM server can use this value in future requests if it wants to retrieve only records added or removed since this request.

### devices

- **Type:** `[Device]`
- **Required:** No

An array of dictionaries that provide information about devices. The devices are sorted in chronological order by the time stamp of the operation performed on the device.

### fetched_until

- **Type:** `string`
- **Required:** No

A date stamp that indicates the progress of the device fetch request, in ISO 8601 format.

### more_to_follow

- **Type:** `boolean`
- **Required:** No

Indicates that the request’s limit and cursor values resulted in only a partial list of devices.

The MDM server should immediately make another request (starting from the newly returned cursor) to obtain additional records.

