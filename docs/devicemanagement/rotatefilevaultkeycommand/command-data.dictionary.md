# RotateFileVaultKeyCommand.Command

The command to change the FileVault primary password on a device.

**Platforms:** macOS 10.9, Device Assignment Services , VPP License Management 

## Properties

### FileVaultUnlock

- **Type:** `RotateFileVaultKeyCommand.Command.FileVaultUnlock`
- **Required:** Yes

A dictionary that contains FileVault unlock options.

### KeyType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `personal`, `institutional`

The type of FileVault key you want to change the password for. Set this value to `personal` and set a value for `Password` in the `FileVaultUnlock` dictionary to enable unlocking a device with a password. Set this value to `institutional` and set values for `PrivateKeyExport` and `PrivateKeyExportPassword` in the `FileVaultUnlock` dictionary.

### NewCertificate

- **Type:** `data`
- **Required:** No

A DER-encoded certificate for creating a new institutional recovery key, which the system requires if `KeyType` is `institutional`.

### ReplyEncryptionCertificate

- **Type:** `data`
- **Required:** No

A DER-encoded certificate for encrypting the new personal recovery key in a wrapper conforming to the IETF Cryptographic Message Syntax (CMS) standard.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `RotateFileVaultKey`

The request type for this command.

## Topics

### Objects

- [RotateFileVaultKeyCommand.Command.FileVaultUnlock](/documentation/devicemanagement/rotatefilevaultkeycommand/command-data.dictionary/filevaultunlock-data.dictionary) - A dictionary that contains FileVault unlock options.

