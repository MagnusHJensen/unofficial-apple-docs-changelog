# ClearPasscodeCommand.Command

The command to remove the passcode from a device.

**Platforms:** iOS 4.0, iPadOS 4.0, visionOS 1.1, watchOS 10.0

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ClearPasscode`

The request type for this command.

### UnlockToken

- **Type:** `data`
- **Required:** Yes

The unlock token value that the device provides in its `TokenUpdateMessage` check-in message.

