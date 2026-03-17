# CertificateTransparency.SubjectPublicKeyInfoHashDict

A dictionary of hashed public keys.

**Platforms:** iOS 12.1.1, iPadOS 12.1.1, macOS 10.14.2, tvOS 12.1.1, visionOS 1.0, watchOS 5.1.1

## Properties

### Algorithm

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `sha256`

The algorithm must be `sha256`.

### Hash

- **Type:** `data`
- **Required:** Yes

The hash of the DER-encoding of the certificate’s `subjectPublicKeyInfo`.

The hash field requires the data (`subjectPublicKeyInfo` hash) in a specific format: a Base64 encoded (binary) SHA-256 hash of the certificate’s public key.

