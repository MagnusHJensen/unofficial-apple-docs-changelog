# Firewall

The payload that configures the firewall.

**Platforms:** macOS 10.12, Device Assignment Services , VPP License Management 

## Properties

### AllowSigned

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the system allows built-in software to receive incoming connections. Available in macOS 12.3 and later.

> 

### AllowSignedApp

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the system allows downloaded signed software to receive incoming connections. Available in macOS 12.3 and later.

> 

### Applications

- **Type:** `[Firewall.ApplicationsItem]`
- **Required:** No

The list of apps with connections that the firewall controls.

### BlockAllIncoming

- **Type:** `boolean`
- **Required:** No

If `true`, the system enables blocking all incoming connections.

### EnableFirewall

- **Type:** `boolean`
- **Required:** Yes

If `true`, the system enables the firewall.

### EnableStealthMode

- **Type:** `boolean`
- **Required:** No

If `true`, the system enables stealth mode.

## Discussion

Specify `com.apple.security.firewall` as the payload type.

The payload needs to exist in a system-scoped profile.

If more than one profile contains this payload, the system uses the most restrictive union of settings.

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
            <key>EnableFirewall</key>
            <true/>
            <key>Applications</key>
            <array>
                <dict>
                    <key>BundleID</key>
                    <string>com.example.myapp</string>
                    <key>Allowed</key>
                    <false/>
                </dict>
            </array>
            <key>PayloadIdentifier</key>
            <string>com.example.myfirewallpayload</string>
            <key>PayloadType</key>
            <string>com.apple.security.firewall</string>
            <key>PayloadUUID</key>
            <string>28b1fef7-ddb6-4d56-9a6a-6bb4e56e7f0b</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Firewall</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>8f2fa915-f2da-4034-9424-2218355a6f3c</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

## Topics

### Objects

- [Firewall.ApplicationsItem](/documentation/devicemanagement/firewall/applicationsitem) - A dictionary of details for apps.

