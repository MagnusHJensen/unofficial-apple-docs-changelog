# NetworkUsageRules.ApplicationRulesItem

The application rules dictionary.

**Platforms:** iOS 9.0, iPadOS 9.0, Mac Catalyst 9.0, Device Assignment Services , VPP License Management 

## Properties

### AllowCellularData

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables cellular data for all matching managed apps.

### AllowRoamingCellularData

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `false`, disables cellular data while roaming for all matching managed apps.

### AppIdentifierMatches

- **Type:** `[string]`
- **Required:** No

A list of managed app identifiers, as strings, that must follow the associated rules. If this key is missing, the rules apply to all managed apps on the device.

Each string in the `AppIdentifierMatches` array may either be an exact app identifier match (for example, `com.mycompany.myapp`) or it may specify a prefix match for the bundle ID by using the * wildcard character. If used, this character must appear after a period (.) and may only appear once, at the end of the string; for example, `com.mycompany.*`.

