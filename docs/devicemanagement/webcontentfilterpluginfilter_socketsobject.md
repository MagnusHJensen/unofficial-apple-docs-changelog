# WebContentFilterPluginFilter_SocketsObject

Settings that control the socket filter. If not present, the system doesn’t use socket filtering.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### Enabled

- **Type:** `boolean`
- **Required:** Yes

If `true`, enables the filtering of socket traffic.

### ProviderComposedIdentifier

- **Type:** `string`
- **Required:** No

The data provider identifier. This string identifies the filter data provider when the filter starts running. Required when Enabled is true.

In iOS and visionOS, the identifier is a bundle ID, for example, “com.example.app”.

In macOS, the identifier is a composed identifier. The format of the composed identifier is “Bundle-ID {Designated-Requirement}”. “Bundle-ID” is the bundle identifier string of the provider. “Designated-Requirement” is the designated requirement string from the code signature of the provider. For example, “com.example.app {anchor apple generic}”.

