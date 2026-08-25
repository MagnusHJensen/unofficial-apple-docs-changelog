# AccountExchangeSMIME_EncryptionObject

Settings for S/MIME encryption. Applicable for “EAS” only.

**Platforms:** iOS 17.0, iPadOS 17.0, Mac Catalyst 17.0, visionOS 1.1

## Properties

### Enabled

- **Type:** `boolean`
- **Required:** Yes

If `true`, the system enables S/MIME encryption by default, which the user can’t override if `PerMessageSwitchEnabled` is `false`. Applicable for “EAS” only.

### IdentityAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration containing the identity required for S/MIME encryption. The system attaches the public certificate to outgoing mail to allow the user to receive encrypted mail. When the user sends encrypted mail, the system uses the public certificate to encrypt the copy of the mail in their Sent mailbox. Applicable for “EAS” only.

### IdentityUserOverrideable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the user can select an S/MIME signing identity in Settings. Applicable for “EAS” only.

### PerMessageSwitchEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the per-message encryption switch in the compose view. Applicable for “EAS” only.

### UserOverrideable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the user can turn S/MIME encryption by default on or off in Settings. Applicable for “EAS” only.

