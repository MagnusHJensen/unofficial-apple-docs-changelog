# ShutDownDeviceCommand.Command

The command to remotely and immediately shut down a device.

**Platforms:** iOS 10.3, iPadOS 10.3, Mac Catalyst 10.3, macOS 10.13, Device Assignment Services , VPP License Management 

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ShutDownDevice`

The request type for this command.

