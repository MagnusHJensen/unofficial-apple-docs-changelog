# WebContentFilterPluginFilter_URLs_Parameters_PIRObject

A dictionary containing Private Information Retrieval server settings.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta)

## Properties

### AuthenticationTokenAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration containing the HTTP bearer token required to authenticate with the service. The bearer token is provided in the `Password` field of the asset data. The system uses this token to attest that it’s a valid user when requesting anonymous authentication tokens for PIR exchanges.

### PrivacyPassIssuerURL

- **Type:** `string`
- **Required:** Yes

The URL containing the domain name of Privacy Pass Issuer.

### ServerURL

- **Type:** `string`
- **Required:** Yes

The URL containing the domain name of the private information retrieval server.

