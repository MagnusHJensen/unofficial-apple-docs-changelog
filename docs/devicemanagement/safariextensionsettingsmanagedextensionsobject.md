# SafariExtensionSettingsManagedExtensionsObject

The dictionary of managed extensions settings. Each key in the dictionary represents a composed identifier for a specific managed extension, or you can specify a single “*” character to match any extension. The dictionary values represent the settings that Safari applies to each extension that matches the key. In order for the extension to be managed, its host app needs to be present on the device.

**Platforms:** iOS 18.0, iPadOS 18.0, Mac Catalyst 18.0, macOS 15.0, visionOS 26.0

## Properties

### ANY

- **Type:** `SafariExtensionSettingsExtensionDictionaryObject`
- **Required:** No

The dictionary that defines the settings for a managed extension. Each key represents a specific managed extension, or you can specify a single “*” character to match any extension.

## Discussion

The composed identifier of a managed extension uses the format “Identifier (TeamIdentifier)”, for example “com.example.app (ABCD1234)”. Use `codesign -dv <path_to_appex>` to show the information you need to generate this string on macOS, using the path to the extension bundle located in the “PlugIns” folder inside the app bundle. For other platforms, request this information from the app developer.

## Topics

### Objects

- [SafariExtensionSettingsExtensionDictionaryObject](/documentation/devicemanagement/safariextensionsettingsextensiondictionaryobject) - The dictionary that defines the settings for a managed extension. Each key represents a specific managed extension, or you can specify a single “*” character to match any extension.

