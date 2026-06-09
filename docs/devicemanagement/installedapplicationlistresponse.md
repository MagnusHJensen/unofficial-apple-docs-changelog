# InstalledApplicationListResponse

A response from the device after it processes the command to get a list of the installed apps on a device.

**Platforms:** iOS 5.0, iPadOS 5.0, Mac Catalyst 5.0, macOS 10.7, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Properties

### CommandUUID

- **Type:** `string`
- **Required:** No

The unique identifier of the command for this response.

### EnrollmentID

- **Type:** `string`
- **Required:** Yes

The per-enrollment identifier for the device. The system requires this value if the enrollment type is a user enrollment.

Available: iOS 13+ | iPadOS 13+ | macOS 10.15+ | visionOS 1.1+

### EnrollmentUserID

- **Type:** `string`
- **Required:** Yes

The per-enrollment identifier for the user. The system requires this value if the enrollment type is a user enrollment on the user channel.

Available: macOS 10.15+

### ErrorChain

- **Type:** `[InstalledApplicationListResponse.ErrorChainItem]`
- **Required:** No

An array of dictionaries that describes any errors that occur.

### InstalledApplicationList

- **Type:** `[InstalledApplicationListResponse.InstalledApplicationListItem]`
- **Required:** Yes

An array of dictionaries that describes each installed app.

### NotOnConsole

- **Type:** `boolean`
- **Required:** Yes

If `true`, the device isn’t on-console.

### Status

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Acknowledged`, `Error`, `CommandFormatError`, `Idle`, `NotNow`

The status of the response, which is one of the following values:

- `Acknowledged`: The device processed the command successfully.
- `Error`: An error occurred. See the `ErrorChain` for more details.
- `CommandFormatError`: A protocol error occurred, which can result from a malformed command.
- `Idle`: The device is idle; there’s no status.
- `NotNow`: The device received the command, but can’t run it.

### UDID

- **Type:** `string`
- **Required:** Yes

The device’s UDID (unique device identifier). The system requires this value if the enrollment type is a device enrollment.

### UserID

- **Type:** `string`
- **Required:** No

For macOS, this value is the ID of the user.

For Shared iPad, this value is `FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF` to indicate that authentication doesn’t occur.

Available: iOS 9.3+ | iPadOS 9.3+ | macOS 10.7+

### UserLongName

- **Type:** `string`
- **Required:** Yes

The full name of the user.

Available: macOS 10.7+

### UserShortName

- **Type:** `string`
- **Required:** No

For macOS, this value is the short name of the user.

For Shared iPad, this value is the Managed Apple Account identifier of the user on Shared iPad. It indicates that the token is for the user channel.

Available: iOS 9.3+ | iPadOS 9.3+ | macOS 10.7+

## Topics

### Objects

- [InstalledApplicationListResponse.ErrorChainItem](/documentation/devicemanagement/installedapplicationlistresponse/errorchainitem) - A dictionary that describes an error chain item.
- [InstalledApplicationListResponse.InstalledApplicationListItem](/documentation/devicemanagement/installedapplicationlistresponse/installedapplicationlistitem) - A dictionary that describes an app list item.

