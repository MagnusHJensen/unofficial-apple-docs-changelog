# GetBootstrapTokenRequest

The get bootstrap token request details.

**Platforms:** iOS 26.0, iPadOS 26.0, Mac Catalyst 26.0, macOS 10.15, visionOS 26.0, Device Assignment Services , VPP License Management 

## Properties

### AwaitingConfiguration

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device is awaiting a [Device Configured](/documentation/devicemanagement/device-configured-command) command before proceeding through Setup Assistant.

### MessageType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `GetBootstrapToken`

The message type, which requires a value of `GetBootstrapToken`.

