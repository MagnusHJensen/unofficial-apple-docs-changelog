# StatusScreenSharingConnectionGroupUnresolvedConnectionsUnresolvedGroupObject

A status item that contains an unresolved connection group.

**Platforms:** macOS 14.1

## Properties

### _removed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system removed the unresolved connection group and only this key and the `identifier` key are present in the status item object.

### identifier

- **Type:** `string`
- **Required:** Yes

The unique `ConnectionGroupUUID` identifier of the connection group.

### unresolved_connections

- **Type:** `[string]`
- **Required:** No

An array of `ConnectionUUID` values specified in the `Members` key in the group’s declaration for the unresolved connections.

