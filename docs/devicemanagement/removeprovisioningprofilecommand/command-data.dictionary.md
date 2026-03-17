# RemoveProvisioningProfileCommand.Command

The command to remove a previously installed provisioning profile from a device.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 11.0, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `RemoveProvisioningProfile`

The request type for this command.

### UUID

- **Type:** `string`
- **Required:** Yes

The unique identifier of the provisioning profile to remove.

