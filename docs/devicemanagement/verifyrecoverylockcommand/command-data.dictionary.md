# VerifyRecoveryLockCommand.Command

The command to verify the device’s Recovery Lock password.

**Platforms:** macOS 11.5, Device Assignment Services , VPP License Management 

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
- **Allowed Values:** `VerifyRecoveryLock`

The request type for this command.

