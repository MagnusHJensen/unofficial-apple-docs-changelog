# ExchangeWebServices

The payload that configures an Exchange Web Services accounts.

**Platforms:** macOS 10.7

## Properties

### allowMailDrop

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables Mail Drop.

### AuthenticationCertificateUUID

- **Type:** `string`
- **Required:** No

The UUID of the certificate payload within the same profile to use for the identity credential. Supported on macOS 10.11 or later. On macOS 10.12 or later use the PayloadCertificateUUID.

### EmailAddress

- **Type:** `string`
- **Required:** No

The full email address for the account. If the email address string isn’t present in the payload, the device prompts for it during profile installation.

### ExternalHost

- **Type:** `string`
- **Required:** No

The external server address.

### ExternalPath

- **Type:** `string`
- **Required:** No

The external server path.

### ExternalPort

- **Type:** `integer`
- **Required:** No

The external server port number.

### ExternalSSL

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the system enables SSL for connections to the external server.

### Host

- **Type:** `string`
- **Required:** No

The Exchange server host name or IP address. Ignored if using OAuth.

### OAuth

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables OAuth for authentication. Don’t specify a password if `OAuth` is `true`. Available in macOS 10.14 and later

### OAuthSignInURL

- **Type:** `string`
- **Required:** No

The URL to load into a web view for authentication through OAuth when autodiscovery isn’t used. This setting requires a `Host` value.

### Password

- **Type:** `string`
- **Required:** No

The password of the account. Use only with encrypted profiles.

### Path

- **Type:** `string`
- **Required:** No

The server path.

### PayloadCertificateUUID

- **Type:** `string`
- **Required:** No

The UUID of the certificate payload within the same profile to use for the identity credential. Supported on macOS 10.12 or later.

### Port

- **Type:** `integer`
- **Required:** No

The server port number.

### SSL

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the system enables SSL.

### UserName

- **Type:** `string`
- **Required:** No

The user name for this Exchange account. Required for noninteractive installation, such as through MDM. If missing, the system prompts the user for it during interactive profile installation.

## Discussion

Specify `com.apple.ews.account` as the payload type.

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
            <key>EmailAddress</key>
            <string>juanchavez4@companyemail.com</string>
            <key>ExternalHost</key>
            <string>host.example.com</string>
            <key>ExternalPath</key>
            <string></string>
            <key>ExternalSSL</key>
            <true/>
            <key>Host</key>
            <string>host.example.com</string>
            <key>MailNumberOfPastDaysToSync</key>
            <integer>0</integer>
            <key>Password</key>
            <string>Password123</string>
            <key>PayloadDisplayName</key>
            <string>Exchange Web Service</string>
            <key>PayloadIdentifier</key>
            <string>com.example.myewspayload</string>
            <key>PayloadType</key>
            <string>com.apple.ews.account</string>
            <key>PayloadUUID</key>
            <string>bb4adbbc-5516-45bb-bdee-cc7e47a5c5b5</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
            <key>SSL</key>
            <true/>
            <key>UserName</key>
            <string>juanchavez4@companyemail.com</string>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Exchange Web Service</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>ccb3eded-4486-4af8-81a3-add2d39cdfe7</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

