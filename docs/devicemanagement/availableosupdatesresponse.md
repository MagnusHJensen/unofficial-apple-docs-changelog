# AvailableOSUpdatesResponse

A response from the device after it processes the command to get a list of available operating-system updates for a device.

**Platforms:** iOS 9.0, iPadOS 9.0, Mac Catalyst 9.0, macOS 10.11, tvOS 12.0, Device Assignment Services , VPP License Management 

## Properties

### AvailableOSUpdates

- **Type:** `[AvailableOSUpdatesResponse.AvailableOSUpdatesItem]`
- **Required:** Yes

An array of dictionaries that contains only the most recent available updates in iOS and tvOS, and possibly multiple available updates in macOS. Follow the instructions in the Managed Apps and Updates section of the Apple Software Lookup Service to find a complete catalog of iOS and tvOS updates.

In macOS 14 and later, `AvailableOSUpdates` doesn’t include InstallAssistant-based, full-replacement installers. It only contains over-the-air (OTA) updates. OTA updates can update or upgrade the OS and support all `InstallAction` options.

If a Software Update is actively managed by a Declarative Device Management Specific Enforcement configuration, the device ignores this command as it applies to the actively managed update. This command can return information for unmanaged updates, such as System Applications and Configuration Data. For information about available updates when using Declarative Device Management, see [Using the Apple Software Lookup Service](https://support.apple.com/guide/deployment/depafd2fad80/web).

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

- **Type:** `[AvailableOSUpdatesResponse.ErrorChainItem]`
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

- [AvailableOSUpdatesResponse.AvailableOSUpdatesItem](/documentation/devicemanagement/availableosupdatesresponse/availableosupdatesitem) - The response dictionary that describes the available operating-system updates item.
- [AvailableOSUpdatesResponse.ErrorChainItem](/documentation/devicemanagement/availableosupdatesresponse/errorchainitem) - A dictionary that describes an error chain item.

