# AssetCredentialCertificate

A reference to a PKCS #1 or PEM encoded certificate.

**Platforms:** iOS 17.0, iPadOS 17.0, macOS 14.0, tvOS 17.0, visionOS 1.1, watchOS 10.0

## Discussion

Specify `com.apple.asset.credential.certificate` as the declaration type.

### Asset example

```json
{
    "Type": "com.apple.asset.credential.certificate",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "Reference": {
            "DataURL": "https://example.com/asset-data/certificates/cert.pem",
            "ContentType": "application/pem"
        }
    }
}
```

## Topics

### Objects

- [AssetCredentialCertificateAuthenticationObject](/documentation/devicemanagement/assetcredentialcertificateauthenticationobject) - The server authentication details for an asset-credential certificate.
- [AssetCredentialCertificateReferenceObject](/documentation/devicemanagement/assetcredentialcertificatereferenceobject) - The external reference for an asset-credential certificate.

