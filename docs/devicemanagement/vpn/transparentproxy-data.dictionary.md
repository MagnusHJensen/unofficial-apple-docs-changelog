# VPN.TransparentProxy

The dictionary to use for a transparent proxy VPN type.

**Platforms:** macOS 14.0, Device Assignment Services , VPP License Management 

## Properties

### AuthenticationMethod

- **Type:** `string`
- **Required:** No
- **Default:** `Password`
- **Allowed Values:** `Password`, `Certificate`, `Password+Certificate`

The type of authentication method to use: `Password`, `Certificate`, or `Password+Certificate`.

Available in macOS 14 and later.

### DisconnectOnIdle

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, the VPN disconnects automatically disconnect after a period defined by `DisconnectOnIdleTimer`.

Available in macOS 14 and later.

### DisconnectOnIdleTimer

- **Type:** `integer`
- **Required:** No

The number of seconds before the VPN disconnects. This value is only used if `DisconnectOnIdle` is `1`.

Available in macOS 14 and later.

### EnforceRoutes

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, then all the VPN’s non-default routes take precedence over any locally-defined routes. If `IncludeAllNetworks` is `1`, the system ignores the value of `EnforceRoutes`.

Available in macOS 14 and later.

### OnDemandEnabled

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, the system brings up the VPN on demand.

Available in macOS 14 and later.

### OnDemandRules

- **Type:** `[VPN.VPN.OnDemandRulesElement]`
- **Required:** No

Determines when and how the system uses an OnDemand VPN.

Available in macOS 14 and later.

### Order

- **Type:** `integer`
- **Required:** No

A positive integer.

Available in macOS 14 and later.

### Password

- **Type:** `string`
- **Required:** No

The password to use for the account credentials. Only used if `AuthenticationMethod` is `Password`.

Available in macOS 14 and later.

### PayloadCertificateUUID

- **Type:** `string`
- **Required:** No

The UUID of the identity certificate as the account credential. If `AuthenticationMethod` is `Certificate`, and extended authentication (EAP) isn’t used, this certificate is sent out for IKE client authentication. If extended authentication is used, this certificate can be used for EAP-TLS.

Available in macOS 14 and later.

### ProviderBundleIdentifier

- **Type:** `string`
- **Required:** No

If the VPNSubType field contains the bundle identifier of an app that contains multiple VPN providers of the same type (app-proxy or packet-tunnel), then the system uses this field to choose which provider to use for this configuration. If the VPN provider is implemented as a System Extension, then this field is required.

Available in macOS 14 and later.

### ProviderDesignatedRequirement

- **Type:** `string`
- **Required:** No

If the VPN provider is implemented as a System Extension, then this field is required.

Available in macOS 14 and later.

### ProviderType

- **Type:** `string`
- **Required:** No
- **Default:** `packet-tunnel`
- **Allowed Values:** `packet-tunnel`, `app-proxy`

If the value of this key is `app-proxy`, the VPN service tunnels traffic at the application layer. If the value of this key is `packet-tunnel`, the VPN service tunnels traffic at the IP layer.

Available in macOS 14 and later.

