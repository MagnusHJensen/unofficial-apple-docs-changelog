# SafariExtensionSettingsExtensionDictionaryObject

The dictionary that defines the settings for a managed extension. Each key represents a specific managed extension, or you can specify a single “*” character to match any extension.

**Platforms:** iOS 18.0, iPadOS 18.0, Mac Catalyst 18.0, macOS 15.0, visionOS 26.0

## Properties

### AllowedDomains

- **Type:** `[string]`
- **Required:** No

Controls the domains and sub-domains the extension is granted access to. The device ignores this key when the extension identifier is a single “*” character.

### DeniedDomains

- **Type:** `[string]`
- **Required:** No

Controls the domains and sub-domains the extension isn’t allowed to access. The device uses this key when the extension identifier is a composed identifier or a single “*” character.

### PrivateBrowsing

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Allowed`, `AlwaysOn`, `AlwaysOff`

Controls whether an extension is allowed in Private Browsing. The device uses this key when the extension identifier is a composed identifier or a single “*” character.

- `Allowed` - The user is allowed to turn the extension on or off in Private Browsing.
- `AlwaysOn` - The extension will always be on in Private Browsing if the extension is on outside of Private Browsing.
- `AlwaysOff` - The extension will never be on in Private Browsing.

### State

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Allowed`, `AlwaysOn`, `AlwaysOff`

Controls whether an extension is allowed. The device uses this key when the extension identifier is a composed identifier or a single “*” character.

- `Allowed` - The user is allowed to turn the extension on or off.
- `AlwaysOn` - The extension will always be on.
- `AlwaysOff` - The extension will always be off.

