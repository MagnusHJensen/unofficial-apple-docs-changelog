# MigrationAssistantSettings

The declaration to configure Migration Assistant settings.

**Platforms:** macOS 26.4 (Beta)

## Properties

### ExcludedAccounts

- **Type:** `[string]`
- **Required:** No

An array of strings that represent the user account short names the system excludes from migration..

### ExcludedPaths

- **Type:** `[string]`
- **Required:** No

An array of strings that represent paths relative to the user’s home directory that the system excludes from migration. For example, to exclude the “Excluded” directory in the “Documents” folder of a user’s home directory, use “Documents/Excluded”.

### RequiredPaths

- **Type:** `[string]`
- **Required:** No

An array of strings that represent paths relative to the user’s home directory that the system excludes from migration. For example, to require the “Required” directory in the “Documents” folder of a user’s home directory, use “Documents/Required”.

### ShouldDoManagedMigration

- **Type:** `boolean`
- **Required:** Yes

If `true`, the device manages Migration Assistant.

### ShouldMigrateSecurityPrivacySettings

- **Type:** `boolean`
- **Required:** Yes

If `true`, the system migrates Security & Privacy settings.

## Discussion

Specify `com.apple.configuration.migration-assistant.settings` as the declaration type.

### Configuration availability

### Configuration example

This configuration provides settings for a Mac to Mac migration.

```json
{
    "Type": "com.apple.configuration.migration-assistant.settings",
    "Identifier": "F3CD2AD7-85AA-4FF3-9264-A737259FB55E",
    "ServerToken": "5AB2B98C-FCE9-4A33-88B3-ADB05F081F77",
    "Payload": {
        "ShouldDoManagedMigration": true,
        "ExcludedAccounts": [
            "admin"
        ],
        "ExcludedPaths": [
            "Documents/Personal Items"
        ],
        "RequiredPaths": [
            "Documents/Work Items"
        ],
        "ShouldMigrateSecurityPrivacySettings": false
    }
}
```

