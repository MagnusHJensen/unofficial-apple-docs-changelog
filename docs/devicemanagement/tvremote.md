# TVRemote

The payload that configures the Apple TV remote.

**Platforms:** iOS 11.3, iPadOS 11.3, tvOS 11.3

## Discussion

Specify `com.apple.tvremote` as the payload type.

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
            <key>AllowedRemotes</key>
            <array>
                <dict>
                    <key>RemoteDeviceID</key>
                    <string>10:10:10:10:10:10</string>
                </dict>
            </array>
            <key>AllowedTVs</key>
            <array/>
            <key>PayloadIdentifier</key>
            <string>com.example.mytvremotepayload</string>
            <key>PayloadType</key>
            <string>com.apple.tvremote</string>
            <key>PayloadUUID</key>
            <string>696eb2c4-df9d-463f-b74a-685dae845fac</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>TV Remote</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>3de47695-ad92-44f5-9bc9-f1b2f2c7727b</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [TVRemote.AllowedRemotesItem](/documentation/devicemanagement/tvremote/allowedremotesitem) - The array of valid devices that Apple TV can connect to.
- [TVRemote.AllowedTVsItem](/documentation/devicemanagement/tvremote/allowedtvsitem) - The array of valid Apple TV identifiers that the remote can connect to.

