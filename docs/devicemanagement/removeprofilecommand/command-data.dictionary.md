# RemoveProfileCommand.Command

The command to remove a previously installed profile from the device.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 9.0, visionOS 1.1, watchOS 10.0

## Properties

### Identifier

- **Type:** `string`
- **Required:** Yes

The identifier of the profile to remove.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `RemoveProfile`

The request type for this command.

