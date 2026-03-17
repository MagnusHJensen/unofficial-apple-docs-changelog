# DNSSettings.OnDemandRulesElement.ActionParameter

A dictionary that provides per-connection rules.

**Platforms:** iOS 14.0, iPadOS 14.0, macOS 11.0, visionOS 1.0

## Properties

### DomainAction

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `NeverConnect`, `ConnectIfNeeded`

The DNS settings behavior for the specified domains. Allowed values:

- ‘NeverConnect’: Don’t use the DNS Settings for the specified domains.
- ‘ConnectIfNeeded’: Allow using the DNS Settings for the specified domains.

### Domains

- **Type:** `[string]`
- **Required:** Yes

The domains for which this evaluation applies.

## Discussion

The keys allowed in each dictionary are described below. Note: This array is only for dictionaries in which `EvaluateConnection` is the `Action` value.

