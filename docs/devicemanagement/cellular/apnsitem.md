# Cellular.APNsItem

A dictionary that contains details about an access point name (APN) configuration.

**Platforms:** iOS 7.0, iPadOS 7.0, watchOS 3.2

## Properties

### AllowedProtocolMask

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `1`, `2`, `3`

The Internet Protocol versions that the system supports. Available in iOS 10.3 and later. Allowed values:

- `1`: IPv4
- `2`: IPv6
- `3`: Both

### AllowedProtocolMaskInDomesticRoaming

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `1`, `2`, `3`

The Internet Protocol versions that the system supports while roaming. Available in iOS 10.3 and later. Allowed values:

- `1`: IPv4
- `2`: IPv6
- `3`: Both

### AllowedProtocolMaskInRoaming

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `1`, `2`, `3`

The Internet Protocol versions that the system supports while roaming. Available in iOS 10.3 and later. Allowed values:

- `1`: IPv4
- `2`: IPv6
- `3`: Both

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

The default Internet Protocol versions. Available in iOS 10.3 but no longer used in iOS 11 and later. Allowed values:

- `1`: IPv4
- `2`: IPv6
- `3`: Both

### EnableXLAT464

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system enables XLAT464. Available in iOS 16 and later and watchOS 9 and later.

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

