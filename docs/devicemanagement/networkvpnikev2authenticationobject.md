# NetworkVPNIKEV2AuthenticationObject

Settings that control authentication.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), tvOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### CredentialsAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the credentials (password) to authenticate with the VPN server. Required when `Authentication.Method` is set to `SharedSecret`.

### ExtendedAuth

- **Type:** `NetworkVPNIKEV2Authentication_ExtendedAuthObject`
- **Required:** No

Specifies details about how the VPN routes different types of network traffic.

### IdentityAssetReference

- **Type:** `string`
- **Required:** No

The identifier of a credential asset declaration that contains the identity that this account requires to authenticate with the VPN server. If the value of `AuthenticationMethod` is `Certificate`, the system sends this certificate out for IKEv2 machine authentication. If extended authentication (EAP) is used, the system sends this certificate out for EAP-TLS authentication. Required when `Authentication.Method` is set to `Certificate`.

### IdentityCertificateType

- **Type:** `string`
- **Required:** No
- **Default:** `RSA`
- **Allowed Values:** `RSA`, `ECDSA256`, `ECDSA384`, `ECDSA521`, `RSA-PSS`

The type of key used by the identity set in the `IdentityAssetReference` to use for IKEv2 machine authentication. If this key is included, the system requires a value for `ServerCertificateIssuerCommonName`.

### Method

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `None`, `SharedSecret`, `Certificate`

The type of authentication method for the VPN.

To enable EAP-only authentication, set this to `None` and `ExtendedAuthEnabled` to `true`. If this is `None` and the `ExtendedAuthEnabled` key isn’t set, the authentication configuration defaults to `SharedSecret`.

## Topics

### Objects

- [NetworkVPNIKEV2Authentication_ExtendedAuthObject](/documentation/devicemanagement/networkvpnikev2authentication_extendedauthobject) - Specifies details about how the VPN routes different types of network traffic.

