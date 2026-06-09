# ManageSubscriptionsRequest

The request for subscription management.

## Properties

### adamIds

- **Type:** `[int64]`
- **Required:** Yes

The set of Adam IDs for the subscriptions to manage.

### clientUserIds

- **Type:** `[string]`
- **Required:** Yes

The set of identifiers for users in your organization.

### renewing

- **Type:** `boolean`
- **Required:** No

A Boolean value that indicates whether the subscription is renewing. Used for association operations.

### deferred

- **Type:** `boolean`
- **Required:** No

A Boolean value that indicates whether the disassociation is deferred. Used for disassociation operations.

