# ExtensibleSSOPlatformSSO_AccessKeyObject

Settings for Access Key authentication.

**Platforms:** macOS 27.0 (Beta)

## Properties

### AllowExpressMode

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system uses the access key in express mode, and doesn’t require authentication before use.

### ReaderGroupIdentifier

- **Type:** `string`
- **Required:** No

The reader group identifier for use with the `AccessKey`. The value needs to match the configured access key. Required if `UserCreation.AuthenticationMethods` contains `AccessKey`.

### ReaderIssuerCertificateAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the certificate for the issuer certificate of the `Terminal` identity of the access key. Other specifications refer to the key as the “Reader CA Public Key”. The key must be an elliptic curve key. Required if `UserCreation.AuthenticationMethods` includes `AccessKey`. The issuer of the Terminal identity of the access key needs to match this certificate, otherwise the device fails the authentication.

### TerminalIdentityAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the identity to use as the Terminal identity of the Access Key. The Access Key needs to trust the identity. Required if `UserCreation.AuthenticationMethods` includes `AccessKey`.

