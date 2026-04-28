# ManagementServerCapabilities

The declaration to configure the server’s feature set.

**Platforms:** iOS 15.0, iPadOS 15.0, Mac Catalyst 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### SupportedFeatures

- **Type:** `ManagementServerCapabilitiesSupportedFeaturesObject`
- **Required:** Yes

A dictionary that contains the server’s optional protocol features.

Each dictionary item uses the key name to represent a feature, and the value to hold the feature’s associated parameters. This protocol reserves keys with a prefix of `com.apple.`, which appear as subkeys in this dictionary.

### Version

- **Type:** `string`
- **Required:** Yes

The server’s protocol version.

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

