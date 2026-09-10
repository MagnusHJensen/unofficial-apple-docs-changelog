# NetworkVPNIPSecOnDemandObject

Specifies details about how the system controls on-demand VPN.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, visionOS 27.0

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

