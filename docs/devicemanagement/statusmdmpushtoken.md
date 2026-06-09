# StatusMDMPushToken

The status item that reports the device’s push token.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), tvOS 27.0 (Beta), visionOS 27.0 (Beta), watchOS 27.0 (Beta)

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

