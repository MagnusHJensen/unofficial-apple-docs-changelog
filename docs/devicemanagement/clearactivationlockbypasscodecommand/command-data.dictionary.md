# ClearActivationLockBypassCodeCommand.Command

The command to clear the Activation Lock bypass code on a device.

**Platforms:** iOS 7.1, iPadOS 7.1, macOS 10.15, visionOS 2.0

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ClearActivationLockBypassCode`

The request type for this command.

