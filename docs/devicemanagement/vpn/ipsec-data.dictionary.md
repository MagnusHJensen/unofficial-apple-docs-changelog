# VPN.IPSec

The dictionary to use for an IPSec VPN type.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, visionOS 1.0

## Properties

### AuthenticationMethod

- **Type:** `string`
- **Required:** No
- **Default:** `SharedSecret`
- **Allowed Values:** `SharedSecret`, `Certificate`

The authentication method for L2TP and Cisco IPSec.

### DisconnectOnIdle

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, disconnect after an on-demand connection idles.

### DisconnectOnIdleTimer

- **Type:** `integer`
- **Required:** No

The length of time to wait before disconnecting an on-demand connection.

### LocalIdentifier

- **Type:** `string`
- **Required:** No

The name of the group. For hybrid authentication, the string needs to end with “hybrid”.

Present only for Cisco IPSec if `AuthenticationMethod` is `SharedSecret`.

### LocalIdentifierType

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `KeyID`

Present only if `AuthenticationMethod` is `SharedSecret`. The value is `KeyID`. The system uses this value for L2TP and Cisco IPSec VPNs.

### OnDemandEnabled

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`

If `1`, enables bringing the VPN connection up on demand.

### OnDemandMatchDomainsAlways

- **Type:** `[string]`
- **Required:** No

Deprecated. A list of domain names. In iOS 7 and later, if this key is present, the system treats associated domain names as though they’re associated with the `OnDemandMatchDomainsOnRetry` key. This behavior can be overridden by `OnDemandRules`.

### OnDemandMatchDomainsNever

- **Type:** `[string]`
- **Required:** No

Deprecated. A list of domain names. In iOS 7 and later, this key is deprecated (but still supported) in favor of `EvaluateConnection` actions in the `OnDemandRules` dictionaries.

### OnDemandMatchDomainsOnRetry

- **Type:** `[string]`
- **Required:** No

Deprecated. A list of domain names. In iOS 7 and later, this field is deprecated (but still supported) in favor of `EvaluateConnection` actions in the `OnDemandRules` dictionaries.

### OnDemandRules

- **Type:** `[VPN.VPN.OnDemandRulesElement]`
- **Required:** No

The on-demand rules dictionary.

### PayloadCertificateUUID

- **Type:** `string`
- **Required:** No

The UUID of the certificate payload within the same profile to use for the account credentials.

Only use this with Cisco IPSec VPNs and if the `AuthenticationMethod` key is to `Certificate`.

### PromptForVPNPIN

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, prompts for a PIN when connecting to Cisco IPSec VPNs.

### RemoteAddress

- **Type:** `string`
- **Required:** No

The IP address or host name of the VPN server.

### SharedSecret

- **Type:** `data`
- **Required:** No

The shared secret for this VPN account.

Only use this with L2TP and Cisco IPSec VPNs and if the `AuthenticationMethod` key is to `SharedSecret`.

### XAuthEnabled

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `0`, `1`

If `1`, enables Xauth for Cisco IPSec VPNs.

### XAuthName

- **Type:** `string`
- **Required:** No

The user name for the VPN account for Cisco IPSec.

### XAuthPassword

- **Type:** `string`
- **Required:** No

The VPN account password for Cisco IPSec.

### XAuthPasswordEncryption

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Prompt`

A string that either has the value “Prompt” or isn’t present.

