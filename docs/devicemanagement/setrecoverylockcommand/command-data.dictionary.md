# SetRecoveryLockCommand.Command

The command to set or clear the Recovery Lock password.

**Platforms:** macOS 11.5

## Properties

### CurrentPassword

- **Type:** `string`
- **Required:** No

If the device has a Recovery Lock password set, the system requires the current password.

### NewPassword

- **Type:** `string`
- **Required:** Yes

The new password for Recovery Lock. Set as an empty string to clear the Recovery Lock password.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `SetRecoveryLock`

The request type for this command.

