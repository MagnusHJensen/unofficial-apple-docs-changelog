# StatusSoftwareUpdateInstallReason

The status item that reports the device’s pending software update reason.

**Platforms:** iOS 17.0, iPadOS 17.0, Mac Catalyst 17.0, macOS 14.0, tvOS 18.4, visionOS 26.0

## Properties

### softwareupdate.install-reason

- **Type:** `StatusSoftwareUpdateInstallReasonDictionaryObject`
- **Required:** Yes

Details about the reason for a pending software update.

## Discussion

### Status item availability

### Status item example

```json
{
    "softwareupdate": {
        "install-reason": {
            "reason": [
                "declaration"
            ],
            "declaration-id": "com.example.softwareupdate-config"
        }
    }
}
```

## Topics

### Objects

- [StatusSoftwareUpdateInstallReasonDictionaryObject](/documentation/devicemanagement/statussoftwareupdateinstallreasondictionaryobject) - Details about the reason for a pending software update.

