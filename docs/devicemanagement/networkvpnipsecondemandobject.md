# NetworkVPNIPSecOnDemandObject

Specifies details about how the system controls on-demand VPN.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### Enabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables VPN On Demand.

### Rules

- **Type:** `[NetworkVPNIPSecRulesElementObject]`
- **Required:** No

An array of dictionaries defining On Demand Rules.

## Topics

### Objects

- [NetworkVPNIPSecRulesElementObject](/documentation/devicemanagement/networkvpnipsecruleselementobject) - An array of dictionaries defining On Demand Rules.

