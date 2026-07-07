# StatusMigrationAssistantState

A status item that shows the device’s current migration state.

**Platforms:** macOS 26.4

## Properties

### migration-assistant.state

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `waiting`, `migrating`, `completed`, `failed`, `cancelled`, `unknown`

The current migration state of the system, which has the following possible values:

- `none`: Migration hasn’t started yet or no migration has taken place.
- `migrating`: Migration is in progress.
- `completed`: Migration has completed successfully.
- `failed`: Migration has failed.
- `cancelled`: The user cancelled migration.
- `unknown`: Migration status is unknown.

## Discussion

### Status item availability

### Status item example

```json
{
    "migration-assistant": {
        "state": "completed"
    }
}
```

