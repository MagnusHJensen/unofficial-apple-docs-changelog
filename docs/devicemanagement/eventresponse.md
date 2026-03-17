# EventResponse

The response that contains the event identifier.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### eventId

- **Type:** `string`
- **Required:** No

The unique identifier for the asynchronous event.

### mdmInfo

- **Type:** `MdmInfo`
- **Required:** No

The current information for the provided token.

The response only includes this field when MDM sets a value using the [Client Config](/documentation/devicemanagement/client-config-4szk1) endpoint.

### tokenExpirationDate

- **Type:** `string`
- **Required:** No

The token expiration date in an ISO-8601 format.

Note: The server shows all dates and times in UTC.

### uId

- **Type:** `string`
- **Required:** No

The unique library identifier. When querying records using multiple tokens that may share libraries, use the `uId` field to filter duplicates and avoid double-counting records when different content managers upload duplicate tokens.

## Topics

### Objects and Data Types

- [MdmInfo](/documentation/devicemanagement/mdminfo) - Information about the MDM client.

