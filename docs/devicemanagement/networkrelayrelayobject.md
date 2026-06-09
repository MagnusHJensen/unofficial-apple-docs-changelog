# NetworkRelayRelayObject

An array of dictionaries that describe one or more relay servers that the system can chain together.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### AdditionalHTTPHeaderFields

- **Type:** `NetworkRelayRelay_AdditionalHTTPHeaderFieldsObject`
- **Required:** No

A dictionary that contains custom HTTP header keys and values to add to each request. The dictionary key name represents the HTTP header field name to use, and the dictionary value is the string to use as the HTTP header field value.

### CredentialAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the identity to install.

### HTTP2RelayURL

- **Type:** `string`
- **Required:** No

The URL or URI template, as defined in RFC 9298, of a relay server that’s reachable using HTTP/2 and supports proxying TCP and UDP using the CONNECT method.

Each relay needs to include either `HTTP2RelayURL` or `HTTP3RelayURL`, or it can include both.

### HTTP3RelayURL

- **Type:** `string`
- **Required:** No

The URL or URI template, as defined in RFC 9298, of a relay server that’s reachable using HTTP/3 and supports proxying TCP and UDP using the CONNECT method.

Each relay needs to include either `HTTP2RelayURL` or `HTTP3RelayURL`, or it can include both.

### PublicKeyData

- **Type:** `[string]`
- **Required:** No

An array of references to data assets containing DER-encoded public key data that the system uses to authenticate the server during a TLS handshake. The server needs to use one of the keys in the handshake to authenticate. If this array is empty, the system uses the default TLS trust evaluation.

## Topics

### Objects

- [NetworkRelayRelay_AdditionalHTTPHeaderFieldsObject](/documentation/devicemanagement/networkrelayrelay_additionalhttpheaderfieldsobject) - A dictionary that contains custom HTTP header keys and values to add to each request. The dictionary key name represents the HTTP header field name to use, and the dictionary value is the string to use as the HTTP header field value.

