# ResponseSubscriptionAdmin

An administrator for a subscription.

## Properties

### adamId

- **Type:** `int64`
- **Required:** No

The Adam ID for the subscription that the user administers.

### clientUserId

- **Type:** `string`
- **Required:** No

The client user identifier for the administrator.

### idHash

- **Type:** `string`
- **Required:** No

The hash of the user’s identifier.

### userStatus

- **Type:** `string`
- **Required:** No

The current association state of the user. The server includes this key only when the request sets the `includeUserState` query parameter to `true`.

