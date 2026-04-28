# SetFirmwarePasswordCommand.Command

The command to change or clear the firmware password on a device.

**Platforms:** macOS 10.13, Device Assignment Services , VPP License Management 

## Properties

### AllowOroms

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enable ROMs.

### CurrentPassword

- **Type:** `string`
- **Required:** No

The current password, which you must set if the device has a firmware password.

### NewPassword

- **Type:** `string`
- **Required:** Yes

The new firmware password. Set to an empty string to clear the password. The characters in this value must consist of low-ASCII, printable characters (`0x20` through `0x7E`) to ensure that all characters are enterable on the EFI login screen.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `SetFirmwarePassword`

The request type for this command.

