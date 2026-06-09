# NetworkVPNIPSecProxies_ProtocolObject

The dictionary to use to configure HTTP servers  for `Proxies` for use with `VPN`.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### CredentialsAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the credentials (user name and password) to authenticate with the proxy server.

### HTTP

- **Type:** `NetworkVPNIPSecProxies_Protocol_HTTPObject`
- **Required:** No

The dictionary to use to configure the HTTP (non-TLS) server.

### HTTPS

- **Type:** `NetworkVPNIPSecProxies_Protocol_HTTPSObject`
- **Required:** No

The dictionary to use to configure the HTTPS (TLS) server.

## Topics

### Objects

- [NetworkVPNIPSecProxies_Protocol_HTTPObject](/documentation/devicemanagement/networkvpnipsecproxies_protocol_httpobject) - The dictionary to use to configure the HTTP (non-TLS) server.
- [NetworkVPNIPSecProxies_Protocol_HTTPSObject](/documentation/devicemanagement/networkvpnipsecproxies_protocol_httpsobject) - The dictionary to use to configure the HTTPS (TLS) server.

