# GoogleAccount

The payload that configures a Google account.

**Platforms:** iOS 9.3, iPadOS 9.3, Mac Catalyst 9.3, visionOS 1.1, Device Assignment Services , VPP License Management 

## Properties

### AccountDescription

- **Type:** `string`
- **Required:** No

A user-visible description of the Google account, shown in the Mail and Settings apps.

### AccountName

- **Type:** `string`
- **Required:** No

The user’s full name for the Google account. This name appears in sent messages.

### CommunicationServiceRules

- **Type:** `GoogleAccount.CommunicationServiceRules`
- **Required:** No

The communication service handler rules for this account.

### EmailAddress

- **Type:** `string`
- **Required:** Yes

The full Google email address for the account.

### VPNUUID

- **Type:** `string`
- **Required:** No

The VPNUUID of the per-app VPN the account uses for network communication. Available in iOS 14 and later.

## Discussion

Specify `com.apple.google-oauth` as the payload type.

You can install multiple Google payloads. Each sets up a Google email address and any other Google services the user enables after authentication.

> 

The payload never contains credentials; the system prompts the user to enter credentials shortly after installation of the payload.

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
            <key>AccountDescription</key>
            <string>Google Account</string>
            <key>AccountName</key>
            <string>Juan Chavez</string>
            <key>EmailAddress</key>
            <string>juanchavez4@example.com</string>
            <key>PayloadIdentifier</key>
            <string>com.example.mygoogleaccountpayload</string>
            <key>PayloadType</key>
            <string>com.apple.google-oauth</string>
            <key>PayloadUUID</key>
            <string>0fe8f4dc-8cf0-4da3-9d5b-f734efb98a59</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>GoogleAccount</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>80141025-50d3-4a38-9387-ca610ff3a247</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [GoogleAccount.CommunicationServiceRules](/documentation/devicemanagement/googleaccount/communicationservicerules-data.dictionary) - The communication service handler rules for this account.

