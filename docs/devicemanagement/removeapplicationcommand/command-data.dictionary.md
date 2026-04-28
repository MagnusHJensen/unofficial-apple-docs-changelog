# RemoveApplicationCommand.Command

The command to remove an app.

**Platforms:** iOS 5.0, iPadOS 5.0, Mac Catalyst 5.0, macOS 11.0, tvOS 10.2, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### Identifier

- **Type:** `string`
- **Required:** Yes

The bundle identifier of the managed app.

> 

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `RemoveApplication`

The request type for this command.

