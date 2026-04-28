# SystemPolicyKernelExtensions

The payload that configures the kernel extension policies.

**Platforms:** macOS 10.13.2, Device Assignment Services , VPP License Management 

## Properties

### AllowedKernelExtensions

- **Type:** `SystemPolicyKernelExtensions.AllowedKernelExtensions`
- **Required:** No

The dictionary that represents a set of kernel extensions that the system always allows to load on the computer. The dictionary maps team identifiers (keys) to arrays of bundle identifiers.

### AllowedTeamIdentifiers

- **Type:** `[string]`
- **Required:** No

The array of team identifiers that define which validly signed kernel extensions can load.

### AllowNonAdminUserApprovals

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, nonadministrative users can approve additional kernel extensions in the Security & Privacy preferences.

Available in macOS 11 and later.

### AllowUserOverrides

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, users can approve additional kernel extensions that configuration profiles don’t explicitly allow.

## Discussion

Specify `com.apple.syspolicy.kernel-extension-policy` as the payload type.

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
            <key>AllowUserOverrides</key>
            <false/>
            <key>AllowedTeamIdentifiers</key>
            <array>
                <string>ABCDE12345</string>
            </array>
            <key>AllowedKernelExtensions</key>
            <dict>
                <key></key>
                <array>
                    <string>com.example.mydriver</string>
                </array>
                <key>ABCDE12345</key>
                <array>
                    <string>com.example.kext.mydriver</string>
                </array>
            </dict>
            <key>PayloadIdentifier</key>
            <string>com.example.mysystempolicykernalextensionspayload</string>
            <key>PayloadType</key>
            <string>com.apple.syspolicy.kernel-extension-policy</string>
            <key>PayloadUUID</key>
            <string>3202f59b-3583-4e6c-82ae-776f3c815df1</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>System Policy Kernal Extension</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>d9fa7f5b-713d-48f8-a8bd-219cf3061873</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [SystemPolicyKernelExtensions.AllowedKernelExtensions](/documentation/devicemanagement/systempolicykernelextensions/allowedkernelextensions-data.dictionary) - The dictionary that represents a set of kernel extensions.

