# VPN.PPP

The dictionary that contains PPP settings.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, visionOS 1.0

## Properties

### AuthEAPPlugins

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `EAP-RSA`, `EAP-TLS`, `EAP-KRB`

An array of authentication plugins. For use of RSA SecurID, this array should only have one value: `EAP-RSA`. This key is for use with L2TP and PPTP networks.

### AuthName

- **Type:** `string`
- **Required:** No

The VPN account user name. This key is for use with L2TP and PPTP networks.

### AuthPassword

- **Type:** `string`
- **Required:** No

If `TokenCard` is `1`, use this password for authentication. This key is for use with L2TP and PPTP networks.

### AuthProtocol

- **Type:** `[string]`
- **Required:** No
- **Allowed Values:** `EAP`

An array of authentication protocols. For use of RSA SecurID, this array should have one value, `EAP`. This key is for use with L2TP and PPTP networks.

### CCPEnabled

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `0`, `1`

If `1`, enables encryption on the connection. This key is for use with PPTP networks.

### CCPMPPE128Enabled

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `0`, `1`

If `1` and `CCPEnabled` is also `1`, enables CCPMPPE40 encryption.

### CCPMPPE40Enabled

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `0`, `1`

If `1` and `CCPEnabled` is also `1`, enables CCPMPPE128 encryption.

### CommRemoteAddress

- **Type:** `string`
- **Required:** No

The IP address or host name of VPN server. This key is for use with L2TP and PPTP networks.

### DisconnectOnIdle

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, disconnects after an on demand connection idles.

### DisconnectOnIdleTimer

- **Type:** `integer`
- **Required:** No

The length of time to wait before disconnecting an on demand connection

### TokenCard

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, uses a token card such as an RSA SecurID card for connecting. This key is for use with L2TP networks.

