# NetworkVPNIKEV2

The declaration to configure a VPN using the IKEv2 sub-type.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, tvOS 27.0, visionOS 27.0

## Properties

### Authentication

- **Type:** `NetworkVPNIKEV2AuthenticationObject`
- **Required:** Yes

Settings that control authentication.

### ChildSecurityAssociationParameters

- **Type:** `NetworkVPNIKEV2SecurityAssociationParametersObject`
- **Required:** No

The `ChildSecurityAssociationParameters` dictionaries.

### DisableMOBIKE

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables MOBIKE.

### DisableRedirect

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system disables IKEv2 redirect. If not set, the system redirects an IKEv2 connection when it receives a redirect request from the server.

### DNS

- **Type:** `NetworkVPNIKEV2DNSObject`
- **Required:** No

A dictionary to use for all VPN types.

### EnableCertificateRevocationCheck

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system performs a certificate revocation check for IKEv2 connections. This is a best-effort revocation check and server response timeouts won’t cause it to fail.

### EnableFallback

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables a tunnel over cellular data to carry traffic that’s eligible for Wi-Fi Assist and also requires VPN.

Enabling fallback requires that the server support multiple tunnels for a single user.

Available: iOS 27+ | iPadOS 27+ | tvOS 27+ | visionOS 27+

### EnablePFS

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`,  enables Perfect Forward Secrecy (PFS) for IKEv2 Connections.

### EnforceStrictAlgorithmSelection

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If set to `true`, the device doesn’t allow DES, 3DES, and Diffie-Hellman groups less than 14. Also the device requires the encryption algorithm specified in `IKESecurityAssociationParameters` to be at least as cryptographically strong as the algorithm specified in `ChildSecurityAssociationParameters`. The device rejects this configuration if these requirements aren’t met.

### HostName

- **Type:** `string`
- **Required:** Yes

The IP address or hostname of the VPN server.

### Idle

- **Type:** `NetworkVPNIKEV2IdleObject`
- **Required:** No

Specifies details about how the system handles idle VPN connections.

### IKESecurityAssociationParameters

- **Type:** `NetworkVPNIKEV2SecurityAssociationParametersObject`
- **Required:** No

These parameters apply to Child Security Association unless `ChildSecurityAssociationParameters` is specified.

### LocalIdentifier

- **Type:** `string`
- **Required:** Yes

Identifier of the IKEv2 client.

### MTU

- **Type:** `integer`
- **Required:** No
- **Default:** `1280`

The Maximum Transmission Unit (MTU) specifies the maximum size in bytes of each packet that the system sends over the IKEv2 VPN interface.

### NetworkRouting

- **Type:** `NetworkVPNIKEV2NetworkRoutingObject`
- **Required:** No

Specifies details about how the VPN routes different types of network traffic.

Available: iOS 27+ | iPadOS 27+ | macOS 27+ | visionOS 27+

### OnDemand

- **Type:** `NetworkVPNIKEV2OnDemandObject`
- **Required:** No

Specifies details about how the system controls on-demand VPN.

### PostQuantumKeyExchange

- **Type:** `NetworkVPNIKEV2PostQuantumKeyExchangeObject`
- **Required:** No

Post Quantum Key Exchange settings.

### Proxies

- **Type:** `NetworkVPNIKEV2ProxiesObject`
- **Required:** No

The dictionary to use to configure `Proxies` for use with `VPN`.

### RemoteIdentifier

- **Type:** `string`
- **Required:** Yes

The remote identifier.

### UseConfigurationAttributeInternalIPSubnet

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, negotiations should use IKEv2 Configuration Attribute `INTERNAL_IP4_SUBNET` and `INTERNAL_IP6_SUBNET`.

### VisibleName

- **Type:** `string`
- **Required:** Yes

The name of the VPN connection that the system displays on the device.

## Discussion

Specify `com.apple.configuration.network.vpn.ikev2` as the declaration type.

### Configuration availability

### Configuration examples

## Topics

### Objects

- [NetworkVPNIKEV2AuthenticationObject](/documentation/devicemanagement/networkvpnikev2authenticationobject) - Settings that control authentication.
- [NetworkVPNIKEV2DNSObject](/documentation/devicemanagement/networkvpnikev2dnsobject) - A dictionary to use for all VPN types.
- [NetworkVPNIKEV2IdleObject](/documentation/devicemanagement/networkvpnikev2idleobject) - Specifies details about how the system handles idle VPN connections.
- [NetworkVPNIKEV2NetworkRoutingObject](/documentation/devicemanagement/networkvpnikev2networkroutingobject) - Specifies details about how the VPN routes different types of network traffic.
- [NetworkVPNIKEV2OnDemandObject](/documentation/devicemanagement/networkvpnikev2ondemandobject) - Specifies details about how the system controls on-demand VPN.
- [NetworkVPNIKEV2PostQuantumKeyExchangeObject](/documentation/devicemanagement/networkvpnikev2postquantumkeyexchangeobject) - Post Quantum Key Exchange settings.
- [NetworkVPNIKEV2ProxiesObject](/documentation/devicemanagement/networkvpnikev2proxiesobject) - The dictionary to use to configure `Proxies` for use with `VPN`.
- [NetworkVPNIKEV2SecurityAssociationParametersObject](/documentation/devicemanagement/networkvpnikev2securityassociationparametersobject) - These parameters apply to Child Security Association unless `ChildSecurityAssociationParameters` is specified.

