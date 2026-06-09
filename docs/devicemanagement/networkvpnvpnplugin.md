# NetworkVPNVPNPlugin

The declaration to configure a VPN using the VPN plugin sub-type.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), tvOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### Authentication

- **Type:** `NetworkVPNVPNPluginAuthenticationObject`
- **Required:** Yes

Settings that control authentication.

### DNS

- **Type:** `NetworkVPNVPNPluginDNSObject`
- **Required:** No

A dictionary to use for all VPN types.

### HostName

- **Type:** `string`
- **Required:** Yes

The IP address or hostname of the VPN server.

### Idle

- **Type:** `NetworkVPNVPNPluginIdleObject`
- **Required:** No

Specifies details about how the system handles idle VPN connections.

### NetworkRouting

- **Type:** `NetworkVPNVPNPluginNetworkRoutingObject`
- **Required:** No

Specifies details about how the VPN routes different types of network traffic.

Available: iOS 27+ | iPadOS 27+ | macOS 27+ | visionOS 27+

### OnDemand

- **Type:** `NetworkVPNVPNPluginOnDemandObject`
- **Required:** No

Specifies details about how the system controls on-demand VPN.

### Provider

- **Type:** `NetworkVPNVPNPluginProviderObject`
- **Required:** No

Specifies details about the provider.

### Proxies

- **Type:** `NetworkVPNVPNPluginProxiesObject`
- **Required:** No

The dictionary to use to configure `Proxies` for use with `VPN`.

### SubType

- **Type:** `string`
- **Required:** Yes

An identifier for a vendor-specified configuration dictionary.

If the configuration targets a VPN solution that uses a VPN plugin, then this field contains the bundle identifier of the plugin. Here are some examples:

- Cisco AnyConnect: `com.cisco.anyconnect.applevpn.plugin`
- Juniper SSL: `net.juniper.sslvpn`
- F5 SSL: `com.f5.F5-Edge-Client.vpnplugin`
- SonicWALL Mobile Connect: `com.sonicwall.SonicWALL-SSLVPN.vpnplugin`
- ``Aruba VIA: `com.arubanetworks.aruba-via.vpnplugin`

If the configuration targets a VPN solution that uses a network extension provider, then this field contains the bundle identifier of the app that contains the provider. Contact the VPN solution vendor for the value of the identifier.

### VendorConfig

- **Type:** `NetworkVPNVPNPluginVendorConfigObject`
- **Required:** No

The vendor-specific configuration dictionary, which the system reads only when `SubType` has a value.

### VisibleName

- **Type:** `string`
- **Required:** Yes

The name of the VPN connection that the system displays on the device.

## Discussion

Specify `com.apple.configuration.network.vpn.vpn-plugin` as the declaration type.

### Configuration availability

### Configuration examples

## Topics

### Objects

- [NetworkVPNVPNPluginAuthenticationObject](/documentation/devicemanagement/networkvpnvpnpluginauthenticationobject) - Settings that control authentication.
- [NetworkVPNVPNPluginDNSObject](/documentation/devicemanagement/networkvpnvpnplugindnsobject) - A dictionary to use for all VPN types.
- [NetworkVPNVPNPluginIdleObject](/documentation/devicemanagement/networkvpnvpnpluginidleobject) - Specifies details about how the system handles idle VPN connections.
- [NetworkVPNVPNPluginNetworkRoutingObject](/documentation/devicemanagement/networkvpnvpnpluginnetworkroutingobject) - Specifies details about how the VPN routes different types of network traffic.
- [NetworkVPNVPNPluginOnDemandObject](/documentation/devicemanagement/networkvpnvpnpluginondemandobject) - Specifies details about how the system controls on-demand VPN.
- [NetworkVPNVPNPluginProviderObject](/documentation/devicemanagement/networkvpnvpnpluginproviderobject) - Specifies details about the provider.
- [NetworkVPNVPNPluginProxiesObject](/documentation/devicemanagement/networkvpnvpnpluginproxiesobject) - The dictionary to use to configure `Proxies` for use with `VPN`.
- [NetworkVPNVPNPluginVendorConfigObject](/documentation/devicemanagement/networkvpnvpnpluginvendorconfigobject) - The vendor-specific configuration dictionary, which the system reads only when `SubType` has a value.

