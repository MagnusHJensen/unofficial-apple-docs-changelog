# AirPlaySecurity

The payload that configures Apple TV for a particular style of AirPlay security.

**Platforms:** tvOS 11.0

## Properties

### AccessType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ANY`, `WIFI_ONLY`

The access policy for AirPlay.

`ANY` allows connections from both Ethernet, Wi-Fi, and Apple Wireless Direct Link.

`WIFI_ONLY` allows connections only from devices on the same Ethernet or Wi-Fi network as Apple TV.

### Password

- **Type:** `string`
- **Required:** No

The AirPlay password; required if `SecurityType` is `PASSWORD`.

### SecurityType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `PASSCODE_ONCE`, `PASSCODE_ALWAYS`, `PASSWORD`

The security policy for AirPlay. Allowed values:

> 

## Discussion

Specify `com.apple.airplay.security` as the payload type.

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
            <key>Password</key>
            <string>MyPassword</string>
            <key>PayloadIdentifier</key>
            <string>com.example.myairplaysecuritypayload</string>
            <key>PayloadType</key>
            <string>com.apple.airplay.security</string>
            <key>PayloadUUID</key>
            <string>c0b60f19-91c7-482e-9a95-6ba71220d93e</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
            <key>SecurityType</key>
            <string>PASSWORD</string>
            <key>AccessType</key>
            <string>ANY</string>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>AirPlay Security</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>40dc47ea-41c9-4c79-8b30-a027cf6eacd6</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

