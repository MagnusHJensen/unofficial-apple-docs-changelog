# SettingsCommand.Command.Settings.MDMOptions.MDMOptions

A dictionary that contains MDM options.

**Platforms:** iOS 7.0, iPadOS 7.0, Mac Catalyst 7.0, macOS 10.15, visionOS 2.0, Device Assignment Services , VPP License Management 

## Properties

### ActivationLockAllowedWhileSupervised

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, a supervised device registers itself with Activation Lock when the user enables Find My. This setting is available for supervised devices in iOS 7 and later, and macOS 10.15 and later.

### BootstrapTokenAllowed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the server supports the Bootstrap Token commands.

### IdleRebootAllowed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device automatically reboots while locked after several days of inactivity. This is set to `false` by default when a supervised device enrolls.

### PromptUserToAllowBootstrapTokenForAuthentication

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, warn the user that they need to reboot into RecoveryOS and allow the MDM server to use the Bootstrap Token for authentication for certain sensitive operations; for example, enabling kernel extensions or installing certain types of software updates. Set this value to `false` if your MDM server doesn’t need to perform these operations. The value provided here overrides the value specified in MDM, and only applies when `BootstrapTokenAllowedForAuthentication` is `true` in the [SecurityInfoResponse.SecurityInfo](/documentation/devicemanagement/securityinforesponse/securityinfo-data.dictionary) response. This value is available for a Mac with Apple silicon in macOS 11 and later.

