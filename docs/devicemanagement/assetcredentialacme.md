# AssetCredentialACME

A reference to an ACME identity.

**Platforms:** iOS 17.0, iPadOS 17.0, Mac Catalyst 17.0, macOS 14.0, tvOS 17.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

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

- **Type:** `AssetCredentialACMEAuthenticationObject`
- **Required:** No

The server authentication details.

### Reference

- **Type:** `AssetCredentialACMEReferenceObject`
- **Required:** Yes

The external reference. Ensure that the asset data:

- Is a JSON document that represents the `com.apple.credential.acme` credential type
- Uses a media type of `application/json`, and if it includes a `ContentType` sub-key, that sub-key media type is also `application/json`

## Discussion

Specify `com.apple.asset.credential.acme` as the declaration type.

### Asset example

```json
{
    "Type": "com.apple.asset.credential.acme",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "Reference": {
            "DataURL": "https://example.com/asset-data/certificates/security_acme.json",
            "ContentType": "application/json"
        }
    }
}
```

## Topics

### Objects

- [AssetCredentialACMEAuthenticationObject](/documentation/devicemanagement/assetcredentialacmeauthenticationobject) - The server authentication details for an ACME asset credential.
- [AssetCredentialACMEReferenceObject](/documentation/devicemanagement/assetcredentialacmereferenceobject) - The external reference for an ACME asset credential.

