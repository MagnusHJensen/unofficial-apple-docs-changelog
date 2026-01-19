# FileProvider

The payload that configures file provider settings.

**Platforms:** macOS 11.0

## Discussion

Specify `com.apple.fileproviderd` as the payload type.

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
            <key>AllowManagedFileProvidersToRequestAttribution</key>
            <true/>
            <key>PayloadIdentifier</key>
            <string>com.example.myfileproviderpayload</string>
            <key>PayloadType</key>
            <string>com.apple.fileproviderd</string>
            <key>PayloadUUID</key>
            <string>3ffb5741-f0f1-4fc2-9566-679ca8b10db1</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>FileProvider</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>8efec75c-f1d3-4aaa-a4ef-104dc25cfc3d</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

