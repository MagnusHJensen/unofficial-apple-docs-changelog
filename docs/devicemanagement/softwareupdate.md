# SoftwareUpdate

The payload that configures the software update policy. Removed: use the declarative management `com.apple.configuration.softwareupdate.settings` configuration.

**Platforms:** macOS 10.7

## Properties

### AllowPreReleaseInstallation

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`





Removed: macOS 27+

### AutomaticallyInstallAppUpdates

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`





Removed: macOS 27+

### AutomaticallyInstallMacOSUpdates

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`





Removed: macOS 27+

### AutomaticCheckEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`





Removed: macOS 27+

### AutomaticDownload

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`





Removed: macOS 27+

### CatalogURL

- **Type:** `string`
- **Required:** No




Removed: macOS 27+

### ConfigDataInstall

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`





Removed: macOS 27+

### CriticalUpdateInstall

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`





Removed: macOS 27+

### restrict-software-update-require-admin-to-install

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`





Removed: macOS 27+

## Discussion

Specify `com.apple.SoftwareUpdate` as the payload type.

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
            <key>AutomaticallyInstallAppUpdates</key>
            <false/>
            <key>PayloadIdentifier</key>
            <string>com.example.mysoftwareupdatepayload</string>
            <key>PayloadType</key>
            <string>com.apple.SoftwareUpdate</string>
            <key>PayloadUUID</key>
            <string>af3c6efa-0dd3-4021-814b-6f2dba91428b</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Software Update</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>8b6061ab-31c6-4eee-ba5b-8a09ea8f5fa7</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

