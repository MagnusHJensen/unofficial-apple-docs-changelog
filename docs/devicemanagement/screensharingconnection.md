# ScreenSharingConnection

The declaration to configure a connection to a screen-sharing host.

**Platforms:** macOS 14.0

## Discussion

Specify `com.apple.configuration.screensharing.connection` as the declaration type.

### Configuration availability

### Configuration example

```json
{
    "Type": "com.apple.configuration.screensharing.connection",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "ConnectionUUID": "FA80E209-B31B-4862-B880-399F79E8FC35",
        "HostName": "example.com",
        "DisplayName": "Host1",
        "DisplayConfiguration": {
            "DisplayType": "Virtual1",
            "UseHDR": true
        }
    }
}
```

## Topics

### Objects

- [ScreenSharingConnectionDisplayConfigurationObject](/documentation/devicemanagement/screensharingconnectiondisplayconfigurationobject) - The display configuration for this connection.

