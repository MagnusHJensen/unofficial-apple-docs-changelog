# NetworkVPNVPNPluginProxiesObject

The dictionary to use to configure `Proxies` for use with `VPN`.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), tvOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### AutoConfigEnable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables automatic proxy configuration.

### AutoConfigURLString

- **Type:** `string`
- **Required:** No

The URL to the location of the proxy auto-configuration file. Used only when `ProxyAutoConfigEnable` is `true`.

### AutoDiscoveryEnable

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, enables proxy auto discovery.

### Protocol

- **Type:** `NetworkVPNVPNPluginProxies_ProtocolObject`
- **Required:** No

The dictionary to use to configure HTTP servers  for `Proxies` for use with `VPN`.

### SupplementalMatchDomains

- **Type:** `[string]`
- **Required:** No

An array of domains that defines which hosts use proxy settings for hosts.

## Topics

### Objects

- [NetworkVPNVPNPluginProxies_ProtocolObject](/documentation/devicemanagement/networkvpnvpnpluginproxies_protocolobject) - The dictionary to use to configure HTTP servers  for `Proxies` for use with `VPN`.

