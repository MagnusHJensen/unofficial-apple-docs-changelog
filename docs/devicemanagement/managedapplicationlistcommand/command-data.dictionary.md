# ManagedApplicationListCommand.Command

The command to get the status of all managed apps on a device.

**Platforms:** iOS 5.0, iPadOS 5.0, Mac Catalyst 5.0, macOS 11.0, tvOS 10.2, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### Identifiers

- **Type:** `[string]`
- **Required:** No

The bundle identifiers of the managed apps to include in the response.

> 

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ManagedApplicationList`

The request type for this command.

