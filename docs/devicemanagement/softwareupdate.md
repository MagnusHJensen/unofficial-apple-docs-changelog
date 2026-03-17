# SoftwareUpdate

The payload that configures the software update policy.

**Platforms:** macOS 10.7

## Properties

### AllowPreReleaseInstallation

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, prerelease software can be installed on this computer.

### AutomaticallyInstallAppUpdates

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, deselects the “Install app updates from the App Store” option and prevents the user from changing the option.

### AutomaticallyInstallMacOSUpdates

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, restricts the “Install macOS Updates” option and prevents the user from changing the option.

### AutomaticCheckEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, deselects the “Check for updates” option and prevents the user from changing the option.

### AutomaticDownload

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, deselects the “Download new updates when available from the App Store” option and prevents the user from changing the option.

### CatalogURL

- **Type:** `string`
- **Required:** No

The URL of the software update catalog. This property is not supported in macOS 11 and later.

### ConfigDataInstall

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, restricts the automatic installation of configuration data.

### CriticalUpdateInstall

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables the automatic installation of critical updates and prevents the user from changing the “Install system data files and security updates” option.

### restrict-software-update-require-admin-to-install

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, restrict app installations to admin users. This key has the same function as the  `restrict-store-require-admin-to-install` key in the `com.apple.appstore` payload.

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

