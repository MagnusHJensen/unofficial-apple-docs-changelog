# ResponseSubscription

A subscription with its assignment counts.

## Properties

### parentAdamId

- **Type:** `int64`
- **Required:** No

The parent Adam ID for the subscription.

### adamId

- **Type:** `int64`
- **Required:** No

The Adam ID for the subscription.

### status

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `ACTIVE`, `INACTIVE`, `PAUSED`

The current state of the subscription.

### periodEndDate

- **Type:** `string`
- **Required:** No

The end date of the current billing period in ISO-8601 calendar date format (`YYYY-MM-DD`).

### counts

- **Type:** `SubscriptionCounts`
- **Required:** No

The assignment counts for the subscription, broken down by renewal state.

## Topics

### Objects and Data Types

- [SubscriptionCounts](/documentation/devicemanagement/subscriptioncounts) - The subscription assignment counts broken down by assigned and available.

