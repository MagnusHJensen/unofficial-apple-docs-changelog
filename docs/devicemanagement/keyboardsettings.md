# KeyboardSettings

The declaration to configure keyboard settings.

**Platforms:** iOS 26.4 (Beta), iPadOS 26.4 (Beta), macOS 26.4 (Beta)

## Discussion

Specify `com.apple.configuration.keyboard.settings` as the declaration type.

### Configuration availability

### Configuration examples

This configuration restricts keyboard features.

```json
{
    "Type": "com.apple.configuration.keyboard.settings",
    "Identifier": "A1B2C3D4-E5F6-4A5B-9C8D-7E6F5A4B3C2D",
    "ServerToken": "F1E2D3C4-B5A6-4D5E-8F9A-0B1C2D3E4F5A",
    "Payload": {
        "AllowAutoCorrection": false,
        "AllowSlideToType": false,
        "AllowDefinitionLookup": false,
        "AllowDictation": false,
        "AllowMathKeyboardSuggestions": false,
        "AllowPredictiveText": false,
        "AllowTextReplacement": false,
        "AllowSpellCheck": false
    }
}
```

