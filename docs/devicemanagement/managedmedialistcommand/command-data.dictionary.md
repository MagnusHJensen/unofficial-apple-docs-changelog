# ManagedMediaListCommand.Command

The command to get a list of the managed books on a device.

**Platforms:** iOS 8.0, iPadOS 8.0, Mac Catalyst 8.0, Device Assignment Services , VPP License Management 

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ManagedMediaList`

The request type for this command.

