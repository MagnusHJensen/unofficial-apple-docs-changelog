# DeviceConfiguredCommand.Command

The command to inform the device that it can allow the user to continue in Setup Assistant.

**Platforms:** iOS 9.0, iPadOS 9.0, macOS 10.11, tvOS 10.2, visionOS 2.0

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `DeviceConfigured`

The request type for this command.

