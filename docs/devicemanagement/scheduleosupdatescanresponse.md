# ScheduleOSUpdateScanResponse

A response from the device after it processes the command to schedule a background scan for operating-system updates on a device. Removed: use the declarative management `com.apple.configuration.softwareupdate.enforcement.specific` configuration.

**Platforms:** macOS 10.11

## Properties

### CommandUUID

- **Type:** `string`
- **Required:** No

The unique identifier of the command for this response.

### EnrollmentID

- **Type:** `string`
- **Required:** Yes





Removed: macOS 27+

### EnrollmentUserID

- **Type:** `string`
- **Required:** Yes





Removed: macOS 27+

### ErrorChain

- **Type:** `[ScheduleOSUpdateScanResponse.ErrorChainItem]`
- **Required:** No

An array of dictionaries that describes any errors that occur.

### NotOnConsole

- **Type:** `boolean`
- **Required:** Yes

If `true`, the device isn’t on-console.

### ScanInitiated

- **Type:** `boolean`
- **Required:** Yes




Removed: macOS 27+

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




Removed: macOS 27+

### UserID

- **Type:** `string`
- **Required:** No







Removed: macOS 27+

### UserLongName

- **Type:** `string`
- **Required:** Yes




Removed: macOS 27+

### UserShortName

- **Type:** `string`
- **Required:** No







Removed: macOS 27+

## Topics

### Objects

- [ScheduleOSUpdateScanResponse.ErrorChainItem](/documentation/devicemanagement/scheduleosupdatescanresponse/errorchainitem) - A dictionary that describes an error chain item.

