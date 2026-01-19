# ManagedPreferences

The payload that configures managed preferences.

**Platforms:** macOS 10.7

## Discussion

Specify `com.apple.ManagedClient.preferences` as the payload type.

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
            <key>PayloadContent</key>
            <dict>
                <key>com.example.myapp</key>
                <dict>
                    <key>Forced</key>
                    <array>
                        <dict>
                            <key>mcx_preference_settings</key>
                            <dict>
                                <key>MySetting</key>
                                <false/>
                            </dict>
                        </dict>
                    </array>
                </dict>
            </dict>
            <key>PayloadIdentifier</key>
            <string>com.example.mymanprefpayload</string>
            <key>PayloadType</key>
            <string>com.apple.ManagedClient.preferences</string>
            <key>PayloadUUID</key>
            <string>83c9f6e8-ef4b-4974-b83b-b2e7fe79b75c</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Managed Preference</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>54e23577-3424-4092-a9b4-a5e5af88fd52</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [ManagedPreferences.PayloadContent](/documentation/devicemanagement/managedpreferences/payloadcontent-data.dictionary) - The dictionary containing app preference domains. The key names are application preference domain identifiers (for example, `com.example.my-app`), or the string `.GlobalPreferences` for the global domain. The values are the corresponding forced and set-once preferences.

