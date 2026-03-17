# StopMirroringCommand.Command

The command to stop mirroring the display to another device.

**Platforms:** iOS 7.0, iPadOS 7.0, macOS 10.10

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `StopMirroring`

The request type for this command.

