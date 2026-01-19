# ProfileRemovalPassword

The payload that configures profile removal.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 9.0

## Discussion

Specify `com.apple.profileRemovalPassword` as the payload type.

This payload provides a password to allow users to remove a locked configuration profile from the device. If this payload is present and has a password value set, the device asks for the password when the user taps a profile’s Remove button. This system encrypts the payload with the rest of the profile.

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
            <key>RemovalPassword</key>
            <string>Password123</string>
            <key>PayloadIdentifier</key>
            <string>com.example.myprofileremovalpasswordprofile</string>
            <key>PayloadType</key>
            <string>com.apple.profileRemovalPassword</string>
            <key>PayloadUUID</key>
            <string>55f87465-a869-4ab0-9031-11ca1073641b</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Password Removal</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>a42bfe59-7f3a-46a0-9908-4a5095fd2c6d</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
    <key>HasRemovalPasscode</key>
    <true/>
</dict>
</plist>
```

