# ScreensaverUser

The payload that configures a user’s screen saver settings.

**Platforms:** macOS 10.11

## Properties

### idleTime

- **Type:** `integer`
- **Required:** No

The number of seconds of inactivity before the screen saver activates (`0` = Never activate).

### moduleName

- **Type:** `string`
- **Required:** Yes

The name of the screen saver module.

### modulePath

- **Type:** `string`
- **Required:** No

A full path to the screen saver module to use.

## Discussion

Specify `com.apple.screensaver.user` as the payload type.

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
            <key>idleTime</key>
            <integer>60</integer>
            <key>modulePath</key>
            <string>/System/Library/Screen Savers/Name.saver</string>
            <key>PayloadIdentifier</key>
            <string>com.example.myscreensaverpayload</string>
            <key>PayloadType</key>
            <string>com.apple.screensaver.user</string>
            <key>PayloadUUID</key>
            <string>c5dceece-f633-44e6-b899-9d46631fd6e5</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Screen Saver User</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>59332e00-d3c5-4d2a-a23a-ba36e45034e9</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

