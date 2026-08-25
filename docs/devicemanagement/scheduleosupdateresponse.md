# ScheduleOSUpdateResponse

A response from the device after it processes the command to schedule an update of the operating system on a device. Removed: use the declarative management `com.apple.configuration.softwareupdate.enforcement.specific` configuration.

**Platforms:** iOS 9.0, iPadOS 9.0, Mac Catalyst 9.0, macOS 10.11, tvOS 12.0

## Properties

### CommandUUID

- **Type:** `string`
- **Required:** No

The unique identifier of the command for this response.

### EnrollmentID

- **Type:** `string`
- **Required:** Yes





Removed: iOS 27+ | iPadOS 27+ | macOS 27+ | tvOS 27+

### EnrollmentUserID

- **Type:** `string`
- **Required:** Yes





Removed: iOS 27+ | iPadOS 27+ | macOS 27+ | tvOS 27+

### ErrorChain

- **Type:** `[ScheduleOSUpdateResponse.ErrorChainItem]`
- **Required:** No

An array of dictionaries that describes any errors that occur.

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




Removed: iOS 27+ | iPadOS 27+ | macOS 27+ | tvOS 27+

### UpdateResults

- **Type:** `[ScheduleOSUpdateResponse.UpdateResultsItem]`
- **Required:** Yes




Removed: iOS 27+ | iPadOS 27+ | macOS 27+ | tvOS 27+

### UserID

- **Type:** `string`
- **Required:** No







Removed: iOS 27+ | iPadOS 27+ | macOS 27+ | tvOS 27+

### UserLongName

- **Type:** `string`
- **Required:** Yes





Removed: iOS 27+ | iPadOS 27+ | macOS 27+ | tvOS 27+

### UserShortName

- **Type:** `string`
- **Required:** No







Removed: iOS 27+ | iPadOS 27+ | macOS 27+ | tvOS 27+

## Topics

### Objects

- [ScheduleOSUpdateResponse.ErrorChainItem](/documentation/devicemanagement/scheduleosupdateresponse/errorchainitem) - A dictionary that describes an error chain item.
- [ScheduleOSUpdateResponse.UpdateResultsItem](/documentation/devicemanagement/scheduleosupdateresponse/updateresultsitem) - The response dictionary that describes the result of processing an operating-system update.

