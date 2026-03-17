# DirectoryService

The payload that configures an Active Directory (AD) domain.

**Platforms:** macOS 10.8

## Properties

### ADAllowMultiDomainAuth

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system allows authentication from any domain in the namespace.

### ADAllowMultiDomainAuthFlag

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the `ADAllowMultiDomainAuth` key.

### ADCreateMobileAccountAtLogin

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system creates a mobile account at login.

### ADCreateMobileAccountAtLoginFlag

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the `ADCreateMobileAccountAtLogin` key.

### ADDefaultUserShell

- **Type:** `string`
- **Required:** No

The default user shell.

### ADDefaultUserShellFlag

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the `ADDefaultUserShell` key.

### ADDomainAdminGroupList

- **Type:** `[string]`
- **Required:** No

The list of Active Directory groups with admin access.

### ADDomainAdminGroupListFlag

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the `ADDomainAdminGroupList` key.

### ADForceHomeLocal

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system forces a local home directory.

### ADForceHomeLocalFlag

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the `ADForceHomeLocal` key.

### ADMapGGIDAttribute

- **Type:** `string`
- **Required:** No

The map group GID to attribute.

### ADMapGGIDAttributeFlag

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the `ADMapGGIDAttributeFlag` key.

### ADMapGIDAttribute

- **Type:** `string`
- **Required:** No

The map GID to attribute.

### ADMapGIDAttributeFlag

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the `ADMapGIDAttribute` key.

### ADMapUIDAttribute

- **Type:** `string`
- **Required:** No

The map UID to attribute.

### ADMapUIDAttributeFlag

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the `ADMapUIDAttribute` key.

### ADMountStyle

- **Type:** `string`
- **Required:** No

The network home protocol to use: `afp` or `smb`.

### ADNamespace

- **Type:** `string`
- **Required:** No

The primary user account naming convention; either `forest` or `domain`.

### ADNamespaceFlag

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the `ADNamespace` key.

### ADOrganizationalUnit

- **Type:** `string`
- **Required:** No

The organizational unit to add the joining computer object to.

### ADPacketEncrypt

- **Type:** `string`
- **Required:** No

The packet encryption policy.

### ADPacketEncryptFlag

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the `ADPacketEncrypt` key.

### ADPacketSign

- **Type:** `string`
- **Required:** No

The packet signing policy.

### ADPacketSignFlag

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the `ADPacketSign` key.

### ADPreferredDCServer

- **Type:** `string`
- **Required:** No

The preferred domain server.

### ADPreferredDCServerFlag

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the `ADPreferredDCServer` key.

### ADRestrictDDNS

- **Type:** `[string]`
- **Required:** No

An array of strings that represent the interfaces allowed for dynamic DNS updates, such as en0 and en1.

### ADRestrictDDNSFlag

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the `ADRestrictDDNS` key.

### ADTrustChangePassIntervalDays

- **Type:** `integer`
- **Required:** No

The number of days before requiring a change of the computer trust account password. Set to `0` to disable the feature.

### ADTrustChangePassIntervalDaysFlag

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the `ADTrustChangePassIntervalDays` key.

### ADUseWindowsUNCPath

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system uses the UNC path from Active Directory to derive the network home location.

### ADUseWindowsUNCPathFlag

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the `ADUseWindowsUNCPath` key.

### ADWarnUserBeforeCreatingMA

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the warning before creating the mobile account.

### ADWarnUserBeforeCreatingMAFlag

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables the `ADWarnUserBeforeCreatingMA` key.

### ClientID

- **Type:** `string`
- **Required:** No

The client’s identifier.

### Description

- **Type:** `string`
- **Required:** No

The directory service description.

### HostName

- **Type:** `string`
- **Required:** Yes

The Active Directory domain to join.

### Password

- **Type:** `string`
- **Required:** No

The password of the account for the domain.

### UserName

- **Type:** `string`
- **Required:** No

The user name of the account for the domain.

## Discussion

Specify `com.apple.DirectoryService.managed` as the payload type.

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
            <key>HostName</key>
            <string>host.example.com</string>
            <key>Password</key>
            <string>Password123</string>
            <key>UserName</key>
            <string>bind</string>
            <key>PayloadIdentifier</key>
            <string>com.example.mydspayload</string>
            <key>PayloadType</key>
            <string>com.apple.DirectoryService.managed</string>
            <key>PayloadUUID</key>
            <string>bb657e20-60b9-4c47-8730-51803ddf69e7</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Active Directory</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>079b6bc3-4356-4d80-89b4-a4b8a82eb739</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

