# SCEP.PayloadContent

The SCEP dictionary.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, tvOS 9.0, visionOS 1.0, watchOS 3.0, Device Assignment Services , VPP License Management 

## Properties

### AllowAllAppsAccess

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, all apps have access to the private key.

### CAFingerprint

- **Type:** `data`
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

Always `RSA`.

### Key Usage

- **Type:** `integer`
- **Required:** No
- **Default:** `0`

A bitmask indicating the use of the key. Possible values:

- `1`: Signing
- `4`: Encryption

Some certificate authorities, such as Windows CA, support only encryption or signing, but not both at the same time.

### KeyIsExtractable

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system disables exporting the private key from the keychain.

### Keysize

- **Type:** `integer`
- **Required:** No
- **Default:** `1024`
- **Allowed Values:** `1024`, `2048`, `4096`

The key size, in bits.

### Name

- **Type:** `string`
- **Required:** No

A string that’s understood by the SCEP server; for example, a domain name like example.org. If a certificate authority has multiple CA certificates, this field can be used to distinguish which is required.

### Retries

- **Type:** `integer`
- **Required:** No
- **Default:** `3`

The number of times the device should retry if the server sends a PENDING response.

### RetryDelay

- **Type:** `integer`
- **Required:** No
- **Default:** `10`

The number of seconds to wait between subsequent retries. The first retry is attempted without this delay.

### Subject

- **Type:** `[[[string]]]`
- **Required:** No

The representation of an X.500 name as an array of OID and value.

For example, `/C=US/O=Apple Inc./CN=foo/1.2.5.3=bar` translates to `[ [ ["C", "US"] ], [ ["O", "Apple Inc."] ], …, [ [ "1.2.5.3", "bar" ] ] ]`.

OIDs can be represented as dotted numbers, with shortcuts for country (C), locality (L), state (ST), organization (O), organizational unit (OU), and common name (CN).

### SubjectAltName

- **Type:** `SCEP.PayloadContent.SubjectAltName`
- **Required:** No

The SCEP payload can specify an optional `SubjectAltName` dictionary that provides values required by the CA for issuing a certificate. You can specify a single string or an array of strings for each key. The values you specify depend on the CA you’re using, but might include DNS name, URL, or email values. For an example, see Sample Configuration Profile or Over-the-Air Profile Delivery and Configuration.

### URL

- **Type:** `string`
- **Required:** Yes

The SCEP URL. See Over-the-Air Profile Delivery and Configuration for more information about SCEP.

## Topics

### Objects

- [SCEP.PayloadContent.SubjectAltName](/documentation/devicemanagement/scep/payloadcontent-data.dictionary/subjectaltname-data.dictionary) - An optional dictionary that provides values required by the CA for issuing a certificate.

