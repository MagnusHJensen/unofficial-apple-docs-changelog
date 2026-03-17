# StatusMigrationAssistantState

The current migration state of the system.

**Platforms:** macOS 26.4 (Beta)

## Properties

### migration-assistant.state

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `waiting`, `migrating`, `completed`, `failed`, `cancelled`, `unknown`

The current migration state of the system, which has the following possible values:

- `none`: Migration has not started yet or no migration has taken place.
- `migrating`: Migration is in progress.
- `completed`: Migration has completed successfully.
- `failed`: Migration has failed.
- `cancelled`: The user cancelled migration.
- `unknown`: Migration status is unknown.

## Discussion

### Status item availability

