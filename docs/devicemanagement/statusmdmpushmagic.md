# StatusMDMPushMagic

The status item that reports the device’s push magic value.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), tvOS 27.0 (Beta), visionOS 27.0 (Beta), watchOS 27.0 (Beta)

## Properties

### mdm.push-magic

- **Type:** `string`
- **Required:** Yes

The push magic value that the device expects the MDM server to include in Apple Push Notification service messages.

## Discussion

### Status item availability

### Status item example

```json
{
    "mdm": {
        "push-magic": "3B5D81A2-9F4E-4B7C-A8D6-1E2F3A4B5C6D"
    }
}
```

