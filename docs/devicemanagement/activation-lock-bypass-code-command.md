# Activation Lock Bypass Code

Get the code to bypass Activation Lock on a device.

**Platforms:** iOS 7.1, iPadOS 7.1, macOS 10.15, visionOS 2.0

## Discussion

This command allows organizations to retrieve the device’s bypass code. Organizations can use the bypass code to remove the Activation Lock from supervised devices prior to device activation without knowing the user’s personal Apple Account and password.

Supervised devices generate a device-specific Activation Lock bypass code. The activation server verifies this code to bypass Activation Lock on the device. For more information, see [Creating and Using Bypass Codes](/documentation/devicemanagement/creating-and-using-bypass-codes).

A device creates a new bypass code when:

- Setting up the device the first time.
- Erasing and not restoring the device from a backup.
- Erasing and restoring the device from a backup from a different device.

### Command availability

### Example request and response

## Topics

### Commands and responses

- [ActivationLockBypassCodeCommand](/documentation/devicemanagement/activationlockbypasscodecommand) - The command to get the code to bypass Activation Lock on a device.
- [ActivationLockBypassCodeResponse](/documentation/devicemanagement/activationlockbypasscoderesponse) - A response from the device after it processes the command to get the code to bypass Activation Lock on a device.

