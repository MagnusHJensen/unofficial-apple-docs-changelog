# StatusSoftwareUpdatePendingVersion

The status item that reports the device’s pending software update version.

**Platforms:** iOS 17.0, iPadOS 17.0, Mac Catalyst 17.0, macOS 14.0, tvOS 18.4, visionOS 26.0

## Properties

### softwareupdate.pending-version

- **Type:** `StatusSoftwareUpdatePendingVersionDictionaryObject`
- **Required:** Yes

A dictionary that contains the build and OS versions of the software update that’s pending on the device.

## Discussion

### Status item availability

### Status item example

```json
{
    "softwareupdate": {
        "pending-version": {
            "os-version": "27.1",
            "build-version": "24B32"
        }
    }
}
```

## Topics

### Objects

- [StatusSoftwareUpdatePendingVersionDictionaryObject](/documentation/devicemanagement/statussoftwareupdatependingversiondictionaryobject) - A dictionary that contains the build and OS versions of the software update that’s pending on the device.

