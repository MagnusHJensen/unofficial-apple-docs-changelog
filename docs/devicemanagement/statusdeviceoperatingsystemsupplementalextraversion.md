# StatusDeviceOperatingSystemSupplementalExtraVersion

The status item that reports the device’s operating system Background Security Improvement version.

**Platforms:** iOS 16.1, iPadOS 16.1, Mac Catalyst 16.1, macOS 13.0, tvOS 16.1, visionOS 1.1, watchOS 10.0

## Properties

### device.operating-system.supplemental.extra-version

- **Type:** `string`
- **Required:** Yes

The operating system’s Background Security Improvement version in use on the device, for example, `a`.

## Discussion

### Status item availability

### Status item example

```json
{
    "device": {
        "operating-system": {
            "supplemental": {
                "extra-version": "a"
            }
        }
    }
}
```

