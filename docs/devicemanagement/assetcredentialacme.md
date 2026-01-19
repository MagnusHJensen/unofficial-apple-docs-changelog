# AssetCredentialACME

A reference to an ACME identity.

**Platforms:** iOS 17.0, iPadOS 17.0, macOS 14.0, tvOS 17.0, visionOS 1.1, watchOS 10.0

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

