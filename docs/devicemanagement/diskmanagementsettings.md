# DiskManagementSettings

The declaration to configure disk management settings on the device.

**Platforms:** macOS 15.0

## Discussion

Specify `com.apple.configuration.diskmanagement.settings` as the declaration type.

### Configuration availability

### Configuration example

This configuration prevents the use of external and network storage devices.

```json
{
    "Type": "com.apple.configuration.diskmanagement.settings",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "Restrictions": {
            "ExternalStorage": "Disallowed",
            "NetworkStorage": "Disallowed"
        }
    }
}
```

## Topics

### Objects

- [DiskManagementSettingsRestrictionsObject](/documentation/devicemanagement/diskmanagementsettingsrestrictionsobject) - The restrictions for the disk.

