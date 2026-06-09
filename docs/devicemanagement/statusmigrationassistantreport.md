# StatusMigrationAssistantReport

The status item that reports the state of a completed migration.

**Platforms:** macOS 26.4

## Properties

### migration-assistant.report

- **Type:** `StatusMigrationAssistantReportMigrationAssistantReportObject`
- **Required:** Yes

The Migration Assistant migration status.

## Discussion

### Status item availability

### Status item example

```json
{
    "migration-assistant": {
        "report": {
            "completed-data-size": 53687091200,
            "completed-file-count": 125000,
            "completion-time": "2025-05-15T14:30:00Z",
            "source-user": "user",
            "start-time": "2025-05-15T12:00:00Z",
            "target-user": "user",
            "total-data-size": 53687091200,
            "total-file-count": 125000
        }
    }
}
```

## Topics

### Objects

- [StatusMigrationAssistantReportMigrationAssistantReportObject](/documentation/devicemanagement/statusmigrationassistantreportmigrationassistantreportobject) - The Migration Assistant migration status.

