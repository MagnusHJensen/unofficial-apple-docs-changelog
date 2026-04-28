# NetworkProxyConfiguration.Proxies

The payload for configuring network proxies.

**Platforms:** macOS 10.7, Device Assignment Services , VPP License Management 

## Properties

### ExceptionsList

- **Type:** `[string]`
- **Required:** No

The list of hosts and domains that should bypass proxy settings.

### FallBackAllowed

- **Type:** `integer`
- **Required:** No

If `1`, enables fallback. Default is `1`.

For managed devices, if not supplied, the default is `0`.

### FTPEnable

- **Type:** `integer`
- **Required:** No

If `true`, enables FTP proxy.

### FTPPassive

- **Type:** `integer`
- **Required:** No

If `true`, enables passive FTP mode.

### FTPPort

- **Type:** `integer`
- **Required:** No

The FTP proxy port.

### FTPProxy

- **Type:** `string`
- **Required:** No

The host name or IP address for the FTP proxy.

### GopherEnable

- **Type:** `integer`
- **Required:** No

If `true`, enables gopher proxy.

### GopherPort

- **Type:** `integer`
- **Required:** No

The gopher proxy port.

### GopherProxy

- **Type:** `string`
- **Required:** No

The host name or IP address for the gopher proxy.

### HTTPEnable

- **Type:** `integer`
- **Required:** No

If `true`, enables web proxy.

### HTTPPort

- **Type:** `integer`
- **Required:** No

The web proxy port.

### HTTPProxy

- **Type:** `string`
- **Required:** No

The host name or IP address for the web proxy.

### HTTPSEnable

- **Type:** `integer`
- **Required:** No

If `true`, enables secure web proxy.

### HTTPSPort

- **Type:** `integer`
- **Required:** No

The secure web proxy port.

### HTTPSProxy

- **Type:** `string`
- **Required:** No

The host name or IP address for the secure web proxy.

### ProxyAutoConfigEnable

- **Type:** `integer`
- **Required:** No

If `true`, enables automatic proxy configuration.

### ProxyAutoConfigURLString

- **Type:** `string`
- **Required:** No

The automatic proxy configuration URL.

### ProxyCaptiveLoginAllowed

- **Type:** `integer`
- **Required:** No

If 1, allows client to log into captive portal network.

### RTSPEnable

- **Type:** `integer`
- **Required:** No

If `true`, enable streaming proxy.

### RTSPPort

- **Type:** `integer`
- **Required:** No

The streaming proxy port.

### RTSPProxy

- **Type:** `string`
- **Required:** No

The host name or IP address for the streaming proxy.

### SOCKSEnable

- **Type:** `integer`
- **Required:** No

If `true`, enable the SOCKS proxy.

### SOCKSPortinteger

- **Type:** `integer`
- **Required:** No

The SOCKS proxy port.

### SOCKSProxy

- **Type:** `string`
- **Required:** No

The host name or IP address for the SOCKS proxy.

