# ManagementServerCapabilities

The declaration to configure the server’s feature set.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0

## Discussion

Specify `com.apple.management.server-capabilities` as the declaration type.

### Management declaration example

```json
{
    "Type": "com.apple.management.server-capabilities",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "Version": "1.0.0",
        "SupportedFeatures": {
            "Example Feature": {
                "parameter1": 1
            }
        }
    }
}
```

## Topics

### Objects

- [ManagementServerCapabilitiesSupportedFeaturesObject](/documentation/devicemanagement/managementservercapabilitiessupportedfeaturesobject) - A dictionary that contains the server’s optional protocol features.

