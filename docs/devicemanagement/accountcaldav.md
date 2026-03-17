# AccountCalDAV

The declaration to configure a Calendar account.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 13.0, visionOS 1.1

## Properties

### AuthenticationCredentialsAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the credentials for this account. Set the corresponding asset type to `CredentialUserNameAndPassword`.

### HostName

- **Type:** `string`
- **Required:** Yes

The hostname or IP address of the CalDAV server.

### Path

- **Type:** `string`
- **Required:** No

The path for the CalDAV server.

### Port

- **Type:** `integer`
- **Required:** No

The port number for the CalDAV server.

### VisibleName

- **Type:** `string`
- **Required:** No

The name that apps show to the user for this calendar account. If not present, the system generates a suitable default.

## Discussion

Specify `com.apple.configuration.account.caldav` as the declaration type.

### Configuration availability

### Configuration example

```json
{
    "Type": "com.apple.configuration.account.caldav",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "VisibleName": "Work Calendar",
        "HostName": "calendar.example.com",
        "Port": 8443,
        "Path": "/principals",
        "AuthenticationCredentialsAssetReference": "64BF8F5C-8CFD-40AA-9082-A0B594D4E100"
    }
}
```

