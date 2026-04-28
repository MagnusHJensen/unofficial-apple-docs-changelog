# MigrationAssistantSettings

The declaration to configure Migration Assistant settings.

**Platforms:** macOS 26.4, Device Assignment Services , VPP License Management 

## Properties

### ExcludedAccounts

- **Type:** `[string]`
- **Required:** No

An array of strings that represent the user account short names the system excludes from migration.

### ExcludedPaths

- **Type:** `[string]`
- **Required:** No

An array of strings that represent files and directories relative to the user’s home directory that the system excludes from migration. Directory paths need to include a trailing “/”. For example, to exclude the “Excluded” directory in the “Documents” folder of a user’s home directory, use “Documents/Excluded/”.

### RequiredPaths

- **Type:** `[string]`
- **Required:** No

An array of strings that represent files and directories relative to the user’s home directory that the system needs to migrate. Directory paths need to include a trailing “/”. For example, to require the “Required” directory in the “Documents” folder of a user’s home directory, use “Documents/Required/”.

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

This declaration allows the device management service to configure Migration Assistant when it runs during Setup Assistant on a Mac. This makes it easy for users to do Mac-to-Mac migrations of enterprise devices when they setup a new Mac.

Configure the device to use the `AwaitingConfiguration` state after it enrolls with the server. The server needs to send the configuration and verify the configuration as both active and valid using the Declarative Device Management status, before it sends the [DeviceConfiguredCommand](/documentation/devicemanagement/deviceconfiguredcommand) command to exit that state.

The device reports Migration Assistant progress using the [StatusMigrationAssistantState](/documentation/devicemanagement/statusmigrationassistantstate) status item, and provides a report when migration completes using the [StatusMigrationAssistantReport](/documentation/devicemanagement/statusmigrationassistantreport) status item.

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
            "Documents/Personal Items/"
        ],
        "RequiredPaths": [
            "Documents/Work Items/"
        ],
        "ShouldMigrateSecurityPrivacySettings": false
    }
}
```

