# ManagementTest

The declaration to test declarative device management.

**Platforms:** iOS 15.0, iPadOS 15.0, Mac Catalyst 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### Echo

- **Type:** `string`
- **Required:** Yes

The string to echo back in a status response reason.

### EchoDataAssetReference

- **Type:** `string`
- **Required:** No

The string to read from a data asset to echo back in status response reason description.

### ReturnStatus

- **Type:** `string`
- **Required:** No
- **Default:** `Installed`
- **Allowed Values:** `Installed`, `Failed`, `Unlocked`

The status the system reports back when the device implements the configuration. Use this to override the normal `success` result.

## Discussion

Specify `com.apple.configuration.management.test` as the declaration type.

### Configuration availability

### Configuration example

```json
{
    "Type": "com.apple.configuration.management.test",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "Echo": "Hello World!"
    }
}
```

