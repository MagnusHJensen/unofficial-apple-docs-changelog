# DeviceLocationCommand.Command

The command to request the location of a device when in Lost Mode.

**Platforms:** iOS 9.3, iPadOS 9.3

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `DeviceLocation`

The request type for this command.

