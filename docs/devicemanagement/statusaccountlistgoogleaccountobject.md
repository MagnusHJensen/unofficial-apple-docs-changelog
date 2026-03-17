# StatusAccountListGoogleAccountObject

A status report of the client’s Google account details.

**Platforms:** iOS 16.0, iPadOS 16.0, macOS 13.0, visionOS 1.1

## Properties

### _removed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the account is removed and the status item object only contains this key and the `identifier` key.

### are-calendars-enabled

- **Type:** `boolean`
- **Required:** No

A Boolean value that indicates whether the Calendar app displays calendars and events for this account.

### are-contacts-enabled

- **Type:** `boolean`
- **Required:** No

A Boolean value that indicates whether the Contacts app displays contacts for this account.

### are-notes-enabled

- **Type:** `boolean`
- **Required:** No

A Boolean value that indicates whether the Notes app displays notes for this account.

### declaration-identifier

- **Type:** `string`
- **Required:** No

The identifier of the declaration that installed the account. Only present if a declaration installed the account.

### identifier

- **Type:** `string`
- **Required:** Yes

The unique identifier for the account.

### is-mail-enabled

- **Type:** `boolean`
- **Required:** No

A Boolean value that indicates whether the Mail app displays mail for this account.

### username

- **Type:** `string`
- **Required:** No

The user name for the account.

### visible-name

- **Type:** `string`
- **Required:** No

The name of the account.

