# NetworkVPNIKEV2Authentication_ExtendedAuthObject

Specifies details about how the VPN routes different types of network traffic.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), tvOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### CredentialsAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the credentials (user name and password) to authenticate with the VPN server. Required when `Enabled` is set to `true`. Implies the use of EAP-MSCHAPv2.

### Enabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables EAP-only authentication.

### ServerCertificateCommonName

- **Type:** `string`
- **Required:** No

The common name of the server certificate. The system uses this name to validate the certificate sent by the IKE server. If not set, the system uses the remote identifier to validate the certificate.

### ServerCertificateIssuerCommonName

- **Type:** `string`
- **Required:** No

Common Name of the server certificate issuer. If set, this field causes IKE to send a certificate request based on this certificate issuer to the server. This key is required if the `IdentityCertificateType` key is included and the `ExtendedAuth.Enabled` key is `true`.

### TLSMaximumVersion

- **Type:** `string`
- **Required:** No
- **Default:** `1.2`
- **Allowed Values:** `1.0`, `1.1`, `1.2`, `1.3`

The maximum TLS version to use with EAP-TLS authentication.

### TLSMinimumVersion

- **Type:** `string`
- **Required:** No
- **Default:** `1.0`
- **Allowed Values:** `1.0`, `1.1`, `1.2`, `1.3`

The minimum TLS version to use with EAP-TLS authentication.

