# SettingsCommand.Command

The command to configure settings on a device.

**Platforms:** iOS 5.0, iPadOS 5.0, Mac Catalyst 5.0, macOS 10.9, tvOS 9.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Settings`

The request type for this command.

### Settings

- **Type:** `SettingsCommand.Command.Settings`
- **Required:** Yes

An array of dictionaries that contains the settings.

## Topics

### Objects

- [SettingsCommand.Command.Settings](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary) - An array of dictionaries that contains the settings.

