# GetReplacementDetailsResponse

Information about a replacement device, including the original device it replaces and the date the replacement occurred.

## Properties

### serial_number

- **Type:** `string`
- **Required:** No

The serial number of the replacement device.

### original_device_serial_number

- **Type:** `string`
- **Required:** No

The serial number of the original device that this device replaces.

### replacement_date

- **Type:** `string`
- **Required:** No

The date when the device replacement occurred, in ISO 8601 format with day granularity in UTC (for example, `2025-01-15`).

