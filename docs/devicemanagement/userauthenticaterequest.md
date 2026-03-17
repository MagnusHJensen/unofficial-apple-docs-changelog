# UserAuthenticateRequest

The user authenticate request details.

**Platforms:** macOS 10.7

## Properties

### DigestResponse

- **Type:** `string`
- **Required:** Yes

A string that the client provides in the second [User Authenticate](/documentation/devicemanagement/user-authenticate) request after receiving `DigestChallenge` from the server on the first [User Authenticate](/documentation/devicemanagement/user-authenticate) request.

### MessageType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `UserAuthenticate`

The message type, which requires a value of `UserAuthenticate`.

### UDID

- **Type:** `string`
- **Required:** Yes

The device’s UDID (unique device identifier). The system requires this value if the enrollment type is a device enrollment.

### UserID

- **Type:** `string`
- **Required:** Yes

The local mobile user’s GUID or the network user’s GUID from an Open Directory record.

