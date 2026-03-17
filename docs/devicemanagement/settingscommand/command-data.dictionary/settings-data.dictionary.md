# SettingsCommand.Command.Settings

An array of dictionaries that contains the settings.

**Platforms:** iOS 5.0, iPadOS 5.0, macOS 10.9, tvOS 9.0, visionOS 1.1, watchOS 10.0

## Properties

### AccessibilitySettings

- **Type:** `SettingsCommand.Command.Settings.AccessibilitySettings`
- **Required:** No

A dictionary that contains accessibility settings. Available in iOS 16 and later.

### AppAnalytics

- **Type:** `SettingsCommand.Command.Settings.AppAnalytics`
- **Required:** No

A dictionary that contains settings for sharing app analytics. This setting is available only for Shared iPad in iOS 9.3.2 and later.

### ApplicationAttributes

- **Type:** `SettingsCommand.Command.Settings.ApplicationAttributes`
- **Required:** No

A dictionary that contains the attributes to apply to the app. Omit this setting to remove existing attributes. This setting supports user enrollment, is available in iOS 7 and later, and tvOS 10.2 and later. This setting fails for apps that Declarative Device Management manages.

### ApplicationConfiguration

- **Type:** `SettingsCommand.Command.Settings.ApplicationConfiguration`
- **Required:** No

A dictionary that contains the configurations to apply to the app. Omit this setting to remove existing configurations. This setting requires the App Management access right, supports user enrollment, and is available in iOS 7 and later, macOS 10.15 and later, and tvOS 10.2 and later. This setting fails for apps that Declarative Device Management manages.

### Bluetooth

- **Type:** `SettingsCommand.Command.Settings.Bluetooth`
- **Required:** No

A dictionary that contains Bluetooth settings. This setting requires the Network Information access right, doesn’t support user enrollment, and is available only on supervised devices. Available in iOS 11.3 and later, and macOS 10.13.4 and later.

### DataRoaming

- **Type:** `SettingsCommand.Command.Settings.DataRoaming`
- **Required:** No

A dictionary that contains data roaming settings. This setting requires the Network Information access right, and doesn’t support user enrollment. Available in iOS 5 and later.

### DefaultApplications

- **Type:** `SettingsCommand.Command.Settings.DefaultApplications`
- **Required:** No

A dictionary that contains default application bundle identifiers for each default application type that can be set.

### DeviceName

- **Type:** `SettingsCommand.Command.Settings.DeviceName`
- **Required:** No

A dictionary that contains device name settings. This setting doesn’t support user enrollment, and is available only on supervised devices. Available in iOS 5 and later, macOS 10.10 and later, and visionOS 2 and later.

### DiagnosticSubmission

- **Type:** `SettingsCommand.Command.Settings.DiagnosticSubmission`
- **Required:** No

A dictionary that contains diagnostic submission settings. This setting is available only for Shared iPad in iOS 9.3 and later.

### HostName

- **Type:** `SettingsCommand.Command.Settings.HostName`
- **Required:** No

A dictionary that contains hostname settings. This setting doesn’t support user enrollment, and is available in macOS 10.11 and later.

### MaximumResidentUsers

- **Type:** `SettingsCommand.Command.Settings.MaximumResidentUsers`
- **Required:** No

A dictionary that contains settings for maximum resident users. Apple deprecated this setting in iOS 13.4. Use ’SharedDeviceConfiguration` instead. This setting is available only for Shared iPad.

### MDMOptions

- **Type:** `SettingsCommand.Command.Settings.MDMOptions`
- **Required:** No

A dictionary that contains settings related to the MDM protocol. This setting doesn’t support user enrollment. Available in iOS 7 and later, macOS 10.15 and later, and visionOS 2 and later.

### OrganizationInfo

- **Type:** `SettingsCommand.Command.Settings.OrganizationInfo`
- **Required:** No

A dictionary that contains settings about the organization operating the MDM server. This setting supports user enrollment. Available in iOS 5 and later.

### PasscodeLockGracePeriod

- **Type:** `SettingsCommand.Command.Settings.PasscodeLockGracePeriod`
- **Required:** No

A dictionary that contains password lock grace period settings. This setting is available only for Shared iPad in iOS 9.3.2 and later. This key is deprecated. Use ‘PasscodeLockGracePeriod’ in SettingsCommand.Command.Settings.SharedDeviceConfiguration.PasscodePolicy instead.

### PersonalHotspot

- **Type:** `SettingsCommand.Command.Settings.PersonalHotspot`
- **Required:** No

A dictionary that contains Personal Hotspot settings. This setting requires the Network Information access right, and doesn’t support user enrollment. Available in iOS 5 and later.

### SharedDeviceConfiguration

- **Type:** `SettingsCommand.Command.Settings.SharedDeviceConfiguration`
- **Required:** No

A dictionary that contains shared device configuration settings. This setting is available only for Shared iPad in iOS 13.4 and later.

### SoftwareUpdateSettings

- **Type:** `SettingsCommand.Command.Settings.SoftwareUpdateSettings`
- **Required:** No

A dictionary that contains software update settings. This setting doesn’t support user enrollment. Available in iOS 14.5 and later.

### TimeZone

- **Type:** `SettingsCommand.Command.Settings.TimeZone`
- **Required:** No

A dictionary that contains time zone settings. This setting is available only on supervised devices and doesn’t support user enrollment. Available in iOS 14 and later, tvOS 14 and later, and visionOS 2 and later.

### VoiceRoaming

- **Type:** `SettingsCommand.Command.Settings.VoiceRoaming`
- **Required:** No

A dictionary that contains voice roaming settings. This setting requires the Network Information access right, and doesn’t support user enrollment. Available in iOS 5 and later.

### Wallpaper

- **Type:** `SettingsCommand.Command.Settings.Wallpaper`
- **Required:** No

A dictionary that contains wallpaper settings. This setting doesn’t support user enrollment. Available in iOS 8 and later. Starting in iOS 16 and iPadOS 17, when setting the wallpaper for the first time, both locations update. After that, you can set either location separately.

## Topics

### Objects

- [SettingsCommand.Command.Settings.AccessibilitySettings](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/accessibilitysettings-data.dictionary) - A dictionary that contains settings for accessibility.
- [SettingsCommand.Command.Settings.AppAnalytics](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/appanalytics-data.dictionary) - A dictionary that contains settings for sharing app analytics.
- [SettingsCommand.Command.Settings.ApplicationAttributes](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/applicationattributes-data.dictionary) - A dictionary that contains the attributes to apply to the app.
- [SettingsCommand.Command.Settings.ApplicationConfiguration](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/applicationconfiguration-data.dictionary) - A dictionary that contains the configurations to apply to the app.
- [SettingsCommand.Command.Settings.Bluetooth](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/bluetooth-data.dictionary) - A dictionary that contains Bluetooth settings.
- [SettingsCommand.Command.Settings.DataRoaming](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/dataroaming-data.dictionary) - A dictionary that contains data roaming settings.
- [SettingsCommand.Command.Settings.DefaultApplications](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/defaultapplications-data.dictionary) - A dictionary that contains default application bundle identifiers for each default application type that can be set.
- [SettingsCommand.Command.Settings.DeviceName](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/devicename-data.dictionary) - A dictionary that contains device name settings.
- [SettingsCommand.Command.Settings.DiagnosticSubmission](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/diagnosticsubmission-data.dictionary) - A dictionary that contains diagnostic submission settings.
- [SettingsCommand.Command.Settings.HostName](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/hostname-data.dictionary) - A dictionary that contains hostname settings.
- [SettingsCommand.Command.Settings.MDMOptions](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/mdmoptions-data.dictionary) - A dictionary that contains settings about the organization operating the MDM server.
- [SettingsCommand.Command.Settings.MaximumResidentUsers](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/maximumresidentusers-data.dictionary) - A dictionary that contains settings for maximum resident users.
- [SettingsCommand.Command.Settings.OrganizationInfo](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/organizationinfo-data.dictionary) - A dictionary that contains settings about the organization operating the MDM server.
- [SettingsCommand.Command.Settings.PasscodeLockGracePeriod](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/passcodelockgraceperiod-data.dictionary) - A dictionary that contains settings for the password lock grace period.
- [SettingsCommand.Command.Settings.PersonalHotspot](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/personalhotspot-data.dictionary) - A dictionary that contains Personal Hotspot settings.
- [SettingsCommand.Command.Settings.SharedDeviceConfiguration](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/shareddeviceconfiguration-data.dictionary) - A dictionary that contains shared device configuration settings.
- [SettingsCommand.Command.Settings.SoftwareUpdateSettings](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/softwareupdatesettings-data.dictionary) - A dictionary that contains software update settings.
- [SettingsCommand.Command.Settings.TimeZone](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/timezone-data.dictionary) - A dictionary that contains time zone settings.
- [SettingsCommand.Command.Settings.VoiceRoaming](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/voiceroaming-data.dictionary) - A dictionary that contains voice roaming settings.
- [SettingsCommand.Command.Settings.Wallpaper](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/wallpaper-data.dictionary) - A dictionary that contains wallpaper settings.

