# SecurityInfoResponse.SecurityInfo.FirmwarePasswordStatus

A dictionary that contains the status of the EFI firmware password.

**Platforms:** macOS 10.13, Device Assignment Services , VPP License Management 

## Properties

### AllowOroms

- **Type:** `boolean`
- **Required:** No

If `true`, enable ROMs.

### ChangePending

- **Type:** `boolean`
- **Required:** No

If `true`, a firmware password change is pending. A device restart is necessary for this change to take effect. Until then, additional attempts to change the password fail.

> 

### PasswordExists

- **Type:** `boolean`
- **Required:** No

If `true`, the device has an EFI firmware password.

