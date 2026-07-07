# FDERecoveryKeyRedirection

The payload that configures FileVault recovery key redirection.

**Platforms:** macOS 10.9

## Properties

### EncryptCertPayloadUUID

- **Type:** `string`
- **Required:** Yes

The UUID of a payload within the same profile that contains a certificate used to encrypt the recovery key when the device sends it to the redirected URL. The referenced payload must be of type `com.apple.security.pkcs1`.

Deprecated: macOS 10.13+

### RedirectURL

- **Type:** `string`
- **Required:** Yes

The URL to which the device sends FDE recovery keys instead of to Apple. The URL must begin with https://.

Deprecated: macOS 10.13+

## Discussion

Specify `com.apple.security.FDERecoveryRedirect` as the payload type.

Although the previous FDE Recovery payload is no longer supported in macOS 10.13 and later, it’s still supported in macOS 10.9 through 10.12. When installed, this payload causes any FDE recovery keys to be redirected to the specified URL instead of being sent to Apple. This requires sites to implement their own HTTPS server to receive the recovery keys through a POST request.

Note these cautions:

- The payload must exist in a system-scoped profile.
- Installing more than one payload of this type per machine results in an error.

### Profile availability

