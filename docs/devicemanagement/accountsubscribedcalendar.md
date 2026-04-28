# AccountSubscribedCalendar

The declaration to configure a subscribed calendar.

**Platforms:** iOS 15.0, iPadOS 15.0, Mac Catalyst 15.0, macOS 14.0, visionOS 1.1, Device Assignment Services , VPP License Management 

## Properties

### AuthenticationCredentialsAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the credentials for this account to authenticate with a calendar server. Set the corresponding asset type to `CredentialUserNameAndPassword`.

### CalendarURL

- **Type:** `string`
- **Required:** Yes

The URL of the subscribed calendar, which needs to start with `https://`.

### VisibleName

- **Type:** `string`
- **Required:** No

The name that apps show to the user for this calendar account. If not present, the system generates a suitable default.

## Discussion

Specify `com.apple.configuration.account.subscribed-calendar` as the declaration type.

### Configuration availability

### Configuration example

```json
{
    "Type": "com.apple.configuration.account.subscribed-calendar",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "VisibleName": "Company Holidays",
        "CalendarURL": "https://calendar.example.com/holidays.ics/"
    }
}
```

