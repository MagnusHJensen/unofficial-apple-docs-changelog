# StatusDeviceOperatingSystemSupplementalBuildVersion

The status item that reports the device’s operating system supplemental build version and Background Security Improvement version.

**Platforms:** iOS 16.1, iPadOS 16.1, Mac Catalyst 16.1, macOS 13.0, tvOS 16.1, visionOS 1.1, watchOS 10.0

## Properties

### device.operating-system.supplemental.build-version

- **Type:** `string`
- **Required:** Yes

The operating system’s build and Background Security Improvement versions in use on the device, for example, `20A123a` or `20B27c`.

## Discussion

### Status item availability

### Status item example

```json
{
    "device": {
        "operating-system": {
            "supplemental": {
                "build-version": "24A346a"
            }
        }
    }
}
```

