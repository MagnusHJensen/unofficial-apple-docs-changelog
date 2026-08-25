# AccountExchangeSMIME_SigningObject

Settings for S/MIME signing. Applicable for “EAS” only.

**Platforms:** iOS 17.0, iPadOS 17.0, Mac Catalyst 17.0, visionOS 1.1

## Properties

### Enabled

- **Type:** `boolean`
- **Required:** Yes

If `true`, the system enables S/MIME signing. Applicable for “EAS” only.

### IdentityAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration containing the identity required for S/MIME signing of messages sent from this account. Applicable for “EAS” only.

### IdentityUserOverrideable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the user can select an S/MIME signing identity in Settings. Applicable for “EAS” only.

### UserOverrideable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the user can turn S/MIME signing on or off in Settings. Applicable for “EAS” only.

