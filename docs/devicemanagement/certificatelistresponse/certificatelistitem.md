# CertificateListResponse.CertificateListItem

A dictionary that contains information about a certificate list item.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 9.0, visionOS 1.1, watchOS 10.0

## Properties

### CommonName

- **Type:** `string`
- **Required:** Yes

The certificate’s common name.

### Data

- **Type:** `data`
- **Required:** Yes

The certificate in DER-encoded X.509 format.

### IsIdentity

- **Type:** `boolean`
- **Required:** Yes

If `true`, this is an identity certificate.

