# NetworkVPNAlwaysOn

The declaration to configure a VPN using the Always On sub-type.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### AllowAllCaptiveNetworkPlugins

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, allows traffic from all captive networking apps outside the VPN tunnel to perform captive network handling.

### AllowCaptiveWebSheet

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, allows traffic from Captive Web Sheet outside the VPN tunnel.

### AllowedCaptiveNetworkPlugins

- **Type:** `[NetworkVPNAlwaysOnAllowedCaptiveNetworkPluginElementObject]`
- **Required:** No

The array of captive networking apps whose traffic is allowed outside the VPN tunnel, to perform captive network handling. Used only when `AllowAllCaptiveNetworkPlugins` is `false`.

### ApplicationExceptions

- **Type:** `[NetworkVPNAlwaysOnApplicationExceptionElementObject]`
- **Required:** No

An array that contains an arbitrary number of apps whose connections occur outside the VPN.

### DNS

- **Type:** `NetworkVPNAlwaysOnDNSObject`
- **Required:** No

A dictionary to use for all VPN types.

### Proxies

- **Type:** `NetworkVPNAlwaysOnProxiesObject`
- **Required:** No

The dictionary to use to configure `Proxies` for use with `VPN`.

### ServiceExceptions

- **Type:** `[NetworkVPNAlwaysOnServiceExceptionElementObject]`
- **Required:** No

An array that contains an arbitrary number of service exceptions.

### TunnelConfigurations

- **Type:** `[NetworkVPNAlwaysOnTunnelConfigurationElementObject]`
- **Required:** Yes

An array that contains an arbitrary number of tunnel configurations.

### UIToggleEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, allows the user to disable the VPN configuration.

### VisibleName

- **Type:** `string`
- **Required:** Yes

The name of the VPN connection that the system displays on the device.

## Discussion

Specify `com.apple.configuration.network.vpn.always-on` as the declaration type.

### Configuration availability

### Configuration example

This configuration sets up an always-on IKEv2 VPN for both Cellular and Wi-Fi interfaces using certificate authentication.

```json
{
    "Type": "com.apple.configuration.network.vpn.always-on",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "VisibleName": "Always-On VPN",
        "UIToggleEnabled": false,
        "TunnelConfigurations": [
            {
                "ProtocolType": "IKEv2",
                "Interfaces": [
                    "Cellular",
                    "WiFi"
                ],
                "IKEV2": {
                    "HostName": "vpn.example.com",
                    "LocalIdentifier": "device@example.com",
                    "RemoteIdentifier": "vpn.example.com",
                    "Authentication": {
                        "Method": "Certificate",
                        "IdentityAssetReference": "CB3E6C7F-2318-437B-8A9E-D50C69376DE4"
                    }
                }
            }
        ]
    }
}
```

## Topics

### Objects

- [NetworkVPNAlwaysOnAllowedCaptiveNetworkPluginElementObject](/documentation/devicemanagement/networkvpnalwaysonallowedcaptivenetworkpluginelementobject) - The array of captive networking apps whose traffic is allowed outside the VPN tunnel, to perform captive network handling. Used only when `AllowAllCaptiveNetworkPlugins` is `false`.
- [NetworkVPNAlwaysOnApplicationExceptionElementObject](/documentation/devicemanagement/networkvpnalwaysonapplicationexceptionelementobject) - An array that contains an arbitrary number of apps whose connections occur outside the VPN.
- [NetworkVPNAlwaysOnDNSObject](/documentation/devicemanagement/networkvpnalwaysondnsobject) - A dictionary to use for all VPN types.
- [NetworkVPNAlwaysOnProxiesObject](/documentation/devicemanagement/networkvpnalwaysonproxiesobject) - The dictionary to use to configure `Proxies` for use with `VPN`.
- [NetworkVPNAlwaysOnServiceExceptionElementObject](/documentation/devicemanagement/networkvpnalwaysonserviceexceptionelementobject) - An array that contains an arbitrary number of service exceptions.
- [NetworkVPNAlwaysOnTunnelConfigurationElementObject](/documentation/devicemanagement/networkvpnalwaysontunnelconfigurationelementobject) - An array that contains an arbitrary number of tunnel configurations.

