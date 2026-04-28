# ManagedApplicationAttributesCommand.Command

The command to query attributes in managed apps on a device.

**Platforms:** iOS 7.0, iPadOS 7.0, Mac Catalyst 7.0, tvOS 10.2, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### Identifiers

- **Type:** `[string]`
- **Required:** Yes

The bundle identifiers of the managed apps.

> 

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ManagedApplicationAttributes`

The request type for this command.

