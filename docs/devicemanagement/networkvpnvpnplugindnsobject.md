# NetworkVPNVPNPluginDNSObject

A dictionary to use for all VPN types.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, tvOS 27.0, visionOS 27.0

## Properties

### DomainName

- **Type:** `string`
- **Required:** No

The primary domain of the tunnel.

### IdentityAssetReference

- **Type:** `string`
- **Required:** No

The identifier of a credential asset declaration that contains the identity that the system uses to authenticate the user to the DNS resolver.

### SearchDomains

- **Type:** `[string]`
- **Required:** No

The list of domain strings used to fully qualify single-label host names.

### ServerAddresses

- **Type:** `[string]`
- **Required:** Yes

The array of DNS server IP address strings. These IP addresses can be a mixture of IPv4 and IPv6 addresses.

### SupplementalMatchDomains

- **Type:** `[string]`
- **Required:** No

The list of domain strings used to determine which DNS queries use the DNS resolver settings in `ServerAddresses`. The system uses this key to create a split DNS configuration where it resolves only hosts in certain domains using the tunnel’s DNS resolver. The system uses the default resolver for hosts that aren’t in one of the domains in this list.

If `SupplementalMatchDomains` contains the empty string it becomes the default domain.

Split-tunnel configurations can direct all DNS queries to the VPN DNS servers before the primary DNS servers. If the VPN tunnel becomes the network’s default route, the servers listed in `ServerAddresses` become the default resolver and the system ignores the `SupplementalMatchDomains` list.

### SupplementalMatchDomainsNoSearch

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, don’t append the domains in the `SupplementalMatchDomains` list to the resolver’s list of search domains.

