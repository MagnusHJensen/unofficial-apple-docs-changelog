# DeviceLockCommand.Command

The command to remotely and immediately lock a device.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, visionOS 2.0, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### Message

- **Type:** `string`
- **Required:** No

The message to display on the Lock Screen of the device. This value doesn’t apply to a Shared iPad device. This value is available in iOS 4 and later, and macOS 10.14 and later.

### PhoneNumber

- **Type:** `string`
- **Required:** No

The phone number to display on the Lock Screen. This value doesn’t apply to a Shared iPad device. This value is available in iOS 7 and later and macOS 11.5 and later (for a Mac with Apple silicon only).

### PIN

- **Type:** `string`
- **Required:** No

The six-character PIN for Find My. This value is available in macOS 10.8 and later.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `DeviceLock`

The request type for this command.

