# NetworkVPNIKEV2OnDemandObject

Specifies details about how the system controls on-demand VPN.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), tvOS 27.0 (Beta), visionOS 27.0 (Beta)

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

- **Type:** `[NetworkVPNIKEV2RulesElementObject]`
- **Required:** No

An array of dictionaries defining On Demand Rules.

## Topics

### Objects

- [NetworkVPNIKEV2RulesElementObject](/documentation/devicemanagement/networkvpnikev2ruleselementobject) - An array of dictionaries defining On Demand Rules.

