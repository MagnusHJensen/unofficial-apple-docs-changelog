# LoginWindow

The payload that configures Login Window behavior.

**Platforms:** macOS 10.7, Device Assignment Services , VPP License Management 

## Properties

### AdminHostInfo

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `HostName`, `SystemVersion`, `IPAddress`

The admin host info. If present in the payload, the system displays its value in the Login Window as additional computer information. Before macOS 10.10, this string could only contain host name, system version, or IP address. After macOS 10.10, setting this key to any value allows the user to click the time area of the menu bar to toggle through various computer information values.

### AllowList

- **Type:** `[string]`
- **Required:** No

The list of user GUIDs or group GUIDs of users that the system allows to log in. An asterisk (`*`) string specifies all users or groups. This only applies to network accounts and mobile accounts.

### AutologinPassword

- **Type:** `string`
- **Required:** No

An optional user password to set up auto login. This must match the `AutologinUsername` user’s current password.

### AutologinUsername

- **Type:** `string`
- **Required:** No

The user short name for an existing user to set up auto login.

### DenyList

- **Type:** `[string]`
- **Required:** No

The list of user GUIDs or group GUIDs of users that the system disallows to log in. This list takes priority over the list in the `AllowList` key. This only applies to network accounts and mobile accounts.

### DisableConsoleAccess

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disregards the `>console` special user name, which provides a command line UI.

### DisableFDEAutoLogin

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables the automatic login option when using FileVault.

### DisableScreenLockImmediate

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables the immediate Screen Lock functions. Available in macOS 10.13 and later.

### HideAdminUsers

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system hides administrator users when showing a user list.

### HideLocalUsers

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system shows only network and system users when showing a user list.

### HideMobileAccounts

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system hides mobile account users in a user list. In some cases, mobile users show up as network users.

### IncludeNetworkUser

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system shows network users when showing a user list.

### LoginwindowText

- **Type:** `string`
- **Required:** No

The text to display in the Login Window.

### LogOutDisabledWhileLoggedIn

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables the Log Out menu item when the user is logged in. Available in macOS 10.13 and later.

### PowerOffDisabledWhileLoggedIn

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables the Power Off menu item when the user is logged in.

### RestartDisabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables the Restart item.

### RestartDisabledWhileLoggedIn

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables the Restart menu item when the user is logged in.

### SHOWFULLNAME

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system shows the name and password dialog. If `false`, the system displays a list of users.

### showInputMenu

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system shows the Input Menu in the Login Window.

### SHOWOTHERUSERS_MANAGED

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system displays “Other…” when it shows a list of users.

### ShutDownDisabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables the Shut Down button.

### ShutDownDisabledWhileLoggedIn

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables the Shut Down menu item when the user is logged in.

### SleepDisabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables the Sleep button.

## Discussion

Specify `com.apple.loginwindow` as the payload type.

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
            <key>AdminHostInfo</key>
            <string>HostName</string>
            <key>AdminMayDisableMCX</key>
            <true/>
            <key>AllowList</key>
            <array/>
            <key>AlwaysShowWorkgroupDialog</key>
            <true/>
            <key>CombineUserWorkgroups</key>
            <true/>
            <key>DenyList</key>
            <array/>
            <key>DisableConsoleAccess</key>
            <true/>
            <key>EnableExternalAccounts</key>
            <false/>
            <key>FlattenUserWorkgroups</key>
            <true/>
            <key>HideAdminUsers</key>
            <false/>
            <key>HideLocalUsers</key>
            <false/>
            <key>HideMobileAccounts</key>
            <false/>
            <key>IncludeNetworkUser</key>
            <true/>
            <key>LocalUserLoginEnabled</key>
            <true/>
            <key>LocalUsersHaveWorkgroups</key>
            <true/>
            <key>RestartDisabled</key>
            <false/>
            <key>RetriesUntilHint</key>
            <integer>0</integer>
            <key>SHOWFULLNAME</key>
            <false/>
            <key>SHOWOTHERUSERS_MANAGED</key>
            <true/>
            <key>ShutDownDisabled</key>
            <false/>
            <key>SleepDisabled</key>
            <false/>
            <key>UseComputerNameForComputerRecordName</key>
            <true/>
            <key>com.apple.login.mcx.DisableAutoLoginClient</key>
            <true/>
            <key>showInputMenu</key>
            <true/>
            <key>PayloadIdentifier</key>
            <string>com.example.myloginwindowpayload</string>
            <key>PayloadType</key>
            <string>com.apple.loginwindow</string>
            <key>PayloadUUID</key>
            <string>fe9ba3c5-0f1a-45c7-b6df-a5f4489695fe</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Login Window</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>61bd7d63-4a4a-4b67-9112-5ceb16afb4dc</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
```

