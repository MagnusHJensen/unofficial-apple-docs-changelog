# StatusAccountListLDAPAccountObject

A status report of the client’s LDAP account details.

**Platforms:** iOS 16.0, iPadOS 16.0, Mac Catalyst 16.0, macOS 13.0, visionOS 1.1, Device Assignment Services , VPP License Management 

## Properties

### _removed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the account is removed and the status item object only contains this key and the `identifier` key.

### declaration-identifier

- **Type:** `string`
- **Required:** No

The identifier of the declaration that installed the account. Only present if a declaration installed the account.

### hostname

- **Type:** `string`
- **Required:** No

The server host name for the account.

### identifier

- **Type:** `string`
- **Required:** Yes

The unique identifier for the account.

### is-enabled

- **Type:** `boolean`
- **Required:** No

A Boolean value that indicates whether the account is enabled for use with the Contacts app.

### port

- **Type:** `integer`
- **Required:** No

The server port for the account.

### username

- **Type:** `string`
- **Required:** No

The user name for the account.

### visible-name

- **Type:** `string`
- **Required:** No

The name of the account.

