# Set Firmware Password

Change or clear the firmware password on a device.

**Platforms:** macOS 10.13

## Discussion

This command has a throttle interval to prevent executing it more frequently than every 30 seconds. Requests that occur within the throttle interval return an error.

> 

After processing the command, the device restarts so that the new firmware password takes effect. This command returns an error and fails if a firmware password is already pending.

Refer to the following sections to determine supported channels and requirements, and to see an example request and response.

This command isn’t supported on a Mac with Apple silicon.

### Command availability

### Example request and response

## Topics

### Commands and responses

- [SetFirmwarePasswordCommand](/documentation/devicemanagement/setfirmwarepasswordcommand) - The command to change or clear the firmware password on a device.
- [SetFirmwarePasswordResponse](/documentation/devicemanagement/setfirmwarepasswordresponse) - A response from the device after it processes the command to change or clear the firmware password on a device.

