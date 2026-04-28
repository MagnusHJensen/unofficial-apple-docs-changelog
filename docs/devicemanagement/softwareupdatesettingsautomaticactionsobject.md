# SoftwareUpdateSettingsAutomaticActionsObject

The object that configures various automatic Software Update functionality.

**Platforms:** iOS 18.0, iPadOS 18.0, Mac Catalyst 18.0, macOS 15.0, tvOS 18.4, visionOS 26.0, Device Assignment Services , VPP License Management 

## Properties

### Download

- **Type:** `string`
- **Required:** No
- **Default:** `Allowed`
- **Allowed Values:** `Allowed`, `AlwaysOn`, `AlwaysOff`

Specifies whether the user can control automatic downloads of available updates:

- `Allowed` - the user can enable or disable automatic downloads.
- `AlwaysOn` - automatic downloads are always enabled.
- `AlwaysOff` - automatic downloads are always disabled.

### InstallOSUpdates

- **Type:** `string`
- **Required:** No
- **Default:** `Allowed`
- **Allowed Values:** `Allowed`, `AlwaysOn`, `AlwaysOff`

Specifies whether the user can control automatic installation of available updates:

- `Allowed` - the user can enable or disable automatic installation.
- `AlwaysOn` - automatic installations are always enabled.
- `AlwaysOff` - automatic installations are always disabled.

### InstallSecurityUpdate

- **Type:** `string`
- **Required:** No
- **Default:** `Allowed`
- **Allowed Values:** `Allowed`, `AlwaysOn`, `AlwaysOff`

Specifies whether the user can control automatic installation of available security updates:

- `Allowed` - the user can enable or disable automatic installation.
- `AlwaysOn` - automatic installations are always enabled.
- `AlwaysOff` - automatic installations are always disabled.

