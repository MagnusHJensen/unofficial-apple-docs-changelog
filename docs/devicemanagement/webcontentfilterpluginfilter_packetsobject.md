# WebContentFilterPluginFilter_PacketsObject

Settings that control the packet filter. If not present, the system doesn’t use packet filtering.

**Platforms:** macOS 27.0 (Beta)

## Properties

### Enabled

- **Type:** `boolean`
- **Required:** Yes

If `true`, the system enables filtering network packets.

### ProviderComposedIdentifier

- **Type:** `string`
- **Required:** No

The packet provider identifier. This string identifies the filter data provider when the filter starts running. Required when Enabled is true.

The identifier is a composed identifier. The format of the composed identifier is “Bundle-ID {Designated-Requirement}”. “Bundle-ID” is the bundle identifier string of the provider. “Designated-Requirement” is the designated requirement string from the code signature of the provider. For example, “com.example.app {anchor apple generic}”.

