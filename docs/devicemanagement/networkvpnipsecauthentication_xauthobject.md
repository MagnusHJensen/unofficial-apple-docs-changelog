# NetworkVPNIPSecAuthentication_XAuthObject

Settings that control XAuth.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### CredentialsAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the credentials (user name and password) required for XAuth. Required when `Enabled` key is set to `true`.

### Enabled

- **Type:** `boolean`
- **Required:** Yes

If `true`, enables Xauth for Cisco IPSec VPNs.

### PasswordEncryption

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Prompt`

A string that either has the value `Prompt` or isn’t present.

