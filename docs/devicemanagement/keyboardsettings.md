# KeyboardSettings

The declaration to configure keyboard settings.

**Platforms:** iOS 26.4, iPadOS 26.4, macOS 26.4

## Properties

### AllowAutoCorrection

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables auto-correction.

### AllowDefinitionLookup

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables definition lookup.

### AllowDictation

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables dictation.

### AllowMathKeyboardSuggestions

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables keyboard suggestions that include math solutions. This key is also supported by the math.settings configuration.

### AllowPredictiveText

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables predictive text.

### AllowSlideToType

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables slide to type.

### AllowSpellCheck

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables spell check.

### AllowTextReplacement

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables text replacement.

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

