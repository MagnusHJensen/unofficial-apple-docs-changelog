# SettingsCommand.Command.Settings.AccessibilitySettings

A dictionary that contains settings for accessibility.

**Platforms:** iOS 16.0, iPadOS 16.0, Mac Catalyst 16.0, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### BoldTextEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables bold text.

### GrayscaleEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables grayscale display.

### IncreaseContrastEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables increase contrast.

### Item

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `AccessibilitySettings`

Sets various accessibility settings. The system allows only keys with explicitly provided values.

### ReduceMotionEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables reduced motion.

### ReduceTransparencyEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables reduced transparency.

### TextSize

- **Type:** `integer`
- **Required:** No
- **Default:** `4`
- **Allowed Values:** `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`

The accessibility text size apps that support dynamic text use. `0` is the smallest value, and `11` is the largest available.

### TouchAccommodationsEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables touch accommodations.

### VoiceOverEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables voiceover.

### ZoomEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables zoom.

