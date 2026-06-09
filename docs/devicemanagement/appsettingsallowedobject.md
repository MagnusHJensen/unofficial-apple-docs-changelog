# AppSettingsAllowedObject

The dictionary of allowed app settings.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), tvOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### AllowedApps

- **Type:** `[string]`
- **Required:** No

If present, the device only shows or launches apps with bundle IDs in the array. Include the value `com.apple.webapp` to allow all webclips. This applies to App Store apps, marketplace apps, and locally installed apps (using Configurator, Xcode, and so forth).

Available: iOS 27+ | iPadOS 27+ | tvOS 27+ | visionOS 27+

### AllowedBinaries

- **Type:** `[AppSettingsAllowed_BinaryIdentifierObject]`
- **Required:** No

If present, the device only allows binaries that match the binary identifier properties to run. A binary is matched only when all the binary identifiers match. The device always runs system critical processes. Use “codesign -dvvv <path_to_binary>” to show the information you need to generate these values.

Available: macOS 27+
Allowed scopes: system

### AlwaysAllowManagedApps

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device implicitly includes managed apps in the effective allow list when `AllowedApps` or `AllowedBinaries` is present.

Available: macOS 27+
Allowed scopes: system

### DeniedApps

- **Type:** `[string]`
- **Required:** No

If present, the device prevents showing or launching apps with bundle IDs in the array. Include the value `com.apple.webapp` to restrict all webclips. This applies to App Store apps, marketplace apps, and locally installed apps (using Configurator, Xcode, and so forth).

> 

Available: iOS 27+ | iPadOS 27+ | tvOS 27+ | visionOS 27+

### DeniedBinaries

- **Type:** `[AppSettingsAllowed_BinaryIdentifierObject]`
- **Required:** No

If present, the device doesn’t allow binaries that match the binary identifier properties to run. A binary is matched only when all the binary identifiers match.

Available: macOS 27+
Allowed scopes: system

## Topics

### Objects

- [AppSettingsAllowed_BinaryIdentifierObject](/documentation/devicemanagement/appsettingsallowed_binaryidentifierobject) - Dictionary containing one or more identifier fields to match a binary.

