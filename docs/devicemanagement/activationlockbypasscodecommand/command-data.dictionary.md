# ActivationLockBypassCodeCommand.Command

The command to get the code to bypass Activation Lock on a device.

**Platforms:** iOS 7.1, iPadOS 7.1, Mac Catalyst 7.1, macOS 10.15, visionOS 2.0, Device Assignment Services , VPP License Management 

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ActivationLockBypassCode`

The request type for this command.

