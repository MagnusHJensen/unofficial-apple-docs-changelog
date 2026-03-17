# SCEPCredential

A SCEP identity that the device generates.

**Platforms:** iOS 17.0, iPadOS 17.0, macOS 14.0, tvOS 17.0, visionOS 1.1, watchOS 10.0

## Properties

### CAFingerprint

- **Type:** `string`
- **Required:** No

The fingerprint of the Certificate Authority certificate.

### Challenge

- **Type:** `string`
- **Required:** No

A preshared secret.

### Key Type

- **Type:** `string`
- **Required:** No
- **Default:** `RSA`

The key type, which always has the value `RSA`.

### Key Usage

- **Type:** `integer`
- **Required:** No
- **Default:** `0`

A bitmask that specifies the use of the key: `1` is signing, `4` is encryption, and `5` is both signing and encryption. Some certificate authorities, such as Windows CA, support only encryption or signing, but not both at the same time.

### Keysize

- **Type:** `integer`
- **Required:** No
- **Default:** `1024`
- **Allowed Values:** `1024`, `2048`, `4096`

The key size in bits, either `1024`, `2048`, or `4096`.

### Name

- **Type:** `string`
- **Required:** No

Any string that the SCEP server recognizes. For example, it could be a domain name such as `example.org`. If a certificate authority has multiple CA certificates, you can use this field to specify the required certificate.

### Retries

- **Type:** `integer`
- **Required:** No
- **Default:** `3`

The number of times the device should retry if the server sends a `PENDING` response.

### RetryDelay

- **Type:** `integer`
- **Required:** No
- **Default:** `10`

The number of seconds to wait between subsequent retries. The system makes the first retry without this delay.

### Subject

- **Type:** `[[[string]]]`
- **Required:** No

The representation of an X.500 name is an array of OID and value. For example, `/C=US/O=Apple Inc./CN=foo/1.2.5.3=bar` corresponds to:

`[ [ ["C", "US"] ], [ ["O", "Apple Inc."] ], [ [ "CN", "foo"] ], [ [ "1.2.5.3", "bar" ] ] ]`

You can represent OIDs as dotted numbers or use shortcuts for country (`C`), locality (`L`), state (`ST`), organization (`O`), organizational unit (`OU`), and common name (`CN`).

### SubjectAltName

- **Type:** `SCEPCredentialSubjectAltNameObject`
- **Required:** No

The subject’s alternative name for the certificate.

### URL

- **Type:** `string`
- **Required:** Yes

The SCEP URL.

## Topics

### Objects

- [SCEPCredentialSubjectAltNameObject](/documentation/devicemanagement/scepcredentialsubjectaltnameobject) - The subject’s alternative name for the certificate.

