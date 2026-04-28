# IntelligenceSettings

The declaration to configure Apple Intelligence settings.

**Platforms:** iOS 26.4, iPadOS 26.4, Mac Catalyst 26.4, macOS 26.4, visionOS 26.4, Device Assignment Services , VPP License Management 

## Properties

### AllowAppleIntelligenceReport

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables Apple Intelligence Report.

### AllowGenmoji

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables Genmoji.

### AllowImagePlayground

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables Image Playground.

### AllowImageWand

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables Image Wand.

### AllowPersonalizedHandwritingResults

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables Personalized Handwriting Results.

### AllowVisualIntelligenceSummary

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables Visual Intelligence Summary.

### AllowWritingTools

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables Writing Tools.

### Apps

- **Type:** `IntelligenceSettingsAppsObject`
- **Required:** No

If present, configures app-specific Intelligence features.

### ForceOnDeviceOnlyDictation

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, forces On-Device Only Dictation.

### ForceOnDeviceOnlyTranslation

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, forces On-Device Only Translation.

## Discussion

Specify `com.apple.configuration.intelligence.settings` as the declaration type.

### Configuration availability

### Configuration examples

This configuration restricts several Apple Intelligence features.

```json
{
    "Type": "com.apple.configuration.intelligence.settings",
    "Identifier": "A1B2C3D4-E5F6-4A5B-9C8D-7E6F5A4B3C2D",
    "ServerToken": "F1E2D3C4-B5A6-4D5E-8F9A-0B1C2D3E4F5A",
    "Payload": {
        "AllowAppleIntelligenceReport": false,
        "AllowGenmoji": false,
        "AllowImagePlayground": false,
        "AllowImageWand": false,
        "Apps": {
            "Mail": {
                "AllowSmartReplies": false,
                "AllowSummary": false
            },
            "Notes": {
                "AllowTranscription": false,
                "AllowTranscriptionSummary": false
            },
            "Safari": {
                "AllowSummary": false
            }
        },
        "AllowPersonalizedHandwritingResults": false,
        "AllowVisualIntelligenceSummary": false,
        "AllowWritingTools": false,
        "ForceOnDeviceOnlyDictation": true,
        "ForceOnDeviceOnlyTranslation": true
    }
}
```

## Topics

### Objects

- [IntelligenceSettingsAppsObject](/documentation/devicemanagement/intelligencesettingsappsobject) - If present, configures app-specific Intelligence features.

