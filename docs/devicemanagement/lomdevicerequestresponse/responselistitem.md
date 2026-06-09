# LOMDeviceRequestResponse.ResponseListItem

A dictionary that describes a response list item.

**Platforms:** macOS 11.0

## Properties

### DeviceRequestReturnError

- **Type:** `string`
- **Required:** No

If present, a description of the error for a failed request.

### DeviceRequestSuccess

- **Type:** `boolean`
- **Required:** Yes

If `true`, the request was successful.

### DeviceRequestUUID

- **Type:** `string`
- **Required:** Yes

The unique identifier of the request for this response list item.

