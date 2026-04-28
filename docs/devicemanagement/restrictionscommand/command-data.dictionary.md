# RestrictionsCommand.Command

The command to get a list of restrictions on the device.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, tvOS 9.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### ProfileRestrictions

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device reports restrictions from each profile. This value is available in iOS 4 and later, and tvOS 6.1 and later.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Restrictions`

The request type for this command.

