# NetworkVPNVPNPluginAuthenticationObject

Settings that control authentication.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, tvOS 27.0, visionOS 27.0

## Properties

### CredentialsAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the credentials (user name and password) to authenticate with the VPN server. Required when `Authentication.Method` is set to `Password`.

### IdentityAssetReference

- **Type:** `string`
- **Required:** No

The identifier of a credential asset declaration that contains the identity that this account requires to authenticate with the VPN server. Required when `Authentication.Method` is set to `Certificate`.

### Method

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Password`, `Certificate`, `Password+Certificate`

The authentication method to use.

