# SetupAssistant

The payload that configures Setup Assistant settings.

**Platforms:** iOS 14.0, iPadOS 14.0, macOS 10.12

## Properties

### SkipAccessibility

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the Accessibility pane.

### SkipAppearance

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the Choose Your Look pane.

### SkipCloudSetup

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the Apple Account setup pane.

### SkipiCloudStorageSetup

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the iCloud Storage pane.

### SkipPrivacySetup

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the Privacy consent pane.

### SkipScreenTime

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the Screen Time pane.

### SkipSetupItems

- **Type:** `[string]`
- **Required:** No

An array of strings that describe the setup items to skip. [SkipKeys](/documentation/devicemanagement/skipkeys) provides a list of valid strings and their meanings. Available in iOS 14 and later, and macOS 15 and later.

### SkipSiriSetup

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the Siri setup pane.

### SkipTouchIDSetup

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the Touch ID setup pane.

### SkipTrueTone

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the True Tone Display pane.

### SkipUnlockWithWatch

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system skips the Unlock With Apple Watch pane.

### SkipWallpaper

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If ‘true’, the system skips the Wallpaper selection window.

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

