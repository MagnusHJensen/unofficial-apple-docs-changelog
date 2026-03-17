# AppManaged

The declaration to configure a managed app.

**Platforms:** iOS 17.2, iPadOS 17.2, macOS 26.0, visionOS 2.4

## Properties

### AppComposedIdentifier

- **Type:** `string`
- **Required:** No

A string that specifies the composed identifier of an existing app that needs to be managed. The device uses this to take over management of an app installed by some other process, for example installed manually by the user, or via a package configuration. If the app isn’t present when the device applies the configuration, the device takes over management of it when it does install.

The following rules apply when the device takes over management:

- If the `InstallBehavior.Install` key is set to `Required`, the device takes over management of the app.
- If the `InstallBehavior.Install` key is set to `Optional`, the device takes over management of the app when the user “installs” it using an MDM management app.

The format of the composed identifier is either “Bundle-ID (Team-ID)” or “Bundle-ID {Designated-Requirement}”. For example, `com.example.app (ABCD1234)` for the team ID format, or `com.example.app {anchor apple generic}` for the designated requirement format. Management of the app occurs only if its code signature matches the composed identifier.

In macOS, only one of `AppStoreID`, `BundleID`, or `AppComposedIdentifier` needs to be present.

Available only in macOS.

### AppConfig

- **Type:** `AppManagedAppConfigDictionaryObject`
- **Required:** No

A dictionary of app config data and credentials.

Available only in iOS and visionOS.

### AppStoreID

- **Type:** `string`
- **Required:** No

The App Store ID of the managed app that is downloaded from the App Store.

Only one of `AppStoreID`, `BundleID`, `ManifestURL`, or `AppComposedIdentifier` needs to be present.

### Attributes

- **Type:** `AppManagedAttributesObject`
- **Required:** No

A dictionary of values to associate with the app.

Available only in iOS and visionOS.

### BundleID

- **Type:** `string`
- **Required:** No

The bundle ID of the managed app that is downloaded from the App Store.

Only one of `AppStoreID`, `BundleID`, `ManifestURL`, or `AppComposedIdentifier` needs to be present.

### ExtensionConfigs

- **Type:** `AppManagedExtensionConfigsObject`
- **Required:** No

A dictionary of extension config data and credentials.

Available only in iOS and visionOS.

### IncludeInBackup

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, backups contain the app and its data.

Available only in iOS and visionOS.

### InstallBehavior

- **Type:** `AppManagedInstallBehaviorObject`
- **Required:** No

A dictionary that describes how and when to install the app.

### iOSApp

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device installs an iOS or iPadOS app that runs on a Mac with Apple Silicon. This is only used when the app is an App Store app.

Available only in macOS.

### LegacyAppConfigAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration containing a reference to the app config data. The device provides the app config data to the app using the MDMv1 behavior. The corresponding asset needs to be of type `com.apple.asset.data`. The referenced data needs to be a property list file, and the asset’s “ContentType” value set to match the data type.

Available only in iOS and visionOS.

### ManifestURL

- **Type:** `string`
- **Required:** No

The URL of the manifest for the managed app that the device downloads from a web site. The manifest is returned as a [ManifestURL](/documentation/devicemanagement/manifesturl) property list.

Only one of `AppStoreID`, `BundleID`, `ManifestURL`, or `AppComposedIdentifier` needs to be present.

Available only in iOS and visionOS.

### UpdateBehavior

- **Type:** `AppManagedUpdateBehaviorObject`
- **Required:** No

A dictionary that specifies how the device updates apps.

## Discussion

Specify `com.apple.configuration.app.managed` as the declaration type.

### Configuration availability

### Configuration examples

## Topics

### Objects

- [AppManagedAppConfigDictionaryObject](/documentation/devicemanagement/appmanagedappconfigdictionaryobject) - A dictionary of app config data and credentials.
- [AppManagedAttributesObject](/documentation/devicemanagement/appmanagedattributesobject) - A dictionary of values associated with an app.
- [AppManagedExtensionConfigsObject](/documentation/devicemanagement/appmanagedextensionconfigsobject) - A dictionary of values associated with an extension config.
- [AppManagedInstallBehaviorObject](/documentation/devicemanagement/appmanagedinstallbehaviorobject) - A dictionary that describes how and when to install an app.
- [AppManagedUpdateBehaviorObject](/documentation/devicemanagement/appmanagedupdatebehaviorobject) - Specifies the update behavior of the apps installed from the App Store. Apps in packages are not automatically updated.

