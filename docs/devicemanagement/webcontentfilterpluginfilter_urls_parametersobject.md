# WebContentFilterPluginFilter_URLs_ParametersObject

A dictionary containing URL filter parameters. Required when `Enabled` is `true`.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta)

## Properties

### FailClosed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system blocks URLs if the filter is enabled, but it fails to make any filtering decision; for example, if there’s a communication failure with the PIR server. If `false`, the system allows URLs if the filter is enabled, but it fails to make any filtering decision.

### PIR

- **Type:** `WebContentFilterPluginFilter_URLs_Parameters_PIRObject`
- **Required:** Yes

A dictionary containing Private Information Retrieval server settings.

### PrefilterFetchFrequency

- **Type:** `integer`
- **Required:** No
- **Default:** `86400`

The time interval in seconds that the system uses to periodically run the `NEURLFilterControlProvider` app extension. The default value is 86400 seconds (1 day). The minimum allowed value is 2700 seconds (45 minutes). The system allows `NEURLFilterControlProvider` implementations to download prefilter Bloom filter data onto the device periodically at the specified interval. Implementations need to allow for a slight difference between the scheduled time and the actual runtime of the task, due to the scheduling mechanism on the system.

### ProviderComposedIdentifier

- **Type:** `string`
- **Required:** Yes

The URL filter control provider identifier. This string identifies the filter data provider when the filter starts running. Required when Enabled is true.

In iOS, the identifier is a bundle ID, for example, “com.example.app”.

In macOS, the identifier is a composed identifier. The format of the composed identifier is “Bundle-ID {Designated-Requirement}”. “Bundle-ID” is the bundle identifier string of the provider. “Designated-Requirement” is the designated requirement string from the code signature of the provider. For example, “com.example.app {anchor apple generic}”.

## Topics

### Objects

- [WebContentFilterPluginFilter_URLs_Parameters_PIRObject](/documentation/devicemanagement/webcontentfilterpluginfilter_urls_parameters_pirobject) - A dictionary containing Private Information Retrieval server settings.

