# AssetCredentialIdentity

A reference to a PKCS #12 password-protected identity.

**Platforms:** iOS 17.0, iPadOS 17.0, macOS 14.0, tvOS 17.0, visionOS 1.1, watchOS 10.0

## Properties

### Accessible

- **Type:** `string`
- **Required:** No
- **Default:** `Default`
- **Allowed Values:** `Default`, `AfterFirstUnlock`

The keychain accessibility that determines when the keychain item is available for use, which has these allowed values:

- `Default`: The most restrictive accessibility that still satisfies all uses of the asset by configurations that reference it.
- `AfterFirstUnlock`: The keychain item is only available after the first unlock of the device.

### Authentication

- **Type:** `AssetCredentialIdentityAuthenticationObject`
- **Required:** No

The server authentication details.

### Reference

- **Type:** `AssetCredentialIdentityReferenceObject`
- **Required:** Yes

The external reference. Ensure that the asset data:

- Is a JSON document that represents the `com.apple.credential.identity` credential type
- Uses a media type of `application/json`, and if it includes a `ContentType` sub-key, that sub-key media type is also `application/json`

## Discussion

Specify `com.apple.asset.credential.identity` as the declaration type.

### Asset example

```json
{
    "Type": "com.apple.asset.credential.identity",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "Reference": {
            "DataURL": "https://example.com/asset-data/certificates/www.example.com.json",
            "ContentType": "application/json"
        }
    }
}
```

## Topics

### Objects

- [AssetCredentialIdentityAuthenticationObject](/documentation/devicemanagement/assetcredentialidentityauthenticationobject) - The server authentication details for an asset-credential identity.
- [AssetCredentialIdentityReferenceObject](/documentation/devicemanagement/assetcredentialidentityreferenceobject) - A dictionary that describes the external reference.

