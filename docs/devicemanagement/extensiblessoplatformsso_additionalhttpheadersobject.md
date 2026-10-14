# ExtensibleSSOPlatformSSO_AdditionalHTTPHeadersObject

A dictionary with keys and values that are strings. The keys are HTTP header names, and the values are HTTP header values. The device adds these headers to nonce, pre-authentication, login, key and key exchange requests for Platform SSO. The device limits the number of headers to 20, the length of the header name to 128 characters, and the length of the header value to 2048 characters.

**Platforms:** macOS 27.2 (Beta)

## Properties

### ANY

- **Type:** `string`
- **Required:** No

Keys and values to pass as HTTP headers to the identity provider.

