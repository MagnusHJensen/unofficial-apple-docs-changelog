# NetworkRelay

The declaration to configure Network Relay settings.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### AllowDNSFailover

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device allows the relay to failover to the default system DNS resolver.

### ExcludedDomains

- **Type:** `[string]`
- **Required:** No

A list of domain strings to exclude from routing through the servers in `Relays`. Any connection that matches a domain in the list exactly or is a subdomain of the listed domain won’t use the relay server.

### ExcludedFQDNs

- **Type:** `[string]`
- **Required:** No

A list of Fully Qualified Domain Names (FQDNs) to exclude from routing through the servers contained in `Relays`. Any connection that matches an FQDN in the list exactly won’t use the relay server. When `MatchDomains` is also present, any FQDN listed in the list should be a subdomain of at least one `MatchDomain` value, otherwise it will not have any effect.

### MatchDomains

- **Type:** `[string]`
- **Required:** No

A list of domain strings that the system uses to determine which connection to route through the servers in `Relays`.

Any connection that matches a domain in the list exactly or is a subdomain of the listed domain uses the relay servers, unless it matches a domain in `ExcludedDomains`.

If this list and `MatchFQDNs` are empty, the system routes traffic to all domains to the relay servers, except those that match an excluded domain or excluded FQDN.

### MatchFQDNs

- **Type:** `[string]`
- **Required:** No

A list of Fully Qualified Domain Names (FQDNs) to be routed through the servers contained in `Relays`. Any connection that matches an FQDN in the list exactly uses the relay servers. If this list and `MatchDomains` are empty, the system routes traffic to all domains to the relay servers, except those that match an excluded domain or excluded FQDN.

### Relays

- **Type:** `[NetworkRelayRelayObject]`
- **Required:** Yes

An array of dictionaries that describe one or more relay servers that the system can chain together.

### RelayUUID

- **Type:** `string`
- **Required:** No

A globally unique identifier for this relay configuration. The system uses this UUID to route managed apps through the servers in `Relays`. This key is required for user enrollment.

Available: iOS 27+ | iPadOS 27+ | visionOS 27+

### UIToggleEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the device allows the user to disable this network relay configuration.

### VisibleName

- **Type:** `string`
- **Required:** Yes

The name of the network relays that the system displays on the device.

## Discussion

Specify `com.apple.configuration.network.relay` as the declaration type.

### Configuration availability

### Configuration examples

## Topics

### Objects

- [NetworkRelayRelayObject](/documentation/devicemanagement/networkrelayrelayobject) - An array of dictionaries that describe one or more relay servers that the system can chain together.

