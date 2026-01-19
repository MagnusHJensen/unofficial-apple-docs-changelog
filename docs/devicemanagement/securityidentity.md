# SecurityIdentity

The declaration to install an identity on the device.

**Platforms:** iOS 17.0, iPadOS 17.0, macOS 14.0, tvOS 17.0, visionOS 1.1, watchOS 10.0

## Discussion

Specify `com.apple.configuration.security.identity` as the declaration type.

### Configuration availability

### Configuration example

```json
{
    "Type": "com.apple.configuration.security.identity",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "CredentialAssetReference": "48FFFD93-AB1E-4726-979E-421FD3265AA8"
    }
}
```

