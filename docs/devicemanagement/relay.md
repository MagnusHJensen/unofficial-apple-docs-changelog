# Relay

The payload that configures relay settings.

**Platforms:** iOS 17.0, iPadOS 17.0, Mac Catalyst 17.0, macOS 14.0, tvOS 17.0, visionOS 1.0

## Properties

### AllowDNSFailover

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device allows the relay to failover to the default system DNS resolver.

Available: iOS 26+ | iPadOS 26+ | macOS 26+ | tvOS 26+ | visionOS 26+

### ExcludedDomains

- **Type:** `[string]`
- **Required:** No

A list of domain strings to exclude from routing through the servers in `Relays`. Any connection that matches a domain in the list exactly or is a subdomain of the listed domain won’t use the relay server.

### ExcludedFQDNs

- **Type:** `[string]`
- **Required:** No

A list of Fully Qualified Domain Names (FQDNs) to exclude from routing through the servers contained in `Relays`. Any connection that matches an FQDN in the list exactly won’t use the relay server. When `MatchDomains` is also present, any FQDN listed in the list should be a subdomain of at least one `MatchDomain` value, otherwise it won’t have any effect.

Available: iOS 18.4+ | iPadOS 18.4+ | macOS 15.4+ | tvOS 18.4+ | visionOS 2.4+

### MatchDomains

- **Type:** `[string]`
- **Required:** No

A list of domain strings that the system uses to determine which connection to route through the servers in `Relays`.

Any connection that matches a domain in the list exactly or is a subdomain of the listed domain uses the relay servers, unless it matches a domain in `ExcludedDomains`.

If this list and `MatchFQDNs` are empty, the system routes traffic to all domains to the relay servers, except those that match an excluded domain or excluded FQDN.

### MatchFQDNs

- **Type:** `[string]`
- **Required:** No

A list of Fully Qualified Domain Names (FQDNs) to route through the servers contained in `Relays`. Any connection that matches an FQDN in the list exactly uses the relay servers. If this list and `MatchDomains` are empty, the system routes traffic to all domains to the relay servers, except those that match an excluded domain or excluded FQDN.

Available: iOS 18.4+ | iPadOS 18.4+ | macOS 15.4+ | tvOS 18.4+ | visionOS 2.4+

### Relays

- **Type:** `[Relay.Relay]`
- **Required:** Yes

An array of dictionaries that describe one or more relay servers that the system can chain together.

### RelayUUID

- **Type:** `string`
- **Required:** No

A globally unique identifier for this relay configuration. The system uses this UUID to route managed apps through the servers in `Relays`. This key is required for user enrollment.

Available: iOS 17+ | iPadOS 17+ | tvOS 17+ | visionOS 1+

### UIToggleEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, the device allows the user to disable this network relay configuration.

Available: iOS 26+ | iPadOS 26+ | macOS 26+ | tvOS 26+ | visionOS 26+

## Discussion

Specify `com.apple.relay.managed` as the payload type.

### Profile availability

## Topics

### Objects

- [Relay.Relay](/documentation/devicemanagement/relay/relay) - A dictionary that describes a relay server.

