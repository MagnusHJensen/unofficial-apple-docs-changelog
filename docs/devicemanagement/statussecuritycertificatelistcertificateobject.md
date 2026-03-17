# StatusSecurityCertificateListCertificateObject

A status report of a security certificate.

**Platforms:** iOS 17.0, iPadOS 17.0, macOS 14.0, tvOS 17.0, visionOS 1.1, watchOS 10.0

## Properties

### _removed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system removed the app and only this key and the `identifier` key are present in the status item object.

### data

- **Type:** `string`
- **Required:** Yes

The certificate data in DER-encoded X.509 format.

### declaration-identifier

- **Type:** `string`
- **Required:** No

The identifier of the asset declaration that installed the certificate, which is only present if a declaration installed the certificate.

### identifier

- **Type:** `string`
- **Required:** Yes

The unique identifier of the certificate which the system uses as the primary key.

### is-identity

- **Type:** `boolean`
- **Required:** Yes

If `true`, the certificate is an identity certificate.

### subject-summary

- **Type:** `string`
- **Required:** Yes

The summary of the certificate’s subject.

