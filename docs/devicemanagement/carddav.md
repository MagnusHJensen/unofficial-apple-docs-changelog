# CardDAV

The payload that configures a Contacts account.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, visionOS 1.1, Device Assignment Services , VPP License Management 

## Properties

### CardDAVAccountDescription

- **Type:** `string`
- **Required:** No

The description of the account.

### CardDAVHostName

- **Type:** `string`
- **Required:** Yes

The server’s address.

### CardDAVPassword

- **Type:** `string`
- **Required:** No

The user’s password. Only use this in encrypted profiles.

### CardDAVPort

- **Type:** `integer`
- **Required:** No

The server’s port.

### CardDAVPrincipalURL

- **Type:** `string`
- **Required:** No

The base URL to the user’s address book.

### CardDAVUsername

- **Type:** `string`
- **Required:** No

The user name for logins.

### CardDAVUseSSL

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the system enables SSL.

### CommunicationServiceRules

- **Type:** `CardDAV.CommunicationServiceRules`
- **Required:** No

An array of communication service rules for this account.

### VPNUUID

- **Type:** `string`
- **Required:** No

The VPNUUID of the per-app VPN the account uses for network communication. Available in iOS 14 and later.

## Discussion

Specify `com.apple.carddav.account` as the payload type.

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
            <key>CardDAVAccountDescription</key>
            <string>My CardDAV Account</string>
            <key>CardDAVHostName</key>
            <string>server.example.com</string>
            <key>CardDAVPassword</key>
            <string>Password123</string>
            <key>CardDAVPort</key>
            <integer>443</integer>
            <key>CardDAVUseSSL</key>
            <true/>
            <key>CardDAVUsername</key>
            <string>juanchavez4</string>
            <key>PayloadIdentifier</key>
            <string>com.example.mycardavpayload</string>
            <key>PayloadType</key>
            <string>com.apple.carddav.account</string>
            <key>PayloadUUID</key>
            <string>b23d14e3-2f9d-4087-a819-747903fbb176</string>
            <key>PayloadVersion</key>
            <real>1</real>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>CardDAV</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>4fdb234d-8c58-48de-a76d-9ed9d241d273</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [CardDAV.CommunicationServiceRules](/documentation/devicemanagement/carddav/communicationservicerules-data.dictionary) - The communication service handler rules for this account.

