# SiriSettings

The declaration to configure Siri settings.

**Platforms:** iOS 26.4, iPadOS 26.4, Mac Catalyst 26.4, macOS 26.4, tvOS 27.0 (Beta), visionOS 26.4, watchOS 26.4

## Properties

### AllowSiriAI

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables the Siri AI specific features of Siri.

Available: iOS 27+ | iPadOS 27+ | macOS 27+ | visionOS 27+

### AllowUserGeneratedContent

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables Siri user-generated content.

Available: iOS 26.4+ | iPadOS 26.4+ | watchOS 26.4+

### AllowWhileLocked

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables Siri while locked.

Available: iOS 26.4+ | iPadOS 26.4+ | watchOS 26.4+

### Enabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables Siri.

Available: iOS 26.4+ | iPadOS 26.4+ | macOS 26.4+ | tvOS 27+ | visionOS 26.4+

### ForceProfanityFilter

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, forces Siri profanity filter.

Available: iOS 26.4+ | iPadOS 26.4+ | macOS 26.4+

### ForceReduceSensitiveContent

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, forces Siri to reduce sensitive content.

Available: iOS 27+ | iPadOS 27+ | macOS 27+

## Discussion

Specify `com.apple.configuration.siri.settings` as the declaration type.

### Configuration availability

### Configuration examples

This configuration restricts Siri features.

```json
{
    "Type": "com.apple.configuration.siri.settings",
    "Identifier": "A1B2C3D4-E5F6-4A5B-9C8D-7E6F5A4B3C2D",
    "ServerToken": "F1E2D3C4-B5A6-4D5E-8F9A-0B1C2D3E4F5A",
    "Payload": {
        "Enabled": false,
        "AllowSiriAI": false,
        "AllowUserGeneratedContent": false,
        "AllowWhileLocked": false,
        "ForceProfanityFilter": true,
        "ForceReduceSensitiveContent": true
    }
}
```

