# VPN.DNS

The dictionary to configure DNS settings for the VPN.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, tvOS 17.0, visionOS 1.0, watchOS 10.0

## Properties

### DNSProtocol

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Cleartext`, `HTTPS`, `TLS`

The transport protocol to communicate with the DNS server.

Available: iOS 14+ | iPadOS 14+ | macOS 11+ | tvOS 17+ | visionOS 1+ | watchOS 10+

### DomainName

- **Type:** `string`
- **Required:** No

The primary domain of the tunnel.

Available: iOS 10+ | iPadOS 10+ | macOS 10.12+ | tvOS 17+ | visionOS 1+ | watchOS 10+

### PayloadCertificateUUID

- **Type:** `string`
- **Required:** No

That UUID that points to an identity certificate payload. The system uses this identity to authenticate the user to the DNS resolver.

Available: iOS 16+ | iPadOS 16+ | macOS 13+ | tvOS 17+ | visionOS 1+ | watchOS 10+

### SearchDomains

- **Type:** `[string]`
- **Required:** No

The list of domain strings used to fully qualify single-label host names.

Available: iOS 10+ | iPadOS 10+ | macOS 10.12+ | tvOS 17+ | visionOS 1+ | watchOS 10+

### ServerAddresses

- **Type:** `[string]`
- **Required:** Yes

The array of DNS server IP address strings. These IP addresses can be a mixture of IPv4 and IPv6 addresses.

Available: iOS 10+ | iPadOS 10+ | macOS 10.12+ | tvOS 17+ | visionOS 1+ | watchOS 10+

### ServerName

- **Type:** `string`
- **Required:** No

The hostname of a DNS-over-TLS server to validate the server certificate, as defined in RFC 7858. If `ServerAddresses` isn’t specified, the system uses the hostname to determine the server addresses. This key is required if the `DNSProtocol` is `TLS`.

Available: iOS 14+ | iPadOS 14+ | macOS 11+ | tvOS 17+ | visionOS 1+ | watchOS 10+

### ServerURL

- **Type:** `string`
- **Required:** No

The URI template of a DNS-over-HTTPS server, as defined in RFC 8484, which needs to use the `https://` scheme. The system uses the hostname or address in the URL to validate the server certificate. If `ServerAddresses` isn’t specified, the system uses the hostname or address in the URL to determine the server addresses. This key is required if the `DNSProtocol` is `HTTPS`.

Available: iOS 14+ | iPadOS 14+ | macOS 11+ | tvOS 17+ | visionOS 1+ | watchOS 10+

### SupplementalMatchDomains

- **Type:** `[string]`
- **Required:** No

The list of domain strings used to determine which DNS queries use the DNS resolver settings in `ServerAddresses`. The system uses this key to create a split DNS configuration where it resolves only hosts in certain domains using the tunnel’s DNS resolver. The system uses the default resolver for hosts that aren’t in one of the domains in this list.

If `SupplementalMatchDomains` contains the empty string it becomes the default domain.

Split-tunnel configurations can direct all DNS queries to the VPN DNS servers before the primary DNS servers. If the VPN tunnel becomes the network’s default route, the servers listed in `ServerAddresses` become the default resolver and the system ignores the `SupplementalMatchDomains` list.

Available: iOS 10+ | iPadOS 10+ | macOS 10.12+ | tvOS 17+ | visionOS 1+ | watchOS 10+

### SupplementalMatchDomainsNoSearch

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `0`, append the domains in the `SupplementalMatchDomains` list to the resolver’s list of search domains.

Available: iOS 10+ | iPadOS 10+ | macOS 10.12+ | tvOS 17+ | visionOS 1+ | watchOS 10+

