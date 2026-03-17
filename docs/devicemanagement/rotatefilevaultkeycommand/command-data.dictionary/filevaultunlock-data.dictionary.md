# RotateFileVaultKeyCommand.Command.FileVaultUnlock

A dictionary that contains FileVault unlock options.

**Platforms:** macOS 10.9

## Properties

### Password

- **Type:** `string`
- **Required:** No

A FileVault user’s password, or if using a CoreStorage volume, the personal recovery key.

### PrivateKeyExport

- **Type:** `data`
- **Required:** No

The data for a .p12 export of the private key for the current institutional recovery key, which requires that `KeyType` is `institutional`. The system ignores this key on APFS volumes.

### PrivateKeyExportPassword

- **Type:** `string`
- **Required:** No

The password for `PrivateKeyExport`. Either `Password` or both `PrivateKeyExport` and `PrivateKeyExportPassword` must be present. The system ignores this key on APFS volumes.

