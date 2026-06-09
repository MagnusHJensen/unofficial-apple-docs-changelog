# NetworkVPNAlwaysOnTunnelConfigurationElement_IKEV2_OnDemandObject

Specifies details about how the system controls on-demand VPN.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### DisableUserOverride

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the Connect On Demand toggle in Settings is disabled for this configuration.

### Enabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables VPN On Demand.

### Rules

- **Type:** `[NetworkVPNAlwaysOnRulesElementObject]`
- **Required:** No

An array of dictionaries defining On Demand Rules.

## Topics

### Objects

- [NetworkVPNAlwaysOnRulesElementObject](/documentation/devicemanagement/networkvpnalwaysonruleselementobject) - An array of dictionaries defining On Demand Rules.

