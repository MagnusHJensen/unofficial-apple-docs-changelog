# OSUpdateStatusCommand.Command

The command to get the status of operating-system updates on a device.

**Platforms:** iOS 9.0, iPadOS 9.0, macOS 10.11.5, tvOS 12.0

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `OSUpdateStatus`

The request type for this command.

