# AppToAppLayerVPNMapping.AppLayerVPNMappingItem

A dictionary defining a per-app VPN relationship.

**Platforms:** macOS 10.9, Device Assignment Services , VPP License Management 

## Properties

### DesignatedRequirement

- **Type:** `string`
- **Required:** Yes

The code signature designated requirement of the app using the per-app VPN.

### Identifier

- **Type:** `string`
- **Required:** Yes

The bundle identifier of the app using the per-app VPN.

### MatchTools

- **Type:** `[AppToAppLayerVPNMapping.AppLayerVPNMappingItem.MatchToolsItem]`
- **Required:** No

An array of dictionaries. Each dictionary specifies a per-app VPN rule. Use this property to restrict this per-app VPN rule to only match the app’s spawned ** network traffic.

For example, to match network traffic that the `curl` command generates when run from the Terminal.app, create an app mapping payload for Terminal.app and set the payload’s `MatchTools` key to an array that contains a dictionary that matches the `curl` command-line tool.

If you don’t specify the `MatchTools` key, this per-app VPN rule matches all network traffic that the matching app and its spawned helper tools generate.

### Path

- **Type:** `string`
- **Required:** No

The file-system path of the executable using the per-app VPN.

### SigningIdentifier

- **Type:** `string`
- **Required:** Yes

The code signature signing identifier of the app using the per-app VPN.

### VPNUUID

- **Type:** `string`
- **Required:** Yes

The identifier of the per-app VPN payload, which defines the per-app VPN that the app uses. See the `VPNUUID` key of the [AppLayerVPN](/documentation/devicemanagement/applayervpn) payload.

## Topics

### Objects

- [AppToAppLayerVPNMapping.AppLayerVPNMappingItem.MatchToolsItem](/documentation/devicemanagement/apptoapplayervpnmapping/applayervpnmappingitem/matchtoolsitem) - Specifies a per-app VPN rule to match network traffic that the app’s spawned command-line tool generates.

