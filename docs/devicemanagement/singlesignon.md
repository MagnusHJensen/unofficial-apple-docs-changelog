# SingleSignOn

The payload that configures single sign-on (SSO).

**Platforms:** iOS 7.0, iPadOS 7.0

## Discussion

Specify `com.apple.sso` as the payload type.

Deprecated in iOS 26. Use the [ExtensibleSingleSignOn](/documentation/devicemanagement/extensiblesinglesignon) payload instead.

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
            <string>com.apple.com</string>
            <key>TeamIdentifier</key>
            <string>RandomTeamID</string>
            <key>Hosts</key>
            <array>
                <string>.com.example.com</string>
            </array>
            <key>Realm</key>
            <string>com.example.com</string>
            <key>Type</key>
            <string>Credential</string>
            <key>PayloadIdentifier</key>
            <string>com.example.myssopayload</string>
            <key>PayloadType</key>
            <string>com.apple.sso</string>
            <key>PayloadUUID</key>
            <string>02cdfc1c-3c53-434d-99db-c55ee62548bd</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>SSO</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>00d92c73-9844-4dc6-b742-eda33efbbf23</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [SingleSignOn.Kerberos](/documentation/devicemanagement/singlesignon/kerberos-data.dictionary) - The Kerberos dictionary.

