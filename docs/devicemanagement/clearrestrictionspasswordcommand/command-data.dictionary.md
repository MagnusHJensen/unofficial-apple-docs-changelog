# ClearRestrictionsPasswordCommand.Command

The command to clear the Screen Time password and the restrictions on a device.

**Platforms:** iOS 8.0, iPadOS 8.0

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ClearRestrictionsPassword`

The request type for this command.

