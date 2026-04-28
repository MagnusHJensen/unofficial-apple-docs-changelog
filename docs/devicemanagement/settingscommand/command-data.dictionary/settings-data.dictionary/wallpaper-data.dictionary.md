# SettingsCommand.Command.Settings.Wallpaper

A dictionary that contains wallpaper settings.

**Platforms:** iOS 8.0, iPadOS 8.0, Mac Catalyst 8.0, Device Assignment Services , VPP License Management 

## Properties

### Image

- **Type:** `data`
- **Required:** Yes

A Base64-encoded image in either PNG or JPG format to use for wallpaper.

### Item

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Wallpaper`

A string that identifies this setting.

### Where

- **Type:** `integer`
- **Required:** Yes
- **Allowed Values:** `1`, `2`, `3`

A number that indicates where to use the wallpaper, which is one of the following values:

- `1`: Lock Screen
- `2`: Home Screen
- `3`: Both Lock and Home Screens. In iOS 16 and later, and iPadOS 17 and later, when you set the wallpaper for the first time, the system sets both the Lock Screen and Home Screen. After that, you can separately set each location.

