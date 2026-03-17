# InstalledApplicationListResponse.InstalledApplicationListItem

A dictionary that describes an app list item.

**Platforms:** iOS 5.0, iPadOS 5.0, macOS 10.7, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Properties

### AdHocCodeSigned

- **Type:** `boolean`
- **Required:** No

If `true`, the app is ad-hoc code signed. This query is available in iOS 11.3 and later, and tvOS 11.3 and later.

### AppStoreVendable

- **Type:** `boolean`
- **Required:** No

If `true`, the app came from the App Store and can participate in store features. For device-based Volume Purchase Program (VPP) apps, this value is `false`. This value is available in iOS 11.3 and later, and tvOS 11.3 and later.

### BetaApp

- **Type:** `boolean`
- **Required:** No

If `true`, the app is part of the Apple Beta Software Program. This value is available in iOS 11.3 and later, and tvOS 11.3 and later.

### BundleSize

- **Type:** `integer`
- **Required:** No

The app’s static bundle size, in bytes. This value is expensive to calculate. Starting in iOS 26, macOS 26, tvOS 26, watchOS 26, and visionOS 26 it isn’t present in the response unless it is included in the `Items` request key. This value is available in iOS 5 and later, and macOS 10.7 and later, tvOS 10.2 and later, watchOS 10 and later, and visionOS 1.1 and later.

### DeviceBasedVPP

- **Type:** `boolean`
- **Required:** No

If `true`, installing the app didn’t require an Apple Account. This value is available in iOS 11.3 and later, and tvOS 11.3 and later.

### DistributorIdentifier

- **Type:** `string`
- **Required:** No

The marketplace hosted application’s distributor ID. This value is available in iOS 17.4 and later.

### DownloadCancelled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the user canceled the download.

### DownloadFailed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the download failed.

### DownloadPaused

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the user paused the download.

### DownloadWaiting

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the app is in the initial state, which is waiting to download.

### DynamicSize

- **Type:** `integer`
- **Required:** No

The size of the app’s file system in bytes, including the Documents, Library, and other directories. This value is expensive to calculate. Starting in iOS 26, tvOS 26, watchOS 26, and visionOS 26 it isn’t present in the response unless it is included in the `Items` request key. This value is available in iOS 5 and later, tvOS 10.2 and later, watchOS 10 and later, and visionOS 1.1 and later.

### ExternalVersionIdentifier

- **Type:** `integer`
- **Required:** No

The app’s external version identifier. You can also retrieve this value from the App Store. For more information, see [Apps and Books for Organizations](/documentation/devicemanagement/apps-and-books-for-organizations).

If the current external version identifier of an app on the App Store doesn’t match the external version identifier reported by the device, there may be an app update available for the device.

> 

### HasUpdateAvailable

- **Type:** `boolean`
- **Required:** No

If `true`, the app has an update available. This key is present only for App Store apps. In macOS, this key is present only for Volume Purchase Program (VPP) apps. This status updates daily and isn’t always up-to-date when installing an app.

### Identifier

- **Type:** `string`
- **Required:** No

The app’s identifier. This key is always be present on iOS and tvOS, but may be missing on macOS.

> 

### Installing

- **Type:** `boolean`
- **Required:** No

If `true`, the app is downloading. If `false`, it’s already installed.

### IsAppClip

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the app is an App Clip. Available in iOS 16 and later.

### IsValidated

- **Type:** `boolean`
- **Required:** No

If `true`, the app is valid and can run on the device. If the app is enterprise-distributed and unvalidated, it won’t be able to run until validation has occurred. This value is available in iOS 9.2 and later, and tvOS 10.2 and later.

### Name

- **Type:** `string`
- **Required:** No

The app’s name.

### ShortVersion

- **Type:** `string`
- **Required:** No

The app’s short version.

### Source

- **Type:** `string`
- **Required:** No

The source of the application. When the app is managed by Declarative Device Management this value is `Declarative Device Management`.

### Version

- **Type:** `string`
- **Required:** No

The app’s version.

