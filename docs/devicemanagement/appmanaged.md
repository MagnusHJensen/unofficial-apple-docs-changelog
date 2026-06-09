# AppManaged

The declaration to configure a managed app.

**Platforms:** iOS 17.2, iPadOS 17.2, Mac Catalyst 17.2, macOS 26.0, visionOS 2.4

## Properties

### AppComposedIdentifier

- **Type:** `string`
- **Required:** No

A string that specifies the composed identifier of an existing app that needs to be managed. The device uses this to take over management of an app installed by some other process, for example installed manually by the user, or via a package configuration. If the app isn’t present when the device applies the configuration, the device takes over management of it when it does install. Management of the app occurs only if its code signature matches the composed identifier.

The following rules apply when the device takes over management:

- If the `InstallBehavior.Install` key is set to `Required`, the device takes over management of the app.
- If the `InstallBehavior.Install` key is set to `Optional`, the device takes over management of the app when the user “installs” it using an MDM management app.

The format of the composed identifier is either “Bundle-ID” or “Bundle-ID (Team-ID)”. “Bundle-ID” is the bundle identifier string of the provider. “Team-ID” is the team identifier from the provider’s code signature. For example, “com.example.app” for the bundle ID format, or “com.example.app (ABCD1234)” for the team ID format.

In macOS, only one of `AppStoreID`, `BundleID`, or `AppComposedIdentifier` needs to be present.

Available: macOS 26+

### AppConfig

- **Type:** `AppManagedAppConfigDictionaryObject`
- **Required:** No

A dictionary of app config data and credentials.

Available: iOS 18.4+ | iPadOS 18.4+ | macOS 27+ | visionOS 2.4+

### AppStoreID

- **Type:** `string`
- **Required:** No

The App Store ID of the managed app that is downloaded from the App Store.

Only one of `AppStoreID`, `BundleID`, `ManifestURL`, or `AppComposedIdentifier` needs to be present.

### Attributes

- **Type:** `AppManagedAttributesObject`
- **Required:** No

A dictionary of values to associate with the app.

Available: iOS 17.2+ | iPadOS 17.2+ | visionOS 2.4+

### BundleID

- **Type:** `string`
- **Required:** No

The bundle ID of the managed app that is downloaded from the App Store.

Only one of `AppStoreID`, `BundleID`, `ManifestURL`, or `AppComposedIdentifier` needs to be present.

### ExtensionConfigs

- **Type:** `AppManagedExtensionConfigsObject`
- **Required:** No

A dictionary of extension config data and credentials.

Available: iOS 18.4+ | iPadOS 18.4+ | macOS 27+ | visionOS 2.4+

### IncludeInBackup

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, backups contain the app and its data.

Available: iOS 17.2+ | iPadOS 17.2+ | visionOS 2.4+

### InstallBehavior

- **Type:** `AppManagedInstallBehaviorObject`
- **Required:** No

A dictionary that describes how and when to install the app.

### iOSApp

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device installs an iOS or iPadOS app that runs on a Mac with Apple silicon. This is only used when the app is an App Store app.

Available: macOS 26+

### LegacyAppConfigAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration containing a reference to the app config data. The device provides the app config data to the app using the MDMv1 behavior. The corresponding asset needs to be of type `com.apple.asset.data`. The referenced data needs to be a property list file, and the asset’s “ContentType” value set to match the data type.

Available: iOS 18.4+ | iPadOS 18.4+ | macOS 27+ | visionOS 2.4+

### ManifestURL

- **Type:** `string`
- **Required:** No

The URL of the manifest for the managed app that the device downloads from a web site. The manifest is a [ManifestURL](/documentation/devicemanagement/manifesturl) property list.

Only one of `AppStoreID`, `BundleID`, `ManifestURL`, or `AppComposedIdentifier` needs to be present.

Available: iOS 17.2+ | iPadOS 17.2+ | visionOS 2.4+

### UpdateBehavior

- **Type:** `AppManagedUpdateBehaviorObject`
- **Required:** No

A dictionary that specifies how the device updates apps.

Available: iOS 26+ | iPadOS 26+ | macOS 26+ | visionOS 26+

## Discussion

Specify `com.apple.configuration.app.managed` as the declaration type.

### Configuration availability

### Configuration examples

## Topics

### Objects

- [AppManagedAppConfigDictionaryObject](/documentation/devicemanagement/appmanagedappconfigdictionaryobject) - A dictionary of app config data and credentials.
- [AppManagedAttributesObject](/documentation/devicemanagement/appmanagedattributesobject) - A dictionary of values to associate with the app.
- [AppManagedExtensionConfigsObject](/documentation/devicemanagement/appmanagedextensionconfigsobject) - A dictionary of extension config data and credentials.
- [AppManagedInstallBehaviorObject](/documentation/devicemanagement/appmanagedinstallbehaviorobject) - A dictionary that describes how and when to install the app.
- [AppManagedUpdateBehaviorObject](/documentation/devicemanagement/appmanagedupdatebehaviorobject) - A dictionary that specifies how the device updates apps.

