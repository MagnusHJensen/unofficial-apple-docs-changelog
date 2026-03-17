# StatusManagementClientCapabilitiesCapabilitiesObject

A collection of the device’s supported features, payloads, and versions.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0

## Properties

### supported-features

- **Type:** `StatusManagementClientCapabilitiesCapabilities_SupportedFeaturesObject`
- **Required:** Yes

A set of optional protocol features that the client supports. Each object’s key represents a feature, and the property value represents the feature’s associated parameters.

### supported-payloads

- **Type:** `StatusManagementClientCapabilitiesCapabilities_SupportedPayloadsObject`
- **Required:** Yes

A set of declaration and status items that the client supports.

### supported-versions

- **Type:** `[string]`
- **Required:** Yes

A list of protocol versions that the client supports.

## Topics

### Objects

- [StatusManagementClientCapabilitiesCapabilities_SupportedFeaturesObject](/documentation/devicemanagement/statusmanagementclientcapabilitiescapabilities_supportedfeaturesobject) - A set of optional protocol features that the client supports.
- [StatusManagementClientCapabilitiesCapabilities_SupportedPayloadsObject](/documentation/devicemanagement/statusmanagementclientcapabilitiescapabilities_supportedpayloadsobject) - The set of declaration and status items that the client supports.

