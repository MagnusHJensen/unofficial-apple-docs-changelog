# StatusMigrationAssistantReportMigrationAssistantReportObject

The Migration Assistant migration status.

**Platforms:** macOS 26.4

## Properties

### completed-data-size

- **Type:** `integer`
- **Required:** No

The total amount of data the system successfully migrated from the source system.

### completed-file-count

- **Type:** `integer`
- **Required:** No

The number of files successfully migrated from the source system.

### completion-time

- **Type:** `string`
- **Required:** No

The RFC 3339 timestamp for when the system completed migration.

### errors

- **Type:** `[string]`
- **Required:** No

The descriptions of migration errors that the system reports.

### source-user

- **Type:** `string`
- **Required:** No

The username of the user that the system migrated from the source system.

### start-time

- **Type:** `string`
- **Required:** No

The RFC 3339 timestamp for when the system started migration.

### target-user

- **Type:** `string`
- **Required:** No

The username of the target user account on the destination system.

### total-data-size

- **Type:** `integer`
- **Required:** No

The total amount of data the system considers in scope for migration from the source system.

### total-file-count

- **Type:** `integer`
- **Required:** No

The number of files the system considers in scope for migration from the source system.

