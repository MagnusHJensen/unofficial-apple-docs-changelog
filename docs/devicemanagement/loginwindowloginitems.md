# LoginWindowLoginItems

The payload that configures login behavior.

**Platforms:** macOS 10.7

## Discussion

Specify `loginwindow` as the payload type.

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
            <key>DisableLoginItemsSuppression</key>
            <true/>
            <key>PayloadIdentifier</key>
            <string>com.example.myloginwindowloginitemspayload</string>
            <key>PayloadType</key>
            <string>loginwindow</string>
            <key>PayloadUUID</key>
            <string>e032844e-db81-4387-98d9-9ee7c6038275</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Login Window Login Items</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>30724e47-e58e-447d-b21c-a65bbe184f98</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

