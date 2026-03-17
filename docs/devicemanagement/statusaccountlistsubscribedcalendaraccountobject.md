# StatusAccountListSubscribedCalendarAccountObject

A status report of the client’s subscribed calendar details.

**Platforms:** iOS 16.0, iPadOS 16.0, macOS 14.0, visionOS 1.1

## Properties

### _removed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the subscribed calendar is removed and the status item object only contains this key and the `identifier` key.

### calendar-url

- **Type:** `string`
- **Required:** No

The URL of the subscribed calendar.

### declaration-identifier

- **Type:** `string`
- **Required:** No

The identifier of the declaration that installed the subscribed calendar. Only present if a declaration installed the subscribed calendar.

### identifier

- **Type:** `string`
- **Required:** Yes

The unique identifier for the subscribed calendar.

### is-enabled

- **Type:** `boolean`
- **Required:** No

A Boolean value that indicates whether the Calendar app displays this subscribed calendar.

### username

- **Type:** `string`
- **Required:** No

The user name for authenticating with the subscribed calendar.

### visible-name

- **Type:** `string`
- **Required:** No

The name of the subscribed calendar.

