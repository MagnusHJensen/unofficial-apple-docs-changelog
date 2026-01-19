# Disable Lost Mode

Take the device out of Lost Mode.

**Platforms:** iOS 9.3, iPadOS 9.3

## Discussion

A device responds with error codes:

- `12067`: If it isn’t in Lost Mode.
- `12069`: If the request to disable Lost Mode failed.
- `12078`: If the command is invalid while in Lost Mode.

Erasing a device also disables Lost Mode. To reenable Lost Mode, the MDM server needs to store the device’s Lost Mode state before erasing it, and restore that state if the device enrolls again.

Refer to the following sections to determine supported channels and requirements, and to see an example request and response.

### Command availability

### Example request and response

## Topics

### Commands and responses

- [DisableLostModeCommand](/documentation/devicemanagement/disablelostmodecommand) - The command to take the device out of Lost Mode.
- [DisableLostModeResponse](/documentation/devicemanagement/disablelostmoderesponse) - A response from the device after it processes the command to take the device out of Lost Mode.

