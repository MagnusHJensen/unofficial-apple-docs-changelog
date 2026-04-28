# VPN.VPN

The dictionary that contains VPN, IPSec, and IKEv2 settings.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, tvOS 17.0, visionOS 1.0, Device Assignment Services , VPP License Management 

## Properties

### AuthenticationMethod

- **Type:** `string`
- **Required:** No
- **Default:** `Password`
- **Allowed Values:** `Password`, `Certificate`, `Password+Certificate`

The authentication method to use.

### AuthName

- **Type:** `string`
- **Required:** No

The VPN account username.

### AuthPassword

- **Type:** `string`
- **Required:** No

The VPN account password. Only use this if `AuthenticationMethod` is set to `Password`.

### DisconnectOnIdle

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, disconnects after an on-demand connection idles.

### DisconnectOnIdleTimer

- **Type:** `integer`
- **Required:** No

The length of time to wait, in seconds, before disconnecting an on-demand connection. In watchOS, the maximum allowed value is `15`.

### EnforceRoutes

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, all the VPN’s non-default routes take precedence over any locally defined routes.

If `IncludeAllNetworks` is `1`, the system ignores the value of `EnforceRoutes`.

Available in iOS 14.2 and later, and macOS 11 and later. Not available in watchOS.

### ExcludeAPNs

- **Type:** `integer`
- **Required:** No
- **Default:** `1`
- **Allowed Values:** `0`, `1`

If `1` and `IncludeAllNetworks` is `1`, then the system excludes the network traffic for the Apple Push Notification service (APNs) from the tunnel. Not available in watchOS.

### ExcludeCellularServices

- **Type:** `integer`
- **Required:** No
- **Default:** `1`
- **Allowed Values:** `0`, `1`

If `1` and `IncludeAllNetworks` is `1`, then the system excludes internet-routable network traffic for cellular services (VoLTE, Wi-Fi Calling, IMS, MMS, Visual Voicemail, etc.) from the tunnel. Note that some cellular carriers route cellular services traffic directly to the carrier network, bypassing the internet. Such cellular services traffic is always excluded from the tunnel. Not available in watchOS.

### ExcludeDeviceCommunication

- **Type:** `integer`
- **Required:** No
- **Default:** `1`
- **Allowed Values:** `0`, `1`

If set to `1` and `IncludeAllNetworks` is set to `1`, the device excludes network traffic used for communicating with devices connected via USB or Wi-Fi from the tunnel.

### ExcludeLocalNetworks

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `0`, `1`

If `1` and `IncludeAllNetworks` is `1`, routes all local network traffic outside the VPN. Not available in watchOS.

### IncludeAllNetworks

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1``, routes all traffic through the VPN, with some exclusions. Several of the exclusions can be controlled with the `ExcludeLocalNetworks`, `ExcludeCellularServices`, `ExcludeAPNs`and`ExcludeDeviceCommunication` properties. The following traffic is always excluded from the tunnel:

- Traffic necessary for connecting and maintaining the device’s network connection, such as DHCP.
- Traffic necessary for connecting to captive networks.
- Certain cellular services traffic that is not routable over the internet and is instead directly routed to the cellular network. See the ExcludeCellularServices property for more details.
- Network communication with a companion device such as a watchOS device.

Not available in watchOS.

### OnDemandEnabled

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, enables VPN On Demand.

### OnDemandMatchDomainsAlways

- **Type:** `[string]`
- **Required:** No

A list of domain names. The system treats associated domain names as though they’re associated with the `OnDemandMatchDomainsOnRetry` key. This behavior can be overridden by `OnDemandRules`.

In iOS 7 and later, this key is deprecated (but still supported) in favor of `EvaluateConnection` actions in the `OnDemandRules` dictionaries.

Not available in watchOS.

### OnDemandMatchDomainsNever

- **Type:** `[string]`
- **Required:** No

A list of domain names. If the host name ends with one of these domain names, the system doesn’t start the VPN automatically. The system uses this value to exclude a subdomain within an included domain.

In iOS 7 and later, this key is deprecated (but still supported) in favor of `EvaluateConnection` actions in the `OnDemandRules` dictionaries.

Not available in watchOS.

### OnDemandMatchDomainsOnRetry

- **Type:** `[string]`
- **Required:** No

A list of domain names. If the host name ends with one of these domain names and a DNS query for that domain name fails, the system starts the VPN automatically.

In iOS 7 and later, this key is deprecated (but still supported) in favor of `EvaluateConnection` actions in the `OnDemandRules` dictionaries.

Not available in watchOS.

### OnDemandRules

- **Type:** `[VPN.VPN.OnDemandRulesElement]`
- **Required:** No

An array of dictionaries defining On Demand Rules.

### OnDemandUserOverrideDisabled

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, the Connect On Demand toggle in Settings is disabled for this configuration. Available in iOS 14 and later. Not available in watchOS.

### PayloadCertificateUUID

- **Type:** `string`
- **Required:** No

The UUID of the certificate payload within the same profile to use for account credentials.

### ProviderBundleIdentifier

- **Type:** `string`
- **Required:** No

The bundle identifier for the VPN provider. Not available in watchOS.

### ProviderDesignatedRequirement

- **Type:** `string`
- **Required:** No

If the VPN provider is implemented as a system extension, this field is required. Not available in watchOS.

### ProviderType

- **Type:** `string`
- **Required:** No
- **Default:** `packet-tunnel`
- **Allowed Values:** `packet-tunnel`, `app-proxy`

The type of VPN service. If the value is `app-proxy`, the service tunnels traffic at the app level. If the value is `packet-tunnel`, the service tunnels traffic at the IP layer. Not available in watchOS.

### RemoteAddress

- **Type:** `string`
- **Required:** Yes

The IP address or hostname of the VPN server.

## Topics

### Objects

- [VPN.VPN.OnDemandRulesElement](/documentation/devicemanagement/vpn/vpn-data.dictionary/ondemandruleselement) - An On Demand rule

