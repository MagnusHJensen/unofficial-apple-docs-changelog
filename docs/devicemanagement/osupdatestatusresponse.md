# OSUpdateStatusResponse

A response from the device after it processes the command to get the status of operating-system updates on a device. Removed: subscribe to the declarative management `softwareupdate.install-state` status item.

**Platforms:** iOS 9.0, iPadOS 9.0, Mac Catalyst 9.0, macOS 10.11.5, tvOS 12.0

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

 


Removed: iOS 27+ | iPadOS 27+ | macOS 27+ | tvOS 27+

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

- [OSUpdateStatusResponse.ErrorChainItem](/documentation/devicemanagement/osupdatestatusresponse/errorchainitem) - A dictionary that describes an error chain item.
- [OSUpdateStatusResponse.OSUpdateStatusItem](/documentation/devicemanagement/osupdatestatusresponse/osupdatestatusitem) - A dictionary that describes the status of a software update.

