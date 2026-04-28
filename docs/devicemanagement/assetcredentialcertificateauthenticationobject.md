# AssetCredentialCertificateAuthenticationObject

The server authentication details for an asset-credential certificate.

**Platforms:** iOS 17.0, iPadOS 17.0, Mac Catalyst 17.0, macOS 14.0, tvOS 17.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### Type

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `MDM`, `None`

The type of authentication, which has these allowed values:

- `MDM`: A request that uses MDM semantics, which includes the device-identity certificate, and any user authentication. This is equivalent to an MDM request made to the `CheckInURL` or `ServerURL`. This option is only available through declarative device management.
- `None`: A standard GET request.

