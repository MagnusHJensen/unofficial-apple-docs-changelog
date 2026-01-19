# SecurityCertificate

The declaration to add a certificate to the device.

**Platforms:** iOS 17.0, iPadOS 17.0, macOS 14.0, tvOS 17.0, visionOS 1.1, watchOS 10.0

## Discussion

Specify `com.apple.configuration.security.certificate` as the declaration type.

### Configuration availability

### Configuration example

```json
{
    "Type": "com.apple.configuration.security.certificate",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "CredentialAssetReference": "A762DE40-2C43-4FE2-BC2E-B6EB2B571CF6"
    }
}
```

