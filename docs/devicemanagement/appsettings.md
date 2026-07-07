# AppSettings

The declaration to configure app settings.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), tvOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### Allowed

- **Type:** `AppSettingsAllowedObject`
- **Required:** No

The dictionary of allowed app settings.

### Privacy

- **Type:** `AppSettingsPrivacyObject`
- **Required:** No

The dictionary of app settings.

Available: iOS 27+ | iPadOS 27+ | macOS 27+
Allowed scopes: iOS: system | macOS: user

## Discussion

Specify `com.apple.configuration.app.settings` as the declaration type.

### Binary identifier rules

The following combinations of binary identifiers are supported for each key:

- `AllowedBinaries`:
- `DeniedBinaries`:

### Privacy permission defaults

Privacy permission defaults allow an organization to suggest a set of privacy permissions for use with an app. When set, the app displays a consent prompt listing all the configured defaults. If the user accepts, the device applies those defaults for the app. If the user declines, no defaults are set and the device prompts the user in the normal way when the app requires permission.

The consent prompt only shows permissions that the user hasn’t previously seen, and won’t appear if the user has seen all permissions. The user can choose from one of two options in the prompt:

- `Allow`: this option sets the app privacy permissions for the specified sub-systems (camera, microphone, and so on) to “Allow”. The device doesn’t prompt the user when the app uses the sub-system.
- `Not Now`: this option ignores the app privacy permission defaults for the specified sub-systems (camera, microphone, and so on). The device prompts the user in the normal way when the app uses the sub-system.

The user can change the app permission privacy settings in Settings.app if they choose.

Only AppKit-based apps on macOS support this feature.

### Configuration availability

### App privacy examples

## Topics

### Objects

- [AppSettingsAllowedObject](/documentation/devicemanagement/appsettingsallowedobject) - The dictionary of allowed app settings.
- [AppSettingsPrivacyObject](/documentation/devicemanagement/appsettingsprivacyobject) - The dictionary of app settings.

