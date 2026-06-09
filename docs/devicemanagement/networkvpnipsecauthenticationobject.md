# NetworkVPNIPSecAuthenticationObject

Settings that control authentication.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### CredentialsAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the credentials (password) to authenticate with the VPN servers.

Only use this with Cisco IPSec VPNs and if the `Authentication.Method` key is to `SharedSecret`.

### IdentityAssetReference

- **Type:** `string`
- **Required:** No

The identifier of a credential asset declaration that contains the identity that this account requires to authenticate with the VPN servers.

Only use this with Cisco IPSec VPNs and if the `Authentication.Method` key is to `Certificate`.

### LocalIdentifier

- **Type:** `string`
- **Required:** No

The name of the group. For hybrid authentication, the string needs to end with “hybrid”.

Present only for Cisco IPSec if `Authentication.Method` is `SharedSecret`.

### LocalIdentifierType

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `KeyID`

Present only if `Authentication.Method` is `SharedSecret`. The value is `KeyID`. The system uses this value for Cisco IPSec VPNs.

### Method

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `SharedSecret`, `Certificate`

The authentication method to use.

### PromptForVPNPIN

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, prompts for a PIN when connecting to Cisco IPSec VPNs.

### XAuth

- **Type:** `NetworkVPNIPSecAuthentication_XAuthObject`
- **Required:** No

Settings that control XAuth.

## Topics

### Objects

- [NetworkVPNIPSecAuthentication_XAuthObject](/documentation/devicemanagement/networkvpnipsecauthentication_xauthobject) - Settings that control XAuth.

