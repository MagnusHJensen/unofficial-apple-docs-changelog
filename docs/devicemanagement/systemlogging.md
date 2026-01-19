# SystemLogging

The payload that configures system logging.

**Platforms:** iOS 9.0, iPadOS 9.0, macOS 10.12, tvOS 9.0, visionOS 1.0, watchOS 3.0

## Discussion

Specify `com.apple.system.logging` as the payload type.

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
            <key>Processes</key>
            <dict/>
            <key>Subsystems</key>
            <dict>
                <key>com.example.app</key>
                <dict>
                    <key>DEFAULT-OPTIONS</key>
                    <dict>
                        <key>Level</key>
                        <dict>
                            <key>Enable</key>
                            <string>Info</string>
                        </dict>
                        <key>Default-Privacy-Setting</key>
                        <string>Public</string>
                    </dict>
                </dict>
            </dict>
            <key>PayloadIdentifier</key>
            <string>com.example.mysystemloggingpayload</string>
            <key>PayloadType</key>
            <string>com.apple.system.logging</string>
            <key>PayloadUUID</key>
            <string>0ff59613-35e9-495c-88c8-01963de4ac80</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>System Logging</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>0fbfa83e-c8ec-49d0-b50c-3acc3c05749c</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [SystemLogging.Processes](/documentation/devicemanagement/systemlogging/processes-data.dictionary) - Not to be used.
- [SystemLogging.Subsystems](/documentation/devicemanagement/systemlogging/subsystems-data.dictionary) - A dictionary enabling the logging level for subsystems. See `Customizing Logging Behavior While Debugging` for more details about the format of the dictionary.
- [SystemLogging.System](/documentation/devicemanagement/systemlogging/system-data.dictionary) - This dictionary has one key, `Enable-Private-Data`. Setting that value to `true` enables private data logging for the entire system.

