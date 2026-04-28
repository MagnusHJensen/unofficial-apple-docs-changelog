# AppManagedAppConfigDictionaryObject

A dictionary of app config data and credentials.

**Platforms:** iOS 18.4, iPadOS 18.4, Mac Catalyst 18.4, visionOS 2.4, Device Assignment Services , VPP License Management 

## Properties

### Certificates

- **Type:** `[AppManagedCredentialConfigObject]`
- **Required:** No

Provides certificates to the managed app or extension. Each element in the array contains a certificate asset reference and an associated identifier which the app or extension uses to look up the certificate.

### DataAssetReference

- **Type:** `string`
- **Required:** No

Specifies the identifier of an asset declaration containing a reference to the app or extension config data. The corresponding asset needs to be of type `com.apple.asset.data`. The referenced data needs to be a property list file, and the asset’s “ContentType” value set to match the data type.

### Identities

- **Type:** `[AppManagedCredentialConfigObject]`
- **Required:** No

Provides identities to the managed app or extension. Each element in the array contains an identity asset reference and an associated identifier which the app or extension uses to look up the identity.

### Passwords

- **Type:** `[AppManagedCredentialConfigObject]`
- **Required:** No

Provides passwords to the managed app or extension. Each element in the array contains a password asset reference and an associated identifier which the app or extension uses to look up the password.

## Topics

### Objects

- [AppManagedCredentialConfigObject](/documentation/devicemanagement/appmanagedcredentialconfigobject) - A dictionary of values associated with a credential config.

