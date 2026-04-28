# AppToAppLayerVPNMapping.AppLayerVPNMappingItem.MatchToolsItem

Specifies a per-app VPN rule to match network traffic that the app’s spawned command-line tool generates.

**Platforms:** macOS 10.15.4, Device Assignment Services , VPP License Management 

## Properties

### DesignatedRequirement

- **Type:** `string`
- **Required:** Yes

The code signature designated requirement of the command-line tool using the per-app VPN.

### Path

- **Type:** `string`
- **Required:** No

The file-system path of the command-line tool using the per-app VPN.

### SigningIdentifier

- **Type:** `string`
- **Required:** Yes

The code signature signing identifier of the command-line tool using the per-app VPN.

