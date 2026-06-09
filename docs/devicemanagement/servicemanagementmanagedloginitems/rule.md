# ServiceManagementManagedLoginItems.Rule

A dictionary that configures a service management rule.

**Platforms:** macOS 13.0

## Properties

### Comment

- **Type:** `string`
- **Required:** No

An optional description of the rule.

### RuleType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `BundleIdentifier`, `BundleIdentifierPrefix`, `Label`, `LabelPrefix`, `TeamIdentifier`

The type of comparison to make.

### RuleValue

- **Type:** `string`
- **Required:** Yes

The value to compare with each login item’s value, to determine if this rule is a match.

### TeamIdentifier

- **Type:** `string`
- **Required:** No

An additional constraint to limit the scope of the rule that the system tests after matching the `RuleType` and `RuleValue`.

