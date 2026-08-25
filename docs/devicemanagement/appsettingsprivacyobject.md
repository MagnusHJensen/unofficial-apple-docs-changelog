# AppSettingsPrivacyObject

The dictionary of app settings.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta)

## Properties

### PermissionDefaults

- **Type:** `AppSettingsPrivacy_PermissionDefaultsObject`
- **Required:** No

The dictionary of app privacy permission defaults. Each key in the dictionary is an app identifier. The dictionary values represent the permission defaults that the device applies for each matching app.

In iOS, the app identifier is a bundle ID, for example, “com.example.app”.

In macOS, the app identifier is a composed identifier. The format of the composed identifier is “Bundle-ID {Designated-Requirement}”. “Bundle-ID” is the bundle identifier string of the app. “Designated-Requirement” is the designated requirement string the device uses to match the code signature of the app. For example, “com.example.app {anchor apple generic}”. The device only applies defaults for an app if its code signature matches the composed identifier.

## Topics

### Objects

- [AppSettingsPrivacy_PermissionDefaultsObject](/documentation/devicemanagement/appsettingsprivacy_permissiondefaultsobject) - The dictionary of app privacy permission defaults. Each key in the dictionary is an app identifier. The dictionary values represent the permission defaults that the device applies for each matching app.

