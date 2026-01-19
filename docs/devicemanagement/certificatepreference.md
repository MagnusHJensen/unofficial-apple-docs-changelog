# CertificatePreference

The payload that configures a certificate preference.

**Platforms:** macOS 10.12

## Discussion

Specify `com.apple.security.certificatepreference` as the payload type.

A `CertificatePreference` payload lets you identify a certificate preference item in the user’s keychain that references a certificate payload included in the same profile. It can only appear in a user profile, not a device profile. You can include multiple `CertificatePreference` payloads as needed.

See also [IdentityPreference](/documentation/devicemanagement/identitypreference)  for information about setting up identity preferences.

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
            <string>Password123</string>
            <key>PayloadContent</key>
            <data>ExampleD</data>
            <key>PayloadIdentifier</key>
            <string>com.example.mypcertprefpayload</string>
            <key>PayloadType</key>
            <string>com.apple.security.certificatepreference</string>
            <key>PayloadUUID</key>
            <string>d669e184-6312-4214-9ebc-00346809f7f0</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
            <key>updated_at_xid</key>
            <integer>23387</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Certificate Preferences</string>
    <key>PayloadIdentifier</key>
    <string>com.example.profile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>89d2e0a7-0821-45c2-a566-c23eb98b137e</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

