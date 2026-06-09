# StatusDeviceModelFamily

The status item that reports the device’s hardware model family.

**Platforms:** iOS 15.0, iPadOS 15.0, Mac Catalyst 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0

## Properties

### device.model.family

- **Type:** `string`
- **Required:** Yes

The hardware family of the device, such as `Mac`, `iPhone`, or `iPad`.

## Discussion

### Status item availability

### Status item example

```json
{
    "device": {
        "model": {
            "family": "iPhone"
        }
    }
}
```

