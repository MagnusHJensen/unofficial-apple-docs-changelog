# NetworkVPNIPSecProxies_Protocol_HTTPObject

The dictionary to use to configure the HTTP (non-TLS) server.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### Enable

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables proxy for HTTP traffic.

### HostName

- **Type:** `string`
- **Required:** No

The host name of the HTTP proxy.

### Port

- **Type:** `integer`
- **Required:** No

The port number of the HTTP proxy. This field is required if `HostName` is specified.

