# Cellular.AttachAPN

A dictionary that contains details about an attach access point name (APN) configuration.

**Platforms:** iOS 7.0, iPadOS 7.0, Mac Catalyst 7.0, watchOS 3.2, Device Assignment Services , VPP License Management 

## Properties

### AllowedProtocolMask

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `1`, `2`, `3`

The Internet Protocol versions that the system supports. Allowed values:

- `1`: IPv4
- `2`: IPv6
- `3`: Both

### AuthenticationType

- **Type:** `string`
- **Required:** No
- **Default:** `PAP`
- **Allowed Values:** `CHAP`, `PAP`

The authentication type.

### Name

- **Type:** `string`
- **Required:** Yes

The name for this configuration.

### Password

- **Type:** `string`
- **Required:** No

The password for the user.

### Username

- **Type:** `string`
- **Required:** No

The user name.

