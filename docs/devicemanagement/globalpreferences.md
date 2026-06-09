# GlobalPreferences

The payload to configure global preferences.

**Platforms:** macOS 10.7

## Properties

### com.apple.autologout.AutoLogOutDelay

- **Type:** `number`
- **Required:** No

The `autologout` delay, in seconds. A value of `0` means `autologout` is off. In some cases, this delay may be restricted to values between 5 minutes and 24 hours.

### MultipleSessionEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables fast user switching.

## Discussion

Specify `.GlobalPreferences` as the payload type.

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
            <key>MultipleSessionEnabled</key>
            <false/>
            <key>com.example.autologout.AutoLogOutDelay</key>
            <real>1800</real>
            <key>LULookupDisabled</key>
            <true/>
            <key>PayloadIdentifier</key>
            <string>com.example.myglobalpayload</string>
            <key>PayloadType</key>
            <string>.GlobalPreferences</string>
            <key>PayloadUUID</key>
            <string>b5033127-c0ef-4055-8fc5-7db5a8216bc8</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Global Preferences</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>52e0e1b6-067f-407a-a36c-7e7b917b2982</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

