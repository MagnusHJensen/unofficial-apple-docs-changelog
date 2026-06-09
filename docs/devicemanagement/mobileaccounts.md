# MobileAccounts

The payload that configures mobile accounts on the device.

**Platforms:** macOS 10.7

## Properties

### cachedaccounts.askForSecureTokenAuthBypass

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system bypasses the secure token authorization dialog. This dialog only appears on APFS volumes.

### cachedaccounts.expiry.delete.disusedSeconds

- **Type:** `integer`
- **Required:** No
- **Default:** `-1`

The minimum number of seconds a mobile account can exist before the system makes an automatic attempt to remove the mobile account. Set to `0` to attempt removing it at the next login or logout. Set to `-1` to never attempt removing the mobile account.

### cachedaccounts.WarnOnCreate.allowNever

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system allows the user to stop the prompts about mobile account creation every time the user logs in. This key is only valid if `com.apple.cachedaccounts.WarnOnCreate` is `true`.

### com.apple.cachedaccounts.CreateAtLogin

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system creates the mobile account at login time.

### com.apple.cachedaccounts.WarnOnCreate

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system asks the user whether to create the mobile account and it allows the user to not create it.

## Discussion

Specify `com.apple.MCX` as the payload type.

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
            <key>com.apple.cachedaccounts.CreateAtLogin</key>
            <true/>
            <key>PayloadIdentifier</key>
            <string>com.example.mymobileccountpayload</string>
            <key>PayloadType</key>
            <string>com.apple.MCX</string>
            <key>PayloadUUID</key>
            <string>93aa2058-4fe5-4f8b-a409-80f05b7fb2f0</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Mobility</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>b89ce975-801b-4994-8f68-dc5cad408ad1</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

