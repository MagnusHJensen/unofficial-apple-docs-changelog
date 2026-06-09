# NetworkDNSSettingsOnDemandRulesElement_ActionParameterObject

A dictionary that provides per-connection rules.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), visionOS 27.0 (Beta)

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

