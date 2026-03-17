# SettingsCommand.Command.Settings.SharedDeviceConfiguration

A dictionary that contains shared device configuration settings.

**Platforms:** iOS 13.4, iPadOS 13.4

## Properties

### AwaitUserConfiguration

- **Type:** `SettingsCommand.Command.Settings.SharedDeviceConfiguration.AwaitUserConfiguration`
- **Required:** No

If enabled, the Shared iPad device enters Setup Assistant after the user triggers a login. The MDM server has a chance to configure the device and user. After configuration, the server needs to send a [User Configured](/documentation/devicemanagement/user-configured-command) command to the user channel to unblock the login. This feature requires the device to have network access during the login process.

Available in iOS 17 and later.

### Item

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `SharedDeviceConfiguration`

A string that identifies this setting.

### ManagedAppleIDDefaultDomains

- **Type:** `[string]`
- **Required:** No

A list of domains that the Shared iPad login screen displays. The user can pick a domain from the list to complete their Managed Apple Account.

If this list contains more than 3 domains, the system picks 3 at random for display. Available in iOS 16 and later.

### OnlineAuthenticationGracePeriod

- **Type:** `integer`
- **Required:** No

A grace period (in days) for Shared iPad online authentication. The Shared iPad only verifies the user’s passcode locally during login for users that already exist on the device. However, the system requires an online authentication (against Apple’s identity server) after the number of days specified by this setting.

Setting this value to 0 enforces online authentication every time.

Available in iOS 16 and later.

### PasscodePolicy

- **Type:** `SettingsCommand.Command.Settings.SharedDeviceConfiguration.PasscodePolicy`
- **Required:** No

A dictionary that contains passcode policies.

### QuotaSize

- **Type:** `integer`
- **Required:** No

The quota size, in megabytes (MB), for each user on the shared device, or if the quota size is too small, the minimum quota size. Available to Temporary Sessions Only guest users on iOS 17+.

### ResidentUsers

- **Type:** `integer`
- **Required:** No

The expected number of users. If this value is greater than the value for the maximum possible number of users that the device supports, the MDM server uses that value instead.

### SkipLanguageAndLocaleSetupForNewUsers

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system picks the system language and locale automatically for the new Shared iPad user.

Available in iOS 16.2 and later.

### TemporarySessionOnly

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the user only sees the Guest Welcome pane and can only log in as a guest user.

If `false`, the user can sign in with a Managed Apple Account (the existing behavior).

Available in iOS 14.5 and later.

### TemporarySessionTimeout

- **Type:** `integer`
- **Required:** No

The timeout, in seconds, for the temporary session. The temporary session logs out automatically after the specified period of inactivity. The minimum value is 30 seconds. Setting this value to `0` removes the timeout.

Available in iOS 14.5 and later.

### UserSessionTimeout

- **Type:** `integer`
- **Required:** No

The timeout, in seconds, for the user session. The user session logs out automatically after the specified period of inactivity. The minimum value is 30 seconds. Setting this value to `0` removes the timeout.

Available in iOS 14.5 and later.

## Topics

### Objects

- [SettingsCommand.Command.Settings.SharedDeviceConfiguration.AwaitUserConfiguration](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/shareddeviceconfiguration-data.dictionary/awaituserconfiguration-data.dictionary) - Enables the user configuration Setup Assistant workflow.
- [SettingsCommand.Command.Settings.SharedDeviceConfiguration.PasscodePolicy](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/shareddeviceconfiguration-data.dictionary/passcodepolicy-data.dictionary) - A dictionary that contains passcode policies.

