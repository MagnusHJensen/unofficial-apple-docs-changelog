# StatusMDMPushToken

The status item that reports the device’s push token.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, tvOS 27.0, visionOS 27.0, watchOS 27.0

## Properties

### mdm.push-token

- **Type:** `string`
- **Required:** Yes

The device push token that the MDM server uses for Apple Push Notification service messages.

## Discussion

### Status item availability

### Status item example

```json
{
    "mdm": {
        "push-token": "4A8B3F2E1D9C7B6A5E4D3C2B1A0F9E8D7C6B5A4E3D2C1B0A"
    }
}
```

