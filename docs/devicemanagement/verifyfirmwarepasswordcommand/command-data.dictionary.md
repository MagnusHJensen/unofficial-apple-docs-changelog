# VerifyFirmwarePasswordCommand.Command

The command to verify the firmware password on a device.

**Platforms:** macOS 10.13

## Properties

### Password

- **Type:** `string`
- **Required:** Yes

The password to verify.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `VerifyFirmwarePassword`

The request type for this command.

