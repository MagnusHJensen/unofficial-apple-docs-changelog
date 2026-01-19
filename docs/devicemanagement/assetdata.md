# AssetData

A reference to arbitrary data with a specific media type.

**Platforms:** iOS 17.0, iPadOS 17.0, macOS 14.0, tvOS 17.0, visionOS 1.1, watchOS 10.0

## Discussion

Specify `com.apple.asset.data` as the declaration type.

### Asset example

```json
{
    "Type": "com.apple.asset.data",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "Reference": {
            "DataURL": "https://example.com/asset-data/data/test.txt",
            "ContentType": "text/plain"
        },
        "Authentication": {
            "Type": "MDM"
        }
    }
}
```

## Topics

### Objects

- [AssetDataAuthenticationObject](/documentation/devicemanagement/assetdataauthenticationobject) - The server authentication details for an asset data.
- [AssetDataReferenceObject](/documentation/devicemanagement/assetdatareferenceobject) - The external reference for an asset data.

