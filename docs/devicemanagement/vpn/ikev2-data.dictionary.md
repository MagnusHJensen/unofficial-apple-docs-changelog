# VPN.IKEv2

The dictionary to use for an IKEv2 VPN type.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 17.0, visionOS 1.0, watchOS 10.0

## Properties

### AllowPostQuantumKeyExchangeFallback

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If set to `0`, the VPN doesn’t establish a connection if the server does not support or doesn’t allow post-quantum key exchanges. Thd device ignores this key if `PostQuantumKeyExchangeMethods` is not present in `IKESecurityAssociationParameters` or `ChildSecurityAssociationParameters`.

### AuthenticationMethod

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `None`, `SharedSecret`, `Certificate`

The type of authentication method for the VPN.

To enable EAP-only authentication, set this to `None` and `ExtendedAuthEnabled` to `1`. If this is `None` and the `ExtendedAuthEnabled` key isn’t set, the authentication configuration defaults to `SharedSecret`.

### AuthName

- **Type:** `string`
- **Required:** No

The user name to use for authentication.

### AuthPassword

- **Type:** `string`
- **Required:** No

The password to use for authentication.

### CertificateType

- **Type:** `string`
- **Required:** No
- **Default:** `RSA`
- **Allowed Values:** `RSA`, `ECDSA256`, `ECDSA384`, `ECDSA521`, `RSA-PSS`

The type of `PayloadCertificateUUID` to use for IKEv2 machine authentication. If this key is included, the system requires a value for `ServerCertificateIssuerCommonName`.

### ChildSecurityAssociationParameters

- **Type:** `VPN.IKEv2.ChildSecurityAssociationParameters`
- **Required:** No

The `ChildSecurityAssociationParameters` dictionaries.

### DeadPeerDetectionRate

- **Type:** `string`
- **Required:** No
- **Default:** `Medium`
- **Allowed Values:** `None`, `Low`, `Medium`, `High`

One of the following:

- `None`: No keepalive.
- `Low`: Send keepalive every 30 minutes.
- `Medium`: Send keepalive every 10 minutes.
- `High`: Send keepalive every 1 minute.

Not available in watchOS.

### DisableMOBIKE

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, the system disables MOBIKE.

### DisableRedirect

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, the system disables IKEv2 redirect. If not set, the system redirects an IKEv2 connection when it receives a redirect request from the server.

### DisconnectOnIdle

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, the VPN disconnects automatically after a period defined by `DisconnectOnIdleTimer`.

### DisconnectOnIdleTimer

- **Type:** `integer`
- **Required:** No

Only used if `DisconnectOnIdle` is `1`. The number of seconds before the VPN disconnects. On watchOS, maximum allowed value is 15 seconds

### EnableCertificateRevocationCheck

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, the system performs a certificate revocation check for IKEv2 connections. This is a best-effort revocation check and server response timeouts won’t cause it to fail.

### EnableFallback

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, the system enables a tunnel over cellular data to carry traffic that’s eligible for Wi-Fi Assist and also requires VPN.

Enabling fallback requires that the server support multiple tunnels for a single user.

This field is available in iOS 13 and later, and tvOS 17 and later. Not available in watchOS.

### EnablePFS

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`,  enables Perfect Forward Secrecy (PFS) for IKEv2 Connections.

### EnforceRoutes

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, all the VPN’s non-default routes take precedence over any locally-defined routes. If `IncludeAllNetworks` is `1`, the system ignores `EnforceRoutes`.

### EnforceStrictAlgorithmSelection

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If set to `1`, the device doesn’t allow DES, 3DES, and Diffie-Hellman groups less than 14. Also the device requires the encryption algorithm specified for the IKE SA to be at least as cryptographically strong as the algorithm specified for the child SA. The device rejects this profile payload if these requirements are not met.

### ExcludeAPNs

- **Type:** `integer`
- **Required:** No
- **Default:** `1`
- **Allowed Values:** `0`, `1`

If `1` and `IncludeAllNetworks` is `1`, the system excludes network traffic for the Apple Push Notification service (APNs) from the tunnel.

### ExcludeCellularServices

- **Type:** `integer`
- **Required:** No
- **Default:** `1`
- **Allowed Values:** `0`, `1`

If `1` and `IncludeAllNetworks` is `1`, the system excludes internet-routable network traffic for cellular services (VoLTE, Wi-Fi Calling, IMS, MMS, Visual Voicemail, etc.) from the tunnel. Note that some cellular carriers route cellular services traffic directly to the carrier network, bypassing the internet. Such cellular services traffic is always excluded from the tunnel.

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

If `1` and either `IncludeAllNetworks` or `EnforceRoutes` are `1`, then the system routes local network traffic outside of the VPN. The default for this value is `0` on macOS and `1` on iOS.

### ExtendedAuthEnabled

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, enables EAP-only authentication.

### IKESecurityAssociationParameters

- **Type:** `VPN.IKEv2.IKESecurityAssociationParameters`
- **Required:** No

These parameters apply to Child Security Association unless `ChildSecurityAssociationParameters` is specified.

### IncludeAllNetworks

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, then the system routes all network traffic through the VPN, with some controllable exclusions, such as `ExcludeLocalNetworks`, `ExcludeCellularServices`, and `ExcludeAPNs` properties. The system always excludes the following traffic from the tunnel:

- Traffic necessary for connecting and maintaining the device’s network connection, such as DHCP.
- Traffic necessary for connecting to captive networks.
- Certain cellular services traffic that’s not routable over the internet and is instead directly routed to the cellular network. See the `ExcludeCellularServices` field for more information.
- Network communication with a companion device such as a watchOS device.

### LocalIdentifier

- **Type:** `string`
- **Required:** Yes

Identifier of the IKEv2 client.

### MTU

- **Type:** `integer`
- **Required:** No
- **Default:** `1280`

The Maximum Transmission Unit (MTU) specifies the maximum size in bytes of each packet that the system sends over the IKEv2 VPN interface. Available in iOS 14 and later, and macOS 11 and later.

### NATKeepAliveInterval

- **Type:** `integer`
- **Required:** No
- **Default:** `20`

The NAT Keepalive interval for Always On VPN IKEv2 connections. This value controls the interval that the device sends keepalive offload packets. The minimum value is 20 seconds. If no key is specified, the default is 20 seconds over Wi-Fi and 110 seconds over a cellular interface.

### NATKeepAliveOffloadEnable

- **Type:** `integer`
- **Required:** No
- **Default:** `1`
- **Allowed Values:** `0`, `1`

If `1`, enables NAT keepalive offload for Always On VPN IKEv2 connections. The device sends keepalive packets to maintain NAT mappings for IKEv2 connections that have a NAT on the path. It sends keepalive packets at regular intervals when the device is awake. If `NATKeepAliveOffloadEnable` is `1`, the system offloads keepalive packets to hardware while the device is asleep.

NAT keepalive offload has an impact on the battery life due to the extra workload during sleep. The default interval for the keepalive offload packets is 20 seconds over Wi-Fi and 110 seconds over Cellular interface. The default NAT keepalive works well on networks with small NAT mapping timeouts but imposes a potential battery impact. If a network has larger NAT mapping timeouts, larger keepalive intervals may be safely used to minimize battery impact. Modify the keepalive interval through the `NATKeepAliveInterval` key.

### OnDemandEnabled

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, enables VPN up on demand.

### OnDemandRules

- **Type:** `[VPN.VPN.OnDemandRulesElement]`
- **Required:** No

A list of rules that determine when and how to use an OnDemand VPN.

### OnDemandUserOverrideDisabled

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, the system disables the Connect On Demand toggle in Settings for this configuration.

### Password

- **Type:** `string`
- **Required:** No

The password to use for the account credentials. Only used if `AuthenticationMethod` is `Password`.

### PayloadCertificateUUID

- **Type:** `string`
- **Required:** No

The UUID of the certificate payload within the same profile to use as the account credential. If the value of `AuthenticationMethod` is `Certificate`, the system sends this certificate out for IKEv2 machine authentication. If extended authentication (EAP) is used, the system sends this certificate out for EAP-TLS authentication.

### PPK

- **Type:** `data`
- **Required:** No

The Post-quantum Pre-shared key (PPK) the device uses for this VPN. This key is is used with VPN servers that support RFC 8784. If this key is present `PPKIdentifier` must also be present.

### PPKIdentifier

- **Type:** `string`
- **Required:** No

The identifier for the Post-quantum Pre-shared key (PPK) the device uses for this VPN. This key is is used with VPN servers that support RFC 8784. If this key is present `PPK` must also be present.

### PPKMandatory

- **Type:** `integer`
- **Required:** No
- **Default:** `1`
- **Allowed Values:** `0`, `1`

If set to `1`, the VPN doesn’t establish a connection if the server doesn’t support RFC 8784 or doesn’t accept the PPK identifier specified in `PPKIdentifier`. The device ignores this key if `PPK` and `PPKIdentifier` are not present.

### ProviderBundleIdentifier

- **Type:** `string`
- **Required:** No

If the VPNSubType field contains the bundle identifier of an app that contains multiple VPN providers of the same type (app-proxy or packet-tunnel), then the system uses this field to choose which provider to use for this configuration. If the VPN provider is implemented as a System Extension, then this field is required.

### ProviderDesignatedRequirement

- **Type:** `string`
- **Required:** No

If the VPN provider is implemented as a System Extension, then this field is required. Available in macOS 10.15 and later, tvOS 17 and later, and watchOS 10 and later.

### ProviderType

- **Type:** `string`
- **Required:** No
- **Default:** `packet-tunnel`
- **Allowed Values:** `packet-tunnel`, `app-proxy`

If the value of this key is `app-proxy`, the VPN service tunnels traffic at the application layer. If the value of this key is `packet-tunnel`, the VPN service tunnels traffic at the IP layer.

### RemoteAddress

- **Type:** `string`
- **Required:** Yes

The IP address or host name of the VPN server.

### RemoteIdentifier

- **Type:** `string`
- **Required:** Yes

The remote identifier.

### ServerCertificateCommonName

- **Type:** `string`
- **Required:** No

The common name of the server certificate. The system uses this name to validate the certificate sent by the IKE server. If not set, the system uses the remote identifier to validate the certificate.

### ServerCertificateIssuerCommonName

- **Type:** `string`
- **Required:** No

Common Name of the server certificate issuer. If set, this field causes IKE to send a certificate request based on this certificate issuer to the server. This key is required if the `CertificateType` key is included and the `ExtendedAuthEnabled` key is `1`.

### SharedSecret

- **Type:** `string`
- **Required:** No

If `AuthenticationMethod` is `SharedSecret`, this value is used for IKE authentication.

### TLSMaximumVersion

- **Type:** `string`
- **Required:** No
- **Default:** `1.2`
- **Allowed Values:** `1.0`, `1.1`, `1.2`

The maximum TLS version to use with EAP-TLS authentication.

### TLSMinimumVersion

- **Type:** `string`
- **Required:** No
- **Default:** `1.0`
- **Allowed Values:** `1.0`, `1.1`, `1.2`

The minimum TLS version to use with EAP-TLS authentication.

### UseConfigurationAttributeInternalIPSubnet

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, negotiations should use IKEv2 Configuration Attribute `INTERNAL_IP4_SUBNET` and `INTERNAL_IP6_SUBNET`.

## Topics

### Objects

- [VPN.IKEv2.ChildSecurityAssociationParameters](/documentation/devicemanagement/vpn/ikev2-data.dictionary/childsecurityassociationparameters-data.dictionary) - The dictionary that contains child security association parameters.
- [VPN.IKEv2.IKESecurityAssociationParameters](/documentation/devicemanagement/vpn/ikev2-data.dictionary/ikesecurityassociationparameters-data.dictionary) - The dictionary that contains security association parameters.

