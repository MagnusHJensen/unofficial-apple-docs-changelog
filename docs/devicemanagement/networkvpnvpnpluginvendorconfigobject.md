# NetworkVPNVPNPluginVendorConfigObject

The vendor-specific configuration dictionary, which the system reads only when `SubType` has a value.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, tvOS 27.0, visionOS 27.0

## Properties

### Group

- **Type:** `string`
- **Required:** No

The group to connect to on the head end. Valid for Cisco AnyConnect and Cisco Legacy AnyConnect.

### LoginGroupOrDomain

- **Type:** `string`
- **Required:** No

The login group or domain. Valid only for SonicWALL Mobile Connect.

### Realm

- **Type:** `string`
- **Required:** No

The Kerberos realm name, which needs to be properly capitalized. Valid only for Juniper SSL and Pulse Secure.

### Role

- **Type:** `string`
- **Required:** No

The role to select when connecting to the server. Valid only for Juniper SSL and Pulse Secure.

