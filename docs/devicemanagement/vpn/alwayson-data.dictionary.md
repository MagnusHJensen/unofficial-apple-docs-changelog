# VPN.AlwaysOn

The dictionary that contains IPSec settings.

**Platforms:** iOS 8.0, iPadOS 8.0, visionOS 1.0

## Properties

### AllowAllCaptiveNetworkPlugins

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, allows traffic from all captive networking apps outside the VPN tunnel to perform captive network handling.

### AllowCaptiveWebSheet

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, allows traffic from Captive Web Sheet outside the VPN tunnel.

### AllowedCaptiveNetworkPlugins

- **Type:** `[VPN.AlwaysOn.AllowedCaptiveNetworkPluginElement]`
- **Required:** No

The array of captive networking apps whose traffic is allowed outside the VPN tunnel, to perform captive network handling. Used only when `AllowAllCaptiveNetworkPlugins` is `false`.

### ApplicationExceptions

- **Type:** `[VPN.AlwaysOn.ApplicationExceptionElement]`
- **Required:** No

An array that contains an arbitrary number of apps whose connections occur outside the VPN.

### ServiceExceptions

- **Type:** `[VPN.AlwaysOn.ServiceExceptionElement]`
- **Required:** No

An array that contains an arbitrary number of service exceptions.

### TunnelConfigurations

- **Type:** `[VPN.AlwaysOn.TunnelConfigurationElement]`
- **Required:** Yes

An array that contains an arbitrary number of tunnel configurations.

### UIToggleEnabled

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, allows the user to disable the VPN configuration.

## Topics

### Objects

- [VPN.AlwaysOn.AllowedCaptiveNetworkPluginElement](/documentation/devicemanagement/vpn/alwayson-data.dictionary/allowedcaptivenetworkpluginelement) - The dictionary for captive network configurations.
- [VPN.AlwaysOn.ApplicationExceptionElement](/documentation/devicemanagement/vpn/alwayson-data.dictionary/applicationexceptionelement) - The dictionary that defines which applications can have traffic outside the VPN tunnel.
- [VPN.AlwaysOn.ServiceExceptionElement](/documentation/devicemanagement/vpn/alwayson-data.dictionary/serviceexceptionelement) - The dictionary that defines service exceptions.
- [VPN.AlwaysOn.TunnelConfigurationElement](/documentation/devicemanagement/vpn/alwayson-data.dictionary/tunnelconfigurationelement) - The dictionary used to configure VPN tunnels.

