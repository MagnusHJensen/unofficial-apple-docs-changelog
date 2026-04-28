# TimeMachine

The payload that configures Time Machine.

**Platforms:** macOS 10.7, Device Assignment Services , VPP License Management 

## Properties

### AutoBackup

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, performs automatic backups at regular intervals.

### BackupAllVolumes

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, backs up only the startup volume by default.

### BackupDestURL

- **Type:** `string`
- **Required:** Yes

The URL of the backup destination.

### BackupSizeMB

- **Type:** `integer`
- **Required:** No
- **Default:** `0`

The backup size limit, in megabytes. Set to 0 for unlimited.

### BackupSkipSys

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, skips system files and folders by default.

### BasePaths

- **Type:** `[string]`
- **Required:** No

The list of paths to back up besides the startup volume.

### MobileBackups

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, create local backup snapshots when not connected to the network.

### SkipPaths

- **Type:** `[string]`
- **Required:** No

The path to skip from start volume.

## Discussion

Specify `com.apple.MCX.TimeMachine` as the payload type.

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
            <key>AutoBackup</key>
            <true/>
            <key>BackupAllVolumes</key>
            <true/>
            <key>BackupDestURL</key>
            <string>server.example.com</string>
            <key>BackupSizeMB</key>
            <integer>1000</integer>
            <key>BackupSkipSys</key>
            <false/>
            <key>MobileBackups</key>
            <true/>
            <key>SkipPaths</key>
            <array>
                <string>/Users/Shared</string>
            </array>
            <key>PayloadIdentifier</key>
            <string>com.example.mytimemachinepayload</string>
            <key>PayloadType</key>
            <string>com.apple.MCX.TimeMachine</string>
            <key>PayloadUUID</key>
            <string>5f0be3a6-c9b8-44db-a2ae-414311772fdb</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Time Machine</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>68ca5f6c-13e8-43c3-b2ee-8bc405f5eed5</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

