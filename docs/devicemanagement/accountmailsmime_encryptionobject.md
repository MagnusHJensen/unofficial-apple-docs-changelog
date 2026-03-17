# AccountMailSMIME_EncryptionObject

Settings for S/MIME encryption.

**Platforms:** iOS 17.0, iPadOS 17.0, visionOS 1.1

## Properties

### Enabled

- **Type:** `boolean`
- **Required:** Yes

If `true`, the system enables S/MIME encryption by default, which the user can’t override if `PerMessageSwitchEnabled` is `false`.

### IdentityAssetReference

- **Type:** `string`
- **Required:** No

Specifies the identifier of an asset declaration containing the identity required for S/MIME encryption. The system attaches the public certificate to outgoing mail to allow the user to receive encrypted mail. When the user sends encrypted mail, the system uses the public certificate to encrypt the copy of the mail in their Sent mailbox.

### IdentityUserOverrideable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the user can select an S/MIME signing identity in Settings.

### PerMessageSwitchEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the per-message encryption switch in the compose view.

### UserOverrideable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the user can set the default value for S/MIME encryption to on or off in Settings.

