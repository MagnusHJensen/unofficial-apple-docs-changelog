# ProvisioningProfileListCommand.Command

The command to get a list of installed provisioning profiles on a device.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 11.0, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Properties

### ManagedOnly

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, only include profiles that MDM has installed. For user enrollments, the device ignores this key and always limits the results to managed profiles. This value is available in iOS 13 and later, and tvOS 13 and later.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ProvisioningProfileList`

The request type for this command.

