# SyncDeviceRequest

The request to sync the list of devices.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### cursor

- **Type:** `string`
- **Required:** No

A hex string that represents the starting position for a request. Use this to retrieve the list of devices that have been added or removed since a previous request. The string can be up to 1000 characters.

### limit

- **Type:** `int32`
- **Required:** No
- **Default:** `100`

The maximum number of entries to return. Optional.

