# ScreenSharingHostSettings

The declaration to configure screen-sharing host settings and restrictions.

**Platforms:** macOS 14.0, Device Assignment Services , VPP License Management 

## Properties

### MaximumVirtualDisplays

- **Type:** `integer`
- **Required:** No

The maximum number of virtual displays to make available to clients.

### PortBase

- **Type:** `integer`
- **Required:** No

The initial UDP port number to connect to the host. Screen sharing requires multiple connections, so the system increments this value by 1 for each additional connection. This doesn’t change the port number that the system uses to initially establish a connection with a host, which is always TCP port 5900.

### PreventCopyFilesFromHost

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system prevents users from copying files from the screen-sharing host.

### PreventCopyFilesToHost

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system prevents users from copying files to the screen-sharing host.

### PreventHighPerformanceConnections

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system prevents clients from establishing high-performance connections to the host.

## Discussion

Specify `com.apple.configuration.screensharing.host.settings` as the declaration type.

### Configuration availability

### Configuration example

```json
{
    "Type": "com.apple.configuration.screensharing.host.settings",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "MaximumVirtualDisplays": 1,
        "PortBase": 1100,
        "PreventCopyFilesFromHost": true,
        "PreventCopyFilesToHost": true,
        "PreventHighPerformanceConnections": true
    }
}
```

