# NetworkVPNVPNPluginOnDemandObject

Specifies details about how the system controls on-demand VPN.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, tvOS 27.0, visionOS 27.0

## Properties

### DisableUserOverride

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device disables the Connect On Demand toggle in Settings for this configuration.

Available: iOS 27+ | iPadOS 27+ | tvOS 27+ | visionOS 27+

### Enabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables VPN On Demand.

### Rules

- **Type:** `[NetworkVPNVPNPluginRulesElementObject]`
- **Required:** No

An array of dictionaries defining On Demand Rules.

## Topics

### Objects

- [NetworkVPNVPNPluginRulesElementObject](/documentation/devicemanagement/networkvpnvpnpluginruleselementobject) - An array of dictionaries defining On Demand Rules.

