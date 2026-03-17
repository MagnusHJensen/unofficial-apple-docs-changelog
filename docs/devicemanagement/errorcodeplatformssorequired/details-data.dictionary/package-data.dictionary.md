# ErrorCodePlatformSSORequired.Details.Package

A dictionary that specifies the package that the device uses to install an app with the SSO app extension used for Platform SSO.

**Platforms:** macOS 26.0

## Properties

### ManifestURL

- **Type:** `string`
- **Required:** Yes

The URL of the app manifest, which needs to begin with `https:`.

### PinningCerts

- **Type:** `[data]`
- **Required:** No

An array of DER-encoded certificates to pin the connection when fetching the `ManifestURL`.

### PinningRevocationCheckRequired

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, certificate revocation checks require a positive response when using certificate pinning with `PinningCerts`.

