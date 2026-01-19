# NetworkUsageRules

The payload that configures network-usage rules.

**Platforms:** iOS 9.0, iPadOS 9.0

## Discussion

Specify `com.apple.networkusagerules` as the payload type.

Network usage rules allow enterprises to specify how devices use networks, such as cellular data networks. iOS 9-12 require the application rules. In iOS 13, application rules, SIM rules, or both must be present.

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
            <key>ApplicationRules</key>
            <array>
                <dict>
                    <key>AllowCellularData</key>
                    <false/>
                    <key>AllowRoamingCellularData</key>
                    <false/>
                    <key>AppIdentifierMatches</key>
                    <array>
                        <string>com.apple.appname</string>
                    </array>
                </dict>
            </array>
            <key>PayloadIdentifier</key>
            <string>com.example.mynetworkusagepayload</string>
            <key>PayloadType</key>
            <string>com.apple.networkusagerules</string>
            <key>PayloadUUID</key>
            <string>'ef0250de-bfd4-4095-9ad3-34cf1281d5da</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Network Usage</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>b724f950-5853-4b78-8f4a-9a37a0ccac1f</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [NetworkUsageRules.ApplicationRulesItem](/documentation/devicemanagement/networkusagerules/applicationrulesitem) - The application rules dictionary.
- [NetworkUsageRules.SIMRulesItem](/documentation/devicemanagement/networkusagerules/simrulesitem) - The policy for individual SIM cards.

