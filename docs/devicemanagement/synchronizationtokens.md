# SynchronizationTokens

The server’s synchronization token.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### DeclarationsToken

- **Type:** `string`
- **Required:** Yes

The synchronization token for declarations.

### Timestamp

- **Type:** `date-time`
- **Required:** Yes

The timestamp for the creation of the set of sync tokens. Clients use this to determine the most recent set of sync tokens when different sources provide the tokens. Use the format `YYYY-mm-ddTHH:MM:SSZ`.

