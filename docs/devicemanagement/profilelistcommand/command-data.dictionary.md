# ProfileListCommand.Command

The command to get a list of installed profiles on a device.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, tvOS 9.0, visionOS 1.1, watchOS 10.0

## Properties

### ManagedOnly

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, only include profiles that MDM has installed. For user enrollments, the device ignores this key and always limits the results to managed profiles.

Available: iOS 13+ | iPadOS 13+ | macOS 10.15+ | tvOS 13+ | visionOS 1.1+ | watchOS 10+

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ProfileList`

The request type for this command.

