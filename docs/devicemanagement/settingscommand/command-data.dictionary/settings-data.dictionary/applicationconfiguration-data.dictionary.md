# SettingsCommand.Command.Settings.ApplicationConfiguration

A dictionary that contains the configurations to apply to the app.

**Platforms:** iOS 7.0, iPadOS 7.0, Mac Catalyst 7.0, macOS 10.15, tvOS 10.2, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### Configuration

- **Type:** `SettingsCommand.Command.Settings.ApplicationConfiguration.Configuration`
- **Required:** No

A dictionary that contains the configurations to apply to the app. Omit this setting to remove existing configurations.

### Identifier

- **Type:** `string`
- **Required:** Yes

The bundle identifier of the managed app.

> 

### Item

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ApplicationConfiguration`

A string that identifies this setting.

## Topics

### Objects

- [SettingsCommand.Command.Settings.ApplicationConfiguration.Configuration](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/applicationconfiguration-data.dictionary/configuration-data.dictionary) - A dictionary that contains the configurations to apply to the app.

