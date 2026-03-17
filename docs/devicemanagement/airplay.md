# AirPlay

The payload that configures AirPlay settings.

**Platforms:** iOS 7.0, iPadOS 7.0, macOS 10.10

## Properties

### AllowList

- **Type:** `[AirPlay.AllowListItem]`
- **Required:** No

If present, only AirPlay destinations in this list are available to the device. This allow list applies to supervised devices.

### Passwords

- **Type:** `[AirPlay.PasswordsItem]`
- **Required:** No

If present, sets passwords for known AirPlay destinations. Using multiple entries for the same destination, whether within the same payload or across multiple installed payloads, is an error and results in undefined behavior.

### Whitelist

- **Type:** `[AirPlay.AllowListItem]`
- **Required:** No

Use `AllowList` instead. This key is deprecated in iOS 14.5 and macOS 11.3.

## Discussion

Specify `com.apple.airplay` as the payload type.

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
            <key>Passwords</key>
            <array>
                <dict>
                    <key>DeviceName</key>
                    <string>Juan&apos;s Apple TV</string>
                    <key>Password</key>
                    <string>Password123</string>
                </dict>
            </array>
            <key>Whitelist</key>
            <array>
                <dict>
                    <key>DeviceID</key>
                    <string>64:C7:53:EB:DB:B4</string>
                </dict>
            </array>
            <key>PayloadIdentifier</key>
            <string>com.example.myairplaysettingspayload</string>
            <key>PayloadType</key>
            <string>com.apple.airplay</string>
            <key>PayloadUUID</key>
            <string>dfb67669-67eb-4ed4-ad4f-1f8861029b42</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>AirPlay</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>c6ec425e-b224-45fb-8889-7475cb3814cb</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [AirPlay.AllowListItem](/documentation/devicemanagement/airplay/allowlistitem) - The dictionary that defines allowed destinations.
- [AirPlay.PasswordsItem](/documentation/devicemanagement/airplay/passwordsitem) - The dictionary that defines passwords for AirPlay destinations.

