# ExtensibleSingleSignOnKerberos

The payload that configures an app extension that performs single sign-on with the Kerberos extension.

**Platforms:** iOS 13.0, iPadOS 13.0, macOS 10.15, visionOS 1.1

## Discussion

Specify `com.apple.extensiblesso` as the payload type.

This is a version of the profile that defines the specific keys and values needed for the Kerberos extension.

The system supports user channel installation in macOS 11 and later.

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
            <key>ExtensionData</key>
            <dict>
                <key>useSiteAutoDiscovery</key>
                <true/>
            </dict>
            <key>ExtensionIdentifier</key>
            <string>com.apple.Extension</string>
            <key>TeamIdentifier</key>
            <string>RandomTeamID</string>
            <key>Hosts</key>
            <array>
                <string>url.example.com</string>
            </array>
            <key>Realm</key>
            <string>COM.URL.COM</string>
            <key>Type</key>
            <string>Credential</string>
            <key>PayloadIdentifier</key>
            <string>com.example.myessokpayload</string>
            <key>PayloadType</key>
            <string>com.apple.extensiblesso</string>
            <key>PayloadUUID</key>
            <string>86c12312-c278-41f1-bbe7-9422a1e40ca2</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Extensible SSO (Kerberos)</string>
    <key>PayloadIdentifier</key>
    <string>com.example.profile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>60bb7b2e-b94b-4f0d-848d-13c3a9857258</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [ExtensibleSingleSignOnKerberos.ExtensionData](/documentation/devicemanagement/extensiblesinglesignonkerberos/extensiondata-data.dictionary) - The additional data to pass to the app extension.

