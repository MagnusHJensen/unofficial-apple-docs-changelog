# AppManagedInstallBehaviorObject

A dictionary that describes how and when to install an app.

**Platforms:** iOS 17.2, iPadOS 17.2, Mac Catalyst 17.2, macOS 26.0, visionOS 2.4, Device Assignment Services , VPP License Management 

## Properties

### AllowDownloadsOverCellular

- **Type:** `string`
- **Required:** No
- **Default:** `StoreSettings`
- **Allowed Values:** `AlwaysOn`, `AlwaysOff`, `StoreSettings`

Indicates how the device uses a cellular network when it downloads the app for automatic install or update operations:

- `AlwaysOn`: The device downloads apps of any size using a cellular network.
- `AlwaysOff`: The device doesn’t download apps using a cellular network. The device pauses the automatic install or update operation until a different network is active.
- `StoreSettings`: The device uses the settings for the corresponding store when downloading apps.

The device always uses the store settings to download apps when the install or update operation is user initiated.

Available only in iOS.

### Install

- **Type:** `string`
- **Required:** No
- **Default:** `Optional`
- **Allowed Values:** `Optional`, `Required`

A string that specifies if the app needs to remain on the device at all times or if the user can freely install and remove it, which is one of the following values:

- `Optional`: The user can install and remove the app after the system activates the configuration.
- `Required`: The system installs the app after it activates the configuration. The user can’t remove the app.

The system automatically installs apps on supervised devices. Otherwise, the device prompts the user to approve installation of the app.

### License

- **Type:** `AppManagedInstallBehavior_LicenseObject`
- **Required:** No

A dictionary that describes the app’s license.

### Version

- **Type:** `integer`
- **Required:** No

The App Store external version identifier (EVID) of the version of the app the device installs. You can retrieve this value from the App Store. For more information, see [Apps and Books for Organizations](/documentation/devicemanagement/apps-and-books-for-organizations). This key is ignored if the app isn’t an App Store app.

The following rules apply when the device applies or updates the configuration:

- If this key isn’t present:
- If this key is present:

> 

## Topics

### Objects

- [AppManagedInstallBehavior_LicenseObject](/documentation/devicemanagement/appmanagedinstallbehavior_licenseobject) - A dictionary that specifies the type of license the app uses.

