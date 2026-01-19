# StatusAppManagedList

The device’s declarative managed apps.

**Platforms:** iOS 17.2, iPadOS 17.2, macOS 26.0, visionOS 2.4

## Discussion

### Status item availability

### Reason Codes

- `Error.AppStoreDisabled`: The App Store is disabled.
- `Error.DownloadFailed`: The app download failed.
- `Error.DuplicateConfiguredApp`: The app is already being managed.
- `Error.InstallFailed`: The app install failed.
- `Error.InvalidAppID`: The app id could not be found.
- `Error.InvalidCodeSignature`: The code signature of the app does not match the composed identifier, and the app cannot be managed.
- `Error.IsSystemApp`: The app is a system app that cannot be managed.
- `Error.LicenseNotFound`: A license for the app was not available.
- `Error.NotAnApp`: The downloaded data is not a valid app.
- `Error.NotSupported`: The app is not supported on this device.
- `Error.UnmanagedAppAlreadyInstalled`: An unmanaged app is already installed and cannot be managed.
- `Error.UserRejected`: The user rejected management of the app.
- `Info.UpdateAvailable`: An update is available for the app.
- `Error.UpdateFailed`: The app update failed.

## Topics

### Objects

- [StatusAppManagedListAppObject](/documentation/devicemanagement/statusappmanagedlistappobject) - A dictionary that describes a declarative managed app.

