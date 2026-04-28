# AccountExchangeSMIME_SigningObject

Settings for S/MIME signing.

**Platforms:** iOS 17.0, iPadOS 17.0, Mac Catalyst 17.0, visionOS 1.1, Device Assignment Services , VPP License Management 

## Properties

### Enabled

- **Type:** `boolean`
- **Required:** Yes

If `true`, the system enables S/MIME signing.

### IdentityAssetReference

- **Type:** `string`
- **Required:** No

Specifies the identifier of an asset declaration containing the identity required for S/MIME signing of messages sent from this account.

### IdentityUserOverrideable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the user can select an S/MIME signing identity in Settings.

### UserOverrideable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the user can turn S/MIME signing on or off in Settings.

