# AccessibilitySettings

The declaration to configure accessibility settings.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### Vision

- **Type:** `AccessibilitySettingsVisionObject`
- **Required:** No

If present, configures vision accessibility settings.

## Discussion

Specify `com.apple.configuration.accessibility.settings` as the declaration type.

### Configuration availability

### Configuration example

This configuration prevents the use of Live Recognition.

```json
{
    "Type": "com.apple.configuration.accessibility.settings",
    "Identifier": "119D31F8-E3A2-454A-A019-FD3F05A008D3",
    "ServerToken": "DC31F056-0ADE-4D01-8E63-A7CA093EC7CD",
    "Payload": {
        "Vision": {
            "AllowLiveRecognition": false
        }
    }
}
```

## Topics

### Objects

- [AccessibilitySettingsVisionObject](/documentation/devicemanagement/accessibilitysettingsvisionobject) - If present, configures vision accessibility settings.

