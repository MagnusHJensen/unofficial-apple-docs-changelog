# Finder

The payload that configures Finder settings.

**Platforms:** macOS 10.7, Device Assignment Services , VPP License Management 

## Properties

### ProhibitBurn

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables the Finder’s burn support.

### ProhibitConnectTo

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables Connect to Server.

### ProhibitEject

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables Eject.

### ProhibitGoToFolder

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables Go to Folder.

### ShowExternalHardDrivesOnDesktop

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system doesn’t show external hard drives on the Desktop.

### ShowHardDrivesOnDesktop

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `false`, the system doesn’t show internal hard drives on the Desktop.

### ShowMountedServersOnDesktop

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `false`, the system doesn’t show mounted file servers on the Desktop.

### ShowRemovableMediaOnDesktop

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system doesn’t show removable media items on the Desktop.

### WarnOnEmptyTrash

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system doesn’t warn the user before emptying the trash.

## Discussion

Specify `com.apple.finder` as the payload type.

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
            <key>InterfaceLevel</key>
            <string>Full</string>
            <key>ShowHardDrivesOnDesktop</key>
            <true/>
            <key>ShowExternalHardDrivesOnDesktop</key>
            <false/>
            <key>ShowRemovableMediaOnDesktop</key>
            <false/>
            <key>ShowMountedServersOnDesktop</key>
            <true/>
            <key>WarnOnEmptyTrash</key>
            <true/>
            <key>ProhibitConnectTo</key>
            <true/>
            <key>ProhibitEject</key>
            <true/>
            <key>ProhibitBurn</key>
            <true/>
            <key>ProhibitGoToFolder</key>
            <true/>
            <key>PayloadIdentifier</key>
            <string>com.example.myfinderpayload</string>
            <key>PayloadType</key>
            <string>com.apple.finder</string>
            <key>PayloadUUID</key>
            <string>feea617a-c8f2-4dce-afae-20b2fe5f9c9b</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Finder</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>2527bd12-fbc4-4957-a9e7-4afeb64e0246</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

