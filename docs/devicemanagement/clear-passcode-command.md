# Clear Passcode

Remove the passcode from a device.

**Platforms:** iOS 4.0, iPadOS 4.0, visionOS 1.1, watchOS 10.0

## Discussion

Clearing the passcode in iOS 16 no longer adds the passcode to the history of passcodes. Therefore, the user can reuse the cleared passcode even when the `Passcode` payload has the `pinHistory` key set.

Refer to the following sections to determine supported channels and requirements, and to see an example request and response.

### Command availability

### Example request and response

## Topics

### Commands and responses

- [ClearPasscodeCommand](/documentation/devicemanagement/clearpasscodecommand) - The command to remove the passcode from a device.
- [ClearPasscodeResponse](/documentation/devicemanagement/clearpasscoderesponse) - A response from the device after it processes the command to remove the passcode from a device.

