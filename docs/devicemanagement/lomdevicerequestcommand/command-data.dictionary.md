# LOMDeviceRequestCommand.Command

The command to send requests to a device using lights-out management (LOM).

**Platforms:** macOS 11.0

## Properties

### RequestList

- **Type:** `[LOMDeviceRequestCommand.Command.RequestListItem]`
- **Required:** Yes

An array of requests to perform.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `LOMDeviceRequest`

The request type for this command.

## Topics

### Objects

- [LOMDeviceRequestCommand.Command.RequestListItem](/documentation/devicemanagement/lomdevicerequestcommand/command-data.dictionary/requestlistitem) - A dictionary that contains a requested action to perform on a device using lights-out management (LOM).

