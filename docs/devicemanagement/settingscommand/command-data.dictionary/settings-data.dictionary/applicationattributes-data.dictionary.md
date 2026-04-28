# SettingsCommand.Command.Settings.ApplicationAttributes

A dictionary that contains the attributes to apply to the app.

**Platforms:** iOS 7.0, iPadOS 7.0, Mac Catalyst 7.0, tvOS 10.2, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### Attributes

- **Type:** `SettingsCommand.Command.Settings.ApplicationAttributes.Attributes`
- **Required:** No

A dictionary that contains the attributes to apply to the app. Omit this setting to remove existing attributes. This setting is available in iOS 7 and later, and tvOS 10.2 and later.

### Identifier

- **Type:** `string`
- **Required:** Yes

The bundle identifier of the app.

> 

### Item

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ApplicationAttributes`

A string that identifies this setting.

## Topics

### Objects

- [SettingsCommand.Command.Settings.ApplicationAttributes.Attributes](/documentation/devicemanagement/settingscommand/command-data.dictionary/settings-data.dictionary/applicationattributes-data.dictionary/attributes-data.dictionary) - A dictionary that contains the attributes to apply to the app.

