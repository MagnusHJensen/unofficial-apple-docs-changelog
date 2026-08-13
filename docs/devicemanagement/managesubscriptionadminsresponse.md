# ManageSubscriptionAdminsResponse

The confirmation response that the server returns after adding or removing subscription administrators.

## Properties

### tokenExpirationDate

- **Type:** `string`
- **Required:** No

The token expiration date in an ISO-8601 format.

Note: The server shows all dates and times in UTC.

### uId

- **Type:** `string`
- **Required:** No

The unique library identifier. When querying records using multiple tokens that may share libraries, use the `uId` field to filter duplicates and avoid double-counting records when different content managers upload duplicate tokens.

### versionId

- **Type:** `string`
- **Required:** No

The current version identifier.

