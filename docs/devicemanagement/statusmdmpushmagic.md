# StatusMDMPushMagic

The status item that reports the device’s push magic value.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, tvOS 27.0, visionOS 27.0, watchOS 27.0

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

