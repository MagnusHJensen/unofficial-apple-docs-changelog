# ProfileListCommand.Command

The command to get a list of installed profiles on a device.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, tvOS 9.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### ManagedOnly

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, only include profiles that MDM has installed. For user enrollments, the device ignores this key and always limits the results to managed profiles. This value is available in iOS 13 and later, macOS 10.5 and later, and tvOS 13 and later.

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

