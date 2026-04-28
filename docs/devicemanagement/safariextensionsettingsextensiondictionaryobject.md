# SafariExtensionSettingsExtensionDictionaryObject

The dictionary that defines managed extensions.

**Platforms:** iOS 18.0, iPadOS 18.0, Mac Catalyst 18.0, macOS 15.0, visionOS 26.0, Device Assignment Services , VPP License Management 

## Properties

### AllowedDomains

- **Type:** `[string]`
- **Required:** No

Controls the domains and sub-domains the extension is granted access to.

### DeniedDomains

- **Type:** `[string]`
- **Required:** No

Controls the domains and sub-domains the extension isn’t allowed to access.

### PrivateBrowsing

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Allowed`, `AlwaysOn`, `AlwaysOff`

Controls whether an extension is allowed in Private Browsing.

- `Allowed` - The user is allowed to turn the extension on or off in Private Browsing.
- `AlwaysOn` - The extension will always be on in Private Browsing if the extension is on outside of Private Browsing.
- `AlwaysOff` - The extension will never be on in Private Browsing.

### State

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Allowed`, `AlwaysOn`, `AlwaysOff`

Controls whether an extension is allowed.

- `Allowed` - The user is allowed to turn the extension on or off.
- `AlwaysOn` - The extension will always be on.
- `AlwaysOff` - The extension will always be off.

