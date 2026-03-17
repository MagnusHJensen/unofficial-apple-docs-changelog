# VPN.Proxies

The dictionary that contains the Proxies settings.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 17.0, visionOS 1.0, watchOS 10.0

## Properties

### HTTPEnable

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, enables proxy for HTTP traffic.

### HTTPPort

- **Type:** `integer`
- **Required:** No

The port number of the HTTP proxy. This field is required if `HTTPProxy` is specified.

### HTTPProxy

- **Type:** `string`
- **Required:** No

The host name of the HTTP proxy.

### HTTPProxyPassword

- **Type:** `string`
- **Required:** No

The password used for authentication.

### HTTPProxyUsername

- **Type:** `string`
- **Required:** No

The user name used for authentication.

### HTTPSEnable

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `true`, enables proxy for HTTPS traffic.

### HTTPSPort

- **Type:** `integer`
- **Required:** No

The port number of the HTTPS proxy. This field is required if `HTTPSProxy` is specified.

### HTTPSProxy

- **Type:** `string`
- **Required:** No

The host name of the HTTPS proxy.

### ProxyAutoConfigEnable

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `0`, `1`

If `true`, enables automatic proxy configuration.

### ProxyAutoConfigURLString

- **Type:** `string`
- **Required:** No

The URL to the location of the proxy auto-configuration file. Used only when `ProxyAutoConfigEnable` is `true`.

### ProxyAutoDiscoveryEnable

- **Type:** `integer`
- **Required:** No
- **Default:** `1`
- **Allowed Values:** `0`, `1`

If `true`, enables proxy auto discovery.

### SupplementalMatchDomains

- **Type:** `[string]`
- **Required:** No

An array of domains that defines which hosts use proxy settings for hosts.

