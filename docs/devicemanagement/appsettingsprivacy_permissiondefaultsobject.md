# AppSettingsPrivacy_PermissionDefaultsObject

The dictionary of app privacy permission defaults. Each key in the dictionary is an app identifier. The dictionary values represent the permission defaults that the device applies for each matching app.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta)

## Properties

### ANY

- **Type:** `AppSettingsAppDictionaryObject`
- **Required:** No

The dictionary that defines the app privacy permission defaults. Each key is an app identifier.

## Discussion

In iOS, the app identifier is a bundle ID, for example, “com.example.app”.

In macOS, the app identifier is a composed identifier. The format of the composed identifier is “Bundle-ID {Designated-Requirement}”. “Bundle-ID” is the bundle identifier string of the app. “Designated-Requirement” is the designated requirement string the device uses to match the code signature of the app. For example, “com.example.app {anchor apple generic}”. The device only applies defaults for an app if its code signature matches the composed identifier.

## Topics

### Objects

- [AppSettingsAppDictionaryObject](/documentation/devicemanagement/appsettingsappdictionaryobject) - The dictionary that defines the app privacy permission defaults. Each key is an app identifier.

