# NetworkVPNIKEV2Proxies_Protocol_HTTPSObject

The dictionary to use to configure the HTTPS (TLS) server.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, tvOS 27.0, visionOS 27.0

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

