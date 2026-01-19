# AutonomousSingleAppMode

The payload that configures Autonomous Single App mode.

**Platforms:** macOS 10.13.4

## Discussion

Specify `com.apple.asam` as the payload type.

The system only allows installation of one profile of this type, and it requires installation through a user-approved MDM server. Apps listed in this profile have low-level access to the system, including, but not limited to, key logging and user interface manipulation outside the app’s context.

> 

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
            <key>AllowedApplications</key>
            <array>
                <dict>
                    <key>BundleIdentifier</key>
                    <string>com.apple.safari</string>
                    <key>TeamIdentifier</key>
                    <string>team-id</string>
                </dict>
            </array>
            <key>PayloadIdentifier</key>
            <string>com.example.myasampayload</string>
            <key>PayloadType</key>
            <string>com.apple.asam</string>
            <key>PayloadUUID</key>
            <string>c324fd3e-d98a-4ea8-818a-5991024cddd0</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Autonomous Single App Mode</string>
    <key>PayloadIdentifier</key>
    <string>com.example.profile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>467ab3a0-9423-4c16-a05c-5c99d771088f</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [AutonomousSingleAppMode.AllowedApplicationsItem](/documentation/devicemanagement/autonomoussingleappmode/allowedapplicationsitem) - A dictionary that specifies an app that can be granted access to the Accessibilty APIs.

