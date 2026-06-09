# StatusDeviceModelMarketingName

The status item that reports the device’s model marketing name.

**Platforms:** iOS 15.0, iPadOS 15.0, Mac Catalyst 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0

## Properties

### device.model.marketing-name

- **Type:** `string`
- **Required:** Yes

The device’s marketing name, such as `iPhone 12`. This value may not always be available.

## Discussion

### Status item availability

### Status item example

```json
{
    "device": {
        "model": {
            "marketing-name": "iPhone 13"
        }
    }
}
```

