# NetworkVPNVPNPluginProxies_Protocol_HTTPSObject

The dictionary to use to configure the HTTPS (TLS) server.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), tvOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### Enable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables proxy for HTTPS traffic.

### HostName

- **Type:** `string`
- **Required:** No

The host name of the HTTPS proxy.

### Port

- **Type:** `integer`
- **Required:** No

The port number of the HTTPS proxy. This field is required if `HostName` is specified.

