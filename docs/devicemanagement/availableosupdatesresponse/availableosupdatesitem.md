# AvailableOSUpdatesResponse.AvailableOSUpdatesItem

The response dictionary that describes the available operating-system updates item.

**Platforms:** iOS 9.0, iPadOS 9.0, macOS 10.11, tvOS 12.0

## Properties

### AllowsInstallLater

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, download the software update and install it later.

### AppIdentifiersToClose

- **Type:** `[string]`
- **Required:** Yes

An array that contains app identifiers of apps to close so you can install the update. This value is available in macOS 10.11 and later.

### Build

- **Type:** `string`
- **Required:** Yes

The build number of the update.

### DeferredUntil

- **Type:** `date`
- **Required:** No

If present, the date when you want the update to install. This value is available in macOS 10.12.4 and later.

### DownloadSize

- **Type:** `integer`
- **Required:** Yes

The storage size necessary to download the software update. Prior to macOS 10.14, this only includes major operating-system updates. In macOS 10.14 and later, this also includes minor updates.

### HumanReadableName

- **Type:** `string`
- **Required:** Yes

The human-readable name of the update in the current user’s current locale.

### HumanReadableNameLocale

- **Type:** `string`
- **Required:** Yes

The locale, in IOS639-1 Alpha-2 code format, of the `HumanReadableName` value. This value is available in macOS 10.11 and later.

### InstallSize

- **Type:** `integer`
- **Required:** Yes

The storage size necessary to install the update. This value is available in iOS 9.0 and later, and tvOS 12.0 and later.

### IsConfigDataUpdate

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, this is an update to a configuration file. This value is available in macOS 10.11 and later.

### IsCritical

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, this is a critical update.

### IsFirmwareUpdate

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, this is an update to firmware. This value is available in macOS 10.11 and later.

### IsMajorOSUpdate

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, this is a major update; for example, 10.15.x to 11. This value is available in macOS 10.11 and later.

### IsSecurityResponse

- **Type:** `boolean`
- **Required:** Yes

If `true`, this update is a Background Security Improvement.

### MetadataURL

- **Type:** `string`
- **Required:** Yes

A URL where the MDM server can request additional localized names for this update. This key isn’t present for certain updates, such as mobile software updates (MSUs) or major OS updates. This value is available in macOS 10.11 and later.

### ProductKey

- **Type:** `string`
- **Required:** Yes

The product key that represents the update.

### ProductName

- **Type:** `string`
- **Required:** Yes

The product name; for example, **. This value is available in iOS 9.0 and later, and tvOS 12.0 and later.

### RequiresBootstrapToken

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device can accept a Bootstrap Token from the MDM server instead of prompting for user authentication prior to installation. This only applies when `BootstrapTokenAllowedForAuthentication` is `true` in the [SecurityInfoResponse.SecurityInfo](/documentation/devicemanagement/securityinforesponse/securityinfo-data.dictionary) response. This value is available for a Mac with Apple silicon in macOS 11 and later.

### RestartRequired

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device restarts after installing the update.

### SupplementalBuildVersion

- **Type:** `string`
- **Required:** No

The build version for the Background Security Improvement update, for example, `13A999`, which is the same as `Build`.

### SupplementalOSVersionExtra

- **Type:** `string`
- **Required:** No

The Background Security Improvement OS version suffix, for example, `(a)`. Only present if this is a Background Security Improvement update.

### Version

- **Type:** `string`
- **Required:** Yes

The version of the update.

