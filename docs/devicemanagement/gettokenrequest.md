# GetTokenRequest

The get token request details.

**Platforms:** iOS 17.0, iPadOS 17.0, macOS 14.0, visionOS 1.1

## Properties

### EnrollmentID

- **Type:** `string`
- **Required:** Yes

The per-enrollment identifier for the device. The system requires this value if the enrollment type is a user enrollment.

### EnrollmentUserID

- **Type:** `string`
- **Required:** Yes

The per-enrollment identifier for the user. The system requires this value if the enrollment type is a user enrollment on the user channel.

### MessageType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `GetToken`

The message type, which requires a value of `GetToken`.

### TokenParameters

- **Type:** `GetTokenRequest.TokenParameters`
- **Required:** No

Parameters that the system uses to generate the token.

### TokenServiceType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `com.apple.maid`, `com.apple.watch.pairing`

A string that specifies the service for the requested token.

### UDID

- **Type:** `string`
- **Required:** Yes

The device’s UDID (unique device identifier). The system requires this value if the enrollment type is a device enrollment.

### UserID

- **Type:** `string`
- **Required:** No

For macOS, this value is the ID of the user.

For Shared iPad, this value is `FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF` to indicate that authentication doesn’t occur.

### UserLongName

- **Type:** `string`
- **Required:** Yes

The full name of the user.

### UserShortName

- **Type:** `string`
- **Required:** No

For macOS, this value is the short name of the user.

For Shared iPad, this value is the Managed Apple Account identifier of the user. When present, it indicates that the token is for the user channel.

## Topics

### Objects

- [GetTokenRequest.TokenParameters](/documentation/devicemanagement/gettokenrequest/tokenparameters-data.dictionary) - Parameters that the system uses to generate the token.

