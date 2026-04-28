# ScreenSharingConnectionGroup

The declaration to configure a group of screen-sharing connections.

**Platforms:** macOS 14.0, Device Assignment Services , VPP License Management 

## Properties

### ConnectionGroupUUID

- **Type:** `string`
- **Required:** Yes

A unique identifier for this connection group.

### GroupName

- **Type:** `string`
- **Required:** Yes

The name of the connection group.

### Members

- **Type:** `[string]`
- **Required:** Yes

An array of `ConnectionUUID`s that represent connections declared in `ScreenSharingConnection` configurations that are members of this group.

## Discussion

Specify `com.apple.configuration.screensharing.connection.group` as the declaration type.

### Configuration availability

### Configuration example

```json
{
    "Type": "com.apple.configuration.screensharing.connection.group",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "ConnectionGroupUUID": "4BF32552-6F85-4F67-999B-7B2494C4DD99",
        "GroupName": "Lab Devices",
        "Members": [
            "7F8F28A5-C024-470B-8166-EC6669A12C3A",
            "EDF04C7F-F9B0-4204-A5EE-34AE094CA7BB"
        ]
    }
}
```

