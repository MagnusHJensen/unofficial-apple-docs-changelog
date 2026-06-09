# NetworkVPNIPSec

The declaration to configure a VPN using the IPSec sub-type.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### Authentication

- **Type:** `NetworkVPNIPSecAuthenticationObject`
- **Required:** Yes

Settings that control authentication.

### DNS

- **Type:** `NetworkVPNIPSecDNSObject`
- **Required:** No

A dictionary to use for all VPN types.

### HostName

- **Type:** `string`
- **Required:** Yes

The IP address or hostname of the VPN server.

### Idle

- **Type:** `NetworkVPNIPSecIdleObject`
- **Required:** No

Specifies details about how the system handles idle VPN connections.

### OnDemand

- **Type:** `NetworkVPNIPSecOnDemandObject`
- **Required:** No

Specifies details about how the system controls on-demand VPN.

### OverridePrimary

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system sends all network traffic over VPN.

### Proxies

- **Type:** `NetworkVPNIPSecProxiesObject`
- **Required:** No

The dictionary to use to configure `Proxies` for use with `VPN`.

### VisibleName

- **Type:** `string`
- **Required:** Yes

The name of the VPN connection that the system displays on the device.

## Discussion

Specify `com.apple.configuration.network.vpn.ipsec` as the declaration type.

### Configuration availability

### Configuration examples

## Topics

### Objects

- [NetworkVPNIPSecAuthenticationObject](/documentation/devicemanagement/networkvpnipsecauthenticationobject) - Settings that control authentication.
- [NetworkVPNIPSecDNSObject](/documentation/devicemanagement/networkvpnipsecdnsobject) - A dictionary to use for all VPN types.
- [NetworkVPNIPSecIdleObject](/documentation/devicemanagement/networkvpnipsecidleobject) - Specifies details about how the system handles idle VPN connections.
- [NetworkVPNIPSecOnDemandObject](/documentation/devicemanagement/networkvpnipsecondemandobject) - Specifies details about how the system controls on-demand VPN.
- [NetworkVPNIPSecProxiesObject](/documentation/devicemanagement/networkvpnipsecproxiesobject) - The dictionary to use to configure `Proxies` for use with `VPN`.

