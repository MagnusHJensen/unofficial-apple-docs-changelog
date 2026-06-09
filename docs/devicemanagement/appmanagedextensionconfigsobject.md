# AppManagedExtensionConfigsObject

A dictionary of extension config data and credentials.

**Platforms:** iOS 18.4, iPadOS 18.4, Mac Catalyst 18.4, macOS 27.0 (Beta), visionOS 2.4

## Properties

### ANY

- **Type:** `AppManagedAppConfigDictionaryObject`
- **Required:** No

A dictionary mapping extension composed identifiers to the extension config data and credentials.

The format of the composed identifier is either “Bundle-ID” or “Bundle-ID (Team-ID)”. “Bundle-ID” is the bundle identifier string of the provider. “Team-ID” is the team identifier from the provider’s code signature. For example, “com.example.app” for the bundle ID format, or “com.example.app (ABCD1234)” for the team ID format.

