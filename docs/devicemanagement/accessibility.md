# Accessibility

The payload that configures the accessibility features of the device.

**Platforms:** macOS 10.9, Device Assignment Services , VPP License Management 

## Properties

### closeViewFarPoint

- **Type:** `integer`
- **Required:** No

The minimum zoom level in the Zoom options.

### closeViewHotkeysEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables “Use keyboard shortcuts” in the Zoom options.

### closeViewNearPoint

- **Type:** `integer`
- **Required:** No

The maximum zoom level in the Zoom options.

### closeViewScrollWheelToggle

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables “Use scroll gesture” in the Zoom options.

### closeViewShowPreview

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables “Show preview rectangle” in the Zoom options. Only available in macOS 10.15 and earlier.

### closeViewSmoothImages

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables “Smooth images” in the Zoom options.

### contrast

- **Type:** `number`
- **Required:** No

The contrast value in the Display options.

### flashScreen

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables “Flash the screen” in the Audio options.

### grayscale

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables “Use grayscale” in the Display options.

This option is deprecated in macOS 11.

### mouseDriver

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables Mouse Keys in the Mouse & Trackpad options.

### mouseDriverCursorSize

- **Type:** `integer`
- **Required:** No

The size of the cursor.

### mouseDriverIgnoreTrackpad

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, ignores the built-in trackpad.

### mouseDriverInitialDelay

- **Type:** `integer`
- **Required:** No

The initial delay before moving the mouse with Mouse Keys.

### mouseDriverMaxSpeed

- **Type:** `integer`
- **Required:** No

The maximum speed for the cursor when using Mouse Keys.

### slowKey

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables “Slow Keys” in the Keyboard options.

### slowKeyBeepOn

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables “click key sounds” for Slow Keys.

### slowKeyDelay

- **Type:** `integer`
- **Required:** No

The acceptance delay, in milliseconds, for Slow Keys.

### stereoAsMono

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, plays stereo audio as mono.

### stickyKey

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables Sticky Keys in the Keyboard options.

### stickyKeyBeepOnModifier

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables the beep when a modifier key is set for Sticky Keys.

### stickyKeyShowWindow

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables “Display pressed keys on screen” for Sticky Keys.

### voiceOverOnOffKey

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables Voice Over.

### whiteOnBlack

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables Invert Colors in Display Accommodations.

## Discussion

Specify `com.apple.universalaccess` as the payload type.

### Profile availability

### Profile example

```plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>PayloadContent</key>
    <array>
        <dict>
            <key>stickyKey</key>
            <true/>
            <key>PayloadIdentifier</key>
            <string>com.example.myaccessibilitypayload</string>
            <key>PayloadType</key>
            <string>com.apple.universalaccess</string>
            <key>PayloadUUID</key>
            <string>bff2939d-cb4c-4f6d-8521-e26bc7c03e96</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Accessibility</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>e7b55cc7-0d94-4045-8868-dcc1b1c58159</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

