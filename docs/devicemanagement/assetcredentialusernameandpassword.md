# AssetCredentialUserNameAndPassword

A reference to data that describes a credential that represents a user name and password.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0

## Discussion

Specify `com.apple.asset.credential.userpassword` as the declaration type.

### Asset example

```json
{
    "Type": "com.apple.asset.credential.userpassword",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "Reference": {
            "DataURL": "https://example.com/asset-data/credential.json",
            "ContentType": "application/json"
        }
    }
}
```

## Topics

### Objects

- [AssetCredentialUserNameAndPasswordAuthenticationObject](/documentation/devicemanagement/assetcredentialusernameandpasswordauthenticationobject) - The server authentication details for an asset-credential user name and password.
- [AssetCredentialUserNameAndPasswordReferenceObject](/documentation/devicemanagement/assetcredentialusernameandpasswordreferenceobject) - The external reference for an asset-credential user name and password.

