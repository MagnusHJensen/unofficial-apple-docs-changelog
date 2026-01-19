# Clear Restrictions Password

Clear the Screen Time password and the restrictions on a device.

**Platforms:** iOS 8.0, iPadOS 8.0

## Discussion

In iOS 11 and earlier, this command clears the restrictions password and all restrictions that the password protects.

In iOS 12.2 and later, if Screen Time uses iCloud to share its settings (Share Across Devices), this command disables Screen Time entirely and clears its restrictions. If the user is a child in an iCloud family, the command fails. Otherwise, if Screen Time isn’t using iCloud, this command clears the passcode, but not the restrictions, and it leaves Screen Time enabled.

Refer to the following sections to determine supported channels and requirements, and to see an example request and response.

### Command availability

### Example request and response

## Topics

### Commands and responses

- [ClearRestrictionsPasswordCommand](/documentation/devicemanagement/clearrestrictionspasswordcommand) - The command to clear the Screen Time password and the restrictions on a device.
- [ClearRestrictionsPasswordResponse](/documentation/devicemanagement/clearrestrictionspasswordresponse) - A response from the device after it processes the command to clear the Screen Time password and the restrictions on a device.

