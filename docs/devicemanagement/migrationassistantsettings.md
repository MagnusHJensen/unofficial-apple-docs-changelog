# MigrationAssistantSettings

The declaration to configure Migration Assistant settings.

**Platforms:** macOS 26.4 (Beta)

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

