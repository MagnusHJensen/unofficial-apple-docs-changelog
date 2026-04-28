# TokenUpdateRequest

The token update request details.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, tvOS 10.2, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### AwaitingConfiguration

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true` from the device channel, the device is awaiting a [Device Configured](/documentation/devicemanagement/device-configured-command) command before proceeding through Setup Assistant.

If `true` from the user channel (Shared iPad only), the device is awaiting a [User Configured](/documentation/devicemanagement/user-configured-command) command before proceeding through Setup Assistant.

### EnrollmentID

- **Type:** `string`
- **Required:** Yes

The per-enrollment identifier for the device. The system requires this value if the enrollment type is a user enrollment.

Available in iOS 13 and later, macOS 10.15 and later, and visionOS 2 and later.

### EnrollmentUserID

- **Type:** `string`
- **Required:** Yes

The per-enrollment identifier for the user. The system requires this value if the enrollment type is a user enrollment on the user channel.

Available in macOS 10.15 and later.

### MessageType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `TokenUpdate`

The message type, which requires a value of `TokenUpdate`.

### NotOnConsole

- **Type:** `boolean`
- **Required:** Yes

If `true`, the device isn’t on-console.

### PushMagic

- **Type:** `string`
- **Required:** Yes

The magic string to include in the push notification message.

### Token

- **Type:** `data`
- **Required:** Yes

The push token for the device.

### Topic

- **Type:** `string`
- **Required:** Yes

The topic the device subscribes to.

### UDID

- **Type:** `string`
- **Required:** Yes

The device’s UDID (unique device identifier). The system requires this value if the enrollment type is a device enrollment.

### UnlockToken

- **Type:** `data`
- **Required:** No

The data to use to unlock the device. If provided, the server needs to retain this data and send it when trying to implement [Clear Passcode](/documentation/devicemanagement/clear-passcode-command).

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

For Shared iPad, this value is the Managed Apple Account identifier of the user on Shared iPad. It indicates that the token is for the user channel.

