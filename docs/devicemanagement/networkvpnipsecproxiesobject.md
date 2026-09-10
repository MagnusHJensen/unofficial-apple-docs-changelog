# NetworkVPNIPSecProxiesObject

The dictionary to use to configure `Proxies` for use with `VPN`.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, visionOS 27.0

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

- **Type:** `NetworkVPNIPSecProxies_ProtocolObject`
- **Required:** No

The dictionary to use to configure HTTP servers  for `Proxies` for use with `VPN`.

### SupplementalMatchDomains

- **Type:** `[string]`
- **Required:** No

An array of domains that defines which hosts use proxy settings for hosts.

## Topics

### Objects

- [NetworkVPNIPSecProxies_ProtocolObject](/documentation/devicemanagement/networkvpnipsecproxies_protocolobject) - The dictionary to use to configure HTTP servers  for `Proxies` for use with `VPN`.

