# NetworkVPNAlwaysOnTunnelConfigurationElementObject

An array that contains an arbitrary number of tunnel configurations.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### IKEV2

- **Type:** `NetworkVPNAlwaysOnTunnelConfigurationElement_IKEV2Object`
- **Required:** No

The IKEv2 configuration for this tunnel.

### Interfaces

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `Cellular`, `WiFi`

The interfaces to apply this configuration to.

### ProtocolType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `IKEv2`

The type of connection, which needs to be `IKEv2`.

## Topics

### Objects

- [NetworkVPNAlwaysOnTunnelConfigurationElement_IKEV2Object](/documentation/devicemanagement/networkvpnalwaysontunnelconfigurationelement_ikev2object) - The IKEv2 configuration for this tunnel.

