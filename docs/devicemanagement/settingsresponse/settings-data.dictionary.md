# SettingsResponse.Settings

A dictionary that describes the results of configuring settings on a device.

**Platforms:** iOS 5.0, iPadOS 5.0, Mac Catalyst 5.0, macOS 10.9, tvOS 9.0, visionOS 1.1, watchOS 10.0, Device Assignment Services , VPP License Management 

## Properties

### ErrorChain

- **Type:** `[SettingsResponse.Settings.ErrorChainItem]`
- **Required:** No

An array of dictionaries that describes any errors that occurred.

### Identifier

- **Type:** `string`
- **Required:** No

The app identifier to which this error applies.

> 

### Status

- **Type:** `string`
- **Required:** Yes

The status of the setting, which is one of the following values:

- `Acknowledged`: The device processed the command successfully.
- `Error`: An error occurred. See the `ErrorChain` for more details.

## Topics

### Objects

- [SettingsResponse.Settings.ErrorChainItem](/documentation/devicemanagement/settingsresponse/settings-data.dictionary/errorchainitem) - A dictionary that describes an error chain item.

