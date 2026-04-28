# SettingsCommand.Command.Settings.SoftwareUpdateSettings

A dictionary that contains software update settings.

**Platforms:** iOS 14.5, iPadOS 14.5, Mac Catalyst 14.5, Device Assignment Services , VPP License Management 

## Properties

### Item

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `SoftwareUpdateSettings`

A string that represents the type of updates that should appear in the Software Update pane in Settings. Supervised only.

### RecommendationCadence

- **Type:** `integer`
- **Required:** Yes
- **Allowed Values:** `0`, `1`, `2`

This value defines how the system presents software updates to the user. When there’s more than one available update for the user, the system behaves as follows:

- `0`: Presents both options to the user.
- `1`: Presents the lower numbered (oldest) software update version.
- `2`: Presents only the highest numbered (most recent) release available for the device.

This value has no effect when there’s only one available update; the system shows the single available update to the user regardless of the value of this setting.

Available in iOS 14.5 and later.

