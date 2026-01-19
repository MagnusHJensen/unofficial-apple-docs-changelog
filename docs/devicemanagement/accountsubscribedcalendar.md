# AccountSubscribedCalendar

The declaration to configure a subscribed calendar.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 14.0, visionOS 1.1

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

