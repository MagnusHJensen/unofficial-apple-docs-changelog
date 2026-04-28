# AppManagedUpdateBehaviorObject

Specifies the update behavior of the apps installed from the App Store. Apps in packages are not automatically updated.

**Platforms:** iOS 26.0, iPadOS 26.0, Mac Catalyst 26.0, macOS 26.0, visionOS 26.0, Device Assignment Services , VPP License Management 

## Properties

### AutomaticAppUpdates

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `AlwaysOn`, `AlwaysOff`, `StoreSettings`

Specifies whether the device automatically updates the app:

- `AlwaysOn`: The device automatically updates the app to the latest version. For App Store apps, the device periodically checks the store for updates. For Enterprise apps, the device periodically downloads the manifest file and compares it to the previous manifest file. If the device detects a change to the bundle version in the manifest, it downloads and updates the app.
- `AlwaysOff`: The device never automatically updates the app.
- `StoreSettings`: The device uses the settings for the corresponding store to determine when to automatically update the app. For Enterprise apps, this setting behaves the same as `AlwaysOff`.

When the `InstallBehavior.Version` key is specified, the device ignores this key and Automatic App Updates are disabled.

In macOS, the device ignores this setting if the `AppComposedIdentifier` key is set in the configuration.

