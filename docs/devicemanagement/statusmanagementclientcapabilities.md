# StatusManagementClientCapabilities

A status report of the client’s protocol capabilities.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0

## Properties

### management.client-capabilities

- **Type:** `StatusManagementClientCapabilitiesCapabilitiesObject`
- **Required:** Yes

An object that contains the client’s protocol capabilities. These typically only change when the device upgrades its software. An implicit status subscription for this status item is always present, so the client always reports changes to the server.

## Discussion

### Status item availability

## Topics

### Objects

- [StatusManagementClientCapabilitiesCapabilitiesObject](/documentation/devicemanagement/statusmanagementclientcapabilitiescapabilitiesobject) - A collection of the device’s supported features, payloads, and versions.

