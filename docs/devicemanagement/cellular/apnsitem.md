# Cellular.APNsItem

A dictionary that contains details about an access point name (APN) configuration.

**Platforms:** iOS 7.0, iPadOS 7.0, Mac Catalyst 7.0, watchOS 3.2

## Properties

### AllowedProtocolMask

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `1`, `2`, `3`

The Internet Protocol versions that the system supports. Allowed values:

- `1`: IPv4
- `2`: IPv6
- `3`: Both

Available: iOS 10.3+ | iPadOS 10.3+ | watchOS 3.2+

### AllowedProtocolMaskInDomesticRoaming

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `1`, `2`, `3`

The Internet Protocol versions that the system supports while roaming. Allowed values:

- `1`: IPv4
- `2`: IPv6
- `3`: Both

Available: iOS 10.3+ | iPadOS 10.3+ | watchOS 3.2+

### AllowedProtocolMaskInRoaming

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `1`, `2`, `3`

The Internet Protocol versions that the system supports while roaming. Allowed values:

- `1`: IPv4
- `2`: IPv6
- `3`: Both

Available: iOS 10.3+ | iPadOS 10.3+ | watchOS 3.2+

### AuthenticationType

- **Type:** `string`
- **Required:** No
- **Default:** `PAP`
- **Allowed Values:** `CHAP`, `PAP`

The authentication type for logging in.

### DefaultProtocolMask

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `1`, `2`, `3`

The default Internet Protocol versions. Allowed values:

- `1`: IPv4
- `2`: IPv6
- `3`: Both

Available: iOS 10.3+ | iPadOS 10.3+ | watchOS 3.2+
Deprecated: iOS 11+ | iPadOS 11+

### EnableXLAT464

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables XLAT464.

Available: iOS 16+ | iPadOS 16+ | watchOS 9+

### Name

- **Type:** `string`
- **Required:** Yes

The name for this configuration.

### Password

- **Type:** `string`
- **Required:** No

The user’s password for the APN.

### ProxyPort

- **Type:** `integer`
- **Required:** No

The proxy server’s port number.

### ProxyServer

- **Type:** `string`
- **Required:** No

The proxy server’s address.

### Username

- **Type:** `string`
- **Required:** No

The user name for the APN.

