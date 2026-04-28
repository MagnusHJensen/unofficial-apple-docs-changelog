# CalDAV

The payload that configures a Calendar account.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, visionOS 1.1, Device Assignment Services , VPP License Management 

## Properties

### CalDAVAccountDescription

- **Type:** `string`
- **Required:** No

The description of the account.

### CalDAVHostName

- **Type:** `string`
- **Required:** Yes

The server’s address.

### CalDAVPassword

- **Type:** `string`
- **Required:** No

The user’s password. Only use this in encrypted profiles.

### CalDAVPort

- **Type:** `integer`
- **Required:** No

The server’s port.

### CalDAVPrincipalURL

- **Type:** `string`
- **Required:** No

The base URL to the user’s calendar.

### CalDAVUsername

- **Type:** `string`
- **Required:** No

The user name for logins. If this profile is part of a non-interactive install, the system requires this field.

### CalDAVUseSSL

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the system enables SSL.

### VPNUUID

- **Type:** `string`
- **Required:** No

The VPNUUID of the per-app VPN the account uses for network communication. Available in iOS 14 and later.

## Discussion

Specify `com.apple.caldav.account` as the payload type.

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
            <key>CalDAVAccountDescription</key>
            <string>My CalDAV Account</string>
            <key>CalDAVHostName</key>
            <string>server.example.com</string>
            <key>CalDAVPassword</key>
            <string>Password123</string>
            <key>CalDAVPort</key>
            <integer>443</integer>
            <key>CalDAVUseSSL</key>
            <true/>
            <key>CalDAVUsername</key>
            <string>juanchavez4@example.com</string>
            <key>PayloadIdentifier</key>
            <string>com.example.mycaldavpayload</string>
            <key>PayloadType</key>
            <string>com.apple.caldav.account</string>
            <key>PayloadUUID</key>
            <string>603409f1-b611-459d-9584-0ed12bc25b5b</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>CalDAV</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>5c8bb406-a74c-4338-93c6-b403a040cc91</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

