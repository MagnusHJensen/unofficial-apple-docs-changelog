# DisableRemoteDesktopCommand.Command

The command to disable Remote Desktop on a device.

**Platforms:** macOS 10.14.4

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `DisableRemoteDesktop`

The request type for this command.

