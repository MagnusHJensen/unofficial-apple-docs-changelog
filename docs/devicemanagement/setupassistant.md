# SetupAssistant

The payload that configures Setup Assistant settings.

**Platforms:** iOS 14.0, iPadOS 14.0, Mac Catalyst 14.0, macOS 10.12

## Properties

### SkipAccessibility

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the Accessibility pane.

Available: macOS 11+
Deprecated: macOS 15+

### SkipAppearance

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the Choose Your Look pane.

Available: macOS 10.14+
Deprecated: macOS 15+

### SkipCloudSetup

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the Apple Account setup pane.

Available: macOS 10.12+
Deprecated: macOS 15+

### SkipiCloudStorageSetup

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the iCloud Storage pane.

Available: macOS 10.13.4+
Deprecated: macOS 15+

### SkipPrivacySetup

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the Privacy consent pane.

Available: macOS 10.13.4+
Deprecated: macOS 15+

### SkipScreenTime

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the Screen Time pane.

Available: macOS 10.15+
Deprecated: macOS 15+

### SkipSetupItems

- **Type:** `[string]`
- **Required:** No

An array of strings that describe the setup items to skip. [SkipKeys](/documentation/devicemanagement/skipkeys) provides a list of valid strings and their meanings.

Available: iOS 14+ | iPadOS 14+ | macOS 15+

### SkipSiriSetup

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the Siri setup pane.

Available: macOS 10.12+
Deprecated: macOS 15+

### SkipTouchIDSetup

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the Touch ID setup pane.

Available: macOS 10.15+
Deprecated: macOS 15+

### SkipTrueTone

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the True Tone Display pane.

Available: macOS 10.13.6+
Deprecated: macOS 15+

### SkipUnlockWithWatch

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the Unlock With Apple Watch pane.

Available: macOS 12+
Deprecated: macOS 15+

### SkipWallpaper

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If ‘true’, the system skips the Wallpaper selection window.

Available: macOS 14.1+
Deprecated: macOS 15+

## Discussion

Specify `com.apple.SetupAssistant.managed` as the payload type.

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
            <key>SkipCloudSetup</key>
            <true/>
            <key>SkipSiriSetup</key>
            <true/>
            <key>SkipPrivacySetup</key>
            <true/>
            <key>SkipiCloudStorageSetup</key>
            <true/>
            <key>SkipTrueTone</key>
            <true/>
            <key>SkipAppearance</key>
            <true/>
            <key>SkipTouchIDSetup</key>
            <true/>
            <key>SkipScreenTime</key>
            <true/>
            <key>SkipAccessibility</key>
            <true/>
            <key>PayloadIdentifier</key>
            <string>com.example.mysetupassistantpayload</string>
            <key>PayloadType</key>
            <string>com.apple.SetupAssistant.managed</string>
            <key>PayloadUUID</key>
            <string>0dfccedc-e28f-4df5-bca7-a0807deab543</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Setup Assistant</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>4a66b685-604a-4558-92c7-ae3e082cf0ae</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

