# Device Location

Request the location of a device when in Lost Mode.

**Platforms:** iOS 9.3, iPadOS 9.3

## Discussion

A device responds with error codes:

- `12067`: If it isn’t in Lost mode.
- `12068`: If its location is unknown.
- `12078`: If the command is invalid while in Lost Mode.

Refer to the following sections to determine supported channels and requirements, and to see an example request and response.

### Command availability

### Example request and response

## Topics

### Commands and responses

- [DeviceLocationCommand](/documentation/devicemanagement/devicelocationcommand) - The command to request the location of a device when in Lost Mode.
- [DeviceLocationResponse](/documentation/devicemanagement/devicelocationresponse) - A response from the device after it processes the command to request the location of a device when in Lost Mode.

