# AppLock

The payload that configures a device to run a single app.

**Platforms:** iOS 6.0, iPadOS 6.0, tvOS 10.2

## Discussion

Specify `com.apple.app.lock` as the payload type.

With an app lock profile, the device locks to the specified app until removal of the profile. The device returns to the app automatically upon wake or restart.

Only use an app lock payload after installing the target app.

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
            <key>App</key>
            <dict>
                <key>Identifier</key>
                <string>com.apple.mobilenotes</string>
            </dict>
            <key>PayloadIdentifier</key>
            <string>com.example.myapplockpayload</string>
            <key>PayloadType</key>
            <string>com.apple.app.lock</string>
            <key>PayloadUUID</key>
            <string>dc0c6fdd-aae0-4fd0-a19c-861ba28f4c55</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Single App Mode</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>736f867e-3972-4889-aa68-7ce5be12eff6</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [AppLock.App](/documentation/devicemanagement/applock/app-data.dictionary) - The only app available for use on the iOS device.

