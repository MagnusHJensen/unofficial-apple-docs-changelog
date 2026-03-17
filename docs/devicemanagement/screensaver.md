# Screensaver

The payload that configures the screen saver.

**Platforms:** macOS 10.11

## Properties

### askForPassword

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the user is prompted for a password when the screen saver is unlocked or stopped. When you use this prompt, you must also provide `askForPasswordDelay`. Available in macOS 10.13 and later.

### askForPasswordDelay

- **Type:** `integer`
- **Required:** No

The number of seconds to delay before the password will be required to unlock or stop the screen saver (the grace period). A value of `2147483647` (for example, `0x7FFFFFFF`) disables this requirement. To use this option, you must set `askForPassword` to `true`. Available in macOS 10.13 and later.

### idleTime

- **Type:** `integer`
- **Required:** No

The number of seconds of inactivity before the screen saver activates (0 = Never activate).

### loginWindowModulePath

- **Type:** `string`
- **Required:** No

The full path to the screen-saver module to use.

### moduleName

- **Type:** `string`
- **Required:** Yes

The name of the screen saver module.

## Discussion

Specify `com.apple.screensaver` as the payload type.

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
            <key>loginWindowIdleTime</key>
            <integer>60</integer>
            <key>loginWindowModulePath</key>
            <string>/System/Library/Screen Savers/Example-Name.saver</string>
            <key>moduleName</key>
            <string>Example Name</string>
            <key>askForPassword</key>
            <true/>
            <key>askForPasswordDelay</key>
            <integer>5</integer>
            <key>PayloadIdentifier</key>
            <string>com.example.myscreensaverpayload</string>
            <key>PayloadType</key>
            <string>com.apple.screensaver</string>
            <key>PayloadUUID</key>
            <string>ba9abec1-ee44-413d-b75f-63748644ca71</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Screem Saver Device</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>4ffe721a-f2e6-4191-a3fe-1d1a463fbbac</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

