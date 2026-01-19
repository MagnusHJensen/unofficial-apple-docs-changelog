# Erase Device

Remotely and immediately erase a device.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Discussion

This command allows the server to immediately erase a device, even a locked device, without warning the user. The device sends a response to the server, but it doesn’t retry if it isn’t successful the first time.

Refer to the following sections to determine supported channels and requirements, and to see an example request and response.

### Command availability

### Example request and response

## Topics

### Commands and responses

- [EraseDeviceCommand](/documentation/devicemanagement/erasedevicecommand) - The command to remotely and immediately erase a device.
- [EraseDeviceResponse](/documentation/devicemanagement/erasedeviceresponse) - A response from the device after it processes the command to remotely and immediately erase a device.

