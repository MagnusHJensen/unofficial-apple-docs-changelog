# OSUpdateStatusResponse

A response from the device after it processes the command to get the status of operating-system updates on a device.

**Platforms:** iOS 9.0, iPadOS 9.0, macOS 10.11.5, tvOS 12.0

## Properties

### CommandUUID

- **Type:** `string`
- **Required:** No

The unique identifier of the command for this response.

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

### ErrorChain

- **Type:** `[OSUpdateStatusResponse.ErrorChainItem]`
- **Required:** No

An array of dictionaries that describes any errors that occur.

### NotOnConsole

- **Type:** `boolean`
- **Required:** Yes

If `true`, the device isn’t on-console.

### OSUpdateStatus

- **Type:** `[OSUpdateStatusResponse.OSUpdateStatusItem]`
- **Required:** Yes

An array of dictionaries that describes the statuses of software updates. The array is empty if there are no software updates currently in progress. This command only returns the status for System Applications and Configuration Data updates when a software update is managed by a Declarative Device Management [SoftwareUpdateEnforcementSpecific](/documentation/devicemanagement/softwareupdateenforcementspecific) configuration.

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

### UserLongName

- **Type:** `string`
- **Required:** Yes

The full name of the user.

### UserShortName

- **Type:** `string`
- **Required:** No

For macOS, this value is the short name of the user.

For Shared iPad, this value is the Managed Apple Account identifier of the user on Shared iPad. It indicates that the token is for the user channel.

## Topics

### Objects

- [OSUpdateStatusResponse.ErrorChainItem](/documentation/devicemanagement/osupdatestatusresponse/errorchainitem) - A dictionary that describes an error chain item.
- [OSUpdateStatusResponse.OSUpdateStatusItem](/documentation/devicemanagement/osupdatestatusresponse/osupdatestatusitem) - A dictionary that describes the status of a software update.

