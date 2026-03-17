# DeviceInformationResponse.QueryResponses.AccessibilitySettings

The response dictionary that contains the devices accessibility settings.

**Platforms:** iOS 16.0, iPadOS 16.0, watchOS 10.0

## Properties

### BoldTextEnabled

- **Type:** `boolean`
- **Required:** No

If `true`, the device has enabled bold text.

### GrayscaleEnabled

- **Type:** `boolean`
- **Required:** No

If `true`, the device has enabled grayscale display.

### IncreaseContrastEnabled

- **Type:** `boolean`
- **Required:** No

If `true`, the device has enabled increase contrast.

### ReduceMotionEnabled

- **Type:** `boolean`
- **Required:** No

If `true`, the device has enabled reduced motion.

### ReduceTransparencyEnabled

- **Type:** `boolean`
- **Required:** No

If `true`, the device has enabled reduced transparency.

### TextSize

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `-1`, `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`

The accessibility text size apps that support dynamic text use. 0 is the smallest value, and 11 is the largest available.

`-1` indicates that the current size is unknown or hasn’t been explicitly set.

### TouchAccommodationsEnabled

- **Type:** `boolean`
- **Required:** No

If `true`, the device has enabled touch accommodations.

### VoiceOverEnabled

- **Type:** `boolean`
- **Required:** No

If `true`, the device has enabled voiceover.

### ZoomEnabled

- **Type:** `boolean`
- **Required:** No

If `true`, the device has enabled zoom.

