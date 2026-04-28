# SubscribedCalendars

The payload that configures subscribed calendars.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, visionOS 1.1, Device Assignment Services , VPP License Management 

## Properties

### SubCalAccountDescription

- **Type:** `string`
- **Required:** No

The description of the account.

### SubCalAccountHostName

- **Type:** `string`
- **Required:** Yes

The server’s address.

### SubCalAccountPassword

- **Type:** `string`
- **Required:** No

The user’s password.

### SubCalAccountUsername

- **Type:** `string`
- **Required:** No

The user’s user name.

### SubCalAccountUseSSL

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables SSL.

### VPNUUID

- **Type:** `string`
- **Required:** No

The VPNUUID of the per-app VPN the account uses for network communication. Available in iOS 14 and later.

## Discussion

Specify `com.apple.subscribedcalendar.account` as the payload type.

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
            <key>SubCalAccountDescription</key>
            <string>US Holiday Calendar</string>
            <key>SubCalAccountHostName</key>
            <string>http://ical.mac.com/ical/US32Holidays.ics</string>
            <key>SubCalAccountUseSSL</key>
            <true/>
            <key>PayloadIdentifier</key>
            <string>com.example.mysubscribedcalpayload</string>
            <key>PayloadType</key>
            <string>com.apple.subscribedcalendar.account</string>
            <key>PayloadUUID</key>
            <string>c23dd040-6f68-4af2-b840-62d1943236b5</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Subscribed Calendars</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>c360335c-3cfd-43d5-88c3-7e58e92821a9</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

