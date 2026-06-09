# NetworkVPNAlwaysOnTunnelConfigurationElement_IKEV2_ProviderObject

Specifies details about the provider.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### ComposedIdentifier

- **Type:** `string`
- **Required:** No

In iOS, tvOS, and visionOS, the identifier is a bundle ID, for example, “com.example.app”.

In macOS, the identifier is a composed identifier. The format of the composed identifier is either “Bundle-ID”, “Bundle-ID (Team-ID)”, or “Bundle-ID {Designated-Requirement}”. “Bundle-ID” is the bundle identifier string of the provider. “Team-ID” is the team identifier from the provider’s code signature. “Designated-Requirement” is the designated requirement string from the code signature of the provider. For example, “com.example.app” for the bundle ID format, “com.example.app (ABCD1234)” for the team ID format, or “com.example.app {anchor apple generic}” for the designated requirement format.

### Type

- **Type:** `string`
- **Required:** No
- **Default:** `packet-tunnel`
- **Allowed Values:** `packet-tunnel`, `app-proxy`

The type of VPN service. If the value is `app-proxy`, the service tunnels traffic at the app level. If the value is `packet-tunnel`, the service tunnels traffic at the IP layer.

