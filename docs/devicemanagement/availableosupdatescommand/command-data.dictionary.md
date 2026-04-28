# AvailableOSUpdatesCommand.Command

The command to get a list of available operating-system updates for a device.

**Platforms:** iOS 9.0, iPadOS 9.0, Mac Catalyst 9.0, macOS 10.11, tvOS 12.0, Device Assignment Services , VPP License Management 

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `AvailableOSUpdates`

The request type for this command.

