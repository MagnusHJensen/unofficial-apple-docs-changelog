# GetSubscriptionAdminsResponse

The response that contains the requested subscription administrators.

## Properties

### admins

- **Type:** `[ResponseSubscriptionAdmin]`
- **Required:** No

The set of requested subscription administrators.

### nextCursor

- **Type:** `string`
- **Required:** No

The cursor for fetching the next page of results.

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

## Topics

### Objects and Data Types

- [ResponseSubscriptionAdmin](/documentation/devicemanagement/responsesubscriptionadmin) - An administrator for a subscription.

