# AssetCredentialSCEP

A reference to a SCEP identity.

**Platforms:** iOS 17.0, iPadOS 17.0, macOS 14.0, tvOS 17.0, visionOS 1.1, watchOS 10.0

## Discussion

Specify `com.apple.asset.credential.scep` as the declaration type.

### Asset example

```json
{
    "Type": "com.apple.asset.credential.scep",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "Reference": {
            "DataURL": "https://example.com/asset-data/certificates/security_scep.json",
            "ContentType": "application/json"
        }
    }
}
```

## Topics

### Objects

- [AssetCredentialSCEPAuthenticationObject](/documentation/devicemanagement/assetcredentialscepauthenticationobject) - The server authentication details for a SCEP asset credential.
- [AssetCredentialSCEPReferenceObject](/documentation/devicemanagement/assetcredentialscepreferenceobject) - The external reference for a SCEP asset credential.

