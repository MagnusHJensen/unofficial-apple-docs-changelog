# FileProvider

The payload that configures file provider settings.

**Platforms:** macOS 11.0

## Properties

### AllowManagedFileProvidersToRequestAttribution

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables file providers access to the path of the requesting process.

### ManagementAllowsExternalVolumeSyncing

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the device only allows File Provider extension volume synchronization for the system “home” volume and any data separated volume, and prevents synchronization with any other volumes. If `true``, the device allows File Provider extension volume synchronization for the system “home” volume, any data separated volume, and any encrypted APFS volumes (on either internal or external media).

### ManagementAllowsKnownFolderSyncing

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the device prevents the File Provider extension from using desktop and documents synchronization in any app. This does not impact the ability for apps to utilize the File Provider extension for file and folder syncing with remote storage.

### ManagementAllowsRemoteSyncing

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the device prevents the File Provider extension from using synchronization in any app. Also, none of the other options will be evaluated. Synchronization will be totally disabled for any application.

### ManagementDomainAutoEnablementList

- **Type:** `[string]`
- **Required:** No

An array of strings representing the composed identifiers of apps. The device automatically enables the File Provider domains for the corresponding apps. The device doesn’t enable existing domains if enrollment happens after they are created. The device doesn’t prevent the user from disabling these File Provider domains. Users need to manually enable File Provider domains in the Finder if their corresponding apps aren’t listed here. The format of the app identifiers is “Bundle-ID (Team-ID)”, for example `com.example.app (ABCD1234)`.

### ManagementExternalVolumeSyncingAllowList

- **Type:** `[string]`
- **Required:** No

An array of strings representing the composed identifiers of apps. The device allows the corresponding apps to use File Provider extension volume synchronization. If present, and `ManagementAllowsExternalVolumeSyncing` is set to `true`, the device allows only the apps in this list to use volume synchronization. This key is ignored if `ManagementAllowsExternalVolumeSyncing` is set to `false`. The format of the app identifiers is “Bundle-ID (Team-ID)”, for example `com.example.app (ABCD1234)`.

### ManagementKnownFolderSyncingAllowList

- **Type:** `[string]`
- **Required:** No

An array of strings representing the composed identifiers of apps. The device allows the corresponding apps to use File Provider extension desktop and documents synchronization. If present, and `ManagementAllowsKnownFolderSyncing` is set to `true`, the device allows only the apps in this list to use desktop and documents synchronization. This key is ignored if `ManagementAllowsKnownFolderSyncing` is set to `false`. This setting does not impact the ability for apps to use File Provider extension volume access. The format of the app identifiers is “Bundle-ID (Team-ID)”, for example `com.example.app (ABCD1234)`.

### ManagementRemoteSyncingAllowList

- **Type:** `[string]`
- **Required:** No

An array of strings representing the composed identifiers of apps. The device allows the corresponding apps to use File Provider extension synchronization. If present, and `ManagementAllowsRemoteSyncing` is set to `true`, the device allows only the apps in this list to use synchronization. This key is ignored if `ManagementAllowsRemoteSyncing` is set to `false`. If present, the other options will only be evaluated for the apps in this list. The format of the app identifiers is “Bundle-ID (Team-ID)”, for example `com.example.app (ABCD1234)`.

## Discussion

Specify `com.apple.fileproviderd` as the payload type.

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
            <key>ManagementAllowsRemoteSyncing</key>
            <true/>
            <key>AllowManagedFileProvidersToRequestAttribution</key>
            <true/>
            <key>ManagementAllowsKnownFolderSyncing</key>
            <true/>
            <key>ManagementKnownFolderSyncingAllowList</key>
            <array>
               <string>com.example.myprovider (ABCD1234)</string>
            </array>
            <key>PayloadIdentifier</key>
            <string>com.example.myfileproviderpayload</string>
            <key>PayloadType</key>
            <string>com.apple.fileproviderd</string>
            <key>PayloadUUID</key>
            <string>3ffb5741-f0f1-4fc2-9566-679ca8b10db1</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>FileProvider</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>8efec75c-f1d3-4aaa-a4ef-104dc25cfc3d</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

