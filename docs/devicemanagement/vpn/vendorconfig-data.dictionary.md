# VPN.VendorConfig

The vendor-specific configuration dictionary.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 17.0, visionOS 1.0

## Properties

### Group

- **Type:** `string`
- **Required:** No

The group to connect to on the head end. Valid for Cisco AnyConnect and Cisco Legacy AnyConnect. Not available in watchOS.

### LoginGroupOrDomain

- **Type:** `string`
- **Required:** No

The login group or domain. Valid only for SonicWALL Mobile Connect. Not available in watchOS.

### Realm

- **Type:** `string`
- **Required:** No

The Kerberos realm name, which needs to be properly capitalized. Valid only for Juniper SSL and Pulse Secure. Not available in watchOS.

### Role

- **Type:** `string`
- **Required:** No

The role to select when connecting to the server. Valid only for Juniper SSL and Pulse Secure. Not available in watchOS.

