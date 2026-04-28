# ScreenSharingConnection

The declaration to configure a connection to a screen-sharing host.

**Platforms:** macOS 14.0, Device Assignment Services , VPP License Management 

## Properties

### AuthenticationCredentialsAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the required credentials for this connection to authenticate with the screen-sharing server. Set the corresponding asset type to `com.apple.asset.credential.userpassword`.

### ConnectionUUID

- **Type:** `string`
- **Required:** Yes

A unique identifier for this connection when it’s in a connection group.

### DisplayConfiguration

- **Type:** `ScreenSharingConnectionDisplayConfigurationObject`
- **Required:** Yes

The display configuration for this connection.

### DisplayName

- **Type:** `string`
- **Required:** Yes

The name of the connection.

### HostName

- **Type:** `string`
- **Required:** Yes

The host name or IP address of the Mac that hosts the screen-sharing connection.

### Port

- **Type:** `integer`
- **Required:** No

The TCP port number on the host to initiate the connection.

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

