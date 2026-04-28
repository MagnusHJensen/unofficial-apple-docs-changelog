# Xsan

The payload that configures an Xsan client system.

**Platforms:** macOS 10.10, Device Assignment Services , VPP License Management 

## Properties

### fsnameservers

- **Type:** `[string]`
- **Required:** No

An array of storage area network (SAN) File System Name Server coordinators. The list should contain the same addresses in the same order as the metadata controller (MDC) `/Library/Preferences/Xsan/fsnameservers` file.

This key is required for StorNext SANs.

### sanAuthMethod

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `auth_secret`

The authentication method for the SAN. This key is required for all Xsan SANs. It’s optional for StorNext SANs but should be set if the StorNext SAN uses an `auth_secret` file.

Only one value is accepted: `auth_secret`

### sanConfigURLs

- **Type:** `[string]`
- **Required:** No

An array of LDAP URLs where Xsan systems can obtain SAN configuration updates. There should be one entry for each Xsan MDC.

This key is required for all Xsan SANs.

Example URL: `ldaps://mdc1.example.com:389`.

### sanName

- **Type:** `string`
- **Required:** Yes

The name of the SAN. This key is required for all Xsan SANs. The name must exactly match the name of the SAN defined in the metadata server.

### sharedSecret

- **Type:** `string`
- **Required:** Yes

The shared secret used for Xsan network authentication. This key is required when the `sanAuthMethod` key is present. The value should equal the content of the MDC’s `/Library/Preferences/Xsan/.auth_secret` file.

## Discussion

Specify `com.apple.xsan` as the payload type.

For more information, see [https://support.apple.com/en-us/HT205333](https://support.apple.com/en-us/HT205333).

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
            <key>fsnameservers</key>
            <array>
                <string>san.example.com</string>
            </array>
            <key>sanAuthMethod</key>
            <string>auth_secret</string>
            <key>sanName</key>
            <string>user_xsan</string>
            <key>sharedSecret</key>
            <string>shared_secret</string>
            <key>PayloadIdentifier</key>
            <string>com.example.myxsanpayload</string>
            <key>PayloadType</key>
            <string>com.apple.xsan</string>
            <key>PayloadUUID</key>
            <string>015b6f29-6a6a-4e23-9305-a4182838516d</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Xsan</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>e93b1eca-d07c-4383-8b4f-ca9db76f4aa8</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

