# SetBootstrapTokenRequest

The set bootstrap token request details.

**Platforms:** iOS 26.0, iPadOS 26.0, Mac Catalyst 26.0, macOS 10.15, visionOS 26.0, Device Assignment Services , VPP License Management 

## Properties

### AwaitingConfiguration

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device is awaiting a [Device Configured](/documentation/devicemanagement/device-configured-command) command before proceeding through Setup Assistant.

### BootstrapToken

- **Type:** `data`
- **Required:** No

The device’s bootstrap token data. If this field is missing or zero length, the server needs to remove the bootstrap token for this device.

### MessageType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `SetBootstrapToken`

The message type, which requires a value of `SetBootstrapToken`.

