# ManagementStatusSubscriptions

The declaration to configure status subscriptions.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0

## Discussion

Specify `com.apple.configuration.management.status-subscriptions` as the declaration type.

### Configuration availability

### Configuration example

```json
{
    "Type": "com.apple.configuration.management.status-subscriptions",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "StatusItems": [
            {
                "Name": "device.identifier.serial-number"
            },
            {
                "Name": "device.identifier.udid"
            },
            {
                "Name": "device.operating-system.build-version"
            },
            {
                "Name": "device.operating-system.version"
            },
            {
                "Name": "device.power.battery-health"
            },
            {
                "Name": "passcode.is-compliant"
            },
            {
                "Name": "security.certificate.list"
            }
        ]
    }
}
```

## Topics

### Objects

- [ManagementStatusSubscriptionsStatusItemObject](/documentation/devicemanagement/managementstatussubscriptionsstatusitemobject) - The declaration for configuring a specific status subscription.

