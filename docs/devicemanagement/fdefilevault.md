# FDEFileVault

The payload that configures FileVault.

**Platforms:** macOS 10.9, Device Assignment Services , VPP License Management 

## Properties

### Certificate

- **Type:** `data`
- **Required:** No

The DER-encoded certificate data if the system creates an institutional recovery key. This key isn’t supported on a Mac with Apple silicon.

### Defer

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system defers enabling FileVault until the designated user logs out. For details, see `fdesetup(8)`. Only a local user or a mobile account user can enable FileVault.

### DeferDontAskAtUserLogout

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system prevents requests to enable FileVault at user logout time.

### DeferForceAtUserLoginMaxBypassAttempts

- **Type:** `integer`
- **Required:** No

The maximum number of times users can bypass enabling FileVault before the system requires the user to enable it to log in. If the value is `0`, the system requires the user to enable FileVault the next time they attempt to log in. Set this key to `-1` to disable this feature.

### Enable

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `On`, `Off`

Set to `On` to enable FileVault and set to `Off` to disable FileVault. Payloads set to `On` sent through MDM need to either include full authentication information in the payload or have the `Defer` option set to `true`. When `Defer` is `true`, the system prompts for the authentication information when the user enables FileVault.

### ForceEnableInSetupAssistant

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, and installation of this payload occurs after enrolling with MDM in Setup Assistant, the system requests Setup Assistant to enable FileVault at setup time.

To use this, enable the Await Device Configured DEP configuration option and send this profile with this key set, before sending the [DeviceConfiguredCommand](/documentation/devicemanagement/deviceconfiguredcommand).

An admin SecureToken user is required, otherwise the FileVault pane does not appear.

### OutputPath

- **Type:** `string`
- **Required:** No

The path to the location of the recovery key and computer information property list.

### Password

- **Type:** `string`
- **Required:** No

The password of the Open Directory user to add to FileVault. Use the `UserEntersMissingInfo` key to prompt for this information.

### PayloadCertificateUUID

- **Type:** `string`
- **Required:** No

The UUID of the payload within the same profile containing the asymmetric recovery key certificate payload.

### ShowRecoveryKey

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, the system prevents display of the personal recovery key to the user after the system enables FileVault.

### UseKeychain

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true` and you don’t include certificate information in this payload, the system uses the keychain created at `/Library/Keychains/FileVaultMaster.keychain` when it adds the institutional recovery key.

### UseRecoveryKey

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the system creates a personal recovery key and displays it to the user.

### UserEntersMissingInfo

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables a prompt for missing user name or password fields.

### Username

- **Type:** `string`
- **Required:** No

The user name of the Open Directory user to add to FileVault.

## Discussion

Specify `com.apple.MCX.FileVault2` as the payload type.

FileVault 2 performs full XTS-AES 128 encryption on the contents of a volume. Removing the FileVault payload doesn’t disable FileVault.

As of macOS 10.15, FileVault settings require supervision or user approval when installed manually.

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
            <key>Defer</key>
            <true/>
            <key>Enable</key>
            <string>On</string>
            <key>ShowRecoveryKey</key>
            <true/>
            <key>UseKeychain</key>
            <false/>
            <key>UseRecoveryKey</key>
            <true/>
            <key>UserEntersMissingInfo</key>
            <false/>
            <key>PayloadIdentifier</key>
            <string>com.example.myfdefilevaultpayload</string>
            <key>PayloadType</key>
            <string>com.apple.MCX.FileVault2</string>
            <key>PayloadUUID</key>
            <string>c2c5f5e9-8ffc-4d8f-9108-fd655b90fdb2</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>FDE File Vault</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>5ba0de0e-ff06-4c0b-8dca-f5b55ed9de86</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

